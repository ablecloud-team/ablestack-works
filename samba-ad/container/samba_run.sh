#!/bin/bash

nohup /usr/local/samba/sbin/samba -i >> /var/log/samba.log 2>> /usr/local/samba/var/samba.err&
nohup /usr/bin/samba-ad >> /var/log/samba-ad.log 2>> /usr/local/samba/var/samba-ad.err
