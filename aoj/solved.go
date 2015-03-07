package aoj

type SolvedSummary struct {
	Id             string `xml:"id"`
	SubmissionDate int64  `xml:"submissiondate"`
	Language       string `xml:"language"`
	CpuTime        int    `xml:"cputime"`
	Memory         int    `xml:"memory"`
	CodeSize       int    `xml:"code_size"`
}
