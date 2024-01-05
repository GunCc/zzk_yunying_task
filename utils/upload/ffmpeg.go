package upload

import (
	"ZZK_YUNYING_TASK/global"
	"fmt"
	"go.uber.org/zap"
	"io"
	"os"
	"os/exec"
)

// @function: SliceVideo
// @description: 截取视频
// @param: file *system.SysVideo
// @return: error
func SliceVideo(fileName string, startTime string, endTime string) (string, error) {
	inputFile := LOCAL_PATH + "/" + fileName
	outputFile := LOCAL_PATH + "/" + "output_" + fileName

	// 设置 ffmpeg命令行参数
	args := []string{
		"-i", inputFile,
		"-ss", startTime, "-to", endTime,
		"-c:v", "libx264", "-crf", "30",
		outputFile}
	cmd := exec.Command("ffmpeg", args...)

	// 获取输出管道
	stdout, err := cmd.StdoutPipe()

	if err != nil {
		global.TASK_LOGGER.Error("管道获取失败：", zap.Error(err))
		return "", err
	}

	// 开始执行命令
	if err := cmd.Start(); err != nil {
		fmt.Println(err)
		return "", err
	}
	// 读取输出
	go func() {
		if _, err := io.Copy(os.Stdout, stdout); err != nil {
			global.TASK_LOGGER.Error("读取输出失败：", zap.Error(err))
			return
		}
	}()

	// 等待命令执行完成
	if err := cmd.Wait(); err != nil {
		global.TASK_LOGGER.Error("转码失败：", zap.Error(err))
		return "", err
	}

	return outputFile, nil
}
