package util

import (
	"crypto/md5"
	"encoding/hex"
	uuid "github.com/satori/go.uuid"
	"math/rand"
	"path"
	"strings"
	"time"
)

//MD5 计算md5
func MD5(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}

//GetRandString 生成随机字符串
func GetRandString(num int) string {
	rand.Seed(time.Now().UnixNano())
	var rangString string
	baseStr := "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	for i := 0; i < num; i++ {
		tmpNum := rand.Intn(len(baseStr) - 1)
		rangString += baseStr[tmpNum : tmpNum+1]
	}
	return rangString
}

//Uuid 生成一个uuid
func Uuid() string {
	return uuid.NewV4().String()
}

//GetFileExt 获取文件或者带文件的路径中的后缀名 并转小写
func GetFileExt(filePath string) string {
	return strings.ToLower(path.Ext(filePath))
}

//InArray 判断传入的值 是否存在一个切片里
func InArray(needle interface{}, slice interface{}) bool {
	switch key := needle.(type) {
	case string:
		for _, item := range slice.([]string) {
			if key == item {
				return true
			}
		}
	case int:
		for _, item := range slice.([]int) {
			if key == item {
				return true
			}
		}
	case int64:
		for _, item := range slice.([]int64) {
			if key == item {
				return true
			}
		}
	default:
		return false
	}
	return false
}