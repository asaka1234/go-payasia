package go_payasia

import (
	"fmt"
	"testing"
)

func TestDeposit(t *testing.T) {

	//构造client
	cli := NewClient(nil, MERCHANT_ID, MERCHANT_TOKEN, SECRECT_CODE, DEPOSIT_URL, WITHDRAW_URL, DEPOSIT_CALLBACK_URL, WITHDRAW_CALLBACK_URL)

	//发请求
	resp, err := cli.Deposit(GenDepositRequestDemo())
	if err != nil {
		fmt.Printf("err:%s\n", err.Error())
		return
	}
	fmt.Printf("resp:%+v\n", resp)
}

func GenDepositRequestDemo() PayAsiaDepositReq {
	return PayAsiaDepositReq{
		MerchantReference: "1234567890", //商户的orderId
		Currency:          "THB",
		Amount:            600.00,            // 使用string避免精度问题，对应Java的BigDecimal
		CustomerIp:        "123.123.123.123", //
		CustomerFirstName: "John",
		CustomerLastName:  "Doe",
		CustomerPhone:     "0123123123",
		CustomerEmail:     "someone@gmail.com",
	}
}
