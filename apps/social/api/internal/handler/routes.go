// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	friend "liteChat/apps/social/api/internal/handler/friend"
	group "liteChat/apps/social/api/internal/handler/group"
	"liteChat/apps/social/api/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/friend/putIn",
				Handler: friend.FriendPutInHandler(serverCtx),
			},
			{
				Method:  http.MethodPut,
				Path:    "/friend/putIn",
				Handler: friend.FriendPutInHandleHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/friend/putIns",
				Handler: friend.FriendPutInListHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/friends",
				Handler: friend.FriendListHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/friends/online",
				Handler: friend.FriendsOnlineHandler(serverCtx),
			},
		},
		rest.WithJwt(serverCtx.Config.JwtAuth.AccessSecret),
		rest.WithPrefix("/v1/social"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/group",
				Handler: group.CreateGroupHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/group/putIn",
				Handler: group.GroupPutInHandler(serverCtx),
			},
			{
				Method:  http.MethodPut,
				Path:    "/group/putIn",
				Handler: group.GroupPutInHandleHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/group/putIns",
				Handler: group.GroupPutInListHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/groups",
				Handler: group.GroupListHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/group/users",
				Handler: group.GroupUserListHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/group/users/online",
				Handler: group.GroupUserOnlineHandler(serverCtx),
			},
		},
		rest.WithJwt(serverCtx.Config.JwtAuth.AccessSecret),
		rest.WithPrefix("/v1/social"),
	)
}
