package api

import (
	"github.com/zakuro9715/AOJRESTAPI/aoj"
)

type ProblemCore struct {
	Id          string
	Name        string
	Judge       string
	TimeLimit   int
	MemoryLimit int
}

type Problem struct {
	*ProblemCore
}

var (
	judge = []string{
		"Not available",
		"Judge",
		"Judge allow precision error",
		"Judge with validator",
		"Reactive judge",
	}
)

func NewProblemCore(ap *aoj.Problem) *ProblemCore {
	// ProblemSearchApi returns value containing newlines like a <id>\n0000\n</id>

	return &ProblemCore{
		Id:          ap.Id,
		Name:        ap.Name,
		Judge:       judge[ap.Available],
		TimeLimit:   ap.ProblemTimeLimit,
		MemoryLimit: ap.ProblemMemoryLimit,
	}
}

func NewProblem(ap *aoj.Problem) *Problem {
	return &Problem{NewProblemCore(ap)}
}

func FetchProblem(id string) (*Problem, error) {
	ap, err := aoj.ProblemSearchApi(id, false)
	return NewProblem(ap), err
}
