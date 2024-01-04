package system

import (
	"ZZK_YUNYING_TASK/global"
	"ZZK_YUNYING_TASK/model/system"
	"ZZK_YUNYING_TASK/utils/upload"
	"fmt"
	"mime/multipart"
	"os/exec"
	"strings"
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
	fmt.Println("video对象:", video)
	if video.StartTime != "" || video.EndTime != "" {
		err = v.SliceVideo(filename, video.StartTime, video.EndTime)

		// 裁剪成功，把源文件删除
		if err == nil {

		}
	}

	// 将上传的文件存到数据库中
	s := strings.Split(header.Filename, ".")
	f := system.SysVideo{
		Url:  filePath,
		Name: header.Filename,
		Tag:  s[len(s)-1],
		Key:  filename,
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
func (v *SysVideoService) SliceVideo(fileName string, startTime string, endTime string) error {
	inputFile := upload.LOCAL_PATH + "/" + fileName
	outputFile := upload.LOCAL_PATH + "/" + "output_" + fileName

	// 设置 ffmpeg命令行参数
	args := []string{"-i", inputFile, "-ss", "00:03", "-to", "00:08", "-c:v", "libx264", "-crf", "30", outputFile}
	cmd := exec.Command("ffmpeg", args...)

	// 运行 ffmpeg 命令
	if err := cmd.Run(); err != nil {
		fmt.Println("ffmpeg转码失败：", err)
		return err
	}
	fmt.Println("ffmpeg转码成功")
	return nil
}
