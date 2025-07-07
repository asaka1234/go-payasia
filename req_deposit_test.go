package go_payasia

import (
	"fmt"
	"testing"
)

type VLog struct {
}

func (l VLog) Debugf(format string, args ...interface{}) {
	fmt.Printf(format+"\n", args...)
}
func (l VLog) Infof(format string, args ...interface{}) {
	fmt.Printf(format+"\n", args...)
}
func (l VLog) Warnf(format string, args ...interface{}) {
	fmt.Printf(format+"\n", args...)
}
func (l VLog) Errorf(format string, args ...interface{}) {
	fmt.Printf(format+"\n", args...)
}

func TestDeposit(t *testing.T) {
	vLog := VLog{}
	//构造client
	cli := NewClient(vLog, &PayAsiaInitParams{MERCHANT_ID, MERCHANT_TOKEN, ACCESS_KEY, DEPOSIT_URL, WITHDRAW_URL, DEPOSIT_FE_BACK_URL, WITHDRAW_BACK_URL})

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
		Amount:            "600.00",          // 使用string避免精度问题，对应Java的BigDecimal
		CustomerIp:        "123.123.123.123", //
		CustomerFirstName: "John",
		CustomerLastName:  "Doe",
		CustomerPhone:     "0123123123",
		CustomerEmail:     "someone@gmail.com",
	}
}
