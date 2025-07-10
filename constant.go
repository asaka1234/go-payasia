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

//-------------------------------------------------

type PayAsiaBankCode struct {
	Currency string `json:"currency"`
	Code     string `json:"code"`
	Name     string `json:"name"`
}

var PayAsiaBankCodes = []PayAsiaBankCode{
	{Currency: "PHP", Code: "ALB", Name: "AllBank"},
	{Currency: "PHP", Code: "AUB", Name: "Asia United Bank CA/SA"},
	{Currency: "PHP", Code: "BDO", Name: "Banco de Oro CA/SA"},
	{Currency: "PHP", Code: "BDON", Name: "BDO Network Bank - instapay"},
	{Currency: "PHP", Code: "BFB", Name: "BPI Family Bank"},
	{Currency: "PHP", Code: "BITC", Name: "COINS.PH"},
	{Currency: "PHP", Code: "BOC", Name: "Bank of Commerce"},
	{Currency: "PHP", Code: "BPI", Name: "BPI CA/SA"},
	{Currency: "PHP", Code: "CBC", Name: "Chinabank CA/SA"},
	{Currency: "PHP", Code: "CBCS", Name: "Chinabank Savings"},
	{Currency: "PHP", Code: "CIMB", Name: "CIMB Bank"},
	{Currency: "PHP", Code: "CITI", Name: "Citibank - instapay"},
	{Currency: "PHP", Code: "CTBC", Name: "Chinatrust"},
	{Currency: "PHP", Code: "DBP", Name: "Development Bank of the Philippines"},
	{Currency: "PHP", Code: "EWB", Name: "EastWest CA/SA"},
	{Currency: "PHP", Code: "GCSH", Name: "GCash"},
	{Currency: "PHP", Code: "GRPY", Name: "GrabPay Wallet - instapay"},
	{Currency: "PHP", Code: "HSBC", Name: "HSBC"},
	{Currency: "PHP", Code: "LBP", Name: "Landbank CA/SA"},
	{Currency: "PHP", Code: "LULU", Name: "Lulu Money"},
	{Currency: "PHP", Code: "MAY", Name: "Maybank"},
	{Currency: "PHP", Code: "MBTC", Name: "Metrobank CA/SA"},
	{Currency: "PHP", Code: "NTB", Name: "Netbank"},
	{Currency: "PHP", Code: "PBCM", Name: "PBCom"},
	{Currency: "PHP", Code: "PLWP", Name: "PalawanPay"},
	{Currency: "PHP", Code: "PNB", Name: "PNB individual CA/SA"},
	{Currency: "PHP", Code: "PSB", Name: "PSBank"},
	{Currency: "PHP", Code: "PVB", Name: "Veterans Bank"},
	{Currency: "PHP", Code: "PYMY", Name: "PayMaya"},
	{Currency: "PHP", Code: "RCBC", Name: "RCBC CA/SA, RCBC Savings Bank CA/SA, RCBC MyWallet"},
	{Currency: "PHP", Code: "RSB", Name: "Robinsons Bank CA/SA"},
	{Currency: "PHP", Code: "SBA", Name: "Sterling Bank"},
	{Currency: "PHP", Code: "SBC", Name: "Security Bank CA/SA"},
	{Currency: "PHP", Code: "SEAB", Name: "Seabank"},
	{Currency: "PHP", Code: "TNIK", Name: "Tonik Bank"},
	{Currency: "PHP", Code: "TYME", Name: "GoTyme"},
	{Currency: "PHP", Code: "UBP", Name: "Unionbank CA/SA, EON"},

	//-----VND-------

	{Currency: "VND", Code: "ABBank", Name: "An Binh Commercial Joint Stock Bank (ABBank)"},
	{Currency: "VND", Code: "ACB", Name: "Asia Commercial Bank (ACB)"},
	{Currency: "VND", Code: "Agribank", Name: "Vietnam Bank For Agriculture and Rural Development (Agribank)"},
	{Currency: "VND", Code: "ANZ", Name: "Australia and New Zealand Banking (ANZ Bank)"},
	{Currency: "VND", Code: "BacABank", Name: "North Asia Commercial Joint Stock Bank (NASB)"},
	{Currency: "VND", Code: "BaoVietBank", Name: "Baoviet Joint Stock Commercial Bank"},
	{Currency: "VND", Code: "BIDV", Name: "Bank for Investment & Dof Vietnam (BIDV)"},
	{Currency: "VND", Code: "CBB", Name: "Vietnam Construction Joint Stock Commercial Bank (VNCB)"},
	{Currency: "VND", Code: "Citibank", Name: "CITIBANK N.A."},
	{Currency: "VND", Code: "DongABank", Name: "DongA Bank"},
	{Currency: "VND", Code: "Eximbank", Name: "Vietnam Export Import Commercial Joint Stock Bank (Eximbank)"},
	{Currency: "VND", Code: "GBBank", Name: "Global Petro Bank (GBBank)"},
	{Currency: "VND", Code: "HDBank", Name: "HoChiMinh City Development Joint Stock Commercial Bank (HDBank)"},
	{Currency: "VND", Code: "HLBVN", Name: "Hong Leong Bank Vietnam"},
	{Currency: "VND", Code: "HSBC", Name: "HSBC Bank (Vietnam) Ltd"},
	{Currency: "VND", Code: "INDOVINA", Name: "Indovina Bank Ltd"},
	{Currency: "VND", Code: "Kienlongbank", Name: "Kien Long Commercial Joint Stock Bank (Kienlongbank)"},
	{Currency: "VND", Code: "LienVietPostBank", Name: "LienVietPostBank"},
	{Currency: "VND", Code: "MBBank", Name: "MB Bank"},
	{Currency: "VND", Code: "MHB", Name: "Mekong Housing Bank (MHB Bank)"},
	{Currency: "VND", Code: "MSB", Name: "Vietnam Maritime Commercial Joint Stock Bank (MaritimeBank)"},
	{Currency: "VND", Code: "NamABank", Name: "Nam A Commercial Joint Stock Bank"},
	{Currency: "VND", Code: "NCB", Name: "National Citizen Bank"},
	{Currency: "VND", Code: "OCB", Name: "Orient Commercial Joint Stock Bank (OCB)"},
	{Currency: "VND", Code: "Oceanbank", Name: "OceanBank"},
	{Currency: "VND", Code: "PBVN", Name: "Public Bank VN"},
	{Currency: "VND", Code: "PGBank", Name: "Petrolimex Group Commercial Joint Stock Bank (PG Bank)"},
	{Currency: "VND", Code: "Pvcombank", Name: "Vietnam Public Joint Stock Commercial Bank (Pvcombank)"},
	{Currency: "VND", Code: "Sacombank", Name: "Saigon Thuong Tin Commercial Joint Stock Bank (Sacombank)"},
	{Currency: "VND", Code: "SCB", Name: "Saigon Commercial Bank (SCB)"},
	{Currency: "VND", Code: "SeABank", Name: "Southeast Asia Commercial Joint Stock Bank (SeABank)"},
	{Currency: "VND", Code: "SGB", Name: "Saigon Bank"},
	{Currency: "VND", Code: "SHB", Name: "Saigon – Hanoi Commercial Joint Stock Bank (SHB)"},
	{Currency: "VND", Code: "Shinhan", Name: "SHINHAN Bank"},
	{Currency: "VND", Code: "Standard_Chartered", Name: "Standard Chartered Bank"},
	{Currency: "VND", Code: "TCB", Name: "Techcombank"},
	{Currency: "VND", Code: "TPBank", Name: "Tien Phong Commercial Joint Stock Bank (TP Bank)"},
	{Currency: "VND", Code: "VCB", Name: "VietcomBank"},
	{Currency: "VND", Code: "VIB", Name: "Vietnam International Commercial Joint Stock Bank (VIB)"},
	{Currency: "VND", Code: "VietABank", Name: "Vietnam Asia Commercial Joint Stock Bank (VietABank)"},
	{Currency: "VND", Code: "Vietbank", Name: "Vietnam Thuong Tin Commercial Joint Stock Bank"},
	{Currency: "VND", Code: "VietCapitalBank", Name: "Viet Capital Commercial Joint Stock Bank"},
	{Currency: "VND", Code: "VietinBank", Name: "Vietnam Bank for Industry and Trade (VietinBank)"},
	{Currency: "VND", Code: "VPBank", Name: "Vietnam Prosperity Bank (VPBank)"},
	{Currency: "VND", Code: "WooriBank", Name: "Woori Bank Vietnam"},

	//---------------IDR------------------------------------
	{Currency: "IDR", Code: "ABN", Name: "Bank ABN Amro"},
	{Currency: "IDR", Code: "ACEH", Name: "Bank Aceh Syariah"},
	{Currency: "IDR", Code: "AE", Name: "American Express Bank LTD"},
	{Currency: "IDR", Code: "AGRIS", Name: "Bank IBK Indonesia / Agris"},
	{Currency: "IDR", Code: "AGRONIAGA", Name: "BRI Agroniaga"},
	{Currency: "IDR", Code: "AKITA", Name: "Bank Akita"},
	{Currency: "IDR", Code: "ALADIN", Name: "Bank Aladin Syariah"},
	{Currency: "IDR", Code: "AMAR", Name: "Bank Amar Indonesia / Anglomas"},
	{Currency: "IDR", Code: "AMERICA_NA", Name: "Bank of America NA"},
	{Currency: "IDR", Code: "ANK", Name: "Bank Arta Niaga Kencana"},
	{Currency: "IDR", Code: "ANTARDAERAH", Name: "Bank Antardaerah"},
	{Currency: "IDR", Code: "ANZ", Name: "ANZ Indonesia"},
	{Currency: "IDR", Code: "ARTAJASA", Name: "Artajasa Pembayaran Elektronik"},
	{Currency: "IDR", Code: "ARTHA", Name: "Bank Artha Graha Internasional"},
	{Currency: "IDR", Code: "ARTOS", Name: "Bank Jago / Artos"},
	{Currency: "IDR", Code: "BALI", Name: "BPD Bali"},
	{Currency: "IDR", Code: "BANGKOK", Name: "Bangkok Bank"},
	{Currency: "IDR", Code: "BANTEN", Name: "BPD Banten"},
	{Currency: "IDR", Code: "BCA", Name: "Bank Central Asia"},
	{Currency: "IDR", Code: "BCA_SYR", Name: "BCA (Bank Central Asia) Syariah"},
	{Currency: "IDR", Code: "BENGKULU", Name: "Bank Bengkulu"},
	{Currency: "IDR", Code: "BISNIS", Name: "Bisnis Internasional"},
	{Currency: "IDR", Code: "BJB", Name: "BJB"},
	{Currency: "IDR", Code: "BJB_SYR", Name: "BJB Syariah"},
	{Currency: "IDR", Code: "BMAN", Name: "Bank Mandiri"},
	{Currency: "IDR", Code: "BNI", Name: "BNI (Bank Negara Indonesia)"},
	{Currency: "IDR", Code: "BNP", Name: "BNP Paribas Indonesia"},
	{Currency: "IDR", Code: "BOC", Name: "Bank of China (Hong Kong) Limited"},
	{Currency: "IDR", Code: "BP", Name: "Bank Permata & Permata Syariah"},
	{Currency: "IDR", Code: "BPD", Name: "Bank BPD DIY Syariah"},
	{Currency: "IDR", Code: "BPR_KS", Name: "BPR KS"},
	{Currency: "IDR", Code: "BRI", Name: "Bank Rakyat Indonesia"},
	{Currency: "IDR", Code: "BSM", Name: "BSI (Bank Syariah Indonesia)"},
	{Currency: "IDR", Code: "BTN", Name: "BTN"},
	{Currency: "IDR", Code: "BTPN_SYR", Name: "Bank BTPN Syariah"},
	{Currency: "IDR", Code: "BUKOPIN", Name: "Bank Bukopin"},
	{Currency: "IDR", Code: "BUKOPIN_SYR", Name: "Bank Bukopin Syariah"},
	{Currency: "IDR", Code: "BUMI_ARTA", Name: "Bank Bumi Arta"},
	{Currency: "IDR", Code: "BWS", Name: "Bank Woori Saudara (HS 1906)"},
	{Currency: "IDR", Code: "CAI", Name: "Bank Credit Agricole Indosuez"},
	{Currency: "IDR", Code: "CAPITAL", Name: "Bank Capital Indonesia"},
	{Currency: "IDR", Code: "CCB", Name: "Bank China Construction Bank Indonesia"},
	{Currency: "IDR", Code: "CHINATRUST", Name: "CTBC (Chinatrust) Indonesia"},
	{Currency: "IDR", Code: "CIMB", Name: "CIMB Niaga Syariah"},
	{Currency: "IDR", Code: "CITIBANK", Name: "Citibank"},
	{Currency: "IDR", Code: "CNB", Name: "Bank CNB (Centratama Nasional Bank)"},
	{Currency: "IDR", Code: "COMMONWEALTH", Name: "Commonwealth Bank"},
	{Currency: "IDR", Code: "DANA", Name: "Dana"},
	{Currency: "IDR", Code: "DANAMON", Name: "Bank Danamon"},
	{Currency: "IDR", Code: "DANAMON_SYR", Name: "Bank Danamon Syariah"},
	{Currency: "IDR", Code: "DBS", Name: "DBS Indonesia"},
	{Currency: "IDR", Code: "DEUTSCHE", Name: "Deutsche Bank AG."},
	{Currency: "IDR", Code: "DKI", Name: "Bank DKI"},
	{Currency: "IDR", Code: "DKI_SYR", Name: "Bank DKI Syariah"},
	{Currency: "IDR", Code: "EKA", Name: "BPR EKA (Bank Eka)"},
	{Currency: "IDR", Code: "EKSPOR", Name: "Bank Ekspor Indonesia"},
	{Currency: "IDR", Code: "FAMA", Name: "Bank Fama Internasional"},
	{Currency: "IDR", Code: "GANESHA", Name: "Bank Ganesha"},
	{Currency: "IDR", Code: "GOPAY", Name: "Gopay"},
	{Currency: "IDR", Code: "HANA", Name: "LINE Bank/KEB Hana"},
	{Currency: "IDR", Code: "HARDA", Name: "Allo Bank/Bank Harda Internasional"},
	{Currency: "IDR", Code: "HARMONI", Name: "Bank Harmoni International"},
	{Currency: "IDR", Code: "HSBC", Name: "HSBC Indonesia"},
	{Currency: "IDR", Code: "ICBC", Name: "ICBC Indonesia"},
	{Currency: "IDR", Code: "INA_PERDANA", Name: "Bank Ina Perdana"},
	{Currency: "IDR", Code: "INDEX_SELINDO", Name: "Bank Index Selindo"},
	{Currency: "IDR", Code: "ING", Name: "ING Indonesia Bank"},
	{Currency: "IDR", Code: "JAMBI", Name: "Bank Jambi"},
	{Currency: "IDR", Code: "JASA_JAKARTA", Name: "Bank Jasa Jakarta"},
	{Currency: "IDR", Code: "JATENG_SYR", Name: "Bank Jateng Syariah"},
	{Currency: "IDR", Code: "JAWA_TENGAH", Name: "Bank Jateng"},
	{Currency: "IDR", Code: "JAWA_TIMUR", Name: "Bank Jatim"},
	{Currency: "IDR", Code: "JPMCI", Name: "JP Morgan Chase Indonesia"},
	{Currency: "IDR", Code: "KALIMANTAN_BARAT", Name: "Bank Kalbar"},
	{Currency: "IDR", Code: "KALIMANTAN_SELATAN", Name: "Bank Kalsel"},
	{Currency: "IDR", Code: "KALIMANTAN_TENGAH", Name: "Bank Kalteng"},
	{Currency: "IDR", Code: "KALTIM", Name: "Bank Kaltim Syariah"},
	{Currency: "IDR", Code: "KEB_DANAMON", Name: "Korea Exchange Bank Danamon"},
	{Currency: "IDR", Code: "KTB", Name: "Bank Keppel Tatlee Buana"},
	{Currency: "IDR", Code: "LAMPUNG", Name: "Bank Lampung"},
	{Currency: "IDR", Code: "LINKAJA", Name: "LinkAja"},
	{Currency: "IDR", Code: "MALUKU", Name: "Bank Maluku"},
	{Currency: "IDR", Code: "MANTAP", Name: "Bank MANTAP (Mandiri Taspen)"},
	{Currency: "IDR", Code: "MAS", Name: "Bank Multi Arta Sentosa / Bank Mas"},
	{Currency: "IDR", Code: "MASPION", Name: "Bank Maspion Indonesia"},
	{Currency: "IDR", Code: "MAYAPADA", Name: "Bank Mayapada"},
	{Currency: "IDR", Code: "MAYBANK", Name: "Maybank Indonesia"},
	{Currency: "IDR", Code: "MAYORA", Name: "Bank Mayora Indonesia"},
	{Currency: "IDR", Code: "MEGA", Name: "Bank Mega"},
	{Currency: "IDR", Code: "MEGA_SYR", Name: "Bank Mega Syariah"},
	{Currency: "IDR", Code: "MESTIKA_DHARMA", Name: "Bank Mestika Dharma"},
	{Currency: "IDR", Code: "MIZUHO", Name: "Bank Mizuho Indonesia"},
	{Currency: "IDR", Code: "MNC_INTERNASIONAL", Name: "Motion/MNC Bank"},
	{Currency: "IDR", Code: "MUAMALAT", Name: "Muamalat"},
	{Currency: "IDR", Code: "MUTIARA", Name: "Bank Mutiara (Jtrust)"},
	{Currency: "IDR", Code: "NATIONALNOBU", Name: "Nobu (Nationalnobu) Bank"},
	{Currency: "IDR", Code: "NUSA_TENGGARA_BARAT", Name: "Bank NTB Syariah"},
	{Currency: "IDR", Code: "NUSA_TENGGARA_TIMUR", Name: "Bank NTT"},
	{Currency: "IDR", Code: "NUSANTARA_PARAHYANGAN", Name: "Bank Nusantara Parahyangan"},
	{Currency: "IDR", Code: "OCBC", Name: "Bank OCBC NISP"},
	{Currency: "IDR", Code: "OKE", Name: "Bank Dinar Indonesia / Oke Bank"},
	{Currency: "IDR", Code: "OVO", Name: "OVO"},
	{Currency: "IDR", Code: "PANIN", Name: "Panin Bank"},
	{Currency: "IDR", Code: "PANIN_SYR", Name: "Panin Dubai Syariah"},
	{Currency: "IDR", Code: "PAPUA", Name: "Bank Papua"},
	{Currency: "IDR", Code: "PERMATA", Name: "Permata Syariah"},
	{Currency: "IDR", Code: "PRIMA", Name: "Bank Prima Master"},
	{Currency: "IDR", Code: "QNB_KESAWAN", Name: "QNB Indonesia"},
	{Currency: "IDR", Code: "RABOBANK", Name: "Rabobank International Indonesia"},
	{Currency: "IDR", Code: "RESONA", Name: "Bank Resona Perdania"},
	{Currency: "IDR", Code: "RIAU_DAN_KEPRI", Name: "Bank Riau Kepri"},
	{Currency: "IDR", Code: "ROYAL", Name: "Blu/BCA Digital"},
	{Currency: "IDR", Code: "SAHABAT_SAMPOERNA", Name: "Bank Sahabat Sampoerna"},
	{Currency: "IDR", Code: "SBI_INDONESIA", Name: "SBI Indonesia"},
	{Currency: "IDR", Code: "SEABANK", Name: "Seabank/Bank Kesejahteraan Ekonomi"},
	{Currency: "IDR", Code: "SHINHAN", Name: "Bank Shinhan Indonesia"},
	{Currency: "IDR", Code: "SHOPEEPAY", Name: "Shopeepay"},
	{Currency: "IDR", Code: "SINARMAS", Name: "Bank Sinarmas"},
	{Currency: "IDR", Code: "SINARMAS_SYR", Name: "Bank Sinarmas Syariah"},
	{Currency: "IDR", Code: "SMBC", Name: "Bank Sumitomo Mitsui Indonesia"},
	{Currency: "IDR", Code: "SP", Name: "Bank Sri Partha"},
	{Currency: "IDR", Code: "STANDARD_CHARTERED", Name: "Standard Chartered Bank"},
	{Currency: "IDR", Code: "SULAWESI_TENGGARA", Name: "Bank Sultra"},
	{Currency: "IDR", Code: "SULSELBAR", Name: "Bank Sulselbar"},
	{Currency: "IDR", Code: "SULTENG", Name: "Bank Sulteng (Sulawesi Tengah)"},
	{Currency: "IDR", Code: "SULUT", Name: "Bank SulutGo"},
	{Currency: "IDR", Code: "SUMBER_SYR", Name: "Bank Sumbar/Nagari Syariah"},
	{Currency: "IDR", Code: "SUMSEL_DAN_BABEL", Name: "Bank Sumsel Babel"},
	{Currency: "IDR", Code: "SUMUT", Name: "Bank Sumut"},
	{Currency: "IDR", Code: "SWADESI", Name: "Bank of India Indonesia / Swadesi"},
	{Currency: "IDR", Code: "TABUNGAN_PENSIUNAN_NASIONAL", Name: "BTPN"},
	{Currency: "IDR", Code: "TOKYO", Name: "Bank of Tokyo Mitsubishi UFJ"},
	{Currency: "IDR", Code: "UOB", Name: "TMRW/UOB"},
	{Currency: "IDR", Code: "VICTORIA_INTERNASIONAL", Name: "Bank Victoria International"},
	{Currency: "IDR", Code: "VICTORIA_SYR", Name: "Bank Victoria Syariah"},
	{Currency: "IDR", Code: "YUDHA_BAKTI", Name: "Neo Commerce/Yudha Bhakti"},
}
