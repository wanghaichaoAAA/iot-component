/*
@Time   : 2020/11/26 9:00
@Author : Haichao Wang
*/
package rsa

import (
	"fmt"
	"testing"
)

func TestGenerateRsa(t *testing.T) {
	err, prvkey, pubkey := GenerateRsaKey()
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println("prvkey>>>", prvkey)
	fmt.Println("pubkey>>>", pubkey)

}
