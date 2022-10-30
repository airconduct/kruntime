package api

type Actor interface {
	Receive(ctx Context)
}
