package go_payasia

// ----------pre order-------------------------

// PayAsiaDepositReq
type PayAsiaDepositReq struct {
	MerchantReference string `json:"merchantReference"`
	Currency          string `json:"currency"`
	Amount            string `json:"amount"` // 使用string避免精度问题，对应Java的BigDecimal
	CustomerIp        string `json:"customerIp"`
	CustomerFirstName string `json:"customerFirstName"`
	CustomerLastName  string `json:"customerLastName"`
	CustomerPhone     string `json:"customerPhone"`
	CustomerEmail     string `json:"customerEmail"`
	Network           string `json:"network"`
	NotifyUrl         string `json:"notify_url"` //回调url

	Url  string `json:"url"` //请求url (该参数不参与sign计算)
	Sign string `json:"sign"`
}

//------------------------------------------------------------

type PayAsiaWithdrawReq struct {
	RequestReference     string `json:"requestReference"`
	BeneficiaryName      string `json:"beneficiaryName"`
	BeneficiaryFirstName string `json:"beneficiaryFirstName"`
	BeneficiaryLastName  string `json:"beneficiaryLastName"`
	BankName             string `json:"bankName"`
	BeneficiaryEmail     string `json:"beneficiaryEmail"`
	BeneficiaryPhone     string `json:"beneficiaryPhone"`
	AccountNumber        string `json:"accountNumber"`
	Currency             string `json:"currency"`
	Amount               string `json:"amount"`
	DatafeedUrl          string `json:"datafeed_url"` //回调url

	Url  string `json:"url"` //请求url (该参数不参与sign计算)
	Sign string `json:"sign"`
}

// ---------------callback-----------------------
type PayAsiaDepositBackReq struct {
	Status            string `json:"status"`
	MerchantReference string `json:"merchant_reference"` //商户订单id
	RequestReference  string `json:"request_reference"`  // pspOrderNo
	Currency          string `json:"currency"`
	Amount            string `json:"amount"` // 使用 string 避免精度问题
	Sign              string `json:"sign"`
}

type PayAsiaWithdrawBackReq struct {
	BatchReference   string `json:"batch_reference"`
	RequestReference string `json:"request_reference"`
	Token            string `json:"token"`
	BeneficiaryName  string `json:"beneficiary_name"`
	BankName         string `json:"bank_name"`
	AccountNumber    string `json:"account_number"`
	OrderCurrency    string `json:"order_currency"`
	OrderAmount      string `json:"order_amount"` // 使用字符串保持精度
	Status           string `json:"status"`
	FailReason       string `json:"fail_reason"`
	CreatedTime      int64  `json:"created_time"`   // Unix 时间戳（毫秒）
	CompletedTime    int64  `json:"completed_time"` // Unix 时间戳（毫秒）
	Sign             string `json:"sign"`
}
