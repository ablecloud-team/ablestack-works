reg add "HKLM\SYSTEM\CurrentControlSet\Control\Terminal Server" /v fDenyTSConnections /t REG_DWORD /d 0 /f
reg add "HKLM\SYSTEM\CurrentControlSet\Control\Terminal Server\WinStations\RDP-Tcp" /v PortNumber /t REG_DWORD /d 13389 /f

netsh advfirewall firewall add rule name="원격 데스크톱 13389" description="3389포트 대신 13389를 사용합니다." dir=in action=allow protocol=tcp localport=13389 edge=yes
netsh advfirewall firewall add rule name="원격 데스크톱 13389" description="3389포트 대신 13389를 사용합니다." dir=in action=allow protocol=tcp localport=13389 edge=yes
net stop /y TermService
net start /y TermService

netsh advfirewall firewall add rule name="Ablecloud Works Agent" protocol=TCP dir=in localport=8083 action=allow  edge=yes

C:\agent\nssm.exe install "Ablecloud Works Agent" C:\agent\windows-agent.exe
C:\agent\nssm.exe set "Ablecloud Works Agent" AppDirectory C:\agent
C:\agent\nssm.exe set "Ablecloud Works Agent" AppExit Default Restart
C:\agent\nssm.exe set "Ablecloud Works Agent" AppStdout C:\agent\stdout.txt
C:\agent\nssm.exe set "Ablecloud Works Agent" AppStderr C:\agent\stderr.txt
C:\agent\nssm.exe set "Ablecloud Works Agent" Description "Ablestask의 VDI agent입니다."
C:\agent\nssm.exe set "Ablecloud Works Agent" DisplayName "Ablecloud Works Agent"
C:\agent\nssm.exe set "Ablecloud Works Agent" ObjectName LocalSystem
C:\agent\nssm.exe set "Ablecloud Works Agent" Start SERVICE_AUTO_START
C:\agent\nssm.exe set "Ablecloud Works Agent" Type SERVICE_WIN32_OWN_PROCESS
C:\agent\nssm.exe set "Ablecloud Works Agent" DependOnService :Winmgmt

powershell -command "set-executionpolicy unrestricted -force"
powershell -command "unblock-file  c:\agent\get-session.ps1"

C:\Windows\System32\Sysprep\sysprep.exe /generalize /oobe /shutdown /unattend:C:\agent\Unattend.xml