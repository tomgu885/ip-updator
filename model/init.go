package model

type NatRule struct {
	RuleNum    int    `json:"rule_num"`
	ListenPort int    `json:"port"`
	ToIp       string `json:"to_ip"`
	ToPort     int    `json:"to_port"`
}

type ReportReq struct {
	Secret string `json:"secret"`
	Name   string `json:"name"`
	ToPort int    `json:"to_port"`
}
