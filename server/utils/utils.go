package utils

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"gopkg.in/go-playground/validator.v9"
	"reflect"
	"strings"
)


func ValidateInputs(dataSet interface{}) error {

	var validate *validator.Validate

	validate = validator.New()

	err := validate.Struct(dataSet)

	if err != nil {
		//Validation syntax is invalid
		if _, ok := err.(*validator.InvalidValidationError); ok {
			return errors.New("validation syntax is invalid")
		}

		reflected := reflect.ValueOf(dataSet)

		for _, err := range err.(validator.ValidationErrors) {

			// Attempt to find field by name and get json tag name
			field, _ := reflected.Type().FieldByName(err.StructField())

			//If json tag doesn't exist, use lower case of name
			name := field.Tag.Get("json")
			if name == "" {
				name = strings.ToLower(err.StructField())
			}

			switch err.Tag() {
			case "required":
				return errors.New(fmt.Sprintf("The %s is required", name))
			case "email":
				return errors.New(fmt.Sprintf("The %s should be a valid email", name))
			default:
				return errors.New(fmt.Sprintf("The %s is invalid", name))
			}
		}
	}

	return nil
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, PATCH, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	}
}


