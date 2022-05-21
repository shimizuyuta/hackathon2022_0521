package lib

import (
	"context"

	"github.com/squadra-ricordo/ent"
)

// 戻り値のない関数
func Register(client *ent.Client, name string, email string, password_hash string, ctx context.Context) {
	client.User.
		Create().
		SetName(name).
		SetEmail(email).
		SetPasswordHash(password_hash).
		Save(ctx)
	return
}
