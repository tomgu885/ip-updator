package server

import (
	"github.com/gin-gonic/gin"
	"iptables/model"
	"net/http"
)

func RunServer() {
	r := gin.Default()
	r.POST("/8809_report", report)
	r.GET("/8809_status", status)
	r.Run(":8099")
}

//
func report(c *gin.Context) {
	var req model.ReportReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.String(http.StatusOK, "绑定错误:"+err.Error())
		return
	}

	clientIp := c.ClientIP()
	rules, err := ListNatWithNumber()
	if err != nil {
		c.String(http.StatusOK, "获取列表失败"+err.Error())
		return
	}

	for _, rule := range rules {
		if rule.ToPort != req.ToPort {
			continue
		}

		if rule.ToPort == req.ToPort {
			if rule.ToIp != clientIp {
				DelNatByLineNumber(rule.RuleNum)
				AddNatRule(model.NatRule{
					RuleNum:    0,
					ListenPort: req.ToPort,
					ToIp:       clientIp,
					ToPort:     req.ToPort,
				})
			}
			break
		}
	}
}

func status(c *gin.Context) {
	rules, err := ListNatWithNumber()
	if err != nil {
		c.String(http.StatusOK, "failed:"+err.Error())
		return
	}

	c.JSON(http.StatusOK, rules)
}
