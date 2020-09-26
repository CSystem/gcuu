# SUBMAIL

发送邮件使用 submail 的 `mail/send` 功能。

## example

```go
config := make(map[string]string)
config["appid"] = "aaa" // submail api id
config["appkey"] = "bbb" // submail api key
config["signType"] = "sha1"
submail := CreateSend(config)
submail.AddTo("someuser@abc.com", "") // 收件人地址
submail.SetSender("no-reply@officaial.com", "Official") // 发件人地址
submail.SetSubject("Email Verification Code") // 邮件标题（200个字符以内）
submail.SetText("Testing Mail") // 纯文本邮件正文（5000个字符以内）
resp := submail.Send()
```

## links

* [submail Mail/send](https://www.mysubmail.com/chs/documents/developer/yR0Ov)
