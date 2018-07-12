package models

// IAMCredentials represents aws iam credentials
type IAMCredentials struct {
	AccessID  string `json:"accessId"`
	SecretKey string `json:"secretKey"`
}
