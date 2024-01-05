package system

import (
	"ZZK_YUNYING_TASK/global"
	"ZZK_YUNYING_TASK/model/commen/request"
	"ZZK_YUNYING_TASK/model/system"
	"ZZK_YUNYING_TASK/utils/upload"
	"errors"
	"mime/multipart"
	"os"
	"os/exec"
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
func (v *SysVideoService) UploadVideo(header *multipart.FileHeader, video system.SysVideo) (file system.SysVideo, err error) {

	oss := upload.NewOOS()
	filePath, filename, uploadErr := oss.UploadFile(header)
	// 如果上传失败
	if uploadErr != nil {
		panic(err)
	}
	if video.StartTime != "" || video.EndTime != "" {
		newFilePath, err := v.SliceVideo(filename, video.StartTime, video.EndTime)
		// 裁剪成功，把源文件删除
		if err == nil {
			err2 := v.DeleteVideo(filename, oss)
			if err2 != nil {
				global.TASK_LOGGER.Error("删除失败文件名：", zap.Error(err))
			}
			filename = "output_" + filename
			filePath = newFilePath
		}
	}

	// 将上传的文件存到数据库中
	s := strings.Split(header.Filename, ".")
	f := system.SysVideo{
		Url:    filePath,
		Name:   header.Filename,
		Tag:    s[len(s)-1],
		Key:    filename,
		UserId: video.UserId,
	}
	err = v.Upload(&f)
	return f, err

}

// @function: Upload
// @description: 创建文件上传记录
// @param: file *system.SysVideo
// @return: error
func (v *SysVideoService) Upload(file *system.SysVideo) error {
	return global.TASK_DB.Create(file).Error
}

// @function: SliceVideo
// @description: 截取视频
// @param: file *system.SysVideo
// @return: error
func (v *SysVideoService) SliceVideo(fileName string, startTime string, endTime string) (string, error) {
	inputFile := upload.LOCAL_PATH + "/" + fileName
	outputFile := upload.LOCAL_PATH + "/" + "output_" + fileName

	// 设置 ffmpeg命令行参数
	args := []string{"-i", inputFile, "-ss", "00:03", "-to", "00:08", "-c:v", "libx264", "-crf", "30", outputFile}
	cmd := exec.Command("ffmpeg", args...)

	// 运行 ffmpeg 命令
	if err := cmd.Run(); err != nil {
		global.TASK_LOGGER.Error("转码失败：", zap.Error(err))
		return outputFile, err
	}
	return outputFile, nil
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
