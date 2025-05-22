package go_payasia

// ----------pre order-------------------------

// PayAsiaDepositReq
type PayAsiaDepositReq struct {
	MerchantReference string  `json:"merchant_reference" mapstructure:"merchant_reference"` //商户的订单id
	Currency          string  `json:"currency" mapstructure:"currency"`
	Amount            float64 `json:"amount" mapstructure:"amount"` // 使用string避免精度问题，对应Java的BigDecimal
	CustomerIp        string  `json:"customer_ip" mapstructure:"customer_ip"`
	CustomerFirstName string  `json:"customer_first_name" mapstructure:"customer_first_name"`
	CustomerLastName  string  `json:"customer_last_name" mapstructure:"customer_last_name"`
	CustomerPhone     string  `json:"customer_phone" mapstructure:"customer_phone"`
	CustomerEmail     string  `json:"customer_email" mapstructure:"customer_email"`

	//这四个字段都不用客户传了. 前2个sdk直接补充设置. 最后一个sdk计算后补充
	//Network string `json:"network"` //值是写死的 DirectDebit
	//NotifyUrl         string `json:"notify_url"` //callback url
	//Url  string `json:"url"` //请求url,是form表单action提交时的具体url (该参数不参与sign计算)
	//Sign string `json:"sign"`
}

//------------------------------------------------------------

type PayAsiaWithdrawReq struct {
	RequestReference     string  `json:"request_reference" mapstructure:"request_reference"`
	BeneficiaryName      string  `json:"beneficiary_name" mapstructure:"beneficiary_name"`
	BeneficiaryFirstName string  `json:"beneficiary_first_name" mapstructure:"beneficiary_first_name"`
	BeneficiaryLastName  string  `json:"beneficiary_last_name" mapstructure:"beneficiary_last_name"`
	BankName             string  `json:"bank_name" mapstructure:"bank_name"`
	BeneficiaryEmail     string  `json:"beneficiary_email" mapstructure:"beneficiary_email"`
	BeneficiaryPhone     string  `json:"beneficiary_phone" mapstructure:"beneficiary_phone"`
	AccountNumber        string  `json:"account_number" mapstructure:"account_number"`
	Currency             string  `json:"currency" mapstructure:"currency"`
	Amount               float64 `json:"amount" mapstructure:"amount"`

	//这三个字段都不用客户传了. 前2个sdk直接补充设置. 最后一个sdk计算后补充
	//DatafeedUrl          string `json:"datafeed_url"` //回调url
	//Url  string `json:"url"` //请求url,是form表单action提交时的具体url (该参数不参与sign计算)
	//Sign string `json:"sign"`
}

// ---------------callback-----------------------
// Method : POST
// Content-Type : application/x-www-form-urlencoded
type PayAsiaDepositBackReq struct {
	Status            string  `json:"status" mapstructure:"status"`
	MerchantReference string  `json:"merchant_reference" mapstructure:"merchant_reference"` //商户订单id
	RequestReference  string  `json:"request_reference" mapstructure:"request_reference"`   // pspOrderNo
	Currency          string  `json:"currency" mapstructure:"currency"`
	Amount            float64 `json:"amount" mapstructure:"amount"` // 使用 string 避免精度问题
	Sign              string  `json:"sign" mapstructure:"sign"`     //对返回数据的签名
}

type PayAsiaWithdrawBackReq struct {
	Status           string `json:"status" mapstructure:"status"`
	BatchReference   string `json:"batch_reference" mapstructure:"batch_reference"`
	RequestReference string `json:"request_reference" mapstructure:"request_reference"`
	Token            string `json:"token" mapstructure:"token"`
	BeneficiaryName  string `json:"beneficiary_name" mapstructure:"beneficiary_name"`
	BankName         string `json:"bank_name" mapstructure:"bank_name"`
	AccountNumber    string `json:"account_number" mapstructure:"account_number"`
	OrderCurrency    string `json:"order_currency" mapstructure:"order_currency"`
	OrderAmount      string `json:"order_amount" mapstructure:"order_amount"` // 使用字符串保持精度
	FailReason       string `json:"fail_reason" mapstructure:"fail_reason"`
	CreatedTime      int64  `json:"created_time" mapstructure:"created_time"`     // Unix 时间戳（毫秒）
	CompletedTime    int64  `json:"completed_time" mapstructure:"completed_time"` // Unix 时间戳（毫秒）
	Sign             string `json:"sign" mapstructure:"sign"`
}
