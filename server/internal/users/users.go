package users

import (
	"log"

	"github.com/neurocome/ine_go/internal/pkg/db/pgsql"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	tableName struct{} `pg:"belivine.users"`
	ID        int      `json:"id" pg:"id"`
	Username  string   `json:"username" pg:"username"`
	Password  string   `json:"password" pg:"password"`
}

func (user *User) Create() {
	hashedPassword, err := HashedPassowrd(user.Password)
	if err != nil {
		log.Fatal(err)
	} else {
		user.Password = hashedPassword
	}
	_, err = pgsql.Db.Model(user).Insert()
	if err != nil {
		log.Fatal(err)
	}
}

func GetUserByUsername(username string) (int, error) {
	user := new(User)
	err := pgsql.Db.Model(user).Where("username = ?", username).Select()
	if err != nil {
		log.Println(err)
		return 0, err
	}

	return user.ID, nil
}

func HashedPassowrd(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func (user *User) Authenticate() bool {
	row := new(User)
	err := pgsql.Db.Model(row).Where("username = ?", user.Username).Select()
	if err != nil {
		log.Println(err)
	}
	return CheckPasswordHash(user.Password, row.Password)
}
