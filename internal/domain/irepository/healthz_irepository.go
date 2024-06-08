package irepository

type IHealthzRepository interface {
	GetHealthz() (string, error)
}
