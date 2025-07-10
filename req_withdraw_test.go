package go_payasia

import (
	"fmt"
	"testing"
)

func TestWithdraw(t *testing.T) {
	vLog := VLog{}
	//构造client
	cli := NewClient(vLog, &PayAsiaInitParams{MERCHANT_ID, MERCHANT_TOKEN, ACCESS_KEY, DEPOSIT_URL, WITHDRAW_URL, DEPOSIT_FE_BACK_URL, WITHDRAW_BACK_URL})

	//发请求
	resp, err := cli.Withdraw(GenPayAsiaWithdrawReqDemo())
	if err != nil {
		fmt.Printf("err:%s\n", err.Error())
		return
	}
	fmt.Printf("resp:%+v\n", resp)
}

func GenPayAsiaWithdrawReqDemo() PayAsiaWithdrawReq {
	return PayAsiaWithdrawReq{
		RequestReference:     "20215071848202615711", //商户的orderId
		BeneficiaryName:      "john wong",
		BeneficiaryFirstName: "john",
		BeneficiaryLastName:  "wong",
		BankName:             "ACB",
		BeneficiaryEmail:     "demo@gmail.com",
		BeneficiaryPhone:     "17927297",
		AccountNumber:        "7207972",
		Currency:             "VND",
		Amount:               "1307900",
	}
}
