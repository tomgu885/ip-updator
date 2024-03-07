package server

import (
	"fmt"
	"github.com/coreos/go-iptables/iptables"
	"iptables/logger"
	"iptables/model"
	"os/exec"
	"strconv"
	"strings"
)

func ListRules(tuleType string) (rules []string, err error) {
	ipt, err := iptables.New()
	if err != nil {
		logger.Errorf("failed to create ipt:%v", err)
		return
	}

	//rules, err = ipt.ListChains(tuleType)
	rules, err = ipt.List("nat", "PREROUTING")
	if err == nil {
		logger.Infof("rules2 got: %d", len(rules))
	}
	return
}

func AddNatRule(rule model.NatRule) (err error) {
	// iptables -t nat -A PREROUTING -i eth0 -p tcp --dport 10024 -j DNAT --to-destination 114.235.208.124:10024
	listenPort := fmt.Sprintf("%d", rule.ListenPort)
	to := fmt.Sprintf("%s:%d", rule.ToIp, rule.ToPort)
	out, err := exec.Command("iptables",
		"-t", "nat", "-A", "PREROUTING", "-i", "eth",
		"-p", "tcp", "--dport", listenPort,
		"-j", "DNAT", "--to-destination", to).Output()

	if err != nil {
		return
	}

	logger.Infof("rule added: %s", string(out))

	return
}

func DelNatPortByToPort(toPort int) (deleted bool, err error) {
	rules, err := ListNatWithNumber()
	if err != nil {
		return
	}

	for _, rule := range rules {
		if toPort == rule.ToPort {
			errDel := DelNatByLineNumber(rule.RuleNum)
			if errDel != nil {
				return
			}
			deleted = true
			return
		}
	}

	return
}

func DelNatByLineNumber(line int) (err error) {
	// #iptables -t nat -D PREROUTING {rule-number-here}
	out, err := exec.Command("iptables", "-t", "nat", "-D", "PREROUTING", fmt.Sprintf("%d", line)).Output()

	if err != nil {
		logger.Errorf("failed to delete rules(%d) err%v", line, err)
		return
	}

	logger.Infof("rule(%d) deleted: out:%s", line, out)

	return
}

func ListNatWithNumber() (rules []model.NatRule, err error) {
	// #iptables -t nat -v -L PREROUTING -n --line-number
	out, err := exec.Command("iptables", "-t", "nat", "-L", "PREROUTING", "-n", "--line-number").Output()
	//out, err := exec.Command("exec")
	if err != nil {
		logger.Errorf("err")
		return
	}

	outlines := strings.Split(string(out), "\n")
	for _, line := range outlines {
		if len(line) == 0 {
			continue
		}

		parts := strings.Fields(line)
		ruleNum, errN := strconv.Atoi(parts[0])
		if errN != nil {
			continue
		}

		portParts := strings.Split(parts[7], ":")
		toParts := strings.Split(parts[8], ":")

		listenPort, _ := strconv.Atoi(portParts[1])
		forwardIp := toParts[1]
		foreardPort, _ := strconv.Atoi(toParts[2])
		//rules[num] = parts[2]
		rules = append(rules, model.NatRule{
			RuleNum:    ruleNum,
			ListenPort: listenPort,
			ToIp:       forwardIp,
			ToPort:     foreardPort,
		})
	}

	return
}
