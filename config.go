package go_payasia

const (
	MERCHANT_ID    = "2021-CPTMARKETS"                      //商户id
	MERCHANT_TOKEN = "496a6764-f46a-4ac8-81ca-cc8b72aa9fe0" //token

	SECRECT_CODE = "b47149cc-5ff2-4178-ab2f-46ec1f2573fc"

	//--------

	DEPOSIT_URL  = "https://payment.pa-sys.com/app/page/496a6764-f46a-4ac8-81ca-cc8b72aa9fe0"          //充值url
	WITHDRAW_URL = "https://gateway.pa-sys.com/496a6764-f46a-4ac8-81ca-cc8b72aa9fe0/payout/v1/request" // 提现url

	DEPOSIT_CALLBACK_URL  = "http://127.0.0.1/order/post" //充值回调url
	WITHDRAW_CALLBACK_URL = "http://127.0.0.1/order/post" //提现回调url
)
