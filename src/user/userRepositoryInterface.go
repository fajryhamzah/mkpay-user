package user

//RepoInterface user repository interface
type RepoInterface interface {
	//	GetBy(where map[string]string) []ModelInterface
	FindByCode(code string) ModelInterface
	//	Save(user ModelInterface) bool
	//	Update(id int, user ModelInterface) bool
	//	Delete(id int) bool
}
