package system

import (
	"ZZK_YUNYING_TASK/model/system"
	"ZZK_YUNYING_TASK/model/system/request"
	"reflect"
	"testing"
)

func TestSysUserService_Login(t *testing.T) {
	type args struct {
		login request.Login
	}
	tests := []struct {
		name         string
		args         args
		wantSys_user *system.SysUser
		wantErr      bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sy := SysUserService{}
			gotSys_user, err := sy.Login(tt.args.login)
			if (err != nil) != tt.wantErr {
				t.Errorf("Login() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotSys_user, tt.wantSys_user) {
				t.Errorf("Login() gotSys_user = %v, want %v", gotSys_user, tt.wantSys_user)
			}
		})
	}
}

func TestSysUserService_Register(t *testing.T) {
	type args struct {
		register request.Register
	}
	tests := []struct {
		name     string
		args     args
		wantUser *system.SysUser
		wantErr  bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sy := SysUserService{}
			gotUser, err := sy.Register(tt.args.register)
			if (err != nil) != tt.wantErr {
				t.Errorf("Register() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotUser, tt.wantUser) {
				t.Errorf("Register() gotUser = %v, want %v", gotUser, tt.wantUser)
			}
		})
	}
}
