package users

import (
	"time"
)

type UserStoreStub struct {
	Users []User
}

func (uss *UserStoreStub) FetchUserByUID(uid UID) (User, error) {
	var user User

	for _, u := range uss.Users {
		if u.UID == uid {
			return u, nil
		}
	}

	return user, ErrUserDontExist

}

func (uss *UserStoreStub) StoreUser(user User) error {
	for _, u := range uss.Users {
		if u.UID == user.UID {
			return ErrUserAlreadyExist
		}
	}

	uss.Users = append(uss.Users, user)

	return nil
}

func NewUserStoreStubWithSampleData() *UserStoreStub {
	store := &UserStoreStub{}

	users := []User{
		{
			UID:           "user-1",
			Username:      "user1",
			Friends:       []string{"user8", "user6", "user10"},
			CreatedAt:     time.Date(2024, 9, 5, 0, 0, 0, 0, time.UTC),
			DayPlansCount: 8,
		},
		{
			UID:           "user-2",
			Username:      "user2",
			Friends:       []string{"user8", "user5", "user10", "user10"},
			CreatedAt:     time.Date(2024, 8, 21, 0, 0, 0, 0, time.UTC),
			DayPlansCount: 18,
		},
		{
			UID:           "user-3",
			Username:      "user3",
			Friends:       []string{"user6", "user8", "user6"},
			CreatedAt:     time.Date(2024, 6, 17, 0, 0, 0, 0, time.UTC),
			DayPlansCount: 2,
		},
		{
			UID:           "user-4",
			Username:      "user4",
			Friends:       []string{"user10", "user10", "user1", "user9"},
			CreatedAt:     time.Date(2024, 2, 5, 0, 0, 0, 0, time.UTC),
			DayPlansCount: 5,
		},
		{
			UID:           "user-5",
			Username:      "user5",
			Friends:       []string{"user4", "user9", "user9", "user2"},
			CreatedAt:     time.Date(2024, 3, 22, 0, 0, 0, 0, time.UTC),
			DayPlansCount: 1,
		},
		{
			UID:           "user-6",
			Username:      "user6",
			Friends:       []string{"user6", "user5"},
			CreatedAt:     time.Date(2024, 4, 6, 0, 0, 0, 0, time.UTC),
			DayPlansCount: 7,
		},
		{
			UID:           "user-7",
			Username:      "user7",
			Friends:       []string{"user9", "user7", "user2"},
			CreatedAt:     time.Date(2024, 11, 23, 0, 0, 0, 0, time.UTC),
			DayPlansCount: 5,
		},
		{
			UID:           "user-8",
			Username:      "user8",
			Friends:       []string{"user5", "user2", "user4", "user3"},
			CreatedAt:     time.Date(2024, 10, 12, 0, 0, 0, 0, time.UTC),
			DayPlansCount: 9,
		},
		{
			UID:           "user-9",
			Username:      "user9",
			Friends:       []string{"user9", "user2", "user9", "user7"},
			CreatedAt:     time.Date(2024, 12, 25, 0, 0, 0, 0, time.UTC),
			DayPlansCount: 7,
		},
		{
			UID:           "user-10",
			Username:      "user10",
			Friends:       []string{"user5", "user8", "user6"},
			CreatedAt:     time.Date(2024, 3, 17, 0, 0, 0, 0, time.UTC),
			DayPlansCount: 13,
		},
	}

	store.Users = users

	return store
}
