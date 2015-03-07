package aoj

import (
	"../util/http"
	"encoding/xml"
	"net/url"
	"strconv"
)

type Problem struct {
	Id                 string          `xml:"id"`
	Name               string          `xml:"name"`
	Available          string          `xml:"available"`
	ProblemTimeLimit   string          `xml:"problemtimelimit"`
	ProblemMemoryLimit string          `xml:"problemmemorylimit"`
	Status             Status          `xml:"status"`
	SolvedList         []SolvedSummary `xml:"solved_list>user"`
}

const (
	ProblemSearchUrl = "http://judge.u-aizu.ac.jp/onlinejudge/webservice/problem"
)

func ProblemSearchApi(id string, status bool) (*Problem, error) {
	q := url.Values{"id": {id}, "status": {strconv.FormatBool(status)}}.Encode()
	data, err := http.Get(ProblemSearchUrl, q)
	if err != nil {
		return nil, err
	}
	p := new(Problem)
	err = xml.Unmarshal(data, p)
	return p, err
}
