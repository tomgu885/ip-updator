package server

import (
	"fmt"
	"iptables/logger"
	"net"
	"net/url"
	"os/exec"
	"strconv"
	"strings"
)

func RunGostPortAndIp(port int, nodeIp string) (updated bool, err error) {
	pid, oldIp, err := GetPidGostByPort(port)
	if err != nil {
		logger.Errorf("GetPidGostByPort(%d) failed:%v", port, err)
		return
	}

	if pid > 0 && nodeIp == oldIp {
		logger.Infof("RunGostPortAndIp the same, not need to run")
		return
	}

	if pid > 0 && nodeIp != oldIp {
		logger.Infof("kill old pid(%d)", pid)
		_, err = exec.Command("kill", fmt.Sprintf("%d", pid)).Output()
		if err != nil {
			return
		}
	}
	// nohup gost -L ss://aes-256-gcm:w0rd2019@:8005 -F relay+wss://jack:w0rd2019@115.195.94.45:8445 > 8005.log 2>&1 &
	cmd := fmt.Sprintf("nohup gost -L ss://aes-256-gcm:w0rd2019@:%d -F relay+wss://jack:w0rd2019@%s:8445 > %d.log 2>&1 &", port, nodeIp, port)
	_cmd := exec.Command("bash", "-c", cmd, "--detached")
	err = _cmd.Start()
	if err != nil {
		fmt.Println("RunGostPortAndIp:", err.Error())
		return
	}
	_cmd.Process.Release()
	fmt.Println("run....ed")
	updated = true
	return
}

func GetPidGostByPort(port int) (pid int, nodeIp string, err error) {
	// ps aux | grep 8003
	cmd := fmt.Sprintf("ps aux | grep gost | grep %d", port)
	fmt.Println("List cmd:", cmd)
	out, err := exec.Command("bash", "-c", cmd).Output()

	if err != nil {
		logger.Errorf("failed: %s", err.Error())
		return
	}

	lines := strings.Split(string(out), "\n")
	//fmt.Println("out:", lines)

	for _, line := range lines {
		parts := strings.Fields(line)
		//fmt.Println("parts", parts)
		//fmt.Println("len(parts)", len(parts))
		if len(parts) < 15 {
			continue
		}

		if parts[10] != "gost" {
			// not gost process
			continue
		}

		fmt.Printf("pid: %s 10: %s|12:%s| 13:%s  |14:%s \n", parts[1], parts[10], parts[12], parts[13], parts[14])
		// relay+wss://jack:w0rd2019@1.2.3.4:8445
		u, errU := url.Parse(parts[14])
		if errU != nil {
			err = errU
			return
		}
		//fmt.Println("u.Host", u.Host)
		host, _, _ := net.SplitHostPort(u.Host)
		fmt.Println("host:", host)

		//_split := strings.Split(parts[14], ":")
		var errP error
		pid, errP = strconv.Atoi(parts[1])
		if errP != nil {
			err = errP
			return
		}

		nodeIp = host
	}

	return
}
