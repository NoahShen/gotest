package aria2rpc

import (
	"fmt"
	"github.com/kdar/httprpc"
	"log"
	"testing"
)

type Args struct {
	Who string
}

type Reply struct {
	Message string
}

func NoTestAddUri(t *testing.T) {
	method := "aria2.addUri"
	var params = make([]interface{}, 2)
	params[0] = make([]string, 1)
	(params[0].([]string))[0] = "https://www.kernel.org/pub/linux/kernel/v3.x/linux-3.8.4.tar.xz"
	params[1] = make(map[string]string)
	(params[1].(map[string]string))["max-download-limit"] = "10K"

	var replyGID string
	err := httprpc.CallJson("2.0", "http://127.0.0.1:6800/jsonrpc", method, params, &replyGID)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(replyGID)
}

func TestGetStatua(t *testing.T) {
	//method = "aria2.getGlobalStat"
	method := "aria2.tellActive"
	var params = make([]string, 0)
	var reply = make([]map[string]interface{}, 1)
	err := httprpc.CallJson("2.0", "http://127.0.0.1:6800/jsonrpc", method, params, &reply)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(reply)
}
