package main

import "time"

const (
	WorkspaceString = "workspace"
	InstanceString  = "instance"
	ServiceDaaS     = "ServiceDaaS"
	ServiceWorks    = "ServiceWorks"
	WorkspaceName   = "WorkspaceName"
	AblecloudWorks  = "Ablecloud.Works"
	ClusterName     = "ClusterName"
	AgentOK         = "AgentOK"
	AgentCheck      = "AgentCheck"
	Enable          = "Enable"
	Disable         = "Disable"
	UserVm          = "UserVm"
	ALL             = "all"

	Trace = "Trace"
	Debug = "Debug"
	Info  = "Info"
	Warn  = "Warn"
	Error = "Error"

	VMStart   = "VMStart"
	VMStop    = "VMStop"
	VMsDeploy = "VMsDeploy"
	VMDestroy = "VMDestroy"

	MsgDBConnectError      = "DB connect error"
	MsgDBConnectOK         = "DB connect success"
	MsgDBNoRows            = "sql: no rows in result set"
	BaseErrorCode          = 9000
	SignatureErrorCode     = 9001
	SQLConnectError        = 9100
	SQLQueryError          = 9101
	NotFound404            = "404 Not Found"
	Conflict409            = "409 Conflict"
	Gone410                = "410 Gone"
	OK200                  = "200 OK"
	Created201             = "201 Created"
	Accepted202            = "202 Accepted"
	NotFound               = "Not Found"
	MessageAccountNotFound = "message.account.not.found"
	MessageSignatureError  = "message.Signature.error"
	MessageAgentUpdateOK   = "message.Agent.update.ok"
)

type ListVirtualMachinesMetrics struct {
	Count          int `json:"count"`
	Virtualmachine []struct {
		Account       string        `json:"account"`
		Affinitygroup []interface{} `json:"affinitygroup"`
		Cpunumber     int           `json:"cpunumber"`
		Cpuspeed      int           `json:"cpuspeed"`
		Cputotal      string        `json:"cputotal"`
		Cpuused       string        `json:"cpuused"`
		Created       string        `json:"created"`
		Details       struct {
			RootDiskController string `json:"rootDiskController"`
		} `json:"details"`
		Diskiopstotal         int    `json:"diskiopstotal"`
		Diskioread            int    `json:"diskioread"`
		Diskiowrite           int    `json:"diskiowrite"`
		Diskkbsread           int    `json:"diskkbsread"`
		Diskkbswrite          int    `json:"diskkbswrite"`
		Diskofferingid        string `json:"diskofferingid"`
		Diskofferingname      string `json:"diskofferingname"`
		Displayname           string `json:"displayname"`
		Domain                string `json:"domain"`
		Domainid              string `json:"domainid"`
		Guestosid             string `json:"guestosid"`
		Haenable              bool   `json:"haenable"`
		Hypervisor            string `json:"hypervisor"`
		Id                    string `json:"id"`
		Ipaddress             string `json:"ipaddress"`
		Isdynamicallyscalable bool   `json:"isdynamicallyscalable"`
		Memory                int    `json:"memory"`
		Memoryintfreekbs      int    `json:"memoryintfreekbs"`
		Memorykbs             int    `json:"memorykbs"`
		Memorytargetkbs       int    `json:"memorytargetkbs"`
		Memorytotal           string `json:"memorytotal"`
		Name                  string `json:"name"`
		Networkkbsread        int    `json:"networkkbsread"`
		Networkkbswrite       int    `json:"networkkbswrite"`
		Networkread           string `json:"networkread"`
		Networkwrite          string `json:"networkwrite"`
		Nic                   []struct {
			Broadcasturi    string        `json:"broadcasturi"`
			Extradhcpoption []interface{} `json:"extradhcpoption"`
			Gateway         string        `json:"gateway"`
			Id              string        `json:"id"`
			Ipaddress       string        `json:"ipaddress"`
			Isdefault       bool          `json:"isdefault"`
			Isolationuri    string        `json:"isolationuri"`
			Macaddress      string        `json:"macaddress"`
			Netmask         string        `json:"netmask"`
			Networkid       string        `json:"networkid"`
			Networkname     string        `json:"networkname"`
			Secondaryip     []interface{} `json:"secondaryip"`
			Traffictype     string        `json:"traffictype"`
			Type            string        `json:"type"`
		} `json:"nic"`
		Osdisplayname       string        `json:"osdisplayname"`
		Ostypeid            string        `json:"ostypeid"`
		Passwordenabled     bool          `json:"passwordenabled"`
		Pooltype            string        `json:"pooltype"`
		Readonlydetails     string        `json:"readonlydetails"`
		Rootdeviceid        int           `json:"rootdeviceid"`
		Rootdevicetype      string        `json:"rootdevicetype"`
		Securitygroup       []interface{} `json:"securitygroup"`
		Serviceofferingid   string        `json:"serviceofferingid"`
		Serviceofferingname string        `json:"serviceofferingname"`
		State               string        `json:"state"`
		Tags                []interface{} `json:"tags"`
		Templatedisplaytext string        `json:"templatedisplaytext"`
		Templateid          string        `json:"templateid"`
		Templatename        string        `json:"templatename"`
		Userid              string        `json:"userid"`
		Username            string        `json:"username"`
		Zoneid              string        `json:"zoneid"`
		Zonename            string        `json:"zonename"`
	} `json:"virtualmachine"`
}

type InstanceInstanceVolumeInfo struct {
	Count  int `json:"count"`
	Volume []struct {
		Account                    string        `json:"account"`
		Created                    string        `json:"created"`
		Deviceid                   int           `json:"deviceid"`
		Diskiopstotal              int           `json:"diskiopstotal"`
		Diskioread                 int           `json:"diskioread"`
		Diskiowrite                int           `json:"diskiowrite"`
		Diskkbsread                int           `json:"diskkbsread"`
		Diskkbswrite               int           `json:"diskkbswrite"`
		Domain                     string        `json:"domain"`
		Domainid                   string        `json:"domainid"`
		Id                         string        `json:"id"`
		Isextractable              bool          `json:"isextractable"`
		Name                       string        `json:"name"`
		Provisioningtype           string        `json:"provisioningtype"`
		Quiescevm                  bool          `json:"quiescevm"`
		Serviceofferingdisplaytext string        `json:"serviceofferingdisplaytext"`
		Serviceofferingid          string        `json:"serviceofferingid"`
		Serviceofferingname        string        `json:"serviceofferingname"`
		Size                       int64         `json:"size"`
		Sizegb                     string        `json:"sizegb"`
		State                      string        `json:"state"`
		Supportsstoragesnapshot    bool          `json:"supportsstoragesnapshot"`
		Tags                       []interface{} `json:"tags"`
		Templatedisplaytext        string        `json:"templatedisplaytext"`
		Templateid                 string        `json:"templateid"`
		Templatename               string        `json:"templatename"`
		Type                       string        `json:"type"`
		Virtualmachineid           string        `json:"virtualmachineid"`
		Vmdisplayname              string        `json:"vmdisplayname"`
		Vmname                     string        `json:"vmname"`
		Vmstate                    string        `json:"vmstate"`
		Zoneid                     string        `json:"zoneid"`
		Zonename                   string        `json:"zonename"`
	} `json:"volume"`
}

type Configuration struct {
	Id, Name, Category, Value, DefaultValue, Description string
	UpdateDate                                           time.Time
}