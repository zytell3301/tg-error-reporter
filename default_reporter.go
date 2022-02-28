package ErrorReporter

import "fmt"

type DefaultReporter struct {
}

func (d DefaultReporter) Report(err Error) {
	fmt.Println(err)
}

func NewDefaultReporter() DefaultReporter {
	return DefaultReporter{}
}
