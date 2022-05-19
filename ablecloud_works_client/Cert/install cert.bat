set CURPATH=%cd%
cd c:\
cd c:
certutil -addstore root %CURPATH%\ablecloud.cer
rmdir %CURPATH% /s /q
pause