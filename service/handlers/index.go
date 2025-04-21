package handlers

import (
	"net/http"

	"h5/model"
	"h5/service"
    "h5/config"
	"encoding/json"
	"strconv"
	"time"
	"gorm.io/gorm"
	"github.com/gin-gonic/gin"
    "os"
    "path/filepath"
    "strings"
    "fmt"
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

func UploadImage(c *gin.Context) {
    // 单文件上传
    file, err := c.FormFile("file")
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "获取上传文件失败: " + err.Error()})
        return
    }

    // 创建上传目录(如果不存在)
    uploadDir := "uploads"
    if err := os.MkdirAll(uploadDir, os.ModePerm); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "创建上传目录失败"})
        return
    }

    // 生成唯一文件名: 时间戳+原始文件名
    filename := fmt.Sprintf("%d-%s", time.Now().Unix(), file.Filename)
    
    // 本地存储路径使用系统分隔符(Windows是\，Linux是/)
    localPath := filepath.Join(uploadDir, filename)
    
    // 保存文件
    if err := c.SaveUploadedFile(file, localPath); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "保存文件失败: " + err.Error()})
        return
    }

    // 使用配置文件中的IP和端口构建URL
    serverIP := config.GlobalConfig.Server.IP
    serverPort := config.GlobalConfig.Server.Port
    
    // 统一转换为URL格式的路径(使用/分隔符)
    urlPath := strings.ReplaceAll(localPath, "\\", "/")
    
    // 返回完整URL和本地路径
    imageURL := fmt.Sprintf("http://%s:%d/%s", serverIP, serverPort, urlPath)
    c.JSON(http.StatusOK, gin.H{
        "data": gin.H{
            "url":  imageURL,   // 使用正斜杠的URL
            "path": localPath,  // 保留原始本地路径(可能包含反斜杠)
        },
    })
}