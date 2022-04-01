package wed

var configVar = `TimeStamp: "2018-07-16 10:23:19"
App: 
  Port: "9000"
  Debug: True
mongodb:
  uri: "mongodb://localhost:27017"
  db: "user"
redis:
  addr: "localhost:6379"
  password: ""
  db: 0
mysql:
  uri:
    "root:123456@tcp(localhost:3306)/user?charset=utf8&parseTime=True&loc=Local&multiStatements=true"
jwt:
  secret: "werwer2323224e4W"
  #签发方
  Issuer: "wegin"
  #有效时间 单位秒
  Time: 5000
AliyunSMS:
  AccessKeyID: "ddddddddddddddddd"
  AccessKeySecret: "dsssssssss"
  SignName: "sd"
  TemplateCode: "dsfsdf"`
