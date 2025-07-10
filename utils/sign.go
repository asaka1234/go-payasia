package utils

import (
	"fmt"
	"github.com/samber/lo"
	"github.com/spf13/cast"
	"net/url"
	"sort"
)

// 计算请求签名
func Sign(paramMap map[string]string, accessKey string) string {

	keys := lo.Keys(paramMap)

	// 1. 参数名按照ASCII码表升序排序
	sort.Strings(keys)

	// Build query string
	//拼接签名原始字符串
	//rawString := ""
	values := url.Values{}
	lo.ForEach(keys, func(x string, index int) {
		value := cast.ToString(paramMap[x])
		values.Add(x, value)
	})
	queryString := values.Encode()

	// 3. 将secretKey拼接到最后
	queryString += accessKey

	fmt.Printf("[rawString]%s\n", queryString)

	//4. 计算签名
	signResult := SHA512(queryString)
	return signResult
}
