package client

import (
	"github.com/imroc/req/v3"
	"iptables/logger"
	"iptables/model"
	"iptables/settings"
)

var (
	client *req.Client
)

func init() {
	client = req.NewClient()
}

func Report() (err error) {
	_type := "iptables"
	if "gost" == settings.GetClient().Type {
		_type = "gost"
	}
	data := model.ReportReq{
		Name:   settings.GetClient().Name,
		ToPort: settings.GetClient().LocalPort,
		Type:   _type,
	}

	data.MakeSign()

	conf := settings.GetClient()
	url := conf.Server + "/8809_report"
	resp, err := client.R().SetBasicAuth("report", settings.GetGlobal().Secret).SetBody(data).Post(url)
	if err != nil {
		return
	}
	//logger.Infof("report status:%d", resp.GetStatusCode())
	logger.Infof("report response:%s", resp.Bytes())
	return
}
