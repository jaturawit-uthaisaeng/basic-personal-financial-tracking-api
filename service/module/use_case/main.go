package use_case

import "basic-personal-financial-tracking-api/service/module/domain"

type newUseCase struct {
	repo domain.PersonalFinancialTrackingRepository
}

func NewUseCase(repo domain.PersonalFinancialTrackingRepository) domain.PersonalFinancialTrackingUseCase {
	return &newUseCase{repo: repo}
}
