# gcrawler
gcrawler是一个基于go的爬虫练手项目。   
1、cnvd  
https://www.cnvd.org.cn/flaw/typeResult?typeId=33&max=20&offset=20
```json
{
	"title": "Huawei Mate 30 Pro授权问题漏洞",
	"time": "https://www.cnvd.org.cn//flaw/show/CNVD-2020-51533",
	"url": "2020-09-09",
	"content": "{\"CNVD-ID\":\"CNVD-2020-51533\",\"CVEID\":\"CVE-2020-1838\",\"公开日期\":\"2020-09-09\",\"危害级别\":\"低(AV:L/AC:M/Au:N/C:N/I:P/A:N)\",\"厂商补丁\":\"HuaweiMate30Pro授权问题漏洞的补丁\",\"参考链接\":\"https://nvd.nist.gov/vuln/detail/CVE-2020-1838\",\"影响产品\":\"HuaweiMate30Pro\\u003c10.1.0.150(C00E136R5P3)\",\"报送时间\":\"2020-07-02\",\"收录时间\":\"2020-09-11\",\"更新时间\":\"2020-09-11\",\"漏洞描述\":\"HuaweiMate30Pro是中国华为（Huawei）公司的一款智能手机。HuaweiMate30Pro10.1.0.150(C00E136R5P3)之前版本中存在安全漏洞，该漏洞源于设备未充分校验用户人脸凭据。攻击者可通过仿造一个用户凭据利用该漏洞通过认证。\",\"漏洞类型\":\"通用型漏洞\",\"漏洞解决方案\":\"目前厂商已发布升级补丁以修复漏洞，补丁获取链接：https://www.huawei.com/cn/psirt/security-advisories/huawei-sa-20200701-03-smartphone-cn\",\"漏洞附件\":\"(无附件)\",\"验证信息\":\"(暂无验证信息)\"}"
}
```

待做
```
https://us-cert.cisa.gov/ics/alerts
https://www.securityweek.com/iot-security
```


参考资料
```
http库：http:https://golang.org/pkg/net/http/
网页解析：https://github.com/PuerkitoBio/goquery#examples
json库：https://blog.golang.org/json
```


相关环境配置问题
```
1、因网络原因导致包下载失败
在goland配置环境，GOPROXY=https://goproxy.io,direct
2、goquery获取a标签内href值
https://stackoverflow.com/questions/32171498/how-to-get-value-of-attribute-href-value-in-go-language
3、map对象转换成json
https://stackoverflow.com/questions/24652775/convert-go-map-to-json/24652881
```