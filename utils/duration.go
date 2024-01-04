package utils

import (
	"strconv"
	"strings"
	"time"
)

func ParseDuration(d string) (time.Duration, error) {
	// 去除空格
	d = strings.TrimSpace(d)
	// 格式化时间
	dr, err := time.ParseDuration(d)

	if err == nil {
		return dr, nil
	}

	if strings.Contains(d, "d") {
		index := strings.Index(d, "d")

		hour, _ := strconv.Atoi(d[:index])
		dr = time.Hour * 24 * time.Duration(hour)
		ndr, err := time.ParseDuration(d[index+1:])
		if err != nil {
			return dr, nil
		}
		return dr + ndr, nil
	}
	dv, err := strconv.ParseInt(d, 10, 64)
	return time.Duration(dv), err
}
