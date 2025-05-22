package utils

import (
	"fmt"
	"testing"
)

// 计算请求签名
func TestSign(t *testing.T) {
	accessKey := "127f7830-b856-4ddf-92b4-a6478e38547b"

	demoMap := map[string]interface{}{
		"merchant_reference":  "1234567890",
		"currency":            "THB",
		"amount":              "600.00", //TODO
		"customer_ip":         "123.123.123.123",
		"customer_first_name": "John",
		"customer_last_name":  "Doe",
		"customer_address":    "1, Bay Street",
		"customer_phone":      "0123123123",
		"customer_email":      "someone@gmail.com",
		"return_url":          "https://demo.shop.com/payment/return",
		"network":             "DirectDebit",
	}

	result := Sign(demoMap, accessKey)
	fmt.Printf("sign: %s\n", result)
}

/*
amount=600.00&currency=THB&customer_address=1%2C+Bay+Street&customer_email=s
omeone%40gmail.com&customer_first_name=John&customer_ip=123.123.123.123&custo
mer_last_name=Doe&customer_phone=0123123123&merchant_reference=1234567890&net
work=DirectDebit&return_url=https%3A%2F%2Fdemo.shop.com%2Fpayment%2Freturn127
f7830-b856-4ddf-92b4-a6478e38547b

// sample hash from example above
e58e599473393fdd38a9fa6fed2977ccd2dbb443a8116be448d0128e77a888542a8d8fde649a0
40339b7602d028d2e4883652082af068573ddf673e417c4eb67
*/
