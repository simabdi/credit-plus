package helper

import (
	"credit-plus/internal/model"
	"encoding/base64"
	"time"
)

func JsonResponse(code int, message string, success bool, error string, data interface{}) model.Response {
	meta := model.Meta{
		Code:    code,
		Status:  success,
		Message: message,
		Error:   error,
	}

	response := model.Response{
		Meta: meta,
		Data: data,
	}

	return response
}

func GetFormattedDate(date time.Time, format string) string {
	return date.Format(format)
}

func GetDate(format string) string {
	date := time.Now()
	return date.Format(format)
}

func ParseDate(s string, format string) time.Time {
	date, _ := time.Parse(format, s)
	return date
}

func Std64Encode(plainText string) string {
	return base64.StdEncoding.EncodeToString([]byte(plainText))
}

func Std64Decode(encoded string) string {
	decodedByte, _ := base64.StdEncoding.DecodeString(encoded)
	return string(decodedByte)
}
