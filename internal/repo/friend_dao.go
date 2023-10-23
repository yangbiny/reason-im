package repo

import (
	"context"
	"fmt"
	mysql2 "reason-im/internal/utils/mysql"
	"reason-im/pkg/model"
)

var friendTableName = "im_friend"
var friendColumns = "id,user_id,friend_id,status,remark,gmt_create,gmt_update"

type Friend = model.Friend

type DeleteFriendCmd struct {
	Id     int64
	Status model.FriendStatus
}

type FriendDao interface {
	NewFriend(friend *Friend) (*Friend, error)
	UpdateFriend(friend *Friend) (bool, error)
	GetFriendInfo(friendId int64) (*Friend, error)
	DeleteFriend(cmd DeleteFriendCmd) (bool, error)
	QueryFriendList(userId int64) ([]*Friend, error)
	QueryFriendInfo(userId int64, friendId int64) (*Friend, error)
}

type FriendDaoImpl struct {
	DatabaseTpl *mysql2.DatabaseTpl
}

func NewFriendDao(tpl *mysql2.DatabaseTpl) FriendDao {
	return FriendDaoImpl{
		DatabaseTpl: tpl,
	}
}

func (dao FriendDaoImpl) UpdateFriend(friend *Friend) (bool, error) {
	var sql = fmt.Sprintf("update %s set status = ?,remark = ?,gmt_update = ? where id = ?", friendTableName)
	rows, err := dao.DatabaseTpl.Update(context.Background(), sql, friend.Status, friend.Remark, friend.GmtUpdate, friend.Id)
	if err != nil {
		return false, err
	}
	if rows == 0 {
		return false, fmt.Errorf("update friend error")
	}
	return true, nil
}

func (dao FriendDaoImpl) NewFriend(friend *Friend) (*Friend, error) {
	var sql = fmt.Sprintf("insert into %s (user_id,friend_id,status,remark,gmt_create,gmt_update) values (?,?,?,?,?,?)", friendTableName)
	insert, err := dao.DatabaseTpl.Insert(context.Background(), sql, friend.UserId, friend.FriendId, friend.Status, friend.Remark, friend.GmtCreate, friend.GmtUpdate)
	if err != nil {
		return nil, err
	}
	friend.Id = insert
	return friend, nil
}

func (dao FriendDaoImpl) GetFriendInfo(friendId int64) (*Friend, error) {
	var sql = fmt.Sprintf("select %s from %s where id = ?", friendColumns, friendTableName)
	one, err := dao.DatabaseTpl.FindOne(context.Background(), sql, Friend{}, friendId)
	if err != nil || one == nil {
		return nil, err
	}
	friend := one.(Friend)
	return &friend, nil
}

func (dao FriendDaoImpl) DeleteFriend(cmd DeleteFriendCmd) (bool, error) {
	var sql = fmt.Sprintf("update %s set status = ? where id = ?", friendTableName)
	rows, err := dao.DatabaseTpl.Update(context.Background(), sql, cmd.Status, cmd.Id)
	if err != nil {
		return false, err
	}
	return rows > 0, nil
}

func (dao FriendDaoImpl) QueryFriendList(userId int64) ([]*Friend, error) {
	var sql = fmt.Sprintf("select %s from %s where user_id = ?", friendColumns, friendTableName)
	rows, err := dao.DatabaseTpl.FindList(context.Background(), sql, Friend{}, userId)
	if err != nil {
		return nil, err
	}
	var friends []*Friend
	for _, row := range rows {
		friend := row.(Friend)
		friends = append(friends, &friend)
	}
	return friends, nil
}

func (dao FriendDaoImpl) QueryFriendInfo(userId int64, friendId int64) (*Friend, error) {
	var sql = fmt.Sprintf("select %s from %s where user_id = ? and friend_id = ?", friendColumns, friendTableName)
	one, err := dao.DatabaseTpl.FindOne(context.Background(), sql, Friend{}, userId, friendId)
	if err != nil || one == nil {
		return nil, err
	}
	friend := one.(Friend)
	return &friend, nil
}
