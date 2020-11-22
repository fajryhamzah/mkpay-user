package jwt

type JwtPayload struct {
	Code string
	Role string
}

func (p JwtPayload) GetCode() string {
	return p.Code
}

func (p JwtPayload) GetRole() string {
	return p.Role
}
