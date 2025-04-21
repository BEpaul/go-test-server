package network

import (
	"sync"

	"github.com/BEpaul/go-test-server/service"
	"github.com/BEpaul/go-test-server/types"
	"github.com/gin-gonic/gin"
)

var (
	userRouterInit     sync.Once
	userRouterInstance *userRouter
)

type userRouter struct {
	router *Network
	// service

	userService *service.User
}

func newUserRouter(router *Network, userService *service.User) *userRouter {
	userRouterInit.Do(func() {
		userRouterInstance = &userRouter{
			router:      router,
			userService: userService,
		}

		router.registerGET("/", userRouterInstance.get)
		router.registerPOST("/", userRouterInstance.create)
		router.registerUPDATE("/", userRouterInstance.update)
		router.registerDELETE("/", userRouterInstance.delete)
	})

	return userRouterInstance
}

// register 유틸 함수들

func (u *userRouter) create(c *gin.Context) {
	var req types.CreateRequest

	// request 검증
	if err := c.ShouldBindJSON(&req); err != nil {
		u.router.failedResponse(c, &types.CreateUserResponse{
			ApiResponse: types.NewApiResponse("바인딩 오류입니다.", -1, err.Error()),
		})
	} else if err = u.userService.Create(req.ToUser()); err != nil {
		u.router.failedResponse(c, &types.CreateUserResponse{
			ApiResponse: types.NewApiResponse("Create 에러입니다.", -1, err.Error()),
		})
	} else {
		u.router.okResponse(c, &types.CreateUserResponse{
			ApiResponse: types.NewApiResponse("성공입니다", 1, nil),
		})
	}
}

func (u *userRouter) get(c *gin.Context) {
	u.router.okResponse(c, &types.GetUserResponse{
		ApiResponse: types.NewApiResponse("성공입니다.", 1, nil),
		Users:       u.userService.Get(),
	})
}

func (u *userRouter) update(c *gin.Context) {
	var req types.UpdateUserRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		u.router.failedResponse(c, &types.CreateUserResponse{
			ApiResponse: types.NewApiResponse("바인딩 오류입니다.", -1, err.Error()),
		})
	} else if err = u.userService.Update(req.Name, req.UpdatedAge); err != nil {
		u.router.failedResponse(c, &types.CreateUserResponse{
			ApiResponse: types.NewApiResponse("Update 에러입니다.", -1, err.Error()),
		})
	} else {
		u.router.okResponse(c, &types.UpdateUserResponse{
			ApiResponse: types.NewApiResponse("성공입니다.", 1, nil),
		})
	}
}

func (u *userRouter) delete(c *gin.Context) {
	var req types.DeleteRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		u.router.failedResponse(c, &types.DeleteUserResponse{
			ApiResponse: types.NewApiResponse("바인딩 오류입니다.", -1, err.Error()),
		})
	} else if err = u.userService.Delete(req.ToUser()); err != nil {
		u.router.failedResponse(c, &types.DeleteUserResponse{
			ApiResponse: types.NewApiResponse("Delete 에러입니다.", -1, err.Error()),
		})
	} else {
		u.router.okResponse(c, &types.DeleteUserResponse{
			ApiResponse: types.NewApiResponse("성공입니다.", 1, nil),
		})
	}
}
