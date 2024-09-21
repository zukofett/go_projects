package models

import (
    "database/sql"
    "errors"
    "strings"
    "time"

    "github.com/go-sql-driver/mysql"
    "golang.org/x/crypto/bcrypt"
)
type User struct {
    ID             int
    Name           string
    Email          string
    HashedPassword []byte
    Created        time.Time
}

type UserModel struct {
    DB *sql.DB
}

type UserModelInterface interface {
    Insert(name, email, password string) error
    Authenticate(email, password string) (int, error)
    Exists(id int) (bool, error)
    Get(id int) (User, error)
    PasswordUpdate(id int, currentPassword, newPassword string) error
}

func (m *UserModel) Insert(name, email, password string) error {
    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 12)
    if err != nil {
        return err
    }

    statement := `INSERT INTO users (name, email, hashed_password, created)
    VALUES(?, ?, ?, UTC_TIMESTAMP())`

     _, err = m.DB.Exec(statement, name, email, string(hashedPassword))
    if err != nil {
        var mySQLError *mysql.MySQLError
        if errors.As(err, &mySQLError) {
            if mySQLError.Number == 1062 && strings.Contains(mySQLError.Message, "users_uc_email") {
                return ErrDuplicateEmail
            }
        }
        return err
    }

    return nil
}

func (model *UserModel) Authenticate(email, password string) (int, error) {
    var id             int
    var hashedPassword []byte

    statement := `SELECT id, hashed_password FROM users WHERE email = ?`

    err := model.DB.QueryRow(statement, email).Scan(&id, &hashedPassword)
    if err != nil {
        if errors.Is(err, sql.ErrNoRows) {
            return 0, ErrNoRecord
        }
        return 0, err
    }

    err = bcrypt.CompareHashAndPassword(hashedPassword, []byte(password))
    if err != nil {
        if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
            return 0, ErrInvalidCredentials
        }
        return 0, err
    }

    return id, nil
}

func (m *UserModel) Exists(id int) (bool, error) {
    var exists bool

    statement := "SELECT EXISTS(SELECT true FROM users WHERE id = ?)"

    err := m.DB.QueryRow(statement, id).Scan(&exists)
    return exists, err
}

func (model *UserModel) Get(id int) (User, error) {
    var user User

    statement := `SELECT id, name, email, created FROM users WHERE id = ?`
    
    err := model.DB.QueryRow(statement, id).Scan(&user.ID, &user.Name, &user.Email, &user.Created)
    if err != nil {
        if errors.Is(err, sql.ErrNoRows) {
            return User{}, ErrNoRecord
        }
        return User{}, err
    }

    return user, nil
}

func (model *UserModel) PasswordUpdate(id int, currentPassword, newPassword string) error {
    var hashedPassword []byte

    statement := `SELECT hashed_password FROM users WHERE id = ?`

    err := model.DB.QueryRow(statement, id).Scan(&hashedPassword)
    if err != nil {
        return err
    }

    err = bcrypt.CompareHashAndPassword(hashedPassword, []byte(currentPassword))
    if err != nil {
        if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
            return ErrInvalidCredentials
        }
        return err
    }

    hashedPassword, err = bcrypt.GenerateFromPassword([]byte(newPassword), 12)
    if err != nil {
        return err
    }

    statement = `UPDATE users SET hashed_password = ? WHERE id = ?`

    _ ,err = model.DB.Exec(statement, string(hashedPassword), id)
    return err
}
