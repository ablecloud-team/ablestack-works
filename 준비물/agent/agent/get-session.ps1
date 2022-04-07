#$user="ycyun"
$server="localhost"
$logons = gwmi win32_loggedonuser
$flag = $false
Write-host "["
foreach ($logon in $logons){
    $username=$logon.antecedent.Split(",")[1].Split('"')[1]
    $domain=$logon.antecedent.Split(",")[0].Split('"')[1]
    #if ($domain -eq "YCY"){
        $logonid = $logon.dependent.split('"')[1]#$logon.dependent.split("=")[1]
        $session = (gwmi win32_logonsession | where-object LogonId -EQ $logonid)
        if ($session.LogonType -eq 3){
            if ($flag){
                Write-host ","
            }
            $time = $session.StartTime.Split(".")[0]
            #$time = "20211228172755"
            $year = $time.Substring(0,4)
            $month = $time.Substring(4,2)
            $day = $time.Substring(6,2)
            $hour = $time.Substring(8,2)
            $min = $time.Substring(10,2)
            $sec = $time.Substring(12,2)
            Write-host "  {"
            Write-host "    ""id"": ""$username"","
            Write-host "    ""domain"": ""$domain"","
            #Write-host " $($session.LogonType),"
            Write-host "    ""time"": ""$time"","
            Write-host "    ""year"": ""$year"","
            Write-host "    ""month"": ""$month"","
            Write-host "    ""day"": ""$day"","
            Write-host "    ""hour"": ""$hour"","
            Write-host "    ""min"": ""$min"","
            Write-host "    ""sec"": ""$sec"","
            Write-host "    ""logonid"": ""$($logonid)"""
            Write-host "  }" -NoNewline
            $flag = $true
            #echo $session
            #echo $logon
        }
    #}
 }
 Write-host
 Write-host "]"