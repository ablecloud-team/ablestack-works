package main

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
