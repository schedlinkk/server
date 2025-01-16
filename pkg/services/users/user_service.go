package users

func NewUserServiceWithStore(store UserStore) *UserService {
	um := &UserService{store: store}
	return um
}

type UserService struct {
	store UserStore
}

func (um UserService) FetchUserByUID(uid UID) (User, error) {
	return um.store.FetchUserByUID(uid)
}

func (um UserService) NewUser(user User) error {
	if err := um.store.StoreUser(user); err != nil {
		return err
	}

	return nil
}
