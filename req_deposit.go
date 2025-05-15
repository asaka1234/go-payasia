package go_payasia

import (
	"github.com/asaka1234/go-payasia/utils"
	"github.com/fatih/structs"
)

// 并不需要发Hhtp请求出去,纯粹是一个计算签名
func (cli *Client) Deposit(req PayAsiaDepositReq) (map[string]interface{}, error) {

	//补充一下
	req.NotifyUrl = cli.DepositCallbackUrl

	//构造请求(加签名)
	paramMap := structs.Map(req)

	//拿到签名的部分
	signKeyList := []string{
		"merchant_reference",
		"currency",
		"amount",
		"customer_ip",
		"customer_first_name",
		"customer_last_name",
		"customer_phone",
		"customer_email",
		"network",
		"notify_url",
	}

	//拿到签名map
	signMap := make(map[string]interface{})
	for _, key := range signKeyList {
		signMap[key] = paramMap[key]
	}

	signStr := utils.Sign(signMap, cli.AccessKey)
	paramMap["sign"] = signStr
	paramMap["url"] = cli.DepositUrl

	return paramMap, nil
}
