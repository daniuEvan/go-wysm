/**
 * @date: 2022/2/24
 * @desc: 网易云信 短信验证码 - test
 * @auth: Evan
 */

package wysm

import (
	"fmt"
	"log"
	"testing"
)

func TestSmClient_SendSmCode(t *testing.T) {
	smClient := NewSmClient()
	smClient.SmConfig.Mobile = "接收验证码的手机号"
	smClient.SmConfig.AppKey = "网易云信AppKey"
	smClient.SmConfig.AppSecret = "网易云信AppSecret"
	//获取验证码
	res, err := smClient.SendSmCode()
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println(res)
}

func TestSmClient_VerifySmCode(t *testing.T) {
	smClient := NewSmClient()
	smClient.SmConfig.Mobile = "接收验证码的手机号"
	smClient.SmConfig.AppKey = "网易云信AppKey"
	smClient.SmConfig.AppSecret = "网易云信AppSecret"
	res, err := smClient.VerifySmCode("4128")
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println(res)
}
