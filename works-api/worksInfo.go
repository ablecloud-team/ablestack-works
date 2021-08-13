package main

import "os"
var (
	DCProtocol = getWorksInfo().DomainController.Protocol
	DCUrl = getWorksInfo().DomainController.Url
	DCPort = getWorksInfo().DomainController.Port
	DCPostfix = getWorksInfo().DomainController.Postfix
	DCInfo = DCProtocol+DCUrl+":"+DCPort+DCPostfix
)

func setenv(){
	os.Setenv("MsqlType", getWorksInfo().Database.TYPE)
	os.Setenv("MysqlUser", getWorksInfo().Database.User.ID)
	os.Setenv("MysqlPassword", getWorksInfo().Database.User.Password)
	os.Setenv("MysqlProtocol", getWorksInfo().Database.Host.Protocol)
	os.Setenv("MysqlServerIp", getWorksInfo().Database.Host.Address)
	os.Setenv("MysqlServerPort", getWorksInfo().Database.Host.Port)
	os.Setenv("MysqlDb", getWorksInfo().Database.DB)
	os.Setenv("DbInfo", getWorksInfo().Database.User.ID +":"+ getWorksInfo().Database.User.Password +"@"+ getWorksInfo().Database.Host.Protocol+"("+ getWorksInfo().Database.Host.Address +":"+ getWorksInfo().Database.Host.Port +")/"+ getWorksInfo().Database.DB)
	os.Setenv("MoldAPIKey", getWorksInfo().Mold.User.ApiKey)
	os.Setenv("MoldSecretKey", getWorksInfo().Mold.User.SecretKey)
}

//var(
//	MoldProtocal = getWorksInfo().Mold.Protocol
//	MoldUrl = getWorksInfo().Mold.Url
//	MoldPostfix = getWorksInfo().Mold.Postfix
//)

