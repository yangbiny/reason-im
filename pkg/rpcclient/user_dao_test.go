package rpcclient

import (
	"context"
	"reason-im/internal/config/mysql"
	mysql2 "reason-im/internal/utils/mysql"
	"reflect"
	"testing"
)

func TestUserClientHandler_GetUserInfo(t *testing.T) {
	value := context.WithValue(nil, "test", "value")
	println(value)
}

func TestUserClientHandler_NewUser1(t *testing.T) {
	datasource := mysql.Datasource()

	type fields struct {
		Client UserDao
	}
	type args struct {
		user *User
	}
	var tests = []struct {
		name   string
		fields fields
		args   args
		want   *User
	}{
		{
			name: "测试创建用户",
			fields: fields{
				Client: UserDaoImpl{
					DatabaseTpl: &mysql2.DatabaseTpl{Db: datasource},
				},
			},
			args: args{
				user: &User{
					Name: "用户名称2",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := tt.fields.Client
			if got := u.NewUser(tt.args.user); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewUser() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserClientHandler_GetUserInfo1(t *testing.T) {
	datasource := mysql.Datasource()
	type fields struct {
		Client UserDao
	}
	type args struct {
		userId int64
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *User
	}{
		{
			name: "测试查询用户",
			fields: fields{
				Client: UserDaoImpl{
					DatabaseTpl: &mysql2.DatabaseTpl{Db: datasource},
				},
			},
			args: args{
				userId: 4,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := tt.fields.Client
			u.GetUserInfo(tt.args.userId)
		})
	}
}
