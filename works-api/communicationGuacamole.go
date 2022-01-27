package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
	"time"
)

type GuacamoleUser struct {
	Username   string `json:"username"`
	Password   string `json:"password"`
	Attributes struct {
		GuacOrganizationalRole string `json:"guac-organizational-role"`
		GuacFullName           string `json:"guac-full-name"`
		Expired                string `json:"expired"`
		AccessWindowStart      string `json:"access-window-start"`
		GuacOrganization       string `json:"guac-organization"`
		AccessWindowEnd        string `json:"access-window-end"`
		Disabled               string `json:"disabled"`
		ValidUntil             string `json:"valid-until"`
		ValidFrom              string `json:"valid-from"`
	} `json:"attributes"`
}

type GuacamoleConnections struct {
	ParentIdentifier string `json:"parentIdentifier,omitempty"`
	Name             string `json:"name,omitempty"`
	Protocol         string `json:"protocol,omitempty"`
	Parameters       struct {
		Port                     string `json:"port,omitempty"`
		ReadOnly                 string `json:"read-only,omitempty"`
		SwapRedBlue              string `json:"swap-red-blue,omitempty"`
		Cursor                   string `json:"cursor,omitempty"`
		ColorDepth               string `json:"color-depth,omitempty"`
		ClipboardEncoding        string `json:"clipboard-encoding,omitempty"`
		DisableCopy              string `json:"disable-copy,omitempty"`
		DisablePaste             string `json:"disable-paste,omitempty"`
		DestPort                 string `json:"dest-port,omitempty"`
		RecordingExcludeOutput   string `json:"recording-exclude-output,omitempty"`
		RecordingExcludeMouse    string `json:"recording-exclude-mouse,omitempty"`
		RecordingIncludeKeys     string `json:"recording-include-keys,omitempty"`
		CreateRecordingPath      string `json:"create-recording-path,omitempty"`
		EnableSftp               string `json:"enable-sftp,omitempty"`
		SftpPort                 string `json:"sftp-port,omitempty"`
		SftpServerAliveInterval  string `json:"sftp-server-alive-interval,omitempty"`
		EnableAudio              string `json:"enable-audio,omitempty"`
		Security                 string `json:"security,omitempty"`
		DisableAuth              string `json:"disable-auth,omitempty"`
		IgnoreCert               string `json:"ignore-cert,omitempty"`
		GatewayPort              string `json:"gateway-port,omitempty"`
		ServerLayout             string `json:"server-layout,omitempty"`
		Timezone                 string `json:"timezone,omitempty"`
		Console                  string `json:"console,omitempty"`
		Width                    string `json:"width,omitempty"`
		Height                   string `json:"height,omitempty"`
		Dpi                      string `json:"dpi,omitempty"`
		ResizeMethod             string `json:"resize-method,omitempty"`
		ConsoleAudio             string `json:"console-audio,omitempty"`
		DisableAudio             string `json:"disable-audio,omitempty"`
		EnableAudioInput         string `json:"enable-audio-input,omitempty"`
		EnablePrinting           string `json:"enable-printing,omitempty"`
		EnableDrive              string `json:"enable-drive,omitempty"`
		CreateDrivePath          string `json:"create-drive-path,omitempty"`
		EnableWallpaper          string `json:"enable-wallpaper,omitempty"`
		EnableTheming            string `json:"enable-theming,omitempty"`
		EnableFontSmoothing      string `json:"enable-font-smoothing,omitempty"`
		EnableFullWindowDrag     string `json:"enable-full-window-drag,omitempty"`
		EnableDesktopComposition string `json:"enable-desktop-composition,omitempty"`
		EnableMenuAnimations     string `json:"enable-menu-animations,omitempty"`
		DisableBitmapCaching     string `json:"disable-bitmap-caching,omitempty"`
		DisableOffscreenCaching  string `json:"disable-offscreen-caching,omitempty"`
		DisableGlyphCaching      string `json:"disable-glyph-caching,omitempty"`
		PreconnectionId          string `json:"preconnection-id,omitempty"`
		Hostname                 string `json:"hostname,omitempty"`
		Username                 string `json:"username,omitempty"`
		Password                 string `json:"password,omitempty"`
		Domain                   string `json:"domain,omitempty"`
		GatewayHostname          string `json:"gateway-hostname,omitempty"`
		GatewayUsername          string `json:"gateway-username,omitempty"`
		GatewayPassword          string `json:"gateway-password,omitempty"`
		GatewayDomain            string `json:"gateway-domain,omitempty"`
		InitialProgram           string `json:"initial-program,omitempty"`
		ClientName               string `json:"client-name,omitempty"`
		PrinterName              string `json:"printer-name,omitempty"`
		DriveName                string `json:"drive-name,omitempty"`
		DrivePath                string `json:"drive-path,omitempty"`
		StaticChannels           string `json:"static-channels,omitempty"`
		RemoteApp                string `json:"remote-app,omitempty"`
		RemoteAppDir             string `json:"remote-app-dir,omitempty"`
		RemoteAppArgs            string `json:"remote-app-args,omitempty"`
		PreconnectionBlob        string `json:"preconnection-blob,omitempty"`
		LoadBalanceInfo          string `json:"load-balance-info,omitempty"`
		RecordingPath            string `json:"recording-path,omitempty"`
		RecordingName            string `json:"recording-name,omitempty"`
		SftpHostname             string `json:"sftp-hostname,omitempty"`
		SftpHostKey              string `json:"sftp-host-key,omitempty"`
		SftpUsername             string `json:"sftp-username,omitempty"`
		SftpPassword             string `json:"sftp-password,omitempty"`
		SftpPrivateKey           string `json:"sftp-private-key,omitempty"`
		SftpPassphrase           string `json:"sftp-passphrase,omitempty"`
		SftpRootDirectory        string `json:"sftp-root-directory,omitempty"`
		SftpDirectory            string `json:"sftp-directory,omitempty"`
	} `json:"parameters"`
	Attributes struct {
		MaxConnections        string `json:"max-connections,omitempty"`
		MaxConnectionsPerUser string `json:"max-connections-per-user,omitempty"`
		Weight                string `json:"weight,omitempty"`
		FailoverOnly          string `json:"failover-only,omitempty"`
		GuacdPort             string `json:"guacd-port,omitempty"`
		GuacdEncryption       string `json:"guacd-encryption,omitempty"`
		GuacdHostname         string `json:"guacd-hostname,omitempty"`
	} `json:"attributes"`
}

func getGuacamoleToken() (*http.Response, error) {
	//var guacamoleUrl = os.Getenv("GuacamoleUrl")
	//TODO 과카몰리 URL을 works URL을 이용하여 통신함
	var guacamoleUrl = "http://" + os.Getenv("WorksIp") + ":8080"
	//var guacamoleUsername = os.Getenv("GuacamoleUsername")
	resource := "/api/tokens"

	params := url.Values{}
	//TODO 과카몰리 계정 및 패스워드 하드코딩

	params.Set("username", "administrator")
	params.Set("password", "Ablecloud1!")
	log.Infof("paramsInfo = [%v]", params)
	u, _ := url.ParseRequestURI(guacamoleUrl)
	u.Path = resource
	urlStr := u.String() // "https://api.com/user/"
	client := &http.Client{
		Timeout: 60 * time.Second,
	}
	//log.Infof("url [%v], username [%v]", urlStr, guacamoleUsername)
	r, err := http.NewRequest("POST", urlStr, strings.NewReader(params.Encode())) // URL-encoded payload
	if err != nil {
		log.Fatal(err)
	}
	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	r.Header.Add("Content-Length", strconv.Itoa(len(params.Encode())))
	resp, err := client.Do(r)
	if err != nil {
		log.Errorf("%v", err)
	} else {
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Errorf("%v", err)
		}
		result := map[string]interface{}{}
		json.Unmarshal(body, &result)
		log.Infof("communication guacamole [%v], result [%v], error [%v]", resp, result, err)
		//return result["authToken"].(string)
		return resp, err

	}

	return nil, nil
}

func getGuacamoleTokenString() string {

	resp, _ := getGuacamoleToken()
	body, _ := ioutil.ReadAll(resp.Body)
	result := map[string]interface{}{}
	json.Unmarshal(body, &result)
	return result["authToken"].(string)
}

func insertGuacamoleUser(username string, password string) *http.Response {
	var guacamoleUrl = os.Getenv("GuacamoleUrl")
	var guacamoleUsername = os.Getenv("GuacamoleUsername")
	resource := "/guacamole/api/session/data/mysql/users"
	token := "?token=" + getGuacamoleTokenString()
	guacamoleUser := GuacamoleUser{}
	guacamoleUser.Username = username
	guacamoleUser.Password = password
	guacamoleUserBytes, _ := json.Marshal(guacamoleUser)
	guacamoleUserBuff := bytes.NewBuffer(guacamoleUserBytes)
	log.Infof("paramsInfo = [%v]", guacamoleUser)
	u, _ := url.ParseRequestURI(guacamoleUrl)
	u.Path = resource
	urlStr := u.String() // "https://api.com/user/"
	client := &http.Client{
		Timeout: 5 * time.Second,
	}
	log.Infof("url [%v], username [%v]", urlStr+token, guacamoleUsername)
	r, err := http.NewRequest("POST", urlStr+token, guacamoleUserBuff) // URL-encoded payload
	if err != nil {
		log.Fatal(err)
	}
	r.Header.Add("Content-Type", "application/json")
	//r.Header.Add("Content-Length", strconv.Itoa(len(params.Encode())))
	//r.Header.Add("Content-Length", strconv.Itoa(len(params.Encode())))
	resp, err := client.Do(r)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	return resp
}

func deleteGuacamoleUser(username string) *http.Response {
	var guacamoleUrl = os.Getenv("GuacamoleUrl")
	var guacamoleUsername = os.Getenv("GuacamoleUsername")
	resource := "/guacamole/api/session/data/mysql/users/" + username
	token := "?token=" + getGuacamoleTokenString()
	log.Infof("paramsInfo = [%v]", username)
	u, _ := url.ParseRequestURI(guacamoleUrl)
	u.Path = resource
	urlStr := u.String() // "https://api.com/user/"
	client := &http.Client{
		Timeout: 5 * time.Second,
	}
	log.Infof("url [%v], username [%v]", urlStr+token, guacamoleUsername)
	r, err := http.NewRequest("DELETE", urlStr+token, nil) // URL-encoded payload
	if err != nil {
		log.Fatal(err)
	}
	r.Header.Add("Content-Type", "application/json")
	resp, err := client.Do(r)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	return resp
}

func insertGuacamoleVMAllocateUser(username string, instanceUuid string) *http.Response {
	var guacamoleUrl = os.Getenv("GuacamoleUrl")
	var guacamoleUsername = os.Getenv("GuacamoleUsername")
	resource := "/guacamole/api/session/data/mysql/connections"
	token := "?token=" + getGuacamoleTokenString()
	instanceList, _ := selectInstanceList(instanceUuid, InstanceString)
	instanceInfo := instanceList[0]
	guacamoleConnections := GuacamoleConnections{}
	//guacamoleConnections.Parameters.ReadOnly = ""
	//guacamoleConnections.Parameters.SwapRedBlue = ""
	//guacamoleConnections.Parameters.Cursor = ""
	//guacamoleConnections.Parameters.ColorDepth = ""
	//guacamoleConnections.Parameters.ClipboardEncoding = ""
	//guacamoleConnections.Parameters.DisableCopy = ""
	//guacamoleConnections.Parameters.DisablePaste = ""
	//guacamoleConnections.Parameters.DestPort = ""
	//guacamoleConnections.Parameters.RecordingExcludeOutput = ""
	//guacamoleConnections.Parameters.RecordingExcludeMouse = ""
	//guacamoleConnections.Parameters.RecordingIncludeKeys = ""
	//guacamoleConnections.Parameters.CreateRecordingPath = ""
	//guacamoleConnections.Parameters.EnableSftp = ""
	//guacamoleConnections.Parameters.SftpPort = ""
	//guacamoleConnections.Parameters.SftpServerAliveInterval = ""
	//guacamoleConnections.Parameters.EnableAudio = ""
	//guacamoleConnections.Parameters.Security = ""
	//guacamoleConnections.Parameters.DisableAuth = ""
	//guacamoleConnections.Parameters.GatewayPort = ""
	//guacamoleConnections.Parameters.ServerLayout = ""
	//guacamoleConnections.Parameters.Timezone = ""
	//guacamoleConnections.Parameters.Console = ""
	//guacamoleConnections.Parameters.Width = ""
	//guacamoleConnections.Parameters.Height = ""
	//guacamoleConnections.Parameters.Dpi = ""
	//guacamoleConnections.Parameters.ResizeMethod = ""
	//guacamoleConnections.Parameters.ConsoleAudio = ""
	//guacamoleConnections.Parameters.DisableAudio = ""
	//guacamoleConnections.Parameters.EnableAudioInput = ""
	//guacamoleConnections.Parameters.EnablePrinting = ""
	//guacamoleConnections.Parameters.EnableDrive = ""
	//guacamoleConnections.Parameters.CreateDrivePath = ""
	//guacamoleConnections.Parameters.EnableWallpaper = ""
	//guacamoleConnections.Parameters.EnableTheming = ""
	//guacamoleConnections.Parameters.EnableFontSmoothing = ""
	//guacamoleConnections.Parameters.EnableFullWindowDrag = ""
	//guacamoleConnections.Parameters.EnableDesktopComposition = ""
	//guacamoleConnections.Parameters.EnableMenuAnimations = ""
	//guacamoleConnections.Parameters.DisableBitmapCaching = ""
	//guacamoleConnections.Parameters.DisableOffscreenCaching = ""
	//guacamoleConnections.Parameters.DisableGlyphCaching = ""
	//guacamoleConnections.Parameters.PreconnectionId = ""
	//guacamoleConnections.Parameters.GatewayHostname = ""
	//guacamoleConnections.Parameters.GatewayUsername = ""
	//guacamoleConnections.Parameters.GatewayPassword = ""
	//guacamoleConnections.Parameters.GatewayDomain = ""
	//guacamoleConnections.Parameters.InitialProgram = ""
	//guacamoleConnections.Parameters.ClientName = ""
	//guacamoleConnections.Parameters.PrinterName = ""
	//guacamoleConnections.Parameters.StaticChannels = ""
	//guacamoleConnections.Parameters.RemoteApp = ""
	//guacamoleConnections.Parameters.RemoteAppDir = ""
	//guacamoleConnections.Parameters.RemoteAppArgs = ""
	//guacamoleConnections.Parameters.PreconnectionBlob = ""
	//guacamoleConnections.Parameters.LoadBalanceInfo = ""
	//guacamoleConnections.Parameters.RecordingPath = ""
	//guacamoleConnections.Parameters.RecordingName = ""
	//guacamoleConnections.Parameters.SftpHostname = ""
	//guacamoleConnections.Parameters.SftpHostKey = ""
	//guacamoleConnections.Parameters.SftpUsername = ""
	//guacamoleConnections.Parameters.SftpPassword = ""
	//guacamoleConnections.Parameters.SftpPrivateKey = ""
	//guacamoleConnections.Parameters.SftpPassphrase = ""
	//guacamoleConnections.Parameters.SftpRootDirectory = ""
	//guacamoleConnections.Parameters.SftpDirectory = ""
	//guacamoleConnections.Attributes.MaxConnections = ""
	//guacamoleConnections.Attributes.MaxConnectionsPerUser = ""
	//guacamoleConnections.Attributes.Weight = ""
	//guacamoleConnections.Attributes.FailoverOnly = ""
	//guacamoleConnections.Attributes.GuacdPort = ""
	//guacamoleConnections.Attributes.GuacdEncryption = ""
	//guacamoleConnections.Attributes.GuacdHostname = ""

	// guacamoleConnections 고정 설정값
	guacamoleConnections.ParentIdentifier = "ROOT"
	guacamoleConnections.Protocol = "rdp"
	guacamoleConnections.Parameters.Port = "3389"
	guacamoleConnections.Parameters.IgnoreCert = "true"
	guacamoleConnections.Parameters.EnableDrive = "true"
	guacamoleConnections.Parameters.CreateDrivePath = "true"
	guacamoleConnections.Parameters.DriveName = "G"
	guacamoleConnections.Parameters.DrivePath = "/share"
	// guacamoleConnections 고정 설정값

	params := []MoldParams{
		{"id": instanceInfo.MoldUuid},
	}
	resultVMInfo := getListVirtualMachinesMetrics(params)
	listVirtualMachinesMetrics := ListVirtualMachinesMetrics{}
	virtualMachineInfo, _ := json.Marshal(resultVMInfo["listvirtualmachinesmetricsresponse"])
	json.Unmarshal([]byte(virtualMachineInfo), &listVirtualMachinesMetrics)

	resultUserInfo := selectUserDBDetail(username)

	guacamoleConnections.Name = instanceInfo.Name
	guacamoleConnections.Parameters.Hostname = listVirtualMachinesMetrics.Virtualmachine[0].Nic[0].Ipaddress
	guacamoleConnections.Parameters.Username = username
	guacamoleConnections.Parameters.Password = resultUserInfo.Password
	guacamoleConnections.Parameters.Domain = os.Getenv("SambaDomain")

	guacamoleConnectionsBytes, _ := json.Marshal(guacamoleConnections)
	//guacamoleConnectionsBytes, _ := json.Marshal("{\n  \"parentIdentifier\": \"ROOT\",\n  \"name\": \"tes1t\",\n  \"protocol\": \"rdp\",\n  \"parameters\": {\n    \"port\": \"\",\n    \"read-only\": \"\",\n    \"swap-red-blue\": \"\",\n    \"cursor\": \"\",\n    \"color-depth\": \"\",\n    \"clipboard-encoding\": \"\",\n    \"disable-copy\": \"\",\n    \"disable-paste\": \"\",\n    \"dest-port\": \"\",\n    \"recording-exclude-output\": \"\",\n    \"recording-exclude-mouse\": \"\",\n    \"recording-include-keys\": \"\",\n    \"create-recording-path\": \"\",\n    \"enable-sftp\": \"\",\n    \"sftp-port\": \"\",\n    \"sftp-server-alive-interval\": \"\",\n    \"enable-audio\": \"\",\n    \"security\": \"\",\n    \"disable-auth\": \"\",\n    \"ignore-cert\": \"\",\n    \"gateway-port\": \"\",\n    \"server-layout\": \"\",\n    \"timezone\": \"\",\n    \"console\": \"\",\n    \"width\": \"\",\n    \"height\": \"\",\n    \"dpi\": \"\",\n    \"resize-method\": \"\",\n    \"console-audio\": \"\",\n    \"disable-audio\": \"\",\n    \"enable-audio-input\": \"\",\n    \"enable-printing\": \"\",\n    \"enable-drive\": \"\",\n    \"create-drive-path\": \"\",\n    \"enable-wallpaper\": \"\",\n    \"enable-theming\": \"\",\n    \"enable-font-smoothing\": \"\",\n    \"enable-full-window-drag\": \"\",\n    \"enable-desktop-composition\": \"\",\n    \"enable-menu-animations\": \"\",\n    \"disable-bitmap-caching\": \"\",\n    \"disable-offscreen-caching\": \"\",\n    \"disable-glyph-caching\": \"\",\n    \"preconnection-id\": \"\",\n    \"hostname\": \"\",\n    \"username\": \"\",\n    \"password\": \"\",\n    \"domain\": \"\",\n    \"gateway-hostname\": \"\",\n    \"gateway-username\": \"\",\n    \"gateway-password\": \"\",\n    \"gateway-domain\": \"\",\n    \"initial-program\": \"\",\n    \"client-name\": \"\",\n    \"printer-name\": \"\",\n    \"drive-name\": \"\",\n    \"drive-path\": \"\",\n    \"static-channels\": \"\",\n    \"remote-app\": \"\",\n    \"remote-app-dir\": \"\",\n    \"remote-app-args\": \"\",\n    \"preconnection-blob\": \"\",\n    \"load-balance-info\": \"\",\n    \"recording-path\": \"\",\n    \"recording-name\": \"\",\n    \"sftp-hostname\": \"\",\n    \"sftp-host-key\": \"\",\n    \"sftp-username\": \"\",\n    \"sftp-password\": \"\",\n    \"sftp-private-key\": \"\",\n    \"sftp-passphrase\": \"\",\n    \"sftp-root-directory\": \"\",\n    \"sftp-directory\": \"\"\n  },\n  \"attributes\": {\n    \"max-connections\": \"\",\n    \"max-connections-per-user\": \"\",\n    \"weight\": \"\",\n    \"failover-only\": \"\",\n    \"guacd-port\": \"\",\n    \"guacd-encryption\": \"\",\n    \"guacd-hostname\": \"\"\n  }\n}")
	guacamoleUserBuff := bytes.NewBuffer(guacamoleConnectionsBytes)
	log.Infof("paramsInfo = [%v]", guacamoleConnections)
	u, _ := url.ParseRequestURI(guacamoleUrl)
	u.Path = resource
	urlStr := u.String() // "https://api.com/user/"
	client := &http.Client{
		Timeout: 5 * time.Second,
	}
	log.Infof("url [%v], username [%v]", urlStr+token, guacamoleUsername)
	log.Infof("guacamoleConnections [%v]", guacamoleConnections)
	r, err := http.NewRequest("POST", urlStr+token, guacamoleUserBuff) // URL-encoded payload
	if err != nil {
		log.Fatal(err)
	}
	r.Header.Add("Content-Type", "application/json")
	resp, err := client.Do(r)
	if err != nil {
		log.Fatal(err)
	}

	return resp
}
