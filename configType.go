package wed

type AllConfig struct {
	TimeStamp string
	App       app
	Mongodb   mongodb
	Redis     redis
	Mysql     mysqlt
	Jwt       jwtt
	AliyunSMS aliyunSMS
}

type app struct {
	Name    string
	Port    string
	Version string
	Debug   bool
}
type mongodb struct {
	Uri string
	DB  string
}
type redis struct {
	Addr      string
	Paswsword string
	DB        int
}
type mysqlt struct {
	Driver string
	Uri    string
}
type jwtt struct {
	Secret string
	Issuer string
	Time   int64
}
type aliyunSMS struct {
	AccessKeyID     string
	AccessKeySecret string
	SignName        string
	TemplateCode    string
}
