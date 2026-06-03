package model

type MasterUiRequest struct {
	ClienId string `json:"clientId"`
	UserId  string `json:"userId"`
	User    User   `json:"user"`
}
