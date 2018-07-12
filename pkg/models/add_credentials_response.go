package models

// AddCredentialsResponse represents aws iam credentials
type AddCredentialsResponse struct {
	Username string `json:"username"`
	UserARN  string `json:"userarn"`
	//cluster credentials also?
}
