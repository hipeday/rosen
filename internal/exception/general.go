package exception

type RosenError interface {
	error
	Status() int
}
