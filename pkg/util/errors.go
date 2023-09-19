package utils

import (
	"errors"
	"gorm.io/gorm"
	"log"
	"strings"
)

func HandleAlreadyRegistered(err string) error {
	if strings.Contains(err, "username") {
		return errors.New("username already taken")
	}

	if strings.Contains(err, "email") {
		return errors.New("email Already Taken")
	}

	if strings.Contains(err, "title") {
		return errors.New("title Already Taken")
	}
	if strings.Contains(err, "hashedPassword") {
		return errors.New("incorrect Password")
	}
	return errors.New("incorrect Details")
}

func (response *Response) HandlePanic(err interface{}, logger *log.Logger) {
	if HandlePanicMacro(err, logger) {
		if response.Status == 0 {
			response.Status = 500
		}
		response.Success = false
		if response.Message == "" {
			response.Message = "Error processing your request at the current time. Kindly contact support"
		}
	}
}

func (response *Response) HandleDbErrorMacro(result *gorm.DB, logger *log.Logger, message string) {
	response.Message, response.Status = HandleDbError(result, logger)

	if response.Status != 0 {
		response.Success, response.Data = false, nil
	} else {
		response.Message, response.Success = message, true
	}
}
