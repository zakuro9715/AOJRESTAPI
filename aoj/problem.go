package aoj

import (
	"encoding/xml"
	"github.com/zakuro9715/AOJRESTAPI/util/http"
	"net/url"
	"strconv"
)

type Problem struct {
	Id                 string          `xml:"id"`
	Name               string          `xml:"name"`
	Available          int             `xml:"available"`
	ProblemTimeLimit   int             `xml:"problemtimelimit"`
	ProblemMemoryLimit int             `xml:"problemmemorylimit"`
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
