package go_payasia

import (
	"github.com/asaka1234/go-payasia/utils"
	"github.com/go-resty/resty/v2"
)

type Client struct {
	Merchant  string
	AccessKey string

	DepositUrl          string
	WithdrawUrl         string
	DepositCallbackUrl  string
	WithdrawCallbackUrl string

	ryClient *resty.Client
	logger   utils.Logger
}

func NewClient(logger utils.Logger, merchant string, accessKey string, depositUrl, withdrawUrl, depositCallbackUrl, withdrawCallbackUrl string) *Client {
	return &Client{
		Merchant:  merchant,
		AccessKey: accessKey,

		DepositUrl:          depositUrl,
		WithdrawUrl:         withdrawUrl,
		DepositCallbackUrl:  depositCallbackUrl,
		WithdrawCallbackUrl: withdrawCallbackUrl,

		ryClient: resty.New(), //client实例
		logger:   logger,
	}
}
