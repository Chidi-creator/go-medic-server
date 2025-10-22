package utils

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

func init() {
	fmt.Println("Initializing validator...")
	validate = validator.New()
	validate.RegisterValidation("roles", IsValidRole)
	validate.RegisterValidation("specialties", isValidSpecialty)
	validate.RegisterValidation("e164", isValidE164)
	validate.RegisterValidation("geopoint", isValidGeoPointType)
}

// ValidateStruct validates any struct using go-playground/validator.
// It returns a map of field errors if validation fails.

func ValidateStruct(s interface{}) string {
	err := validate.Struct(s)

	if err == nil {
		return ""
	}

	// Check if it's a validation error
	if _, ok := err.(validator.ValidationErrors); !ok {
		return err.Error()
	}

	var messages []string

	for _, e := range err.(validator.ValidationErrors) {
		field := e.Field()
		tag := e.Tag()
		var message string

		// Customize certain messages for readability
		switch tag {
		case "required":
			message = fmt.Sprintf("%s is required", field)
		case "email":
			message = fmt.Sprintf("%s must be a valid email address", field)
		case "min":
			message = fmt.Sprintf("%s must meet the minimum length", field)
		case "max":
			message = fmt.Sprintf("%s exceeds the maximum length", field)
		default:
			message = fmt.Sprintf("%s failed on the '%s' validation", field, tag)
		}

		messages = append(messages, message)
	}
	return strings.Join(messages, ", ")
}

func IsValidRole(fl validator.FieldLevel) bool {
	role := fl.Field().String()
	for _, r := range ValidRoles {
		if string(r) == role {
			return true
		}
	}
	return false
}

func isValidSpecialty(fl validator.FieldLevel) bool {
	role := fl.Field().String()
	for _, r := range ValidSpecialties {
		if string(r) == role {
			return true
		}
	}
	return false
}

func isValidE164(fl validator.FieldLevel) bool {
	phone := fl.Field().String()
	// E.164 format: + followed by up to 15 digits
	e164Regex := regexp.MustCompile(`^\+[1-9]\d{1,14}$`)
	return e164Regex.MatchString(phone)
}

func isValidGeoPointType(fl validator.FieldLevel) bool {
	geoType := fl.Field().String()
	// For GeoJSON, type should always be "Point"
	return geoType == "Point"
}
