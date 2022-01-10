package tg_error_reporter

type Reporter interface {
	Report(Error)
}
