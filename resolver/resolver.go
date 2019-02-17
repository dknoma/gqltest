package resolver

import (
	"context"
	"errors"
	"fmt"
	. "github.com/dknoma/gqltest/generated"
	. "github.com/dknoma/gqltest/models"
	"math/rand"
	"time"
)

//go:generate go run ../scripts/gqlgen.go -v

type Resolver struct {
	todos []Todo
}

func (r *Resolver) Mutation() MutationResolver {
	return &mutationResolver{r}
}
func (r *Resolver) Query() QueryResolver {
	return &queryResolver{r}
}
func (r *Resolver) Todo() TodoResolver {
	return &todoResolver{r}
}

type mutationResolver struct{ *Resolver }

func (r *mutationResolver) CreateTodo(ctx context.Context, input NewTodo) (Todo, error) {
	todo := Todo{
		Text:   input.Text,
		ID:     rand.Intn(10000), // Only for testing dummy ids. NOT FOR PRODUCTION USE
		UserID: input.UserID,     // Only for testing dummy ids. NOT FOR PRODUCTION USE
	}
	r.todos = append(r.todos, todo)
	return todo, nil
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) Todos(ctx context.Context) ([]Todo, error) {
	return r.todos, nil
}

func (r *queryResolver) Todo(ctx context.Context, id int) (*Todo, error) {
	time.Sleep(220 * time.Millisecond)

	if id == 666 {
		panic("critical failure")
	}

	for _, todo := range r.todos {
		fmt.Printf("id: %v\n", todo.ID) // Testing for dummy values; NOT FOR PRODUCTION USE
		if todo.ID == id {
			return &todo, nil
		}
	}
	return nil, errors.New("not found")
}

type todoResolver struct{ *Resolver }

func (r *todoResolver) User(ctx context.Context, obj *Todo) (User, error) {
	return User{ID: obj.UserID, Name: fmt.Sprintf("user %d", obj.UserID)}, nil
}
