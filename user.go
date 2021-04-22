package futureauth

import (
	"github.com/go-playground/validator/v10"
	"github.com/pkg/errors"
)

func CreateUser(signupData SignupData) (error, *User) {
	validate := validator.New()
	if err := validate.Struct(&signupData); err != nil {
		validationErrors := err.(validator.ValidationErrors)
		if err != nil {
			return errors.Wrap(err, validationErrors.Error()), &User{}
		}
	}

	if newUser, err := NewUser(signupData.Email, signupData.Tenant, signupData.Password, signupData.Confirmed, signupData.IsAdmin, signupData.Disabled, signupData.AppMetaData, signupData.UserMetaData, signupData.Roles); err != nil {
		return errors.Wrap(err, "User creation failed"), &User{}
	} else {
		return nil, newUser
	}
}

func NewUser(email, tenant, password string, confirmed, isAdmin, disabled bool, appMetaData, userMetaData string, roles []string) (*User, error) {
	newUser := User{
		Email:     email,
		Password:  password,
		Confirmed: confirmed,
		IsAdmin:   isAdmin,
		Disabled:  disabled,
		Roles:     roles,
	}

	return &newUser, nil
}
