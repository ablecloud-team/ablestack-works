# ABLESTACK Desktop
 ABLESTACK HCI를 기반으로 Desktop as a Service를 제공하기 위한 플랫폼

## ABLESTACK Works Template 구성방법

### Works Service 등록  

  1. **/** 에 **works** 폴더 생성 후 **works-shell** 파일 복사  

  ![Screenshot_2022-03-17 18 12 28_BZI8Lu](https://user-images.githubusercontent.com/5581961/158776681-0a54afdb-980b-4627-af5a-29cc81fe9522.png)  

  2. **/usr/lib/systemd/system** 폴더에 **works.service** 파일 복사

  ![Screenshot_2022-03-17 18 13 05_iw1qgC](https://user-images.githubusercontent.com/5581961/158776815-e4ea3956-0c2d-420e-b967-2182b6b5e65d.png)

  3. **systemctl enable works.service** 명령어를 이용하여 **works.service** 등록 및 시작 

### Works Container 등록  

 1. **podman image load -i gucamole.tar** 명령어를 이용하여 container 등록
 2. **podman image load -i works-db.tar** 명령어를 이용하여 container 등록
 3. **podman image load -i works-api.tar** 명령어를 이용하여 container 등록
 4. **podman image load -i works-ui.tar** 명령어를 이용하여 container 등록
 5. **podman image load -i works-samba.tar** 명령어를 이용하여 container 등록
