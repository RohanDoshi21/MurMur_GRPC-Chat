package service

import (
	"context"
	"database/sql"

	"github.com/RohanDoshi21/messaging-platform/models"
	T "github.com/RohanDoshi21/messaging-platform/types"
)

type Messages struct {
	UserID string
}

func GetUserMessages(UserMessages *Messages, dbTrx *sql.Tx) (models.MessageSlice, *T.ServiceError) {
	ctx := context.Background()

	messages, err := models.Messages(models.MessageWhere.Receiver.EQ(UserMessages.UserID)).All(ctx, dbTrx)

	if err != nil {
		return nil, &T.ServiceError{
			Code:    500,
			Message: "Failed to fetch messages",
			Error:   err,
		}
	}

	return messages, nil

}
