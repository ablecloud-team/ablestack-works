FROM centos:8
MAINTAINER Jin Seong CHOI

RUN dnf -y install nginx

#RUN mkdir /usr/share/nginx/html/js
#RUN mkdir /usr/share/nginx/html/img
#RUN mkdir /usr/share/nginx/html/css

COPY ./nginx.conf /etc/nginx/nginx.conf
#COPY ./dist/index.html /usr/share/nginx/html/index.html
#COPY ./dist/favicon.ico /usr/share/nginx/html/favicon.ico

#ADD ./dist/js/* /usr/share/nginx/html/js
#ADD ./dist/img/* /usr/share/nginx/html/img
#ADD ./dist/css/*  /usr/share/nginx/html/css

EXPOSE 8080
CMD ["/usr/sbin/nginx", "-g", "daemon off;"]
