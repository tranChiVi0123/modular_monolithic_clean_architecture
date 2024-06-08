package repository

type HealthzRepository struct {
}

func NewHealthzRepository() HealthzRepository {
	return HealthzRepository{}
}

func (hr HealthzRepository) GetHealthz() (string, error) {
	return "OK", nil
}
