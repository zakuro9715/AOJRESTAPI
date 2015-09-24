package api

import (
	"github.com/zakuro9715/AOJRESTAPI/aoj"
	timeutil "github.com/zakuro9715/AOJRESTAPI/util/time"
	"time"
)

type User struct {
	Id              string
	Name            string
	Affliation      string
	RegisteredAt    time.Time
	LastSubmittedAt time.Time
}

func NewUser(au *aoj.User) *User {
	registerD := time.Duration(au.RegisterDate) * time.Millisecond
	submitD := time.Duration(au.LastSubmitDate) * time.Millisecond
	u := new(User)
	u.Id = au.Id
	u.Name = au.Name
	u.Affliation = au.Affliation
	u.RegisteredAt = timeutil.FromUnixTime(&registerD, timeutil.JST)
	u.LastSubmittedAt = timeutil.FromUnixTime(&submitD, timeutil.JST)
	return u
}

func FetchUser(userId string) (*User, error) {
	aojUser, err := aoj.UserSearchApi(userId)
	return NewUser(aojUser), err
}
