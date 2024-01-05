package test

import (
	"ZZK_YUNYING_TASK/model/system"
	"ZZK_YUNYING_TASK/model/system/request"
	system2 "ZZK_YUNYING_TASK/service/system"
	"reflect"
	"testing"
)

func TestSysUserService_Login(t *testing.T) {
	type args struct {
		login request.Login
	}
	tests := []struct {
		name          string
		args          args
		wantInnerUser system.SysUser
		wantErr       bool
	}{
		{
			name: "login valid",
			args: args{
				login: request.Login{
					NickName: "string",
					Password: "string",
				},
			},
			wantErr: true,
		},
		{
			name: "login valid",
			args: args{
				login: request.Login{
					NickName: "string",
					Password: "123456",
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sy := system2.SysUserService{}
			gotInnerUser, err := sy.Login(tt.args.login)
			if (err != nil) != tt.wantErr {
				t.Errorf("Register() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotInnerUser, tt.wantInnerUser) {
				t.Errorf("Register() gotInnerUser = %v, want %v", gotInnerUser, tt.wantInnerUser)
			}
		})
	}
}

func TestSysUserService_Register(t *testing.T) {
	type args struct {
		register request.Register
	}
	tests := []struct {
		name          string
		args          args
		wantInnerUser system.SysUser
		wantErr       bool
	}{
		// TODO: Add test cases.
		{
			name: "register valid",
			args: args{
				register: request.Register{
					NickName: "string",
					Password: "string",
				},
			},
			wantErr: false,
		},
		{
			name: "register valid",
			args: args{
				register: request.Register{
					NickName: "asd",
					Password: "123456",
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sy := system2.SysUserService{}
			gotInnerUser, err := sy.Register(tt.args.register)
			if (err != nil) != tt.wantErr {
				t.Errorf("Register() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotInnerUser, tt.wantInnerUser) {
				t.Errorf("Register() gotInnerUser = %v, want %v", gotInnerUser, tt.wantInnerUser)
			}
		})
	}
}
