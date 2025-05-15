package utils

import (
	"crypto/sha512"
	"encoding/hex"
)

func SHA512(text string) string {
	// 创建SHA512哈希对象
	hasher := sha512.New()

	// 写入数据
	hasher.Write([]byte(text))

	// 计算哈希值
	hashBytes := hasher.Sum(nil)

	// 转换为十六进制字符串
	return hex.EncodeToString(hashBytes)
}
