package go_payasia

import (
	"github.com/asaka1234/go-payasia/utils"
	"github.com/go-resty/resty/v2"
)

type Client struct {
	MerchantID    string
	MerchantToken string //更多类似applicationId
	AccessKey     string //用来计算签名的

	DepositUrl          string
	WithdrawUrl         string
	DepositCallbackUrl  string
	WithdrawCallbackUrl string

	ryClient *resty.Client
	logger   utils.Logger
}

func NewClient(logger utils.Logger, merchantId string, merchantToken string, accessKey string, depositUrl, withdrawUrl, depositCallbackUrl, withdrawCallbackUrl string) *Client {
	return &Client{
		MerchantID:    merchantId,
		MerchantToken: merchantToken,
		AccessKey:     accessKey,

		DepositUrl:          depositUrl,
		WithdrawUrl:         withdrawUrl,
		DepositCallbackUrl:  depositCallbackUrl,
		WithdrawCallbackUrl: withdrawCallbackUrl,

		ryClient: resty.New(), //client实例
		logger:   logger,
	}
}
