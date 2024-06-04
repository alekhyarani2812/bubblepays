package learning

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPasswordVaidation_validPwd(t *testing.T) {
	t.Run("TestPasswordVaidation_validPwd", func(t *testing.T) {
		isValidPwd, err := ValidatePassword("Alekhya@123")
		//expected := true
		assert.Equal(t, true, isValidPwd, "Expcted result is true. But got false.")
		assert.Equal(t, nil, err, fmt.Sprintf("Expcted error should be nil. But got an error : '%v'", err))
	})

}
func TestPasswordVaidation_emptyPwd(t *testing.T) {
	t.Run("TestPasswordVaidation_validPwd", func(t *testing.T) {
		isValidPwd, err := ValidatePassword("")
		assert.Equal(t, false, isValidPwd, "Expcted result is false. But got true.")
		assert.NotNil(t, err, "ValidatePassword should not return nil")
		assert.Equal(t, "password must be atleast 8 characters length", err.Error(),
			fmt.Sprintf("Expcted error is true. But got an error : '%v'", err))
	})
}

func TestPasswordVaidation_lowercasePwd(t *testing.T) {
	t.Run("TestPasswordVaidation_validPwd", func(t *testing.T) {
		isValidPwd, err := ValidatePassword("alekhya123")
		assert.Equal(t, false, isValidPwd, "Expcted result is false. But got true.")
		assert.NotNil(t, err, "ValidatePassword should not return nil")
		assert.Equal(t, "password must have atleast 1 upper case letter", err.Error(),
			fmt.Sprintf("Expcted error is true. But got an error : '%v'", err))
	})
}
func TestPasswordVaidation_SpeacialCharPwd(t *testing.T) {
	t.Run("TestPasswordVaidation_validPwd", func(t *testing.T) {
		isValidPwd, err := ValidatePassword("Alekhya123")
		assert.Equal(t, false, isValidPwd, "Expcted result is false. But got true.")
		assert.NotNil(t, err, "ValidatePassword should not return nil")
		if err != nil {
			assert.Equal(t, "password must have atleast 1 special character", err.Error(),
				fmt.Sprintf("Expcted error is true. But got an error : '%v'", err))
		}
	})
}
func TestPasswordVaidation_uppercasePwd(t *testing.T) {
	t.Run("TestPasswordVaidation_validPwd", func(t *testing.T) {
		isValidPwd, err := ValidatePassword("ALEKHYA123")
		assert.Equal(t, false, isValidPwd, "Expcted result is false. But got true.")
		assert.NotNil(t, err, "ValidatePassword should not return nil")
		if err != nil {
			assert.Equal(t, "password must have atleast 1 lower case letter", err.Error(),
				fmt.Sprintf("Expcted error is true. But got an error :'%v'", err))
		}
	})
}
