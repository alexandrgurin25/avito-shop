package auth_handler

type AuthDtoIn struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type AuthDtoOut struct {
	AccessToken string `json:"accessToken"`
}
