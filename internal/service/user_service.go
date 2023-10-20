package service

import (
	"github.com/gin-gonic/gin"
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

func (userService *UserService) NewUser(ctx *gin.Context, cmd *NewUserCmd) (*repo.User, error) {
	// 创建用户信息
	user := repo.User{
		Name: cmd.Name,
	}
	return userService.UserDao.NewUser(&user)
}

func (userService *UserService) GetUserInfo(ctx *gin.Context, cmd *QueryUserCmd) (*vo.UserRelationVo, error) {
	info, err := userService.UserDao.GetUserInfo(cmd.QueryUid)
	if err != nil {
		return nil, err
	}
	id := info.Id
	friendInfo, err := userService.UserFriendDao.QueryFriendInfo(cmd.UserId, id)
	if err != nil {
		return nil, err
	}
	return &vo.UserRelationVo{
		Id:        id,
		Name:      info.Name,
		IsFriend:  friendInfo != nil,
		GmtCreate: info.GmtCreate,
	}, nil
}

func (userService *UserService) Login(c *gin.Context, user *UserLoginCmd) (bool, error) {
	queryUser, err := userService.UserDao.QueryUserByName(user.Name)
	if err != nil {
		return false, err
	}
	if queryUser == nil {
		return false, nil
	}
	err = web.GenerateJwtToken(c, queryUser.Name, queryUser.Id)
	if err != nil {
		return false, err
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
