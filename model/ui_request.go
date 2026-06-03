package model

type MasterUiRequest struct {
	TenantId string `json:"tenantId"`
	UserId   string `json:"userId"`
}

type UserRequest struct {
	*MasterUiRequest `json:"masterRequest"`
	User             User `json:"user"`
}

// CategoryRequest accepts tenant/user at the root plus a category object.
// Embedding MasterUiRequest allows `tenantId` and `userId` to be provided
// at the top-level of the JSON payload.
type CategoryRequest struct {
	MasterUiRequest
	Category Category `json:"category"`
}
