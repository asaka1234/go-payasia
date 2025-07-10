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
		MerchantReference: "20250707082184802657", //商户的orderId
		Currency:          "VND",
		Amount:            "1307900",   // 使用string避免精度问题，对应Java的BigDecimal
		CustomerIp:        "127.0.0.1", //
		//CustomerFirstName: "John",
		//CustomerLastName:  "Doe",
		CustomerPhone: "19831000682",
		CustomerEmail: "jane.y@logtec.com",
	}
}
