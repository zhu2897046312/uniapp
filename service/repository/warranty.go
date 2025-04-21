package repository

import (

	"gorm.io/gorm"
	"h5/model"
)

type WarrantyCardRepo struct {
	db *gorm.DB
}

func NewWarrantyCardRepo(db *gorm.DB) *WarrantyCardRepo {
	return &WarrantyCardRepo{db: db}
}

// Create 
func (r *WarrantyCardRepo) Create(card *model.WarrantyCard) error {
    return r.db.Create(card).Error
}

func (r *WarrantyCardRepo) GetList(params map[string]interface{}) ([]model.WarrantyCard, error) {
	var cards []model.WarrantyCard
	
	query := r.db.Model(&model.WarrantyCard{})
	
	if serialNumber, ok := params["serial_number"]; ok {
		query = query.Where("serial_number = ?", serialNumber)
	}
	
	if err := query.Find(&cards).Error; err != nil {
		return nil, err
	}
	
	return cards, nil
}

func (r *WarrantyCardRepo) GetByID(id uint64) (*model.WarrantyCard, error) {
	var card model.WarrantyCard
	if err := r.db.First(&card, id).Error; err != nil {
		return nil, err
	}
	return &card, nil
}

type InstitutionRepo struct {
	db *gorm.DB
}

func NewInstitutionRepo(db *gorm.DB) *InstitutionRepo {
	return &InstitutionRepo{db: db}
}

// 根据token获取机构认证信息
func (r *InstitutionRepo) GetAuthByToken(token string) (*model.InstitutionAuth, error) {
	var auth model.InstitutionAuth
	if err := r.db.Where("token = ?", token).First(&auth).Error; err != nil {
		return nil, err
	}
	return &auth, nil
}

// 根据ID获取机构信息
func (r *InstitutionRepo) GetInstitutionByID(id uint64) (*model.Institution, error) {
	var institution model.Institution
	if err := r.db.First(&institution, id).Error; err != nil {
		return nil, err
	}
	return &institution, nil
}