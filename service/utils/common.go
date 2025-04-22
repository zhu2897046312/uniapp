package utils
import (
	// "net/http"
)
// ApiResponse 基础响应结构体
type ApiResponse[T any] struct {
    Code    int    `json:"code"`    // 状态码
    Message string `json:"message"` // 消息
    Result T      `json:"result"`  // 数据结果
}

// 成功响应构造函数
func SuccessResponse[T any](data T) ApiResponse[T] {
    return ApiResponse[T]{
        Code:    0,
        Message: "success",
        Result:  data,
    }
}

// 错误响应构造函数
func ErrorResponse(code int, message string) ApiResponse[any] {
    return ApiResponse[any]{
        Code:    code,
        Message: message,
        Result:  nil,
    }
}

// 带数据的错误响应构造函数
func ErrorResponseWithData[T any](code int, message string, data T) ApiResponse[T] {
    return ApiResponse[T]{
        Code:    code,
        Message: message,
        Result:  data,
    }
}