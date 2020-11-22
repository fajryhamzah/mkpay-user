package user

//RepoInterface user repository interface
type RepoInterface interface {
	FindByCode(code string) ModelInterface
	FindByEmail(email string) ModelInterface
	Save(user ModelInterface) error
	Update(id uint32, user ModelInterface) error
	Delete(id uint32) error
}
