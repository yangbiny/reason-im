package rpcclient

type User struct {
	Id   int64
	Name string
}

type UserClient interface {
	NewUser(user User) User
	GetUserInfo(userId int64) User
}

type UserClientHandler struct {
}

func (client *UserClientHandler) NewUser(user User) User {
	//TODO implement me
	panic("implement me")
}

func (client *UserClientHandler) GetUserInfo(userId int64) User {
	//TODO implement me
	panic("implement me")
}
