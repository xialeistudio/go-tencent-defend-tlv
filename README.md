# 腾讯天御业务安全系统TLV编码库
最近在基于腾讯天御业务安全系统做一个消息过滤的功能，其中有用到TLV编码，故开发本程序。

1. 本程序TLV编码可能和标准的有出入，仅通过了腾讯云天御业务安全系统接口测试。
2. 本程序编解码都是基于[]byte，但是腾讯云要求使用base64编码。使用的时候请自行编码

[腾讯天御业务安全系统文档](https://cloud.tencent.com/document/product/295/6603)


## 功能
+ [x] TLV编码
+ [x] TLV解码

## 单元测试
```
go test *.go
```

## 快速开始

### 编码
```go
buf, err = Encode(1, "天气不错")
if err != nil {
    fmt.Println(buf)
    return
}
fmt.Println(buf)
```

### 解码

```go
buf, err := base64.StdEncoding.DecodeString("AAAAAQAAAAzlpKnmsJTkuI3plJk=")
if err != nil {
    fmt.Println(buf)
    return
}
tag, data, err := Decode(buf)
if err != nil {
    fmt.Println(buf)
    return
}
fmt.Printf("tag=%d,data=%s",tag,string(data))
```