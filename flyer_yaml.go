package flyer

var config AllConfig

type AllConfig struct {
	App       app_type
	Mongodb   mongodb_type
	Redis     redis_type
	Mysql     mysqlt_type
	Jwt       jwt_type
	AliyunSMS aliyunSMS_type
	Wechat    wechat_type
}

type app_type struct {
	Port    string
	Debug   bool
	Swagger bool
}
type mongodb_type struct {
	Uri string
	DB  string
}
type redis_type struct {
	Addr      string
	Paswsword string
	DB        int
}
type mysqlt_type struct {
	Uri         string
	TablePrefix string
}
type jwt_type struct {
	Secret string
	Issuer string
	Time   int64
}
type aliyunSMS_type struct {
	AccessKeyID     string
	AccessKeySecret string
	SignName        string
	TemplateCode    string
}

type wechat_type struct {
	AppID     string
	AppSecret string
}

var configYaml = `TimeStamp: "2022-04-03"
App: 
  port: "9000"
  debug: True
  swagger: True
mongodb:
  uri: "mongodb://localhost:27017"
  db: "user"
redis:
  addr: "localhost:6379"
  password: ""
  db: 0
mysql:
  uri: "root:123456@tcp(localhost:3306)/user?charset=utf8&parseTime=True&loc=Local&multiStatements=true"
  TablePrefix: "wed_"
jwt:
  secret: "werwer2323224e4W"
  #签发方
  Issuer: "wegin"
  #有效时间 单位秒
  Time: 5000
wechat:
  appID:
  appSecret: 
AliyunSMS:
  AccessKeyID: "ddddddddddddddddd"
  AccessKeySecret: "dsssssssss"
  SignName: "sd"
  TemplateCode: "dsfsdf"`
