#!/bin/sh
#Author  YeoCheon Yun
#Email:  ycyun@ablecloud.io

10.10.1.100

/samba4 경로에 guacSchema1.ldif, guacSchema2.ldif,  guacSchema3.ldif

1. Samba 스키마 파일 생성 (/samba4 경로 guacSchema1.ldif, guacSchema2.ldif,  guacSchema3.ldif)

**** guacSchema1.ldif



**** guacSchema2.ldif




**** guacSchema3.ldif



2. Samba 스키마 파일 추가

#/usr/local/samba/bin/ldbadd -H /usr/local/samba/private/sam.ldb --option="dsdb:schema update allowed"=true /samba4/guacSchema1.ldif
#/usr/local/samba/bin/ldbadd -H /usr/local/samba/private/sam.ldb --option="dsdb:schema update allowed"=true /samba4/guacSchema2.ldif
#/usr/local/samba/bin/ldbadd -H /usr/local/samba/private/sam.ldb --option="dsdb:schema update allowed"=true /samba4/guacSchema3.ldif


podman run -d -t --net host --name samba localhost/smb:test-09 config <dnsip(ex: 8.8.8.8)> <domain name(ex: mydomain)> <administrator password(ex: Ablecloud1!)> <adserver hostname(ex: ad)>

podman run -d -t --net host --name samba localhost/smb:flat config 8.8.8.8 testdomain Ablecloud1 ad3

podman run -d -t --net host --name samba localhost/works-ad:v3 config 8.8.8.8 mydomain Ablecloud1! ad