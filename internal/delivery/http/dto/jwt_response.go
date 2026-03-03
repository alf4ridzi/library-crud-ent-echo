package dto

type AuthJwtResponse struct {
	Access  string `json:"access"`
	Refresh string `json:"refresh"`
}

type AuthJwt struct {
	Token AuthJwtResponse `json:"token"`
}
