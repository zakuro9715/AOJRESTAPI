package aoj

import (
	"encoding/xml"
	"github.com/zakuro9715/AOJRESTAPI/util/http"
	"net/url"
)

type User struct {
	Id             string          `xml:"id"`
	Name           string          `xml:"name"`
	Affliation     string          `xml:"affliation"`
	RegisterDate   int64           `xml:"registerdate"`
	LastSubmitDate int64           `xml:"lastsubimtdate"`
	Status         Status          `xml:"status"`
	SolvedList     []SolvedSummary `xml:"solved_list>problem"`
}

const (
	UserSearchUrl = "http://judge.u-aizu.ac.jp/onlinejudge/webservice/user"
)

func UserSearchApi(userId string) (*User, error) {
	data, err := http.Get(UserSearchUrl, url.Values{"id": {userId}}.Encode())
	if err != nil {
		return nil, err
	}

	u := new(User)
	err = xml.Unmarshal(data, u)
	return u, err
}
