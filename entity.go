package go_payasia

type PayAsiaInitParams struct {
	MerchantId       string `json:"merchantId" mapstructure:"merchantId" config:"merchantId"  yaml:"merchantId"`             //商户号
	MerchantToken    string `json:"merchantToken" mapstructure:"merchantToken" config:"merchantToken"  yaml:"merchantToken"` //applicationId
	AccessKey        string `json:"accessKey" mapstructure:"accessKey" config:"accessKey"  yaml:"accessKey"`
	DepositUrl       string `json:"depositUrl" mapstructure:"depositUrl" config:"depositUrl"  yaml:"depositUrl"`
	WithdrawUrl      string `json:"withdrawUrl" mapstructure:"withdrawUrl" config:"withdrawUrl"  yaml:"withdrawUrl"`
	DepositFeBackUrl string `json:"depositFeBackUrl" mapstructure:"depositFeBackUrl" config:"depositFeBackUrl"  yaml:"depositFeBackUrl"` //充值的前端回跳地址
	WithdrawBackUrl  string `json:"withdrawBackUrl" mapstructure:"withdrawBackUrl" config:"withdrawBackUrl"  yaml:"withdrawBackUrl"`     //提现回调
}

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
	//ReturnUrl          string `json:"return_url "` //前端回跳地址
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

//amount=311760.00&currency=VND&merchant_reference=202507061318100463&request_reference=eac9f860-9f2b-445f-a332-bc66e27b45cc&status=2&sign=5d9a711dcc98b5e7b34560019a09b88637ba733556ba702af7f748ee9162297c98a622b1685d9c5392e00c1e2b8253d5fb65228dc30703942aa312b0d3f358f2
//amount=311760.00&currency=VND&merchant_reference=202507061318100463&request_reference=eac9f860-9f2b-445f-a332-bc66e27b45cc&status=2&b47149cc-5ff2-4178-ab2f-46ec1f2573fc

type PayAsiaDepositBackReq struct {
	Status            string `form:"status" json:"status" mapstructure:"status"`
	MerchantReference string `form:"merchant_reference" json:"merchant_reference" mapstructure:"merchant_reference"` //商户订单id
	RequestReference  string `form:"request_reference" json:"request_reference" mapstructure:"request_reference"`    // pspOrderNo
	Currency          string `form:"currency" json:"currency" mapstructure:"currency"`
	Amount            string `form:"amount" json:"amount" mapstructure:"amount"` // 使用 string 避免精度问题
	Sign              string `form:"sign" json:"sign" mapstructure:"sign"`       //对返回数据的签名
}

type PayAsiaWithdrawBackReq struct {
	Status           string `form:"status" json:"status" mapstructure:"status"`
	BatchReference   string `form:"batch_reference" json:"batch_reference" mapstructure:"batch_reference"`
	RequestReference string `form:"request_reference" json:"request_reference" mapstructure:"request_reference"`
	Token            string `form:"token" json:"token" mapstructure:"token"`
	BeneficiaryName  string `form:"beneficiary_name" json:"beneficiary_name" mapstructure:"beneficiary_name"`
	BankName         string `form:"bank_name" json:"bank_name" mapstructure:"bank_name"`
	AccountNumber    string `form:"account_number" json:"account_number" mapstructure:"account_number"`
	OrderCurrency    string `form:"order_currency" json:"order_currency" mapstructure:"order_currency"`
	OrderAmount      string `form:"order_amount" json:"order_amount" mapstructure:"order_amount"` // 使用字符串保持精度
	FailReason       string `form:"fail_reason" json:"fail_reason" mapstructure:"fail_reason"`
	CreatedTime      int64  `form:"created_time" json:"created_time" mapstructure:"created_time"`       // Unix 时间戳（毫秒）
	CompletedTime    int64  `form:"completed_time" json:"completed_time" mapstructure:"completed_time"` // Unix 时间戳（毫秒）
	Sign             string `form:"sign" json:"sign" mapstructure:"sign"`
}
