package usecase

import "github.com/FlezzProject/platform-api/internal/domain/irepository"

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
