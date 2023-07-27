package responses

import "github.com/gin-gonic/gin"

type BaseResponse struct {
	Code      int
	message   string
	isSucceed bool
	data      gin.H
}

func (receiver *BaseResponse) CreateErrorResponse(err error) *BaseResponse {
	receiver.Code = 400
	receiver.message = err.Error()
	return receiver
}

func (receiver *BaseResponse) CreateErrorMessage(message string) *BaseResponse {
	receiver.Code = 400
	receiver.message = message
	return receiver
}

func (receiver *BaseResponse) CreateSuccessResponse(message string) *BaseResponse {
	receiver.Code = 200
	receiver.message = message
	return receiver
}

func (receiver *BaseResponse) SetData(key string, value any) {
	receiver.data[key] = value
}

func (receiver *BaseResponse) GetJson() (response *gin.H) {
	response = &gin.H{"message": receiver.message, "succeed": receiver.isSucceed, "data": receiver.data}
	return
}
