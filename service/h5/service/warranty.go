package service

import (
	"h5/model"
	"h5/repository"
)

type WarrantyCardService struct {
	repo *repository.WarrantyCardRepo
}

func NewWarrantyCardService(repo *repository.WarrantyCardRepo) *WarrantyCardService {
	return &WarrantyCardService{repo: repo}
}

func (s *WarrantyCardService) CreateCard(card *model.WarrantyCard) error {
	return s.repo.Create(card)
}

func (s *WarrantyCardService) GetCardList(params map[string]interface{})([]model.WarrantyCard, error) {
	return s.repo.GetList(params)
}

func (s *WarrantyCardService) GetCardByID(id uint64) (*model.WarrantyCard, error) {
	return s.repo.GetByID(id)
}

type InstitutionService struct {
	repo *repository.InstitutionRepo
}

func NewInstitutionService(repo *repository.InstitutionRepo) *InstitutionService {
	return &InstitutionService{repo: repo}
}

// 根据token获取机构信息
func (s *InstitutionService) GetInstitutionByToken(token string) (*model.Institution, error) {
	// 1. 先通过token获取认证信息
	auth, err := s.repo.GetAuthByToken(token)
	if err != nil {
		return nil, err
	}

	// 2. 再通过认证信息中的机构ID获取机构详情
	return s.repo.GetInstitutionByID(auth.InstitutionID)
}