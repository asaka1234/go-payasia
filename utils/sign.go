package utils

import (
	"fmt"
	"github.com/samber/lo"
	"github.com/spf13/cast"
	"sort"
	"strings"
)

// 计算请求签名
func Sign(paramMap map[string]interface{}, accessKey string) string {

	keys := lo.Keys(paramMap)

	// 1. 参数名按照ASCII码表升序排序
	sort.Strings(keys)

	// Build query string
	var sb strings.Builder
	for index, k := range keys {
		value := cast.ToString(paramMap[k])
		sb.WriteString(fmt.Sprintf("%s=%s&", k, value))
	}
	sb.WriteString(accessKey)
	queryString := sb.String()
	queryString := strings.TrimSuffix(queryString, "&")
	fmt.Printf("[rawString]%s\n", queryString)

	//4. 计算签名
	signResult := SHA512(queryString)
	return signResult
}
