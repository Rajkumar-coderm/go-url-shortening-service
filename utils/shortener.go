package utils

import (
	"errors"
	"fmt"
	"math/rand"
	"strings"

	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/mongo"
)

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func GenerateShortCode(length int) string {
	code := make([]byte, length)
	for i := range code {
		code[i] = charset[rand.Intn(len(charset))]
	}
	return string(code)
}

type ApiError struct {
	Field string `json:"field"`
	Msg   string `json:"message"`
}

func FormatValidationError(err error) []ApiError {
	var ve validator.ValidationErrors
	var out []ApiError

	if errors.As(err, &ve) {
		for _, fe := range ve {
			msg := getCustomMsg(fe)
			out = append(out, ApiError{
				Field: fe.Field(),
				Msg:   fmt.Sprintf("The field '%s' %s", fe.Field(), msg),
			})
		}
		return out
	}

	return []ApiError{{Field: "request", Msg: "Invalid JSON format"}}
}

func FormatDBError(err error) (int, map[string]string) {
	if mongo.IsDuplicateKeyError(err) {
		return 409, map[string]string{
			"error": "Conflict",
			"msg":   "This resource (or short code) already exists. Please try again.",
		}
	}

	msg := err.Error()

	if strings.Contains(msg, "NotFound") || err == mongo.ErrNoDocuments {
		return 404, map[string]string{
			"error": "Not Found",
			"msg":   "The requested resource was not found.",
		}
	}

	return 500, map[string]string{
		"error": "Internal Server Error",
		"msg":   "An unexpected database error occurred.",
	}
}

func getCustomMsg(fe validator.FieldError) string {
	switch fe.Tag() {
	case "required":
		return "is mandatory and cannot be empty"
	case "url":
		return "must be a valid URL"
	case "oneof":
		return "must be one of the following: " + fe.Param()
	default:
		return "is invalid"
	}
}
