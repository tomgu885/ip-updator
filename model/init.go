package model

import (
	"crypto/sha1"
	"fmt"
	"iptables/logger"
	"time"
)

type NatRule struct {
	RuleNum    int    `json:"rule_num"`
	ListenPort int    `json:"port"`
	ToIp       string `json:"to_ip"`
	ToPort     int    `json:"to_port"`
}

type ReportReq struct {
	Timestamp int64  `json:"timestamp"`
	Sign      string `json:"sign"`
	Name      string `json:"name"`
	ToPort    int    `json:"to_port"`
}

func (r *ReportReq) MakeSign() {
	if 0 == r.Timestamp {
		r.Timestamp = time.Now().Unix()
	}

	r.Sign = r._reportSign()
}

func (r *ReportReq) SignValid() bool {
	_now := time.Now().Unix()
	diff := r.Timestamp - _now
	if diff < 0 {
		diff = -diff
	}
	if diff > 30 {
		logger.Infof("timestamp is too large r.Time:%d , current:%d diff:%d", r.Timestamp, _now, diff)
		return false
	}

	return r.Sign == r._reportSign()
}

func (r ReportReq) _reportSign() string {
	toSign := fmt.Sprintf("%d_%s_%d", r.Timestamp, r.Name, r.ToPort)
	//logger.Infof("_reportSign: %s", toSign)
	h := sha1.New()
	h.Write([]byte(toSign))
	return fmt.Sprintf("%x", h.Sum(nil))
}
