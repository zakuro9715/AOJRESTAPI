package api

import (
	"github.com/zakuro9715/AOJRESTAPI/aoj"
	timeutil "github.com/zakuro9715/AOJRESTAPI/util/time"
	"time"
)

type UserCore struct {
	Id              string
	Name            string
	Affliation      string
	RegisteredAt    time.Time
	LastSubmittedAt time.Time
}

type User struct {
	*UserCore
}

func NewUserCore(au *aoj.User) *UserCore {
	registerD := time.Duration(au.RegisterDate) * time.Millisecond
	submitD := time.Duration(au.LastSubmitDate) * time.Millisecond
	return &UserCore{
		Id:              au.Id,
		Name:            au.Name,
		Affliation:      au.Affliation,
		RegisteredAt:    timeutil.FromUnixTime(&registerD, timeutil.JST),
		LastSubmittedAt: timeutil.FromUnixTime(&submitD, timeutil.JST),
	}
}

func NewUser(au *aoj.User) *User {
	return &User{NewUserCore(au)}
}

func FetchUser(userId string) (*User, error) {
	aojUser, err := aoj.UserSearchApi(userId)
	return NewUser(aojUser), err
}
