package rpcclient

import (
	"context"
	"reflect"
	"testing"
)

func TestUserClientHandler_GetUserInfo(t *testing.T) {
	value := context.WithValue(nil, "test", "value")
	println(value)
}

func TestUserClientHandler_NewUser1(t *testing.T) {
	type fields struct {
		Client UserClient
	}
	type args struct {
		user *User
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *User
	}{
		{
			name: "测试创建用户",
			fields: fields{
				Client: UserClientHandler{},
			},
			args: args{
				user: &User{
					Name: "用户名称",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := UserClientHandler{
				Client: tt.fields.Client,
			}
			if got := u.NewUser(tt.args.user); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewUser() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserClientHandler_GetUserInfo1(t *testing.T) {
	type fields struct {
		Client UserClient
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
				Client: UserClientHandler{},
			},
			args: args{
				userId: 1,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := UserClientHandler{
				Client: tt.fields.Client,
			}
			if got := u.GetUserInfo(tt.args.userId); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetUserInfo() = %v, want %v", got, tt.want)
			}
		})
	}
}
