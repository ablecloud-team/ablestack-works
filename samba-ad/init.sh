#!/bin/sh
#Author  Achu Abebe
#Email:  Achusime@gmail.com


# # You should run this with root privilage.
# Variables
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




echo ================NETWORK CONFIGURATION=======================
# Configuring the hosts file
# sed -i  "s/127.0.1.1	$HN"/"$IP	$HN"."$FQDN	$HN"/"" /etc/hosts
echo "$IP $HN"."$FQDN $HN" >> /etc/hosts
# Configuration hostname file
echo $HN'.'$FQDN > /etc/hostname

sleep 10
echo ================PREREQUISITES SETUP=====================
# mise a jour
apt-get update && apt-get upgrade -y
# Installing the necessary packages
apt-get install dialog git build-essential libacl1-dev libattr1-dev libblkid-dev net-tools iproute2 libreadline-dev python-dev python-dnspython gdb pkg-config libpopt-dev libldap2-dev dnsutils libbsd-dev attr krb5-user docbook-xsl libcups2-dev libpam0g-dev ntp -y

echo ================SAMBA4 DOWNLOAD AND SETUP=====================
cd /
# Download Samba from git
git clone -b v4-14-stable https://git.samba.org/samba.git/ samba4
cd samba4
git clean -x -f -d # Disposal of obsolete files eventually
./configure --enable-debug --enable-selftest #Configuration
make # Compilation
make install # Installation
echo ========================PLEASE REBOOT NOW and Have some Ethiopian Coffee==========================
exit

