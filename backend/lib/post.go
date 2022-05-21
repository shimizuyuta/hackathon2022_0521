package lib

import (
	"context"
	"time"

	"github.com/squadra-ricordo/ent"
)

// 戻り値のない関数
func CreatePost(client *ent.Client, title string, description string, start_at time.Time, end_at time.Time, user_id int, ctx context.Context) {
	client.Post.
		Create().
		SetTitle(title).
		SetDescription(description).
		SetStartAt(start_at).
		SetEndAt(end_at).
		SetUserID(user_id).
		Save(ctx)
	return
}
