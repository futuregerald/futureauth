package futureauth_test

import (
	"testing"

	"github.com/futuregerald/futureauth"
	"github.com/stretchr/testify/assert"
)

func TestCreateNewUser(t *testing.T) {

	userSignupData := futureauth.SignupData{
		Email:    "test@user.com",
		Password: "123sdf234dsf",
	}

	newUser, err := futureauth.CreateUser(userSignupData)
	assert.NoError(t, err)
	assert.Equal(t, userSignupData.Email, newUser.Email)
	assert.NotEqual(t, userSignupData.Password, newUser.Password)
}

func TestVerifyEncryptedPassword(t *testing.T) {

	userSignupData := futureauth.SignupData{
		Email:    "test@user.com",
		Password: "123sdf234dsf",
	}

	newUser, err := futureauth.CreateUser(userSignupData)
	assert.NoError(t, err)
	isValidPassword, err := newUser.VerifyPassword("123sdf234dsf")
	assert.NoError(t, err)
	assert.True(t, isValidPassword)
	assert.NotEqual(t, userSignupData.Password, newUser.Password)
}
