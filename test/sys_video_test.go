package test

import (
	"ZZK_YUNYING_TASK/model/commen/request"
	"ZZK_YUNYING_TASK/model/system"
	sysReq "ZZK_YUNYING_TASK/model/system/request"
	system2 "ZZK_YUNYING_TASK/service/system"
	"mime/multipart"
	"reflect"
	"testing"
)

func TestSysVideoService_DownloadVideo(t *testing.T) {
	type args struct {
		id string
	}
	tests := []struct {
		name      string
		args      args
		wantVideo *system.SysVideo
		wantErr   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := &system2.SysVideoService{}
			gotVideo, err := v.DownloadVideo(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("DownloadVideo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotVideo, tt.wantVideo) {
				t.Errorf("DownloadVideo() gotVideo = %v, want %v", gotVideo, tt.wantVideo)
			}
		})
	}
}

func TestSysVideoService_GetVideoListByUserId(t *testing.T) {
	type args struct {
		info    request.ListInfo
		user_id uint
	}
	tests := []struct {
		name      string
		args      args
		wantList  interface{}
		wantTotal int64
		wantErr   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := &system2.SysVideoService{}
			gotList, gotTotal, err := v.GetVideoListByUserId(tt.args.info, tt.args.user_id)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetVideoListByUserId() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotList, tt.wantList) {
				t.Errorf("GetVideoListByUserId() gotList = %v, want %v", gotList, tt.wantList)
			}
			if gotTotal != tt.wantTotal {
				t.Errorf("GetVideoListByUserId() gotTotal = %v, want %v", gotTotal, tt.wantTotal)
			}
		})
	}
}

func TestSysVideoService_UploadVideo(t *testing.T) {
	type args struct {
		header *multipart.FileHeader
		video  sysReq.UploadVideoParams
	}
	tests := []struct {
		name     string
		args     args
		wantFile *system.SysVideo
		wantErr  bool
	}{
		// TODO: Add test cases.

	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := &system2.SysVideoService{}
			gotFile, err := v.UploadVideo(tt.args.header, tt.args.video)
			if (err != nil) != tt.wantErr {
				t.Errorf("UploadVideo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotFile, tt.wantFile) {
				t.Errorf("UploadVideo() gotFile = %v, want %v", gotFile, tt.wantFile)
			}
		})
	}
}
