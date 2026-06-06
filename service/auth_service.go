package service

import (
	"errors"
	"os"
	"strconv"
	"strings"
	"time"
	"user-management/model"
	"user-management/repository"

	"github.com/golang-jwt/jwt/v5"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

type AuthService interface {
	Register(req model.RegisterRequest) (model.AuthResponse, error)
	Login(req model.LoginRequest) (model.AuthResponse, error)
	ValidateToken(token string) (*model.AuthClaims, error)
}

type authService struct {
	userRepo  repository.UserRepository
	jwtSecret string
	tokenTTL  time.Duration
}

func NewAuthService(userRepo repository.UserRepository, jwtSecret string) AuthService {
	expiryMinutes := 15
	if raw := os.Getenv("JWT_EXPIRY_MINUTES"); raw != "" {
		if parsed, err := strconv.Atoi(raw); err == nil && parsed > 0 {
			expiryMinutes = parsed
		}
	}

	return &authService{
		userRepo:  userRepo,
		jwtSecret: jwtSecret,
		tokenTTL:  time.Duration(expiryMinutes) * time.Minute,
	}
}

func (s *authService) Register(req model.RegisterRequest) (model.AuthResponse, error) {
	if req.Name == "" || req.Email == "" || req.Password == "" {
		return model.AuthResponse{}, errors.New("name, email and password are required")
	}

	email := strings.TrimSpace(strings.ToLower(req.Email))
	existing, err := s.userRepo.GetUserByEmail(email)
	if err == nil && existing.ID != primitive.NilObjectID {
		return model.AuthResponse{}, errors.New("user already exists")
	}
	if err != nil && !errors.Is(err, mongo.ErrNoDocuments) {
		return model.AuthResponse{}, err
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return model.AuthResponse{}, err
	}

	user := model.User{
		ID:           primitive.NewObjectID(),
		Name:         strings.TrimSpace(req.Name),
		Email:        email,
		PasswordHash: string(hash),
		Role:         "user",
		CreatedAt:    time.Now().UTC(),
		UpdatedAt:    time.Now().UTC(),
	}

	if _, err := s.userRepo.SaveNewUser(user); err != nil {
		return model.AuthResponse{}, err
	}

	token, expiresAt, err := s.createToken(user)
	if err != nil {
		return model.AuthResponse{}, err
	}

	return model.AuthResponse{Token: token, ExpiresAt: expiresAt, User: user}, nil
}

func (s *authService) Login(req model.LoginRequest) (model.AuthResponse, error) {
	if req.Email == "" || req.Password == "" {
		return model.AuthResponse{}, errors.New("email and password are required")
	}

	email := strings.TrimSpace(strings.ToLower(req.Email))
	user, err := s.userRepo.GetUserByEmail(email)
	if err != nil {
		return model.AuthResponse{}, errors.New("invalid email or password")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(req.Password)); err != nil {
		return model.AuthResponse{}, errors.New("invalid email or password")
	}

	token, expiresAt, err := s.createToken(user)
	if err != nil {
		return model.AuthResponse{}, err
	}

	return model.AuthResponse{Token: token, ExpiresAt: expiresAt, User: user}, nil
}

func (s *authService) ValidateToken(tokenString string) (*model.AuthClaims, error) {
	claims := &model.AuthClaims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(s.jwtSecret), nil
	})
	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, errors.New("invalid token")
	}

	return claims, nil
}

func (s *authService) createToken(user model.User) (string, int64, error) {
	expiresAt := time.Now().Add(s.tokenTTL).UTC()
	claims := &model.AuthClaims{
		UserID: user.ID.Hex(),
		Email:  user.Email,
		Role:   user.Role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expiresAt),
			IssuedAt:  jwt.NewNumericDate(time.Now().UTC()),
			Subject:   user.ID.Hex(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signed, err := token.SignedString([]byte(s.jwtSecret))
	if err != nil {
		return "", 0, err
	}

	return signed, expiresAt.Unix(), nil
}
