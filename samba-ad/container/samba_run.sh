#!/bin/bash

nohup /usr/local/samba/sbin/samba -i -d &>> /var/log/samba.log&
/usr/bin/samba-ad &>> /var/log/samba-ad.log