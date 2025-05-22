package go_payasia

import (
	"github.com/asaka1234/go-payasia/utils"
	"github.com/mitchellh/mapstructure"
)

// 并不需要发Hhtp请求出去,纯粹是一个计算签名
func (cli *Client) Withdraw(req PayAsiaWithdrawReq) (map[string]interface{}, error) {
	var paramMap map[string]interface{}
	mapstructure.Decode(req, &paramMap)

	//补充公共字段
	paramMap["datafeed_url"] = cli.WithdrawCallbackUrl //回调地址

	signStr := utils.Sign(paramMap, cli.AccessKey)
	paramMap["sign"] = signStr
	paramMap["url"] = cli.WithdrawUrl //实际前端post from action的地址

	return paramMap, nil
}
