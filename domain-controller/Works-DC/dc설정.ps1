#rsat 설치
#설정 -> 앱 -> 선택적 기능 -> 기능 추가 -> rsat: dns 서버 도구, rsat: 그룹 정책 관리 도구

#Add-WindowsCapability –online –Name Rsat.ActiveDirectory.DS-LDS.Tools~~~~0.0.1.0
#Add-WindowsCapability –online –Name Rsat.Dns.Tools~~~~0.0.1.0
#Add-WindowsCapability -Online -Name Rsat.FileServices.Tools~~~~0.0.1.0
#Add-WindowsCapability -Online -Name Rsat.GroupPolicy.Management.Tools~~~~0.0.1.0
#Add-WindowsCapability -Online -Name Rsat.IPAM.Client.Tools~~~~0.0.1.0
#Add-WindowsCapability -Online -Name Rsat.LLDP.Tools~~~~0.0.1.0
#Add-WindowsCapability -Online -Name Rsat.NetworkController.Tools~~~~0.0.1.0
#Add-WindowsCapability -Online -Name Rsat.NetworkLoadBalancing.Tools~~~~0.0.1.0
#Add-WindowsCapability -Online -Name Rsat.BitLocker.Recovery.Tools~~~~0.0.1.0
#Add-WindowsCapability -Online -Name Rsat.CertificateServices.Tools~~~~0.0.1.0
#Add-WindowsCapability -Online -Name Rsat.DHCP.Tools~~~~0.0.1.0
#Add-WindowsCapability -Online -Name Rsat.FailoverCluster.Management.Tools~~~~0.0.1.0
#Add-WindowsCapability -Online -Name Rsat.RemoteAccess.Management.Tools~~~~0.0.1.0
#Add-WindowsCapability -Online -Name Rsat.RemoteDesktop.Services.Tools~~~~0.0.1.0
#Add-WindowsCapability -Online -Name Rsat.ServerManager.Tools~~~~0.0.1.0
#Add-WindowsCapability -Online -Name Rsat.Shielded.VM.Tools~~~~0.0.1.0
#Add-WindowsCapability -Online -Name Rsat.StorageMigrationService.Management.Tools~~~~0.0.1.0
#Add-WindowsCapability -Online -Name Rsat.StorageReplica.Tools~~~~0.0.1.0
#Add-WindowsCapability -Online -Name Rsat.SystemInsights.Management.Tools~~~~0.0.1.0
#Add-WindowsCapability -Online -Name Rsat.VolumeActivation.Tools~~~~0.0.1.0
#Add-WindowsCapability -Online -Name Rsat.WSUS.Tools~~~~0.0.1.0


Add-WindowsCapability -Online -Name Rsat.ActiveDirectory.DS-LDS.Tools~~~~0.0.1.0
Add-WindowsCapability -Online -Name Rsat.Dns.Tools~~~~0.0.1.0
Add-WindowsCapability -Online -Name Rsat.GroupPolicy.Management.Tools~~~~0.0.1.0

netsh advfirewall firewall add rule name="Ablecloud Works DC" protocol=TCP dir=in localport=8083 action=allow  edge=yes
C:\Works-DC\nssm.exe install Works-DC C:\Works-DC\Abledc.exe
C:\Works-DC\nssm.exe set Works-DC AppDirectory C:\Works-DC
C:\Works-DC\nssm.exe set Works-DC AppExit Default Restart
C:\Works-DC\nssm.exe set Works-DC AppStdout C:\Works-DC\stdout.txt
C:\Works-DC\nssm.exe set Works-DC AppStderr C:\Works-DC\stderr.txt
C:\Works-DC\nssm.exe set Works-DC Description "Ablestask의 DC api server입니다."
C:\Works-DC\nssm.exe set Works-DC DisplayName Works-DC
C:\Works-DC\nssm.exe set Works-DC ObjectName LocalSystem
C:\Works-DC\nssm.exe set Works-DC Start SERVICE_AUTO_START
C:\Works-DC\nssm.exe set Works-DC Type SERVICE_WIN32_OWN_PROCESS