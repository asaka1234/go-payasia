package go_payasia

import (
	"github.com/asaka1234/go-payasia/utils"
	"github.com/mitchellh/mapstructure"
)

// 并不需要发Hhtp请求出去,纯粹是一个计算签名
func (cli *Client) Deposit(req PayAsiaDepositReq) (map[string]interface{}, error) {

	var paramMap map[string]interface{}
	mapstructure.Decode(req, &paramMap)

	//补充公共字段
	paramMap["return_url"] = cli.Params.DepositFeBackUrl //回调地址
	paramMap["network"] = "DirectDebit"                  //写死

	signStr := utils.Sign(paramMap, cli.Params.AccessKey)
	paramMap["sign"] = signStr
	paramMap["url"] = cli.Params.DepositUrl //实际前端post from action的地址
	paramMap["paySendPost"] = true
	
	return paramMap, nil
}
