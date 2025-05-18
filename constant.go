package go_payasia

type PayAsiaStatus struct {
	Code string
	Name string
	Desc string
}

func (p PayAsiaStatus) GetName() string {
	return p.Name
}

func (p PayAsiaStatus) GetValue() string {
	return p.Code
}

func (p PayAsiaStatus) Eq(value string) bool {
	return p.Code == value
}

var (
	PayAsiaStatusPending    = PayAsiaStatus{"0", "pending", "pending"}
	PayAsiaStatusSuccess    = PayAsiaStatus{"1", "success", "success"}
	PayAsiaStatusFail       = PayAsiaStatus{"2", "fail", "fail"}
	PayAsiaStatusAuthorized = PayAsiaStatus{"3", "authorized", "authorized"}
	PayAsiaStatusProcessing = PayAsiaStatus{"4", "processing", "processing"}
	PayAsiaStatusCancelled  = PayAsiaStatus{"8", "cancelled", "cancelled"}
)

// GetByCode returns the PayAsiaStatus by code
func GetByCode(code string) (PayAsiaStatus, bool) {
	switch code {
	case PayAsiaStatusPending.Code:
		return PayAsiaStatusPending, true
	case PayAsiaStatusSuccess.Code:
		return PayAsiaStatusSuccess, true
	case PayAsiaStatusFail.Code:
		return PayAsiaStatusFail, true
	case PayAsiaStatusAuthorized.Code:
		return PayAsiaStatusAuthorized, true
	case PayAsiaStatusProcessing.Code:
		return PayAsiaStatusProcessing, true
	case PayAsiaStatusCancelled.Code:
		return PayAsiaStatusCancelled, true
	default:
		return PayAsiaStatus{}, false
	}
}
