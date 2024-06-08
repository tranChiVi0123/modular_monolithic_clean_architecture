package usecase

import "github.com/tranChiVi0123/modular_monolithic_clean_architecture/internal/domain/irepository"

type HealthzUsecase struct {
	healthzRepository irepository.IHealthzRepository
}

func NewHealthzUsecase(healthzRepository irepository.IHealthzRepository) HealthzUsecase {
	return HealthzUsecase{
		healthzRepository: healthzRepository,
	}
}

func (u HealthzUsecase) GetHealthz() (string, error) {
	healthz, err := u.healthzRepository.GetHealthz()
	if err != nil {
		return "", err
	}

	return healthz, nil
}
