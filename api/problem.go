package api

import (
	"github.com/zakuro9715/AOJRESTAPI/aoj"
	"strconv"
	"strings"
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
	judgeNum, _ := strconv.Atoi(strings.Trim(ap.Available, "\n"))
	timeLimit, _ := strconv.Atoi(strings.Trim(ap.ProblemTimeLimit, "\n"))
	memoryLimit, _ := strconv.Atoi(strings.Trim(ap.ProblemMemoryLimit, "\n"))

	return &ProblemCore{
		Id:          strings.Trim(ap.Id, "\n"),
		Name:        strings.Trim(ap.Name, "\n"),
		Judge:       judge[judgeNum],
		TimeLimit:   timeLimit,
		MemoryLimit: memoryLimit,
	}
}

func NewProblem(ap *aoj.Problem) *Problem {
	return &Problem{NewProblemCore(ap)}
}

func FetchProblem(id string) (*Problem, error) {
	ap, err := aoj.ProblemSearchApi(id, false)
	return NewProblem(ap), err
}
