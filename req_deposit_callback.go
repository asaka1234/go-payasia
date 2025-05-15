package go_payasia

// 充值的回调处理(传入一个处理函数)
func (cli *Client) DepositCallback(req PayAsiaDepositBackReq, processor func(PayAsiaDepositBackReq) error) error {
	//验证签名
	//TODO

	//开始处理
	return processor(req)
}
