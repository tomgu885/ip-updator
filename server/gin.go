package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"iptables/logger"
	"iptables/model"
	"iptables/settings"
	"net/http"
	"time"
)

func RunServer() {
	r := gin.Default()

	r.SetTrustedProxies([]string{"127.0.0.1"})

	r.NoRoute(func(c *gin.Context) {
		c.String(http.StatusOK, "veni vidi vici")
	})
	router := r.Use(gin.BasicAuth(map[string]string{
		"report": settings.GetGlobal().Secret,
	}))
	router.POST("/8809_report", report)
	router.GET("/8809_status", status)
	addr := fmt.Sprintf(":%d", settings.GetServer().Port)
	err := r.Run(addr)
	if err != nil {
		panic(err)
	}
}

//
func report(c *gin.Context) {
	var req model.ReportReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.String(http.StatusOK, "绑定错误:"+err.Error())
		return
	}

	clientIp := c.ClientIP()

	if !req.SignValid() {
		logger.Errorf("timestamp wrong ip: %s:%d , r.Timestamp:%d, current:%d", clientIp, req.ToPort, req.Timestamp, time.Now().Unix())
		c.String(http.StatusOK, "timestamp is too large")
		return
	}
	logger.Infof("report (%s:%s) %s:%d", req.Name, req.Type, clientIp, req.ToPort)

	if "gost" == req.Type {
		updated, errGost := RunGostPortAndIp(req.ToPort, clientIp)
		if errGost != nil {
			logger.Errorf("RunGostPortAndIp(%d, %s) failed: %v", req.ToPort, clientIp, errGost)
			c.String(http.StatusOK, "gost更新失败:"+errGost.Error())
		}

		if updated {
			c.String(http.StatusOK, "update gost")
		} else {
			c.String(http.StatusOK, "gost stay the same")
		}
	} else {
		updated, errIpdate := updateIptables(req, clientIp)
		if errIpdate != nil {
			logger.Errorf("updateIptables(%s  %s:%d) failed: %v", req.Name, clientIp, req.ToPort, errIpdate)
			c.String(http.StatusOK, "failed:"+errIpdate.Error())
			return
		}

		if updated {
			c.String(http.StatusOK, "update iptables")
		} else {
			c.String(http.StatusOK, "iptables stay the same")
		}
	}
}

func updateIptables(req model.ReportReq, clientIp string) (updated bool, err error) {
	rules, err := ListNatWithNumber()
	if err != nil {
		//c.String(http.StatusOK, "获取列表失败"+err.Error())
		return
	}

	for _, rule := range rules {
		if rule.ToPort != req.ToPort {
			continue
		}

		if rule.ToPort == req.ToPort {
			if rule.ToIp != clientIp {
				errDel := DelNatByLineNumber(rule.RuleNum)
				if errDel != nil {
					logger.Errorf("fail to del old nat rule: %v", errDel)
					err = errDel
					return
				}
				break
			} else {
				logger.Infof("report|%s|%s:%d the same", req.Name, clientIp, req.ToPort)
				updated = false
				return
			}
		}
	}

	errAdd := AddNatRule(model.NatRule{
		RuleNum:    0,
		ListenPort: req.ToPort,
		ToIp:       clientIp,
		ToPort:     req.ToPort,
	})

	if errAdd != nil {
		logger.Errorf("failed to add :%v", errAdd)
		err = errAdd
		return
	}
	updated = true
	logger.Infof("report|%s|%s:%d updated", req.Name, clientIp, req.ToPort)
	logger.Infof("report|%s|%s:%d rules not found", req.Name, clientIp, req.ToPort)
	return
}

func status(c *gin.Context) {
	rules, err := ListNatWithNumber()
	if err != nil {
		c.String(http.StatusOK, "failed:"+err.Error())
		return
	}

	c.JSON(http.StatusOK, rules)
}
