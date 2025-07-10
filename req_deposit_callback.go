package go_payasia

import (
	"errors"
	"fmt"
	"github.com/asaka1234/go-payasia/utils"
	"github.com/mitchellh/mapstructure"
)

// 充值的回调处理(传入一个处理函数)
func (cli *Client) DepositCallback(req PayAsiaDepositBackReq, processor func(PayAsiaDepositBackReq) error) error {
	//验证
	var paramMap map[string]interface{}
	mapstructure.Decode(req, &paramMap)

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
