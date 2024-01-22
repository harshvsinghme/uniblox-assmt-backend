package utils

import (
	"github.com/gin-gonic/gin"
	"github.com/harshvsinghme/uniblox-assmt-backend/models"
	"github.com/twinj/uuid"
)

func GetError(er *models.Error, Message string, Code int) {
	er.Message = Message
	er.Code = Code
	er.TraceId = uuid.NewV4().String()
}

func GetValueFromContext(c *gin.Context, key string) string {
	value, exists := c.Get(key)
	if !exists {
		return ""
	}

	if strValue, ok := value.(string); ok {
		return strValue
	}

	return ""
}
