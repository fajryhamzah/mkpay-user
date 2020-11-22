package auth

type authHandlerTransformer struct {
	Token   string `json:"access_token"`
	Refresh string `json:"refresh_token"`
}
