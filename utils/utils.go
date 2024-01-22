package utils

import (
	"github.com/harshvsinghme/uniblox-assmt-backend/models"
	"github.com/twinj/uuid"
)

func GetError(er *models.Error, Message string, Code int) {
	er.Message = Message
	er.Code = Code
	er.TraceId = uuid.NewV4().String()
}
