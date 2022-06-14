FROM centos:8
MAINTAINER [ablecloud-team] Jin Seong CHOI


RUN cd /etc/yum.repos.d/
RUN sed -i 's/mirrorlist/#mirrorlist/g' /etc/yum.repos.d/CentOS-*
RUN sed -i 's|#baseurl=http://mirror.centos.org|baseurl=http://vault.centos.org|g' /etc/yum.repos.d/CentOS-*
RUN dnf -y install nginx

# RUN mkdir /usr/share/nginx/html/js
# RUN mkdir /usr/share/nginx/html/img
# RUN mkdir /usr/share/nginx/html/css

COPY ./nginx.conf /etc/nginx/nginx.conf
COPY ./dist/ /usr/share/nginx/html/

# COPY ./dist/index.html /usr/share/nginx/html/index.html
# COPY ./dist/favicon.ico /usr/share/nginx/html/favicon.ico
# COPY ./dist/js/* /usr/share/nginx/html/js
# COPY ./dist/img/* /usr/share/nginx/html/img
# COPY ./dist/css/*  /usr/share/nginx/html/css

EXPOSE 8080
CMD ["/usr/sbin/nginx", "-g", "daemon off;"]
