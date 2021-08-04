#!/bin/sh
# Author  Achu Abebe
# Email:  Achusime@gmail.com

# # You should run this with root privilage


#Variables
FQDN=dc1.local
DN=mydomain
NSE=8.8.8.8 #NS externe/forwarder
APW=Ablecloud1!
SHARE=/mnt/scripts/
IP=10.1.1.8
NM=255.255.255.0
NW=10.1.1.0
BC=10.1.1.255
GW=10.1.1.1
NSI=10.1.1.1 #NS interne
# NSI2= X.X.X.X
NSE=8.8.8.8 #NS externe/forwarder
# NSE2= X.X.X.X
FQDN=dc1.local
HN=mydomain
ETH=eth0

echo "$IP $HN"."$FQDN $HN" >> /etc/hosts
# Configuration hostname file
echo $HN'.'$FQDN > /etc/hostname

#provision of field
/usr/local/samba/bin/samba-tool domain provision --realm=$FQDN --domain=$DN --adminpass="$APW" --server-role=dc --dns-backend=SAMBA_INTERNAL
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
echo nameserver $IP  >> /etc/resolv.conf
sed -i "s/.*dns forwarder =.*/	dns forwarder = $NSE/g" /usr/local/samba/etc/smb.conf # positioning forwarder
echo ====================DNS TEST==========================
#verification of the operation of DNS
host -t SRV _ldap._tcp.$FQDN.
host -t SRV _kerberos._udp.$FQDN.
host -t A $FQDN.
sleep 5
echo ==================KERBEROS CONFIGURATION==============
# # deKerberos configuration
Newvariable=$( echo "$FQDN" | tr -s  '[:lower:]'  '[:upper:]' )
sed -i "s/\${REALM\([^}]*\)}"/$Newvariable"/g" /usr/local/samba/share/setup/krb5.conf
# #verification of operation of Kerberos
kinit administrator@$Newvariable
klist -e
