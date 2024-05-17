package learning

import (
	"errors"
	"unicode"

	"github.com/bubblepays/log"
)

var passwordLen, hasUppercase, hasLowercase, hasDigit, hasSpecialChar bool

func ValidatePassword(password string) (bool, error) {

	if len(password) >= 8 {
		log.Infof("len(password) %v", len(password))
		passwordLen = true
	} else {
		return false, errors.New("password must be atleast 8 characters length")
	}
	for _, char := range password {
		switch {
		case unicode.IsUpper(char):
			hasUppercase = true
		case unicode.IsLower(char):
			hasLowercase = true
		case unicode.IsDigit(char):
			hasDigit = true
		case unicode.IsPunct(char) || unicode.IsSymbol(char):
			hasSpecialChar = true
		}
	}
	if !passwordLen {
		return false, errors.New("password must be at least 8 characters long")
	}
	if !hasLowercase {
		return false, errors.New("password must have atleast 1 lower case letter")
	}
	if !hasUppercase {
		return false, errors.New("password must have atleast 1 upper case letter")
	}
	if !hasDigit {
		return false, errors.New("password must have atleast 1 digit")
	}
	if !hasSpecialChar {
		return false, errors.New("password must have atleast 1 special character")
	}
	return true, nil
}
// func ValidatePassword(password string) {
// 	isValid, err := ValidatePasswordLogic(password)
// 	if err != nil {
// 		log.Error(err.Error())
// 	} else {
// 		log.Infof("password is Valid it returns %v ", isValid)
// 	}
// }
