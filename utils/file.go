package utils

import (
	"crypto/md5"
	"encoding/hex"
	"io"
	"strings"
)

// ExtractFileNameAndType 提取文件名和文件类型
func ExtractFileNameAndType(str string) (string, string) {
	fileType := ""
	splitFile := strings.Split(str, ".")
	fileName := strings.Join(splitFile[0:len(splitFile)-1], ".")
	if len(splitFile) > 0 {
		fileType = splitFile[len(splitFile)-1]
	}
	return fileName, fileType
}

// CalcFileMD5 计算文件的MD5值
func CalcFileMD5(file io.Reader) (string, error) {
	hash := md5.New()
	_, _ = io.Copy(hash, file)
	return hex.EncodeToString(hash.Sum(nil)), nil
}
