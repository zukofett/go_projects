package models

import (
	"testing"

	"snippetbox.zukofett.net/internal/assert"
)

func TestUserModelExists(t *testing.T) {
    if testing.Short() {
        t.Skip("models: skiping integration test")
    }

    tests := []struct {
        name   string
        userID int
        want   bool
    }{
        {
            name:   "Valid ID",
            userID: 1,
            want:   true,
        },{
            name:   "Zero ID",
            userID: 0,
            want:   false,
        },{
            name:   "Non-existant ID",
            userID: 2,
            want:   false,
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            db    := newTestDB(t)
            model := UserModel{db}

            exists, err := model.Exists(tt.userID)

            assert.Equal(t, exists, tt.want)
            assert.NilError(t, err)
        })
    }
}
