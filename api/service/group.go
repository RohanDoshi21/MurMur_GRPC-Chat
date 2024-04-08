package service

import (
	"context"
	"database/sql"

	"github.com/RohanDoshi21/messaging-platform/models"
	M "github.com/RohanDoshi21/messaging-platform/models"
	T "github.com/RohanDoshi21/messaging-platform/types"
	U "github.com/RohanDoshi21/messaging-platform/util"
	"github.com/gofiber/fiber/v2"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

type GroupBody struct {
	GroupId string `json:"group_id"`
}

type GroupCreateBody struct {
	GroupName  string   `json:"group_name"`
	GroupUsers []string `json:"group_users"`
	GroupOwner string   `json:"group_owner"`
}

type GroupJoinBody struct {
	GroupId string `json:"group_id"`
	UserId  string `json:"user_id"`
}

type GroupAddUserBody struct {
	GroupId string   `json:"group_id"`
	UserId  []string `json:"user_id"`
}

func GetGroupDetails(groupBody *GroupBody, dbTrx *sql.Tx) (interface{}, *T.ServiceError) {

	ctx := context.Background()

	isExist, err := M.GroupExists(ctx, dbTrx, groupBody.GroupId)
	if err != nil {
		return nil, &T.ServiceError{
			Code:    500,
			Message: "Failed to fetch group",
			Error:   err,
		}
	}

	if !isExist {
		return nil, &T.ServiceError{
			Code:    fiber.ErrBadRequest.Code,
			Message: "Group not found",
			Error:   err,
		}
	}

	group, err := models.Groups(models.GroupWhere.ID.EQ(groupBody.GroupId)).One(ctx, dbTrx)
	if err != nil {
		return nil, &T.ServiceError{
			Code:    fiber.ErrInternalServerError.Code,
			Message: "Failed to fetch group",
			Error:   err,
		}
	}

	return group, nil

}

func CreateGroup(createGrpBody *GroupCreateBody, dbTrx *sql.Tx) (string, string, *T.ServiceError) {

	dbCtx := context.Background()

	group := &models.Group{
		ID:    U.UUID(),
		Name:  createGrpBody.GroupName,
		Owner: createGrpBody.GroupOwner,
		Users: []string{createGrpBody.GroupOwner},
	}

	err := group.Insert(dbCtx, dbTrx, boil.Infer())
	if err != nil {
		return "", "", &T.ServiceError{
			Code:    fiber.ErrInternalServerError.Code,
			Message: "Failed to create group",
			Error:   err,
		}
	}

	invite := &models.Invite{
		ID:        U.UUID(),
		GroupID:   group.ID,
		Sender:    createGrpBody.GroupOwner,
		Receiver:  createGrpBody.GroupUsers,
		TimesUsed: 0,
	}
	err = invite.Insert(dbCtx, dbTrx, boil.Infer())

	if err != nil {
		return "", "", &T.ServiceError{
			Code:    fiber.ErrInternalServerError.Code,
			Message: "Failed to create invite",
			Error:   err,
		}
	}

	return group.ID, invite.ID, nil

}

func JoinGroup(groupBody *GroupJoinBody, dbTrx *sql.Tx) *T.ServiceError {

	ctx := context.Background()

	isExist, err := M.GroupExists(ctx, dbTrx, groupBody.GroupId)
	if err != nil {
		return &T.ServiceError{
			Code:    500,
			Message: "Failed to fetch group",
			Error:   err,
		}
	}

	if !isExist {
		return &T.ServiceError{
			Code:    fiber.ErrBadRequest.Code,
			Message: "Group not found",
			Error:   err,
		}
	}
	group, err := models.Groups(models.GroupWhere.ID.EQ(groupBody.GroupId)).One(ctx, dbTrx)
	if err != nil {
		return &T.ServiceError{
			Code:    fiber.ErrInternalServerError.Code,
			Message: "Failed to fetch group",
			Error:   err,
		}
	}
	for _, user := range group.Users {
		if user == groupBody.UserId {
			return &T.ServiceError{
				Code:    fiber.ErrBadRequest.Code,
				Message: "User already in group",
				Error:   nil,
			}
		}
	}

	inviteDetails, err := models.Invites(models.InviteWhere.GroupID.EQ(groupBody.GroupId)).All(ctx, dbTrx)

	if err != nil {
		return &T.ServiceError{
			Code:    fiber.ErrInternalServerError.Code,
			Message: "Failed to fetch invite",
			Error:   err,
		}
	}
	// if user in the invite list append the user to the group
	for _, invite := range inviteDetails {
		for _, receiver := range invite.Receiver {
			if receiver == groupBody.UserId {
				group.Users = append(group.Users, groupBody.UserId)
				_, err := group.Update(ctx, dbTrx, boil.Infer())
				if err != nil {
					return &T.ServiceError{
						Code:    fiber.ErrInternalServerError.Code,
						Message: "Failed to update group",
						Error:   err,
					}
				}
				indexToRemove := U.FindIndex(invite.Receiver, receiver)
				invite.Receiver = append(invite.Receiver[:indexToRemove], invite.Receiver[indexToRemove+1:]...)
				_, err = invite.Update(ctx, dbTrx, boil.Infer())
				if err != nil {
					return &T.ServiceError{
						Code:    fiber.ErrInternalServerError.Code,
						Message: "Failed to update invite",
						Error:   err,
					}
				}
				return nil
			} else {
				return &T.ServiceError{
					Code:    fiber.ErrBadRequest.Code,
					Message: "User not invited",
					Error:   nil,
				}
			}

		}
	}

	return nil

}

func AddUserToGroup(addUserBody *GroupAddUserBody, dbTrx *sql.Tx) *T.ServiceError {

	ctx := context.Background()

	isExist, err := M.GroupExists(ctx, dbTrx, addUserBody.GroupId)
	if err != nil {
		return &T.ServiceError{
			Code:    500,
			Message: "Failed to fetch group",
			Error:   err,
		}
	}

	if !isExist {
		return &T.ServiceError{
			Code:    fiber.ErrBadRequest.Code,
			Message: "Group not found",
			Error:   err,
		}
	}

	group, err := models.Groups(models.GroupWhere.ID.EQ(addUserBody.GroupId)).One(ctx, dbTrx)
	if err != nil {
		return &T.ServiceError{
			Code:    fiber.ErrInternalServerError.Code,
			Message: "Failed to fetch group",
			Error:   err,
		}
	}
	for _, user := range addUserBody.UserId {
		for _, groupUser := range group.Users {
			if user == groupUser {
				return &T.ServiceError{
					Code:    fiber.ErrBadRequest.Code,
					Message: "User already in group",
					Error:   nil,
				}
			}
		}
	}

	invite := &models.Invite{
		ID:       U.UUID(),
		GroupID:  addUserBody.GroupId,
		Sender:   group.Owner,
		Receiver: addUserBody.UserId,
	}

	err = invite.Insert(ctx, dbTrx, boil.Infer())

	if err != nil {
		return &T.ServiceError{
			Code:    fiber.ErrInternalServerError.Code,
			Message: "Failed to create invite",
			Error:   err,
		}
	}

	return nil

}
