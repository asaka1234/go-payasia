package go_payasia

import (
	"crypto/tls"
	"fmt"
	"github.com/asaka1234/go-payasia/utils"
	"github.com/mitchellh/mapstructure"
	"github.com/samber/lo"
	"github.com/spf13/cast"
)

func (cli *Client) Withdraw(req PayAsiaWithdrawReq) (*PayAsiaWithdrawResponse, error) {
	rawURL := cli.Params.WithdrawUrl

	//----------------------判断bank code的正确性------------------
	_, ok := lo.Find(PayAsiaBankCodes, func(i PayAsiaBankCode) bool {
		return i.Code == req.BankName && i.Currency == req.Currency
	})
	if !ok {
		return nil, fmt.Errorf("bank code %s error", req.BankName)
	}
	//---------------------------------------------------------

	var paramMap map[string]interface{}
	mapstructure.Decode(req, &paramMap)

	//补充公共字段
	paramMap["datafeed_url"] = cli.Params.WithdrawBackUrl //回调地址

	signStr := utils.Sign(paramMap, cli.Params.AccessKey)
	paramMap["sign"] = signStr

	//-----构造一个form表单------------
	var result PayAsiaWithdrawResponse
	resp, err := cli.ryClient.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true}).
		SetCloseConnection(true).
		R().
		SetHeaders(getHeaders()).
		SetFormData(ConvertInterfaceMapToStringMap(paramMap)).
		SetDebug(cli.debugMode).
		SetResult(&result).
		Post(rawURL)

	if err != nil {
		return nil, err
	}

	if resp.StatusCode() != 200 {
		return nil, fmt.Errorf("status code: %d", resp.StatusCode())
	}

	if resp.Error() != nil {
		//反序列化错误会在此捕捉
		return nil, fmt.Errorf("%v, body:%s", resp.Error(), resp.Body())
	}

	//说明ajax请求失败了
	if result.Response.Code != "200" {
		return nil, fmt.Errorf("code=%s", result.Response.Code)
	}

	//说明请求成功了
	var payload PayAsiaWithdrawResponsePayload
	mapstructure.Decode(result.Payload, &payload)
	result.PayloadOptimize = payload

	return &result, nil
}

func ConvertInterfaceMapToStringMap(m map[string]interface{}) map[string]string {
	result := make(map[string]string)
	for k, v := range m {
		result[k] = cast.ToString(v)
	}
	return result
}
