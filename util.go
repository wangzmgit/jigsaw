package jigsaw

import (
	"math/rand"
	"os"
	"time"
)

/**
 * 获取指定范围的随机数
 * param: min 随机数最小值
 * param: max 随机数最大值
 * return: 生成的随机数
 */
func getRandInt(min, max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max-min) + min
}

/**
 * 判断文件是否存在
 * param: path 文件路径
 * return: 文件是否存在，错误信息
 */
func pathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}
