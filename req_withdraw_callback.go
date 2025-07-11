package go_payasia

import (
	"errors"
	"fmt"
	"github.com/asaka1234/go-payasia/utils"
	"net/url"
)

// 提现的回调处理(传入一个处理函数)
func (cli *Client) WithdrawCallback(body string, req PayAsiaWithdrawBackReq, processor func(PayAsiaWithdrawBackReq) error) error {
	//不能用req这个结构体, 因为它的签名不过滤value为null/""的key,导致三方传参我们是无法判断的

	values, _ := url.ParseQuery(body)
	// 转换为单值 map
	paramMap := make(map[string]interface{})
	for key, val := range values {
		if len(val) > 0 {
			paramMap[key] = val[0]
		}
	}
	//-----------------------------

	sign := paramMap["sign"] //收到的签名
	delete(paramMap, "sign")
	expectedSign := utils.Sign(paramMap, cli.Params.AccessKey)
	if sign != expectedSign {
		fmt.Println("签名验证失败")
		return errors.New("sign verify failed!")
	}

	//开始处理
	return processor(req)
}
