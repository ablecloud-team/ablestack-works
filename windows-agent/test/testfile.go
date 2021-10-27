package main

import "fmt"

func main() {
	SambaDomain := "testdomain.local"
	MyDomain := "testdomain"
	VmName := "www992"
	SambaIP := "10.1.1.59"
	WorksIP := "10.1.1.59"
	WorksPort := "8083"
	Type := "server"
	InstanceUuid := "3713d4c3-48f7-4b83-b892-d9bb2335d15d"

	payload := "<powershell>\n" +
		"Invoke-WebRequest -Uri http://10.1.1.1/latest/user-data >> \"c:\\agent\\userdata.txt\"\n" +
		"date > \"c:\\agent\\installed.txt\"\n" +
		"$dnsip = \"" + SambaIP + "\"\n" +
		"echo $dnsip >> \"c:\\agent\\installed.txt\"\n" +
		"set-DnsClientServerAddress -InterfaceIndex (Get-NetAdapter |Select-Object InterfaceAlias, interfaceindex).interfaceindex -ServerAddresses " + WorksIP + "\n" +
		"$password = \"Ablecloud1!\" | ConvertTo-SecureString -asPlainText -Force\n" +
		"$username = \"$" + MyDomain + "\\administrator\" \n" +
		"$credential = New-Object System.Management.Automation.PSCredential($username,$password)\n" +
		"echo Rename-Computer >> \"c:\\agent\\installed.txt\"\n" +
		"Rename-Computer -NewName " + VmName + "\n" +
		"echo Add-Computer >> \"c:\\agent\\installed.txt\"\n" +
		"Add-Computer -DomainName " + SambaDomain + " -Credential $credential -NewName " + VmName + "\n" +
		"echo Add-Computer end>> \"c:\\agent\\installed.txt\"\n" +
		"$conf = '{\"WorksServer\": \"" + WorksIP + "\", \"WorksPort\": " + WorksPort + ", \"Type\": \"" + Type + "\", \"UUID\": \"" + InstanceUuid + "\"}'\n" +
		"echo $conf| Out-File -Encoding ascii \"c:\\agent\\conf.json\"\n" +
		"echo $conf\n" +
		"echo $conf >> \"c:\\agent\\installed.txt\"\n" +
		"C:\\agent\\nssm.exe set \"Ablecloud Works Agent\" start SERVICE_DELAYED_AUTO_START\n" +
		"C:\\agent\\nssm.exe restart \"Ablecloud Works Agent\"\n" +
		"date >> \"c:\\agent\\installed.txt\"\n" +
		"</powershell>"
	fmt.Println(payload)
}
