package settings

import (
	"fmt"
	"testing"
)

func TestSetup(t *testing.T) {
	err := Setup("../app.yaml")
	if err != nil {
		t.Errorf("failed: %v", err)
		return
	}

	fmt.Println("loaded")
	fmt.Printf("global.secret: |%s|\n", GetGlobal().Secret)

	fmt.Println("====== Server =====")
	fmt.Printf("server.port:%d\n", GetServer().Port)

	fmt.Println("-======= Client ======-")
	fmt.Printf("client.Server: |%s|\n", GetClient().Server)
	fmt.Printf("localPort:%d\n", GetClient().LocalPort)
	fmt.Printf("name: %s\n", GetClient().Name)
	// 1709727521_test-1_10145
	// 1709727521_test-1_10145
}
