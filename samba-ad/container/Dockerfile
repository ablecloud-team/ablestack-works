FROM docker.io/library/ubuntu:20.04
LABEL MAINTAINER="YeoCheon Yun <ycyun@ablecloud.io>"

ENV TERM=xterm-color
ENV DEBIAN_FRONTEND=noninteractive
RUN apt-get -y update

RUN apt-get install -y \
    acl \
    apt-utils \
    attr \
    autoconf \
    bind9utils \
    binutils \
    bison \
    build-essential \
    ccache \
    chrpath \
    curl \
    debhelper \
    dialog \
    dnsutils \
    docbook-xml \
    docbook-xsl \
    flex \
    gcc \
    gdb \
    git \
    glusterfs-common \
    gzip \
    heimdal-multidev \
    hostname \
    htop \
    iproute2 \
    krb5-config \
    krb5-kdc \
    krb5-user \
    language-pack-en \
    lcov \
    libacl1-dev \
    libarchive-dev \
    libattr1-dev \
    libavahi-common-dev \
    libblkid-dev \
    libbsd-dev \
    libcap-dev \
    libcephfs-dev \
    libcups2-dev \
    libdbus-1-dev \
    libglib2.0-dev \
    libgnutls28-dev \
    libgpgme11-dev \
    libicu-dev \
    libjansson-dev \
    libjs-jquery \
    libjson-perl \
    libkrb5-dev \
    libldap2-dev \
    liblmdb-dev \
    libncurses5-dev \
    libpam0g-dev \
    libparse-yapp-perl \
    libpcap-dev \
    libpopt-dev \
    libreadline-dev \
    libsystemd-dev \
    libtasn1-bin \
    libtasn1-dev \
    libtracker-sparql-2.0-dev \
    libunwind-dev \
    lmdb-utils \
    locales \
    lsb-release \
    make \
    mawk \
    mingw-w64 \
    net-tools \
    ntp \
    patch \
    perl \
    perl-modules \
    pkg-config \
    procps \
    psmisc \
    python-dev \
    python-dnspython \
    python3 \
    python3-cryptography \
    python3-dbg \
    python3-dev \
    python3-dnspython \
    python3-gpg \
    python3-iso8601 \
    python3-markdown \
    python3-matplotlib \
    python3-pexpect \
    python3-pip \
    python3-pyasn1 \
    python3-setproctitle \
    rng-tools \
    rsync \
    sed \
    sudo \
    tar \
    tree \
    uuid-dev \
    vim \
    wget \
    xfslibs-dev \
    xsltproc \
    zlib1g-dev


RUN apt-get update && apt-get upgrade -y

RUN apt-get -y autoremove
RUN apt-get -y autoclean
RUN apt-get -y clean

# RUN pip3 install

RUN git clone -b v4-14-stable https://git.samba.org/samba.git/ samba4
WORKDIR samba4
RUN git clean -x -f -d # Disposal of obsolete files eventually
RUN ./configure --enable-debug --enable-selftest #Configuration
RUN make -j4 # Compilation
RUN make install # Installation

COPY samba-ad /usr/bin/
COPY entrypoint.py /usr/local/bin/
COPY samba_run.sh /usr/local/bin/
COPY samba_config.sh /usr/local/bin/
#   BINDIR: /usr/local/samba/bin
#   SBINDIR: /usr/local/samba/sbin
#   CONFIGFILE: /usr/local/samba/etc/smb.conf
#   NCALRPCDIR: /usr/local/samba/var/run/ncalrpc
#   LOGFILEBASE: /usr/local/samba/var
#   LMHOSTSFILE: /usr/local/samba/etc/lmhosts
#   DATADIR: /usr/local/samba/share
#   MODULESDIR: /usr/local/samba/lib
#   LOCKDIR: /usr/local/samba/var/lock
#   STATEDIR: /usr/local/samba/var/locks
#   CACHEDIR: /usr/local/samba/var/cache
#   PIDDIR: /usr/local/samba/var/run
#   PRIVATE_DIR: /usr/local/samba/private
#   CODEPAGEDIR: /usr/local/samba/share/codepages
#   SETUPDIR: /usr/local/samba/share/setup
#   WINBINDD_SOCKET_DIR: /usr/local/samba/var/run/winbindd
#   NTP_SIGND_SOCKET_DIR: /usr/local/samba/var/lib/ntp_signd
ENV PATH $PATH:/usr/local/samba/bin:/usr/local/samba/sbin
ENV ETH eth0
# VOLUME ["/etc/samba", "/var/lib/samba", "/var/run/samba"]
VOLUME ["/usr/local/samba/etc/", "/usr/local/samba/var/run","/usr/local/samba/var", "/usr/local/samba/share"]
EXPOSE 22 53 389 88 135 139 138 445 464 3268 3269
# CMD ["app:start"]
ENTRYPOINT "/bin/bash"
