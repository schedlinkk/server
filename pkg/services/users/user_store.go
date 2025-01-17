package users

type UserStore interface {
	StoreUser(User) error
	FetchUserByUID(UID) (User, error)
}
