package api

import (
	"github.com/zakuro9715/AOJRESTAPI/aoj"
	"strconv"
	"strings"
)

type Problem struct {
	Id          string
	Name        string
	Judge       string
	TimeLimit   int
	MemoryLimit int
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

func NewProblem(ap *aoj.Problem) *Problem {
	p := new(Problem)

	// ProblemSearchApi returns value containing newlines like a <id>\n0000\n</id>

	judgeNum, _ := strconv.Atoi(strings.Trim(ap.Available, "\n"))
	timeLimit, _ := strconv.Atoi(strings.Trim(ap.ProblemTimeLimit, "\n"))
	memoryLimit, _ := strconv.Atoi(strings.Trim(ap.ProblemMemoryLimit, "\n"))

	p.Id = strings.Trim(ap.Id, "\n")
	p.Name = strings.Trim(ap.Name, "\n")
	p.Judge = judge[judgeNum]
	p.TimeLimit = timeLimit
	p.MemoryLimit = memoryLimit
	return p
}

func FetchProblem(id string) (*Problem, error) {
	ap, err := aoj.ProblemSearchApi(id, false)
	return NewProblem(ap), err
}
