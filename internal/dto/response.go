package dto

import (
	"github.com/gin-gonic/gin"
	ctx2 "github.com/hipeday/rosen/internal/ctx"
)

// GeneralResponse 通用响应结构
type GeneralResponse struct {
	RequestID string `json:"request_id"` // 请求 ID
	Timestamp int64  `json:"timestamp"`  // 请求时间戳
}

// ErrorResponse 错误响应结构
type ErrorResponse struct {
	Error     string `json:"error"`
	RequestId string `json:"request_id"`
}

func NewErrorResponse(message string, c *gin.Context) ErrorResponse {
	requestId, _ := ctx2.GetRequestId(c.Request.Context())
	return ErrorResponse{
		Error:     message,
		RequestId: requestId,
	}
}
