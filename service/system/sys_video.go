package system

import (
	"ZZK_YUNYING_TASK/global"
	"ZZK_YUNYING_TASK/model/commen/request"
	"ZZK_YUNYING_TASK/model/system"
	sysReq "ZZK_YUNYING_TASK/model/system/request"
	"ZZK_YUNYING_TASK/utils/upload"
	"errors"
	"mime/multipart"
	"os"
	"strings"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

type SysVideoService struct {
}

// @function: UploadVideo
// @description: 上传视频
// @param: header *multipart.FileHeader, noSave string
// @return: file system.SysVideo, err error
func (v *SysVideoService) UploadVideo(header *multipart.FileHeader, video sysReq.UploadVideoParams) (file *system.SysVideo, err error) {

	oss := upload.NewOOS()
	filePath, inputFileName, uploadErr := oss.UploadFile(header)

	// 如果上传失败
	if uploadErr != nil {
		panic(err)
	}

	s := strings.Split(header.Filename, ".")
	// 1. 上传成功后保存基本信息
	f := system.SysVideo{
		Url:           filePath,
		InputFileName: inputFileName,
		UserId:        video.UserId,
		Status:        0,
		Name:          header.Filename,
		Tag:           s[len(s)-1],
	}

	err = v.Create(&f)
	if err != nil {
		return nil, errors.New("视频创建失败")
	}

	// 2. 视频裁剪转码
	newFilePath, err := upload.SliceVideo(inputFileName, video.StartTime, video.EndTime)
	outputFileName := "output_" + inputFileName
	// 3. 错误处理
	if err != nil {
		f.Status = 2
		err = v.Update(&f)
		if err == nil {
			err = errors.New("裁剪失败")
		}
	} else {
		// 4. 成功处理
		f.Status = 1
		f.OutputFileName = outputFileName
		f.Url = newFilePath
		// Save 0 表示删除源文件
		if video.Save == "0" {
			err = v.DeleteVideo(inputFileName, oss)
			if err != nil {
				global.TASK_LOGGER.Error("删除失败文件名：", zap.Error(err))
			}
			// 删除了源文件没有保存就不显示了
			inputFileName = ""
		}
		err = v.Update(&f)
	}

	return &f, err
}

// @function: Create
// @description: 创建文件上传记录
// @param: file *system.SysVideo
// @return: error
func (v *SysVideoService) Create(file *system.SysVideo) error {
	return global.TASK_DB.Create(file).Error
}

// @function: Update
// @description: 修改文件上传记录
// @param: file *system.SysVideo
// @return: error
func (v *SysVideoService) Update(file *system.SysVideo) error {
	return global.TASK_DB.Updates(file).Error
}

// @function: DeleteVideo
// @description: 删除视频
// @param: file *system.SysVideo
// @return: error
func (v *SysVideoService) DeleteVideo(fileName string, oss upload.OOS) (err error) {
	if err = oss.DeleteFile(fileName); err != nil {
		return err
	}
	return err
}

// @function: GetVideoListByUserId
// @description: 根据用户ID获取视频列表
// @param: file *system.SysVideo
// @return: error
func (v *SysVideoService) GetVideoListByUserId(info request.ListInfo, user_id uint) (list interface{}, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)

	db := global.TASK_DB.Model(&system.SysVideo{})
	var videoList []system.SysVideo

	db = db.Limit(limit).Offset(offset).Where("user_id = ?", user_id)
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Find(&videoList).Error

	return videoList, total, err
}

// @function: DownloadVideo
// @description: 根据视频ID下载视频
// @param: file *system.SysVideo
// @return: error
func (v *SysVideoService) DownloadVideo(id string) (video *system.SysVideo, err error) {
	db := global.TASK_DB.Model(&system.SysVideo{})
	if errors.Is(db.Where("id = ?", id).Find(&video).Error, gorm.ErrRecordNotFound) {
		return nil, errors.New("视频不存在")
	}
	_, err = os.Stat(video.Url)
	if err != nil {
		global.TASK_LOGGER.Error("视频不存在!", zap.Error(err))
		return nil, errors.New("视频不存在")
	}

	return video, nil
}
