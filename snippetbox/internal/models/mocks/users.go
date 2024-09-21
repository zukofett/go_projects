package mocks

import (
	"time"

	"snippetbox.zukofett.net/internal/models"
)

type UserModel struct{}

var mockUser = models.User{
    ID :            1,
    Name:           "Alice",
    Email:          "alice@example.com",
    Created:        time.Now(),
}

func (model *UserModel) Insert(name, email, password string) error {
    switch email {
    case "dupe@example.com":
        return models.ErrDuplicateEmail
    default:
        return nil
    }
}

func (model *UserModel) Authenticate(email, password string) (int, error) {
    if email == "alice@example.com" && password == "pa$$word" {
        return 1, nil
    }
    return 0, models.ErrInvalidCredentials
}

func (model *UserModel) Exists(id int) (bool, error) {
    switch id {
    case 1:
        return true, nil
    default:
        return false, nil
    }
}

func (model *UserModel) Get(id int) (models.User, error) {
    switch id {
    case 1:
        return mockUser, nil
    default:
        return models.User{}, models.ErrNoRecord
    }
}

func (model *UserModel) PasswordUpdate(id int, currentPassword, newPassword string) error {
    switch id {
    case 1:
        if currentPassword == "pa$$word" {
            return nil
        }
        return models.ErrInvalidCredentials
    default:
        return models.ErrNoRecord
    }
}
