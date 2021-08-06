package main

var (
	MsqlType       	= getWorksInfo().Database.TYPE
	MysqlUser      	= getWorksInfo().Database.User.ID
	MysqlPassword   = getWorksInfo().Database.User.Password
	MysqlProtocol   = getWorksInfo().Database.Host.Protocol
	MysqlServerIp   = getWorksInfo().Database.Host.Address
	MysqlServerPort = getWorksInfo().Database.Host.Port
	MysqlDb         = getWorksInfo().Database.DB
	DbInfo = MysqlUser +":"+ MysqlPassword +"@"+ MysqlProtocol +"("+ MysqlServerIp +":"+ MysqlServerPort +")/"+ MysqlDb
)

var (
	DCProtocol = getWorksInfo().DomainController.Protocol
	DCUrl = getWorksInfo().DomainController.Url
	DCPort = getWorksInfo().DomainController.Port
	DCPostfix = getWorksInfo().DomainController.Postfix
	DCInfo = DCProtocol+DCUrl+":"+DCPort+DCPostfix
)