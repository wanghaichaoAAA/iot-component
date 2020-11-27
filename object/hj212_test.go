/*
@Time   : 2020/11/26 10:24
@Author : Haichao Wang
*/
package object

import (
	"encoding/json"
	"fmt"
	"github.com/wanghaichaoAAA/iot-component/rsa"
	"testing"
)

var cpStr = "DataTime=20201126100000;w00000-ID=0,w00000-Cou=7.381000,w00000-Min=0,w00000-Avg=0.002,w00000-Max=3.413,w00000-Flag=N;w01001-Cou=43.12,w01001-Min=0,w01001-Avg=7.19,w01001-Max=7.25,w01001-Flag=N;w21003-ID=0,w21003-Cou=0.000038,w21003-Min=0,w21003-Avg=0.005,w21003-Max=0.006,w21003-Flag=N;w01018-ID=0,w01018-Cou=0.020336,w01018-Min=0,w01018-Avg=2.945,w01018-Max=3.743,w01018-Flag=N"

func Test_GenerateRsaMessage(t *testing.T) {
	err, Pirvatekey, Pubkey := rsa.GenerateRsaKey()
	if err != nil {
		t.Error("生成公钥私钥失败:", err)
		return
	}
	message := MakeRsaMessage("20201126102553427", "32", "2011", "123456", "888888", cpStr, "1", "1", "true", Pubkey)
	fmt.Println(message)

	hj212Message, err := NewRsaMessage(message[0], Pirvatekey)
	bytes, _ := json.Marshal(&hj212Message)
	fmt.Println(string(bytes))
}
