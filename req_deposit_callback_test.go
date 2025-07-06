package go_payasia

import (
	"fmt"
	"github.com/asaka1234/go-payasia/utils"
	"github.com/mitchellh/mapstructure"
	"testing"
)

//amount=311760.00&currency=VND&merchant_reference=202507061318100463&request_reference=eac9f860-9f2b-445f-a332-bc66e27b45cc&status=2&sign=5d9a711dcc98b5e7b34560019a09b88637ba733556ba702af7f748ee9162297c98a622b1685d9c5392e00c1e2b8253d5fb65228dc30703942aa312b0d3f358f2

// 充值的回调处理(传入一个处理函数)
func TestDepositCallback(t *testing.T) {

	req := PayAsiaDepositBackReq{
		Status:            "2",
		MerchantReference: "202507061318100463",
		RequestReference:  "eac9f860-9f2b-445f-a332-bc66e27b45cc",
		Currency:          "VND",
		Amount:            "311760.00",
		Sign:              "5d9a711dcc98b5e7b34560019a09b88637ba733556ba702af7f748ee9162297c98a622b1685d9c5392e00c1e2b8253d5fb65228dc30703942aa312b0d3f358f2",
	}
	//验证
	var paramMap map[string]interface{}
	mapstructure.Decode(req, &paramMap)

	sign := paramMap["sign"].(string) //收到的签名
	delete(paramMap, "sign")
	expectedSign := utils.Sign(paramMap, "b47149cc-5ff2-4178-ab2f-46ec1f2573fc")
	if sign != expectedSign {
		fmt.Println("签名验证失败")
	}
}
