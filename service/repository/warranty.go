package repository

import (
	
	"h5/model"
	
	"gorm.io/gorm"
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
// Update 更新质保卡信息
func (r *WarrantyCardRepo) Update(card *model.WarrantyCard) error {
    return r.db.Model(&model.WarrantyCard{}).
        Where("card_number = ?", card.CardNumber).
        Updates(card).Error
}

// GetByCardNumber 根据卡号获取质保卡
func (r *WarrantyCardRepo) GetByCardNumber(cardNumber string) (*model.WarrantyCard, error) {
    var card model.WarrantyCard
    if err := r.db.Where("card_number = ?", cardNumber).First(&card).Error; err != nil {
        return nil, err
    }
    return &card, nil
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

func (r *InstitutionRepo) UpdateById(pwd string, id uint) error {
    if err := r.db.Model(&model.Institution{}).Where("id = ?", id).Update("password", pwd).Error; err != nil {
        return err
    }
    return nil
}

// 根据ID获取机构信息
func (r *InstitutionRepo) GetInstitutionByID(id uint64) (*model.Institution, error) {
	var institution model.Institution
	if err := r.db.First(&institution, id).Error; err != nil {
		return nil, err
	}
	return &institution, nil
}

// 根据账户名获取机构信息
func (r *InstitutionRepo) GetInstitutionByAccount(accountName string) (*model.Institution, error) {
    var institution model.Institution
    if err := r.db.Where("account_name = ?", accountName).First(&institution).Error; err != nil {
        return nil, err
    }
    return &institution, nil
}

// 创建认证信息
func (r *InstitutionRepo) CreateAuth(auth *model.InstitutionAuth) error {
    return r.db.Create(auth).Error
}

// 删除旧的认证信息
func (r *InstitutionRepo) DeleteAuthByInstitutionID(institutionID uint64) error {
    return r.db.Where("institution_id = ?", institutionID).Delete(&model.InstitutionAuth{}).Error
}