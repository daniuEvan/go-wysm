# Golang 网易云信(短信验证码)  SDK
- [Python SDK 网易云信(短信验证码)](https://github.com/daniuEvan/pywyyx)


- [网易云信官网地址]( https://yunxin.163.com)
- [网易云信短信功能开通说明](https://doc.yunxin.163.com/docs/DI1Mzc2NTU/TE1ODQ0NDY?platformId=120002)
- [官方短信验证码接口文档](https://doc.yunxin.163.com/docs/DI1Mzc2NTU/zA2MjExNzY?platformId=120002) 

## 安装

```go
go get github.com/daniuEvan/go-wysm/wysm
```

## 创建客户端

```go
smClient := wysm.NewSmClient()
```

## 初始化配置

```go
smClient.SmConfig.Mobile = "接收验证码的手机号"
smClient.SmConfig.AppKey = "网易云信AppKey"
smClient.SmConfig.AppSecret = "网易云信AppSecret"
smClient.SmConfig.SmTemplateCode = 19506299         // 验证码模板 默认为 19506299
smClient.SmConfig.CodeLen = 4                       // 验证码长度默认为 4
```

## 发送短信验证码

```go
resJson, err := smClient.SendSmCode()  // 返回值resJson类型为json
```

- 成功回resJson值, msg字段表示此次发送的sendid；obj字段表示此次发送的验证码。

  ```go
  {
    "code": 200,   // 状态码
    "msg": "88",   
    "obj": "1908"  // 验证码
  }
  ```

- 更多返回码请参考: https://doc.yunxin.163.com/docs/TM5MzM5Njk/Tk5ODIzNjk

- 完整发送短信验证码demo

  ```go
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
  ```

## 校验短信验证码

```go
resJson, err := smClient.VerifySmCode("获取的验证码")   // 返回值resJson类型为json
```

- 成功resJson返回值:

  ```go
  {
    "code":200
  }
  ```

- 更多返回码请参考: https://doc.yunxin.163.com/docs/TM5MzM5Njk/Tk5ODIzNjk

- 完整校验短信验证码示例

  ```golang
  	smClient := wysm.NewSmClient()
  	smClient.SmConfig.Mobile = "接收验证码的手机号"
  	smClient.SmConfig.AppKey = "网易云信AppKey"
  	smClient.SmConfig.AppSecret = "网易云信AppSecret"
  	res, err := smClient.VerifySmCode("获取的验证码")
  	if err != nil {
  		log.Fatal(err.Error())
  	}
  	fmt.Println(res)
  ```



## 测试用例代码

- 完整代码: https://github.com/daniuEvan/go-wysm/blob/master/wysm/wysm_test.go

  ```go
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
  
  ```

  