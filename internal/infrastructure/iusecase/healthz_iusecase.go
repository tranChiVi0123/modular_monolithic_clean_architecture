package iusecase

type IHealthzUsecase interface {
  GetHealthz() (string, error)
}
