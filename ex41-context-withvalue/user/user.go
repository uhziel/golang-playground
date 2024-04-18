package user

import (
	"context"
)

type User struct {
	ID   string
	Name string
}

type key int

const userKey key = 0

func NewContext(ctx context.Context, user *User) context.Context {
	return context.WithValue(ctx, userKey, user)
}

func FromContext(ctx context.Context) (*User, bool) {
	user, ok := ctx.Value(userKey).(*User)
	return user, ok
}
