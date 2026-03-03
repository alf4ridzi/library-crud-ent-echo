package dto

type AuthJwtResponse struct {
	Auth    string `json:"auth"`
	Refresh string `json:"refresh"`
}

type AuthJwt struct {
	Token AuthJwtResponse `json:"token"`
}
