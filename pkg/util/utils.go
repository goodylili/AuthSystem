package utils

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
	"log"
	"runtime/debug"
	"strings"
)

type Response struct {
	Status  int         `json:"-"`
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func HandlePanicMacro(err interface{}, logger *log.Logger) bool {
	if err != nil {
		if logger == nil {
			logger = log.Default()
		}
		logger.Println(err)
		logger.Println(string(debug.Stack()))
		return true
	}
	return false
}

func HandleDbError(result *gorm.DB, logger *log.Logger) (string, int) {
	if result.Error != nil {
		logger.Println("Database Error: ", result.Error)

		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			//Handle Records Not Found

			return "Internal Server Error: Record(s) not Found", 404
		} else if strings.Contains(result.Error.Error(), "Error 1062") {
			//Handle Duplicate Key Error for Mariadb

			fieldName := strings.Split(result.Error.Error(), "for key")
			return fmt.Sprintf("Internal Server Error: Field %s already exists", fieldName[1]), 400
		} else {
			//Catch all other errors

			return "Internal Server Error, Kindly Contact Support", 500
		}
	}

	return "", 0
}
