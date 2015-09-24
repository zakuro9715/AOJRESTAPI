package aoj

type Status struct {
	Submission   int `xml:"submission"`
	Solved       int `xml:"solved"`
	Accepted     int `xml:"accepted"`
	WrongAnswer  int `xml:"wronganswer"`
	TimeLimit    int `xml:"timelimit"`
	MemoryLimit  int `xml:"memorylimt"`
	OutputLimit  int `xml:"outputlimit"`
	RuntimeError int `xml:"runtimeerror"`
	CompileError int `xml:"compileerror"`
}
