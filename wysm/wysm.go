/**
 * @date: 2022/2/24
 * @desc: 网易云信 短信验证码
 * @auth: Evan
 */

package wysm

import (
	"crypto/sha1"
	"fmt"
	"io"
	"io/ioutil"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
	"time"
)

//
// smConfig
// @Description: sm参数配置
//
type smConfig struct {
	Mobile          string
	contentType     string
	sendSmBaseUrl   string
	AppSecret       string
	AppKey          string
	SmTemplateCode  int
	CodeLen         int
	verifySmBaseUrl string
}

type smClient struct {
	SmConfig *smConfig
}

const (
	contentType     = "application/x-www-form-urlencoded;charset=utf-8" // Content-Type
	sendSmBaseUrl   = "https://api.netease.im/sms/sendcode.action"      // 网易云信获取验证码 url
	smTemplateCode  = 19506299                                          // 网易云信 smTemplateCode
	codeLen         = 4                                                 // 验证码的长度
	verifySmBaseUrl = "https://api.netease.im/sms/verifycode.action"    // 网易云信校验验证码 url
)

func initSmConfig() *smConfig {
	return &smConfig{
		contentType:     contentType,
		sendSmBaseUrl:   sendSmBaseUrl,
		SmTemplateCode:  smTemplateCode,
		CodeLen:         codeLen,
		verifySmBaseUrl: verifySmBaseUrl,
	}
}

//
// NewSmClient
// @Description: 构建 sm client
// @return *smClient:
//
func NewSmClient() *smClient {
	return &smClient{
		SmConfig: initSmConfig(),
	}
}

//
// SendSmCode
// @Description: 发送验证码
//
func (s *smClient) SendSmCode() (resObj string, err error) {
	config := s.SmConfig
	rand.Seed(time.Now().UnixNano())
	curTime, nonce, checkSum, err := s.buildCheckArgs()
	if err != nil {
		return "", err
	}
	client := http.Client{}
	body := strings.NewReader(fmt.Sprintf(`codeLen=%d&templateid=%d&mobile=%s`, config.CodeLen, config.SmTemplateCode, config.Mobile))
	request, err := http.NewRequest("POST", config.sendSmBaseUrl, body)
	// 增加 head 选项
	request.Header.Add("Content-Type", config.contentType)
	request.Header.Add("AppKey", config.AppKey)
	request.Header.Add("CurTime", curTime)
	request.Header.Add("Nonce", nonce)
	request.Header.Add("CheckSum", checkSum)
	response, err := client.Do(request)
	if err != nil {
		return "", err
	}
	defer response.Body.Close()
	resMsg, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", err
	}
	return string(resMsg), err
}

//
// VerifySmCode
// @Description: 校验验证码
//
func (s *smClient) VerifySmCode(code string) (resObj string, err error) {
	config := s.SmConfig
	curTime, nonce, checkSum, err := s.buildCheckArgs()
	if err != nil {
		return "", err
	}
	client := http.Client{}
	body := strings.NewReader(fmt.Sprintf(`mobile=%s&code=%s`, config.Mobile, code))
	request, err := http.NewRequest("POST", config.verifySmBaseUrl, body)
	// 增加 head 选项
	request.Header.Add("Content-Type", config.contentType)
	request.Header.Add("AppKey", config.AppKey)
	request.Header.Add("CurTime", curTime)
	request.Header.Add("Nonce", nonce)
	request.Header.Add("CheckSum", checkSum)
	response, err := client.Do(request)
	if err != nil {
		return "", err
	}
	defer response.Body.Close()
	resMsg, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", err
	}
	return string(resMsg), nil
}

//
// buildCheckArgs
// @Description: 构建使用参数
// @return curTime:
// @return nonce:
// @return checkSum:
// @return err:
//
func (s *smClient) buildCheckArgs() (curTime, nonce, checkSum string, err error) {
	rand.Seed(time.Now().UnixNano())
	curTime = strconv.FormatInt(time.Now().Unix(), 10)
	nonce = strconv.FormatInt(int64(rand.Intn(10)), 10)
	tempStr := s.SmConfig.AppSecret + nonce + curTime
	t := sha1.New()
	_, err = io.WriteString(t, tempStr)
	if err != nil {
		return "", "", "", err
	}
	return curTime, nonce, fmt.Sprintf("%x", t.Sum(nil)), nil
}
