package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"log"

	"github.com/neurocome/ine_go/graph/generated"
	"github.com/neurocome/ine_go/graph/model"
	"github.com/neurocome/ine_go/graph/models"
	"github.com/neurocome/ine_go/internal/application/times"
	"github.com/neurocome/ine_go/internal/auth"
	"github.com/neurocome/ine_go/internal/pkg/jwt"
	"github.com/neurocome/ine_go/internal/users"
)

func (r *mutationResolver) CreateUser(ctx context.Context, input model.NewUser) (string, error) {
	var user users.User
	user.Username = input.Username
	user.Password = input.Password
	user.Email = input.Email
	err := user.Create()
	if err != nil {
		return "", err
	}
	token, err := jwt.GenerateToken(user.Username)
	if err != nil {
		return "", err
	}
	return token, nil
}

func (r *mutationResolver) Login(ctx context.Context, input model.Login) (string, error) {
	var user users.User
	user.Username = input.Username
	user.Password = input.Password
	correct := user.Authenticate()
	if !correct {
		return "", &users.WrongUsernameOrPasswordError{}
	}

	token, err := jwt.GenerateToken(user.Username)
	if err != nil {
		return "", err
	}
	// cookie := &http.Cookie{Name: "belivine", Value: token, Expires: time.Now().Add(24 * time.Hour)}
	// headers := context.Context.Value(ctx, "belivine").(http.Header)
	// headers.Add("Set-Cookie", cookie.String())
	return token, nil
}

func (r *mutationResolver) RefreshToken(ctx context.Context, input model.RefreshTokenInput) (string, error) {
	username, err := jwt.ParseToken(input.Token)
	if err != nil {
		return "", fmt.Errorf("access denied")
	}
	token, err := jwt.GenerateToken(username)
	if err != nil {
		return "", err
	}
	return token, nil
}

func (r *mutationResolver) CreateTime(ctx context.Context, input model.NewTime) (*models.Time, error) {
	var time times.InputTime
	time.Task_id = *input.TaskID
	time.User_id = input.UserID
	data, err := time.Create()
	if err != nil {
		return nil, err
	}
	return &data, nil
}

func (r *mutationResolver) UpdateTime(ctx context.Context, input model.UpdateTime) (*models.Time, error) {
	var time times.UpdateTime
	time.Time_id = input.TimeID
	data, err := time.Update()
	if err != nil {
		return nil, err
	}
	return &data, nil
}

func (r *mutationResolver) CreateTask(ctx context.Context, input *model.Task) (*models.Task, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Time(ctx context.Context, id int64) ([]*models.Time, error) {
	data, err := times.GetAll(id)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (r *queryResolver) GetProfile(ctx context.Context) (*model.User, error) {
	user := auth.ForContext(ctx)
	log.Println(user)
	if user == nil {
		return nil, fmt.Errorf("access denied")
	}
	log.Println("sdfdsfdsf")
	data, err := users.GetProfile(user.ID)
	if err != nil {
		return nil, err
	}
	return &data, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
