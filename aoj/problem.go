package aoj

import (
	"encoding/xml"
	"github.com/zakuro9715/AOJRESTAPI/util/http"
	"net/url"
	"strconv"
)

type Problem struct {
	Id                 string       `xml:"id"`
	Name               string       `xml:"name"`
	Available          int          `xml:"available"`
	ProblemTimeLimit   int          `xml:"problemtimelimit"`
	ProblemMemoryLimit int          `xml:"problemmemorylimit"`
	Status             Status       `xml:"status"`
	SolvedList         []SolvedUser `xml:"solved_list>user"`
}

type SolvedUser struct {
	Id             string `xml:"id"`
	SubmissionDate int64  `xml:"submissiondate"`
	Language       string `xml:"language"`
	CpuTime        int    `xml:"cuptime"`
	Memory         int    `xml:"memory"`
	CodeSize       int    `xml:"code_size"`
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
