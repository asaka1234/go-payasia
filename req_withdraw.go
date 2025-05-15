package go_payasia

import (
	"github.com/asaka1234/go-payasia/utils"
	"github.com/fatih/structs"
)

// 并不需要发Hhtp请求出去,纯粹是一个计算签名
func (cli *Client) Withdraw(req PayAsiaWithdrawReq) (map[string]interface{}, error) {

	req.DatafeedUrl = cli.WithdrawCallbackUrl
	//构造请求(加签名)
	paramMap := structs.Map(req)

	//拿到签名的部分
	signKeyList := []string{
		"request_reference",
		"beneficiary_name",
		"beneficiary_first_name",
		"beneficiary_last_name",
		"bank_name",
		"beneficiary_email",
		"beneficiary_phone",
		"account_number",
		"currency",
		"amount",
		"datafeed_url",
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
