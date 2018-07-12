package models

// AuthReponse represents login attempt response
type AuthReponse struct {
	Username      string `json:"username"`
	Authenticated bool   `json:"authenticated"`
	GivenName     string `json:"givenName"`
	Picture       string `json:"pictureUrl"`
	Cluster       string `json:"clusterName"`
}
