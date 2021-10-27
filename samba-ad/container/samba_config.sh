#!/bin/sh
# Author  YeoCheon Yun
# Email:  ycyun@ablecloud.io

#Variables
#IP=10.1.1.166
#NS=8.8.8.8 #NS externe/forwarder
#DN=testdomain

#APW=Ablecloud1!
#HN=ad
# NM=255.255.255.0
# NW=10.10.1.0
# GW=10.10.1.1
# NSI=8.8.8.8 #NS interne
# NSI2= X.X.X.X
# NSE2= X.X.X.X

IP=$(ip a|grep inet | grep -v inet6 | grep -v host | awk '{print $2}' | cut -d / -f 1)
NS=$1
DN=$2
APW=$3
HN=$4
FQDN=$DN.local

echo "$IP $HN.$FQDN $HN" >> /etc/hosts
# Configuration hostname file
echo "$HN.$FQDN" > /etc/hostname

#provision of field
/usr/local/samba/bin/samba-tool domain provision --realm=$FQDN --domain="$DN" --adminpass="$APW" --server-role=dc --dns-backend=SAMBA_INTERNAL
#execution samba
/usr/local/samba/sbin/samba
#verification of versions samba and smbclient
/usr/local/samba/sbin/samba -V
/usr/local/samba/bin/smbclient -V
sleep 2
#Listing list of Samba users
/usr/local/samba/bin/smbclient -L localhost -U%
sleep 2
#verification of how authentication by connecting to the NETLOGON share with the administrator account
/usr/local/samba/bin/smbclient //localhost/netlogon -UAdministrator%"$APW" -c 'ls'

echo ================DNS CONFIGURATION SAMBA_INTERNAL=====================
# Configuration of Resolv.conf
echo domain $FQDN  > /etc/resolv.conf
echo nameserver "$IP"  >> /etc/resolv.conf
sed -i "s/.*dns forwarder =.*/	dns forwarder = $NS/g" /usr/local/samba/etc/smb.conf # positioning forwarder
sed -i "/workgroup/ a\\\tldap server require strong auth = no" /usr/local/samba/etc/smb.conf
#"        ldap server require strong auth = no"
echo ====================DNS TEST==========================
#verification of the operation of DNS
host -t SRV _ldap._tcp.$FQDN.
host -t SRV _kerberos._udp.$FQDN.
host -t A $FQDN.
sleep 5
echo ==================KERBEROS CONFIGURATION==============
# # deKerberos configuration
Newvariable=$( echo "$FQDN" | tr -s  '[:lower:]'  '[:upper:]' )
sed -i "s/\${REALM\([^}]*\)}/$Newvariable/g" /usr/local/samba/share/setup/krb5.conf
sed -i "s/DOMAINNAME/$DN/g" /opt/able-ad/*.ldif
/usr/local/samba/bin/ldbadd -H /usr/local/samba/private/sam.ldb --option="dsdb:schema update allowed"=true /opt/able-ad/guacSchema1.ldif
/usr/local/samba/bin/ldbadd -H /usr/local/samba/private/sam.ldb --option="dsdb:schema update allowed"=true /opt/able-ad/guacSchema2.ldif
/usr/local/samba/bin/ldbadd -H /usr/local/samba/private/sam.ldb --option="dsdb:schema update allowed"=true /opt/able-ad/guacSchema3.ldif
date > /usr/local/samba/etc/installed
killall samba
# #verification of operation of Kerberos
# kinit administrator@"$Newvariable"
# klist -e
