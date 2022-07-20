#!/bin/bash

podman network create --subnet 10.88.2.0/24 works
podman volume create works-db

podman pull docker.io/library/mysql:latest
sleep 5
podman run --name works-db --net works --ip 10.88.2.12 -p 3306:3306 -v works-db:/var/lib/mysql -e MYSQL_ROOT_PASSWORD=Ablecloud1! -d mysql:latest

podman pull docker.io/ablecloudteam/works-ad:latest && \
podman pull docker.io/ablecloudteam/works-api:latest && \
podman pull docker.io/ablecloudteam/works-lite:latest && \
podman pull docker.io/ablecloudteam/works-ui:latest && \
podman pull docker.io/ablecloudteam/works-ad:latest && \
podman pull docker.io/guacamole/guacd:latest


podman run --name works-api --net works --ip 10.88.2.13 -p 8082:8080 -d docker.io/ablecloudteam/works-api:latest
podman run --name works-ui --net works --ip 10.88.2.11 -p 8081:8080 -d docker.io/ablecloudteam/works-ui:latest
podman run --name works-lite --net works --ip 10.88.2.15 -p 8088:8080 -d docker.io/ablecloudteam/works-lite:latest
podman run --name works-guacd --net works --ip 10.88.2.14 -d docker.io/guacamole/guacd:latest

sleep 5
echo 'CREATE DATABASE works start'
#podman exec -it works-db mysql -uroot -pAblecloud1! -e 'CREATE DATABASE works'
podman exec works-db sh -c "mysql -uroot -pAblecloud1! -e 'CREATE DATABASE works'"
echo 'CREATE DATABASE works end'

sleep 5
echo 'CREATE USER "works" IDENTIFIED BY "Ablecloud1!" start'
#podman exec -it works-db mysql -uroot -pAblecloud1! -e 'CREATE USER "works" IDENTIFIED BY "Ablecloud1!"'
podman exec works-db sh -c "mysql -uroot -pAblecloud1! -e 'CREATE USER \"works\" IDENTIFIED BY \"Ablecloud1!\"'"
echo 'CREATE USER "works" IDENTIFIED BY "Ablecloud1!" end'
sleep 5
echo 'GRANT SELECT,INSERT,UPDATE,DELETE ON works.* TO "works" start'
#podman exec -it works-db mysql -uroot -pAblecloud1! -e 'GRANT SELECT,INSERT,UPDATE,DELETE ON works.* TO "works"'
podman exec works-db sh -c "mysql -uroot -pAblecloud1! -e 'GRANT SELECT,INSERT,UPDATE,DELETE ON works.* TO \"works\"'"
echo 'GRANT SELECT,INSERT,UPDATE,DELETE ON works.* TO "works" end'
sleep 5
echo 'FLUSH PRIVILEGES start'
#podman exec -it works-db mysql -uroot -pAblecloud1! -e 'FLUSH PRIVILEGES'
podman exec works-db sh -c "mysql -uroot -pAblecloud1! -e 'FLUSH PRIVILEGES'"
echo 'FLUSH PRIVILEGES end'

podman cp /root/works_schema.sql works-db:/mnt/works_schema.sql
podman exec works-db sh -c "mysql -uroot -pAblecloud1! works < /mnt/works_schema.sql"