package ErrorReporter

type DefaultReporter struct {
}

func (d DefaultReporter) Report(err Error) {
}

func NewDefaultReporter() DefaultReporter {
	return DefaultReporter{}
}
