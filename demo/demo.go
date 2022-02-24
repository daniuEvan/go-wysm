/**
 * @date: 2022/2/24
 * @desc: 网易云信 短信验证码
 */

package main

import (
	"fmt"
	"go-wysm/wysm"
	"log"
)

//
// sendSmCodeDemo
// @Description: 发送验证码 demo
//
func sendSmCodeDemo() {
	smClient := wysm.NewSmClient()
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

//
// verifySmCodeDemo
// @Description: 校验验证码
//
func verifySmCodeDemo() {
	smClient := wysm.NewSmClient()
	smClient.SmConfig.Mobile = "接收验证码的手机号"
	smClient.SmConfig.AppKey = "网易云信AppKey"
	smClient.SmConfig.AppSecret = "网易云信AppSecret"
	res, err := smClient.VerifySmCode("获取的验证码")
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println(res)
}

func main() {
	sendSmCodeDemo()
	verifySmCodeDemo()
}
