package controller

import (
	S "github.com/RohanDoshi21/messaging-platform/api/service"
	H "github.com/RohanDoshi21/messaging-platform/handler"
	U "github.com/RohanDoshi21/messaging-platform/util"
	"github.com/gofiber/fiber/v2"
)

func GetGroupDetails(ctx *fiber.Ctx) error {
	var groupId string

	if groupId = ctx.Params("id"); groupId == "" {
		return H.BuildError(ctx, "Group ID is missing", fiber.ErrBadRequest.Code, nil)
	}

	groupBody := &S.GroupBody{
		GroupId: groupId,
	}

	pgTrx := U.GetPGTrxFromFiberCtx(ctx)

	group, err := S.GetGroupDetails(groupBody, pgTrx)
	if err != nil {
		return H.BuildError(ctx, "Error getting group details", fiber.ErrInternalServerError.Code, fiber.ErrInternalServerError)
	}

	return H.Success(ctx, fiber.Map{
		"ok":    1,
		"group": group,
	})

}

func CreateGroup(ctx *fiber.Ctx) error {
	createGrpBody := ctx.Locals("body").(*S.GroupCreateBody)

	pgTrx := U.GetPGTrxFromFiberCtx(ctx)

	user := U.GetAuthUser(ctx)
	createGrpBody.GroupOwner = user.Id
	groupid, inviteid, err := S.CreateGroup(createGrpBody, pgTrx)

	if err != nil {
		return H.BuildError(ctx, "Error creating group", fiber.ErrInternalServerError.Code, fiber.ErrInternalServerError)
	}

	return H.Success(ctx, fiber.Map{
		"ok":        1,
		"group id":  groupid,
		"invite id": inviteid,
	})
}

func JoinGroup(ctx *fiber.Ctx) error {
	var groupId string

	if groupId = ctx.Params("id"); groupId == "" {
		return H.BuildError(ctx, "Group ID is missing", fiber.ErrBadRequest.Code, nil)
	}

	user := U.GetAuthUser(ctx)
	groupBody := &S.GroupJoinBody{
		GroupId: groupId,
		UserId:  user.Id,
	}

	pgTrx := U.GetPGTrxFromFiberCtx(ctx)

	err := S.JoinGroup(groupBody, pgTrx)
	if err != nil {
		return H.BuildError(ctx, err.Message, fiber.ErrInternalServerError.Code, fiber.ErrInternalServerError)
	}

	return H.Success(ctx, fiber.Map{
		"ok":    1,
		"group": "1",
	})

}

func AddUserToGroup(ctx *fiber.Ctx) error {

	addUserBody := ctx.Locals("body").(*S.GroupAddUserBody)

	pgTrx := U.GetPGTrxFromFiberCtx(ctx)

	invite, group_id, err := S.AddUserToGroup(addUserBody, pgTrx)
	if err != nil {
		H.BuildError(ctx, err.Message, fiber.ErrInternalServerError.Code, fiber.ErrInternalServerError)
	}

	return H.Success(ctx, fiber.Map{
		"ok":        1,
		"invite_id": invite,
		"group_id":  group_id,
	})
}
