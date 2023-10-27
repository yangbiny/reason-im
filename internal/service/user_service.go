package service

import (
	"github.com/gin-gonic/gin"
	apierror "github.com/yangbiny/reason-commons/err"
	"reason-im/internal/config/web"
	"reason-im/internal/repo"
	"reason-im/pkg/dto/vo"
)

type UserService struct {
	UserDao       repo.UserDao
	UserFriendDao repo.FriendDao
}

func NewUserService(userDao repo.UserDao, fiendDao repo.FriendDao) UserService {
	return UserService{
		UserDao:       userDao,
		UserFriendDao: fiendDao,
	}
}

func (userService *UserService) NewUser(ctx *gin.Context, cmd *NewUserCmd) (*repo.User, *apierror.ApiError) {
	ctx2 := ctx.Request.Context()

	// 创建用户信息
	user := repo.User{
		Name: cmd.Name,
	}
	newUser, err := userService.UserDao.NewUser(&ctx2, &user)
	if err != nil {
		return nil, apierror.WhenServiceError(err)
	}
	return newUser, nil
}

func (userService *UserService) GetUserInfo(ctx *gin.Context, cmd *QueryUserCmd) (*vo.UserRelationVo, *apierror.ApiError) {
	ctx2 := ctx.Request.Context()

	info, err := userService.UserDao.GetUserInfo(&ctx2, cmd.QueryUid)
	if err != nil {
		return nil, apierror.WhenServiceError(err)
	}
	id := info.Id
	friendInfo, err := userService.UserFriendDao.QueryFriendInfo(&ctx2, cmd.UserId, id)
	if err != nil {
		return nil, apierror.WhenServiceError(err)
	}
	return &vo.UserRelationVo{
		Id:        id,
		Name:      info.Name,
		IsFriend:  friendInfo != nil,
		GmtCreate: info.GmtCreate,
	}, nil
}

func (userService *UserService) Login(ctx *gin.Context, user *UserLoginCmd) (bool, *apierror.ApiError) {
	ctx2 := ctx.Request.Context()

	queryUser, err := userService.UserDao.QueryUserByName(&ctx2, user.Name)
	if err != nil {
		return false, apierror.WhenServiceError(err)
	}
	if queryUser == nil {
		return false, nil
	}
	err = web.GenerateJwtToken(ctx, queryUser.Name, queryUser.Id)
	if err != nil {
		return false, apierror.WhenServiceError(err)
	}
	return true, nil
}

type UserLoginCmd struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

type NewUserCmd struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

type QueryUserCmd struct {
	UserId   int64 `login_user_id:"user_id"`
	QueryUid int64 `uri:"query_uid"`
}
