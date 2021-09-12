package users

import (
	"context"
	"errors"
	"log"

	"github.com/neurocome/ine_go/graph/model"
	"github.com/neurocome/ine_go/internal/pkg/db/pgsql"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID       int
	Username string
	Email    string
	Password string
}

func (user *User) Create() error {
	hashedPassword, err := HashedPassowrd(user.Password)
	if err != nil {
		log.Fatal(err)
	} else {
		user.Password = hashedPassword
	}
	_, err = pgsql.Db.Exec(context.Background(), "INSERT INTO users (username, password, email) values($1, $2, $3)", user.Username, user.Password, user.Email)
	if err != nil {
		return err
	}
	return nil
}

func GetUserByUsername(username string) (int, error) {
	var ID int
	err := pgsql.Db.QueryRow(context.Background(), "SELECT id FROM users WHERE username=$1", username).Scan(&ID)
	if err != nil {
		log.Println(err)
		return 0, err
	}

	return ID, nil
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
	// row := new(User)
	var password string

	err := pgsql.Db.QueryRow(context.Background(), "SELECT password FROM users WHERE username=$1 OR email=$1", user.Username).Scan(&password)
	log.Println(password)
	if err != nil {
		log.Println("dsfdsf", err)
	}
	return CheckPasswordHash(user.Password, password)
}

func GetProfile(id int) (model.User, error) {
	log.Println(id)
	var user model.User
	err := pgsql.Db.QueryRow(context.Background(), "SELECT id,username,email FROM users where id = $1", id).Scan(&user.ID, &user.Username, &user.Email)
	if err != nil {
		return user, errors.New("User not found, please try to login again")
	}

	return user, nil
}
