package model

import (
	"time"

	"gorm.io/gorm"
)

type WarrantyCard struct {
	ID              uint64         `gorm:"primaryKey;autoIncrement;column:id" json:"id"`
	CardNumber      string         `gorm:"uniqueIndex;size:100;not null;column:card_number" json:"cardNumber"`
	EntryState      uint8          `gorm:"type:tinyint unsigned;not null;default:1;column:entry_state" json:"entryState"`           // 1=未录入,2=已录入
	ActivationState uint8          `gorm:"type:tinyint unsigned;not null;default:1;column:activation_state" json:"activationState"` // 1=未激活,2=已激活
	SerialNumber    string         `gorm:"uniqueIndex;size:200;column:serial_number" json:"serialNumber"`
	PatientName     string         `gorm:"size:200;column:patient_name" json:"patientName"`
	PatientAge      int16          `gorm:"column:patient_age" json:"patientAge"`
	InstitutionID   uint64         `gorm:"index;column:institution_id" json:"institutionId"`
	SurgeryDate     *time.Time     `gorm:"type:date;column:surgery_date" json:"surgeryDate"`
	SurgeonName     string         `gorm:"size:200;column:surgeon_name" json:"surgeonName"`
	Password        string         `gorm:"size:255;column:password" json:"-"`
	BindMobile      string         `gorm:"index;size:20;column:bind_mobile" json:"bindMobile"`
	EntryTime       *time.Time     `gorm:"column:entry_time" json:"entryTime"`
	ActivationTime  *time.Time     `gorm:"column:activation_time" json:"activationTime"`
	Remark          string         `gorm:"type:text;column:remark" json:"remark"`
	Credential      []byte         `gorm:"type:json;column:credential" json:"credential"`
	CreatedTime     time.Time      `gorm:"autoCreateTime;column:created_time" json:"createdTime"`
	UpdatedTime     time.Time      `gorm:"autoUpdateTime;column:updated_time" json:"updatedTime"`
	DeletedTime     gorm.DeletedAt `gorm:"index;column:deleted_time" json:"-"`
}

// TableName 设置表名
func (WarrantyCard) TableName() string {
	return "opd_warranty_card"
}


type Institution struct {
	ID             uint64         `gorm:"primaryKey;autoIncrement;column:id" json:"id"`
	InstitutionName string        `gorm:"size:200;not null;column:institution_name" json:"institutionName"`
	AccountName    string         `gorm:"size:200;not null;uniqueIndex;column:account_name" json:"accountName"`
	Mobile         string         `gorm:"size:20;not null;column:mobile" json:"mobile"`
	ContactPerson  string         `gorm:"size:200;not null;column:contact_person" json:"contactPerson"`
	State          uint8          `gorm:"type:tinyint unsigned;not null;default:1;column:state" json:"state"`
	Password       string         `gorm:"size:255;not null;column:password" json:"-"`
	CreatedTime    time.Time      `gorm:"autoCreateTime;column:created_time" json:"createdTime"`
	UpdatedTime    time.Time      `gorm:"autoUpdateTime;column:updated_time" json:"updatedTime"`
	DeletedTime    gorm.DeletedAt `gorm:"index;column:deleted_time" json:"-"`
}

func (Institution) TableName() string {
	return "opd_institution"
}

type InstitutionAuth struct {
	ID            uint64    `gorm:"primaryKey;autoIncrement;column:id" json:"id"`
	InstitutionID uint64    `gorm:"not null;index;column:institution_id" json:"institutionId"`
	Token         string    `gorm:"size:255;not null;uniqueIndex;column:token" json:"token"`
	IP            string    `gorm:"size:255;not null;column:ip" json:"ip"`
	UserAgent     string    `gorm:"size:255;not null;column:user_agent" json:"userAgent"`
	ExpiresTime   time.Time `gorm:"column:expires_time" json:"expiresTime"`
	CreatedTime   time.Time `gorm:"autoCreateTime;column:created_time" json:"createdTime"`
	UpdatedTime   time.Time `gorm:"autoUpdateTime;column:updated_time" json:"updatedTime"`
}

func (InstitutionAuth) TableName() string {
	return "opd_institution_auth"
}
