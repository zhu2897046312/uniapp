package service

import (
	"h5/model"
	"h5/repository"
	"errors"
    "time"
    "github.com/golang-jwt/jwt/v5"
    "crypto/md5"
    "gorm.io/gorm"
	"encoding/hex"
)

type WarrantyCardService struct {
	repo *repository.WarrantyCardRepo
}

func NewWarrantyCardService(repo *repository.WarrantyCardRepo) *WarrantyCardService {
	return &WarrantyCardService{repo: repo}
}

func (s *WarrantyCardService) CreateCard(card *model.WarrantyCard) error {
	// 检查是否已存在相同卡号的记录
	_, err := s.repo.GetByCardNumber(card.CardNumber)
    if err != nil {
        if errors.Is(err, gorm.ErrRecordNotFound) {
            return errors.New("卡号不存在")
        }
        return err
    }

    // 更新现有记录
    updateData := &model.WarrantyCard{
        CardNumber:  card.CardNumber, // 保持卡号不变
        EntryState:  2,               // 设置为已录入状态
        PatientName: card.PatientName,
        PatientAge:  card.PatientAge,
        SurgeonName: card.SurgeonName,
        SurgeryDate: card.SurgeryDate,
        Remark:     card.Remark,
        Credential: card.Credential,
    }
    return s.repo.Update(updateData)

}

func (s *WarrantyCardService) GetCardList(params map[string]interface{})([]model.WarrantyCard, error) {
	return s.repo.GetList(params)
}

func (s *WarrantyCardService) GetCardByID(id uint64) (*model.WarrantyCard, error) {
	return s.repo.GetByID(id)
}
func (s *WarrantyCardService) GetCardByNumber(cardNumber string) (*model.WarrantyCard, error) {
    if cardNumber == "" {
        return nil, errors.New("卡号不能为空")
    }
    
    card, err := s.repo.GetByCardNumber(cardNumber)
    if err != nil {
        if errors.Is(err, gorm.ErrRecordNotFound) {
            return nil, errors.New("质保卡不存在")
        }
        return nil, err
    }
    
    return card, nil
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

// 定义JWT claims结构
type jwtClaims struct {
    InstitutionID uint64 `json:"institutionId"`
    jwt.RegisteredClaims
}

var jwtSecret = []byte("your-secret-key") // 应该从配置中读取

// 密码比较
// 密码比较 (修改为MD5验证)
func comparePasswords(hashedPwd string, plainPwd string) bool {
	// 对明文密码进行MD5加密
	hasher := md5.New()
	hasher.Write([]byte(plainPwd))
	md5Pwd := hex.EncodeToString(hasher.Sum(nil))
	
	// 比较MD5加密后的密码与数据库中存储的密码
	return hashedPwd == md5Pwd
}
// 生成JWT token
func generateToken(institutionID uint64) (string, error) {
    claims := jwtClaims{
        InstitutionID: institutionID,
        RegisteredClaims: jwt.RegisteredClaims{
            ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)), // 24小时过期
            IssuedAt:  jwt.NewNumericDate(time.Now()),
        },
    }
    
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    return token.SignedString(jwtSecret)
}

// 登录
func (s *InstitutionService) Login(AccountName string, Password string) (string, error) {
    // 1. 根据账户名获取机构
    institution, err := s.repo.GetInstitutionByAccount(AccountName)
    if err != nil {
        return "", errors.New("账户不存在")
    }
    
    // 2. 检查密码
    if !comparePasswords(institution.Password, Password) {
        return "", errors.New("密码错误")
    }
    
    // 3. 检查机构状态
    if institution.State != 1 {
        return "", errors.New("机构已停用")
    }
    
    // 4. 生成token
    token, err := generateToken(institution.ID)
    if err != nil {
        return "", errors.New("生成token失败")
    }
    
    // 5. 保存认证信息(先删除旧的)
    err = s.repo.DeleteAuthByInstitutionID(institution.ID)
    if err != nil {
        return "", errors.New("清理旧token失败")
    }
    
    auth := &model.InstitutionAuth{
        InstitutionID: institution.ID,
        Token:         token,
        IP:            "", // 可以从上下文中获取
        UserAgent:     "", // 可以从上下文中获取
        ExpiresTime:   time.Now().Add(24 * time.Hour),
    }
    
    if err := s.repo.CreateAuth(auth); err != nil {
        return "", errors.New("保存认证信息失败")
    }
    
    // 6. 返回响应
    return token,nil
}

func (s *InstitutionService) UpdatePWD(newPassword string, token string) error {
    // 1. 根据token获取认证信息
    auth, err := s.repo.GetAuthByToken(token)
    if err != nil {
        return errors.New("无效的token")
    }

    // 2. 检查token是否过期
    if time.Now().After(auth.ExpiresTime) {
        return errors.New("token已过期")
    }

    // 3. 对密码进行MD5加密
    hasher := md5.New()
    hasher.Write([]byte(newPassword))
    md5Pwd := hex.EncodeToString(hasher.Sum(nil))

    // 4. 更新密码
    if err := s.repo.UpdateById(md5Pwd, uint(auth.InstitutionID)); err != nil {
        return errors.New("密码更新失败")
    }

    return nil
}