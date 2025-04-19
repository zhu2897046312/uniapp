package handlers

import (
	"net/http"

	"h5/model"
	"h5/service"
	"encoding/json"
	"strconv"
	"time"
	"gorm.io/gorm"
	"github.com/gin-gonic/gin"
)

var warrantyService *service.WarrantyCardService

func InitHandlers(service *service.WarrantyCardService) {
	warrantyService = service
}

func CreateWarrantyCard(c *gin.Context) {
    var request struct {
        ID              uint64         `json:"id"`
        CardNumber      string         `json:"card_number"`
        EntryState      uint8          `json:"entry_state"`
        ActivationState uint8          `json:"activation_state"`
        SerialNumber    string         `json:"serial_number"`
        PatientName     string         `json:"patient_name"`
        PatientAge      int16          `json:"patient_age"`
        InstitutionID   uint64         `json:"institution_id"`
        SurgeryDate     string     `json:"surgery_date"`
        SurgeonName     string         `json:"surgeon_name"`
        Password        string         `json:"-"`
        BindMobile      string         `json:"bind_mobile"`
        EntryTime       *time.Time     `json:"entry_time"`
        ActivationTime  *time.Time     `json:"activation_time"`
        Remark          string         `json:"remark"`
        Credential      []string       `json:"credential"`
        CreatedTime     time.Time      `json:"created_time"`
        UpdatedTime     time.Time      `json:"updated_time"`
        DeletedTime     gorm.DeletedAt `json:"-"`
    }

    if err := c.ShouldBindJSON(&request); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
	// 处理 surgery_date 字符串转 time.Time
    var surgeryDate *time.Time
    if request.SurgeryDate != "" {
        parsedDate, err := time.Parse("2006-01-02", request.SurgeryDate)
        if err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": "invalid surgery_date format, expected YYYY-MM-DD"})
            return
        }
        surgeryDate = &parsedDate
    }

    // 将 Credential 从 []string 转换为 []byte
    credentialBytes, err := json.Marshal(request.Credential)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to marshal credential data"})
        return
    }

    card := model.WarrantyCard{
        ID:              request.ID,
        CardNumber:      request.CardNumber,
        EntryState:      request.EntryState,
        ActivationState: request.ActivationState,
        SerialNumber:    request.SerialNumber,
        PatientName:     request.PatientName,
        PatientAge:      request.PatientAge,
        InstitutionID:   request.InstitutionID,
        SurgeryDate:     surgeryDate,
        SurgeonName:     request.SurgeonName,
        Password:        request.Password,
        BindMobile:      request.BindMobile,
        EntryTime:       request.EntryTime,
        ActivationTime:  request.ActivationTime,
        Remark:          request.Remark,
        Credential:      credentialBytes,
        CreatedTime:     request.CreatedTime,
        UpdatedTime:     request.UpdatedTime,
        DeletedTime:     request.DeletedTime,
    }

    if err := warrantyService.CreateCard(&card); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusCreated, gin.H{"data": card})
}

func GetWarrantyCardList(c *gin.Context) {
	params := make(map[string]interface{})
	// Get filter conditions from query parameters
	if serialNumber := c.Query("serialNumber"); serialNumber != "" {
		params["serial_number"] = serialNumber
	}

	// Call the service to get the list of warranty cards
	cards, err := warrantyService.GetCardList(params)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": cards})
}

func GetWarrantyCardByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid ID format"})
		return
	}

	card, err := warrantyService.GetCardByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "card not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": card})
}

var institutionService *service.InstitutionService

func InitInstitutionHandlers(service *service.InstitutionService) {
	institutionService = service
}

func GetInstitutionByToken(c *gin.Context) {
	token := c.GetHeader("Authorization")
	if token == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization token is required"})
		return
	}

	institution, err := institutionService.GetInstitutionByToken(token)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": institution})
}