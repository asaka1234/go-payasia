package go_payasia

import (
	"github.com/asaka1234/go-payasia/utils"
	"github.com/go-resty/resty/v2"
)

type Client struct {
	Params *PayAsiaInitParams

	ryClient *resty.Client
	logger   utils.Logger
}

func NewClient(logger utils.Logger, params *PayAsiaInitParams) *Client {
	return &Client{
		Params: params,

		ryClient: resty.New(), //client实例
		logger:   logger,
	}
}
