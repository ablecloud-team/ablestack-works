create table application
(
    id                 int auto_increment
        primary key,
    uuid               varchar(40)             not null comment 'application Uuid',
    workspace_uuid     varchar(40) default '1' not null,
    application_server varchar(40)             null comment 'application server UUID',
    path               varchar(100)            not null comment 'application 경로',
    description        varchar(100)            null comment 'application 설명',
    working_directory  varchar(100)            null comment 'application working_directory',
    create_date        datetime                not null comment 'applicatin 등록일',
    removed            datetime                null comment 'application 삭제일'
);

create table async_job
(
    id          varchar(40)       not null comment 'async_job UUID'
        primary key,
    name        varchar(40)       not null comment 'async_job 이름(async_job 실행 func)
상태값 : VMStart, VMStop, VMDeploy, VMDestroy',
    exec_uuid   varchar(40)       not null comment 'async_job 실행 func UUID',
    ready       tinyint default 0 not null comment 'async_job 실행 준비 여부',
    parameter   varchar(100)      null comment 'func 실행을 위한 parameter',
    create_date datetime          not null comment 'async_job 생성일'
)
    comment 'Works async 실행을 위한 테이블
해당 테이블은 VM 의 시작, 정지, 생성, 삭제를 등록하여 비동기 실행을 휘한 테이블 입니다.';

create table configuration
(
    id            int auto_increment
        primary key,
    name          varchar(100)                    not null comment '설정 이름',
    category      varchar(100) default 'Advanced' not null comment '설정 카테고리',
    value         varchar(256)                    not null comment '설정의 값',
    default_value varchar(256)                    null comment '설정 기본 값',
    description   varchar(1024)                   not null comment '설정 설명',
    update_date   datetime                        null comment '설정 업데이트 일자',
    input_type    varchar(20)  default 'input'    not null comment '설정 입력값 타입',
    input_value   varchar(100)                    null comment '설정 입력 값'
)
    comment 'works API 설정값';

create table guacamole_configuration
(
    id int auto_increment
        primary key
);

create table policy_list
(
    id             int auto_increment
        primary key,
    name           varchar(40)              not null comment '정책 이름',
    description    varchar(140)             not null comment '정책 설명',
    i18n           varchar(100)             not null comment '정책 i18n 값',
    value          varchar(100) default '0' null comment '정책 기본값',
    type           varchar(2)   default 'R' not null comment '정책적용 type',
    settable_type  varchar(2)   default 'S' not null comment '정책에서 설정값의 input 유형 (''S''=input select 값, ''T''=input text 값)',
    settable_value varchar(1024)            null comment '정책에서 설정 가능한 값',
    ad_group_name  varchar(40)              null
);

create table port_forwarding_map
(
    port_forwarding         int         not null comment '포워딩 포트 번호',
    instance_uuid           varchar(40) null comment '포트 포워딩된 인스턴스 UUID',
    port_forwarding_rule_id varchar(40) null comment '포트 포워딩 룰 아이디'
);

create table users
(
    id          int auto_increment
        primary key,
    uuid        varchar(40)             not null comment '유저 UUID',
    user_name   varchar(100)            not null comment '유저 계정 이름',
    create_date datetime                not null comment '유저 비밀번호',
    password    varchar(40) default '0' not null,
    removed     timestamp               null
);

create table vm_instances
(
    id                      int auto_increment
        primary key,
    name                    varchar(255)                    not null comment '가상머신 name',
    uuid                    varchar(40)                     not null comment '가상머신 UUID',
    workspace_uuid          varchar(40)                     not null comment '워크스페이스 UUID',
    workspace_name          varchar(50)                     null comment '워크스페이스 이름',
    mold_uuid               varchar(40)                     null comment 'Mold 의 UUID',
    status                  varchar(20) default 'Not Ready' not null comment '가상머신의 실행 여부',
    handshake_status        varchar(40) default 'Not Ready' null comment '핸드쉐이커 상태값',
    owner_account_id        varchar(100)                    null comment '가상머신 할당된 사용자 계정',
    ipaddress               varchar(255)                    null comment 'VDI 의 아이피',
    checked                 tinyint     default 0           not null,
    connected               int         default 0           not null comment '가상머신 접속여부',
    create_date             datetime                        null comment '가상머신 생성일',
    removed                 datetime                        null comment '가상머신 삭제일',
    windows_uuid            varchar(40)                     null,
    port_forwarding_check   int         default 0           not null,
    checked_date            datetime                        null,
    port_forwarding_rule_id int                             null,
    rdp_connected_check     int         default 0           not null,
    hash                    varchar(40)                     null comment '가상머신 hash 값'
)
    comment '데스크탑 및 앱용 가상머신 정보 테이블';

create table workspaces
(
    id                    int auto_increment comment '워크스페이스 ID'
        primary key,
    name                  varchar(255)                         not null comment '워크스페이스 이름',
    description           varchar(255)                         null comment '워크스페이스 설명',
    uuid                  varchar(40)                          not null comment '워크스페이스 UUID',
    state                 varchar(10) default 'Enable'         not null comment '워크스페이스 상태',
    workspace_type        varchar(20)                          not null comment '워크스페이스 타입',
    template_ok_check     varchar(20) default 'BeforeStarting' not null comment 'Template 이 정상적으로 사용가능 여부 확인하는 컬럼
BeforeStarting - VM 생성전
Starting - VM 시작중
AgentCheck - Agent 체크 중
AgentOK - Agent OK',
    quantity              int         default 0                not null comment '워크스페이스에 생성된 Desktop 수량.
Desktop Type만 수량 입력',
    network_uuid          varchar(48)                          not null comment '워크스페이스에서 사용할 network uuid',
    compute_offering_uuid varchar(48)                          not null comment '워크스페이스에서 사용할 compute offering uuid',
    template_uuid         varchar(48)                          not null comment '워크스페이스에서 사용할 template uuid',
    master_template_name  varchar(255)                         null comment '마스터 템플릿 이름',
    rbd_access_allow      int         default 0                not null comment 'RDP 접속 허용 여부
0 : 허용 안함
1 : 허용함',
    postfix               int         default 0                not null comment '워크스페이스 postfix',
    shared                tinyint     default 0                not null,
    create_date           datetime                             not null comment '워크스페이스 생성일',
    removed               datetime                             null comment '워크스페이스 삭제일.
해당 컬럼에 값이 있으면 삭제된 워크스페이스임'
)
    comment '워크스페이스 관련 테이블';

create table workspaces_policy
(
    id              int auto_increment
        primary key,
    name            varchar(40)  not null comment '워크스페이스에 적용된 정책 이름',
    workspaces_uuid varchar(40)  not null comment '워크스페이스 UUID',
    value           varchar(100) not null comment '정책 적용 값'
)
    comment '워크스페이스 정책 테이블';


-- configuration table default value insert
INSERT INTO works.configuration (id, name, category, value, default_value, description, update_date, input_type, input_value) VALUES (1, 'mold.default.domain.account', 'Advanced', 'admin', 'C:\\Program Files (x86)\\Microsoft\\Edge\\Application\\msedge.exe', 'Mold 의 domain 계정', null, 'input', null);
INSERT INTO works.configuration (id, name, category, value, default_value, description, update_date, input_type, input_value) VALUES (2, 'mold.default.domain.id', 'Advanced', '244b7491-5cb4-11ec-b9cd-002481021c72', null, 'Mold 의 domain ID', null, 'input', null);
INSERT INTO works.configuration (id, name, category, value, default_value, description, update_date, input_type, input_value) VALUES (3, 'mold.default.api.key', 'Advanced', 'XaGp5JK8exSk-nMDoEIAKFwAXPMI0V05hASU2E-0w7BibKGKVS32ye2U7vZL5A_FalIDyoyMqMJYY_44dnUHng', null, 'Mold 의 도메인 관리자의 API Key 를 설정하는 항목입니다.', null, 'input', null);
INSERT INTO works.configuration (id, name, category, value, default_value, description, update_date, input_type, input_value) VALUES (4, 'mold.default.secret.key', 'Advanced', 'rT4GcYhiyXfu-O6L-9naVXMWLdilCCkmHmyclADU0uxikOcdc01gqPX_JakvG8S-83mHfqmIbZZ7ZJKr1OELxA', null, 'Mold 의 도메인 관리자의 Secret Key 를 설정하는 항목입니다.', null, 'input', null);
INSERT INTO works.configuration (id, name, category, value, default_value, description, update_date, input_type, input_value) VALUES (5, 'mold.default.protocol', 'Advanced', 'http://', null, 'Mold 의 프로토콜 타입', null, 'select', '{http://,https://}');
INSERT INTO works.configuration (id, name, category, value, default_value, description, update_date, input_type, input_value) VALUES (6, 'mold.default.url', 'Advanced', '10.10.1.10', null, 'Mold 의 URL', null, 'input', null);
INSERT INTO works.configuration (id, name, category, value, default_value, description, update_date, input_type, input_value) VALUES (7, 'mold.default.port', 'Advanced', '8080', null, 'Mold 의 Port', null, 'input', null);
INSERT INTO works.configuration (id, name, category, value, default_value, description, update_date, input_type, input_value) VALUES (8, 'mold.default.url.postfix', 'Advanced', '/client/api', null, 'Mold 의 URL 의 Postfix', null, 'input', null);
INSERT INTO works.configuration (id, name, category, value, default_value, description, update_date, input_type, input_value) VALUES (9, 'dc.default.protocol', 'Advanced', 'http://', null, 'Domain Controller 의 프로토콜 타입', null, 'select', '{http://,https://}');
INSERT INTO works.configuration (id, name, category, value, default_value, description, update_date, input_type, input_value) VALUES (10, 'dc.default.url', 'Advanced', '10.1.1.202', null, 'Domain Controller 의 URL', null, 'input', null);
INSERT INTO works.configuration (id, name, category, value, default_value, description, update_date, input_type, input_value) VALUES (11, 'dc.default.port', 'Advanced', '8083', null, 'Domain Controller 의 Port', null, 'input', null);
INSERT INTO works.configuration (id, name, category, value, default_value, description, update_date, input_type, input_value) VALUES (12, 'dc.default.url.postfix', 'Advanced', '/api', null, 'Domain Controller 의 URL 의 Postfix', null, 'input', null);
INSERT INTO works.configuration (id, name, category, value, default_value, description, update_date, input_type, input_value) VALUES (13, 'mold.network.uuid', 'Advanced', '246', null, 'Works 에서 사용할 Mold network uuid', null, 'input', null);
INSERT INTO works.configuration (id, name, category, value, default_value, description, update_date, input_type, input_value) VALUES (14, 'mold.zone.id', 'Advanced', '3d1447eb-b65c-44a5-9fae-e9d5a57dc31b', null, 'Works 가 생성된 Zone ID', null, 'input', null);
INSERT INTO works.configuration (id, name, category, value, default_value, description, update_date, input_type, input_value) VALUES (15, 'works.default.url', 'Advanced', '10.1.1.201', null, 'Works 의 URL', null, 'input', null);
INSERT INTO works.configuration (id, name, category, value, default_value, description, update_date, input_type, input_value) VALUES (16, 'works.default.port', 'Advanced', '8082', null, 'Works 의 Port', null, 'input', null);
INSERT INTO works.configuration (id, name, category, value, default_value, description, update_date, input_type, input_value) VALUES (17, 'samba.default.url', 'Advanced', '10.1.1.201', null, 'Samba 의 URL', null, 'input', null);
INSERT INTO works.configuration (id, name, category, value, default_value, description, update_date, input_type, input_value) VALUES (18, 'samba.default.domain', 'Advanced', 'able', null, 'Samba 의 Domain', null, 'input', null);
INSERT INTO works.configuration (id, name, category, value, default_value, description, update_date, input_type, input_value) VALUES (19, 'guacamole.default.protocol', 'Advanced', 'http://', null, 'guacamole 의 protocol', null, 'select', '{http://,https://}');
INSERT INTO works.configuration (id, name, category, value, default_value, description, update_date, input_type, input_value) VALUES (20, 'guacamole.default.url', 'Advanced', '192.168.0.101', null, 'guacamole 의 URL', null, 'input', null);
INSERT INTO works.configuration (id, name, category, value, default_value, description, update_date, input_type, input_value) VALUES (21, 'guacamole.default.port', 'Advanced', '8080', null, 'guacamole 의 Port', null, 'input', null);
INSERT INTO works.configuration (id, name, category, value, default_value, description, update_date, input_type, input_value) VALUES (22, 'guacamole.default.url.postfix', 'Advanced', '/guacamole/api', null, 'guacamole 의 URL 의 Postfix', null, 'input', null);
INSERT INTO works.configuration (id, name, category, value, default_value, description, update_date, input_type, input_value) VALUES (23, 'guacamole.default.username', 'Advanced', 'administrator', null, 'guacamole 의 기본 도메인', null, 'input', null);
INSERT INTO works.configuration (id, name, category, value, default_value, description, update_date, input_type, input_value) VALUES (24, 'cluster.default.name', 'Advanced', 'able', null, 'Cluster 의 이름', null, 'input', null);
INSERT INTO works.configuration (id, name, category, value, default_value, description, update_date, input_type, input_value) VALUES (25, 'log.default.level', 'Advanced', 'logrus.DebugLevel', null, 'Works-API 의 Log Level', null, 'input', null);
INSERT INTO works.configuration (id, name, category, value, default_value, description, update_date, input_type, input_value) VALUES (26, 'rdp.default.port', 'Advanced', '3389', null, 'RDP 의 기본 Port', null, 'input', null);
INSERT INTO works.configuration (id, name, category, value, default_value, description, update_date, input_type, input_value) VALUES (27, 'rdp.default.forwarding.range', 'Advanced', '20001-20100', null, 'RDP 의 포트포워딩 범위', null, 'input', 'int');


-- policy_list table default value insert
INSERT INTO works.policy_list (id, name, description, i18n, value, type, settable_type, settable_value, ad_group_name) VALUES (1, 'rdp_port', 'rdp 접속 port', 'label.policy.desc.rdp.port', '3389', 'R', 'T', '', null);
INSERT INTO works.policy_list (id, name, description, i18n, value, type, settable_type, settable_value, ad_group_name) VALUES (2, 'rdp_access_allow', 'RDP 접속 허용 여부', 'label.policy.desc.rdp.access.allow', 'false', 'R', 'S', 'true,false', null);
INSERT INTO works.policy_list (id, name, description, i18n, value, type, settable_type, settable_value, ad_group_name) VALUES (3, 'clipboard_redirection', '클립보드 리디렉션 허용 여부', 'label.policy.desc.clipboard.redirection', 'false', 'R', 'S', 'true,false', 'clipboard_block');
INSERT INTO works.policy_list (id, name, description, i18n, value, type, settable_type, settable_value, ad_group_name) VALUES (4, 'windows_installer', '사용자 프로그램 설치 허용 여부', 'label.policy.desc.windows.installer', 'false', 'R', 'S', 'true,false', 'install_block');
INSERT INTO works.policy_list (id, name, description, i18n, value, type, settable_type, settable_value, ad_group_name) VALUES (5, 'network_file_share', '네트워크 파일 공유 허용 여부', 'label.policy.desc.network.file.share', 'false', 'R', 'S', 'true,false', 'share_block');
INSERT INTO works.policy_list (id, name, description, i18n, value, type, settable_type, settable_value, ad_group_name) VALUES (6, 'network_printer_share', '네트워크 프린터 공유 허용 여부', 'label.policy.desc.network.printer.share', 'false', 'R', 'S', 'true,false', 'network_printer_block');
INSERT INTO works.policy_list (id, name, description, i18n, value, type, settable_type, settable_value, ad_group_name) VALUES (7, 'local_printer', '로컬 프린터 허용 여부', 'label.policy.desc.local.printer', 'false', 'R', 'S', 'true,false', 'local_printer_block');
INSERT INTO works.policy_list (id, name, description, i18n, value, type, settable_type, settable_value, ad_group_name) VALUES (8, 'windows_auto_update', '윈도우 자동 업데이트 허용 여부', 'label.policy.desc.windows.auto.update', 'false', 'R', 'S', 'true,false', 'update_block');
INSERT INTO works.policy_list (id, name, description, i18n, value, type, settable_type, settable_value, ad_group_name) VALUES (9, 'removable_storage', '미동식 저장소 허용 여부', 'label.policy.desc.removable.storage', 'false', 'R', 'S', 'true,false', 'usb_block');
INSERT INTO works.policy_list (id, name, description, i18n, value, type, settable_type, settable_value, ad_group_name) VALUES (10, 'cmd_use', 'CMD 허용 여부', 'label.policy.desc.cmd.use', 'false', 'R', 'S', 'true,false', 'cmd_block');
INSERT INTO works.policy_list (id, name, description, i18n, value, type, settable_type, settable_value, ad_group_name) VALUES (11, 'settings_use', '설정 허용 여부', 'label.policy.desc.settings.use', 'false', 'R', 'S', 'true,false', 'setting_block');
INSERT INTO works.policy_list (id, name, description, i18n, value, type, settable_type, settable_value, ad_group_name) VALUES (12, 'password_expiration_date', '비밀번호 만료일 설정', 'label.policy.desc.password_expiration_date', '30', 'R', 'S', '0,30,60,90,180,365', 'passwordchangeday');
INSERT INTO works.policy_list (id, name, description, i18n, value, type, settable_type, settable_value, ad_group_name) VALUES (13, 'pc-power-use', '사용자 데스크탑 온오프 사용 여부', 'label.policy.desc.pc.power.use', 'false', 'R', 'S', 'true,false', 'logoff_block');
INSERT INTO works.policy_list (id, name, description, i18n, value, type, settable_type, settable_value, ad_group_name) VALUES (14, 'remotefx', 'remotefx 설정 여부', 'label.policy.desc.remotefx', 'true', 'R', 'S', 'true,false', 'remotefx');
INSERT INTO works.policy_list (id, name, description, i18n, value, type, settable_type, settable_value, ad_group_name) VALUES (15, 'console', 'RDP 서버의 콘솔 세션 연결 여부', 'label.policy.desc.console', 'false', 'C', 'S', 'true,false', null);
INSERT INTO works.policy_list (id, name, description, i18n, value, type, settable_type, settable_value, ad_group_name) VALUES (16, 'initial-program', 'Client에서 실행할 프로그램의 전체 경로', 'label.policy.desc.initial.program', null, 'C', 'T', null, null);
INSERT INTO works.policy_list (id, name, description, i18n, value, type, settable_type, settable_value, ad_group_name) VALUES (17, 'server-layout', '서버측 키보드 레이아웃 설정 값', 'label.policy.desc.server.layout', null, 'C', 'T', null, null);
INSERT INTO works.policy_list (id, name, description, i18n, value, type, settable_type, settable_value, ad_group_name) VALUES (18, 'color-depth', '요청할 픽셀당 비트수 값', 'label.policy.desc.color.depth', '24', 'C', 'S', '8,16,24', null);
INSERT INTO works.policy_list (id, name, description, i18n, value, type, settable_type, settable_value, ad_group_name) VALUES (19, 'width', '디스플레이의 너비(픽셀)', 'label.policy.desc.width', null, 'C', 'T', null, null);
INSERT INTO works.policy_list (id, name, description, i18n, value, type, settable_type, settable_value, ad_group_name) VALUES (20, 'height', '디스플레이의 높이(픽셀)', 'label.policy.desc.height', null, 'C', 'T', null, null);
INSERT INTO works.policy_list (id, name, description, i18n, value, type, settable_type, settable_value, ad_group_name) VALUES (21, 'dpi', '디스플레이의 해상도(DPI)', 'label.policy.desc.dpi', null, 'C', 'T', null, null);
INSERT INTO works.policy_list (id, name, description, i18n, value, type, settable_type, settable_value, ad_group_name) VALUES (22, 'resize-method', '디스플레이 크기 변경시 RDP 서버를 업데이트 하는 방법', 'label.policy.desc.resize.method', 'display-update', 'C', 'S', 'display-update,reconnect', null);
INSERT INTO works.policy_list (id, name, description, i18n, value, type, settable_type, settable_value, ad_group_name) VALUES (23, 'force-lossless', '그래픽 압축 무손실 사용 여부', 'label.policy.desc.force.lossless', 'false', 'C', 'S', 'true,false', null);
INSERT INTO works.policy_list (id, name, description, i18n, value, type, settable_type, settable_value, ad_group_name) VALUES (24, 'disable-audio', '오디오 출력 사용 여부', 'label.policy.desc.disable.audio', 'false', 'C', 'S', 'true,false', null);
INSERT INTO works.policy_list (id, name, description, i18n, value, type, settable_type, settable_value, ad_group_name) VALUES (25, 'enable-audio-input', '오디오 입력 사용 여부', 'label.policy.desc.enable.audio.input', 'false', 'C', 'S', 'true,false', null);
INSERT INTO works.policy_list (id, name, description, i18n, value, type, settable_type, settable_value, ad_group_name) VALUES (26, 'enable-printing', '클라이언트에서 PDF 프린터 사용 여부', 'label.policy.desc.enable.printing', 'false', 'C', 'S', 'true,false', null);
INSERT INTO works.policy_list (id, name, description, i18n, value, type, settable_type, settable_value, ad_group_name) VALUES (27, 'printer-name', 'RDP 세션을 통해 전달되는 리디렉션된 프린터 장치의 이름 값', 'label.policy.desc.printer.name', 'VDI-PRINTER', 'C', 'T', null, null);
INSERT INTO works.policy_list (id, name, description, i18n, value, type, settable_type, settable_value, ad_group_name) VALUES (28, 'enable-drive', '클라이언트에서 Guacamole 서버에 있는 가상 드라이브와 파일 전송 허용 여부', 'label.policy.desc.enable.drive', 'true', 'C', 'S', 'true,false', null);
INSERT INTO works.policy_list (id, name, description, i18n, value, type, settable_type, settable_value, ad_group_name) VALUES (29, 'disable-download', '원격 서버에서 클라이언트로의 다운로드 사용 여부', 'label.policy.desc.disable.download', 'false', 'C', 'S', 'true,false', null);
INSERT INTO works.policy_list (id, name, description, i18n, value, type, settable_type, settable_value, ad_group_name) VALUES (30, 'disable-upload', '클라이언트에서 원격 서버 위치로의 업로드 허용 여부', 'label.policy.desc.disable.upload', 'false', 'C', 'S', 'true,false', null);
INSERT INTO works.policy_list (id, name, description, i18n, value, type, settable_type, settable_value, ad_group_name) VALUES (31, 'drive-name', 'RDP 세션으로 전달될 때 사용되는 파일 시스템 이름', 'label.policy.desc.drive.name', 'VDI-DRIVE', 'C', 'T', null, null);
INSERT INTO works.policy_list (id, name, description, i18n, value, type, settable_type, settable_value, ad_group_name) VALUES (32, 'drive-path', '전송된 파일이 저장되어야 하는 Guacamole 서버의 디렉토리', 'label.policy.desc.drive.path', '/share', 'C', 'T', null, null);
INSERT INTO works.policy_list (id, name, description, i18n, value, type, settable_type, settable_value, ad_group_name) VALUES (33, 'create-drive-path', 'drive_path 정책에서 적용된 디렉토리가 아직 존재 하지 않는경우 자동생성 여부', 'label.policy.desc.create.drive.path', 'false', 'C', 'S', 'true,false', null);
INSERT INTO works.policy_list (id, name, description, i18n, value, type, settable_type, settable_value, ad_group_name) VALUES (34, 'console-audio', 'RDP 서버의 콘솔 세션에서 오디오가 활성화 여부', 'label.policy.desc.console.audio', 'false', 'C', 'S', 'true,false', null);
INSERT INTO works.policy_list (id, name, description, i18n, value, type, settable_type, settable_value, ad_group_name) VALUES (35, 'enable-wallpaper', '바탕화면 랜더링 사용 여부', 'label.policy.desc.enable.wallpaper', 'false', 'C', 'S', 'true,false', null);
INSERT INTO works.policy_list (id, name, description, i18n, value, type, settable_type, settable_value, ad_group_name) VALUES (36, 'enable-theming', '창 및 컨트롤ㄹ의 테마 사용 여부', 'label.policy.desc.enable.theming', 'false', 'C', 'S', 'true,false', null);
INSERT INTO works.policy_list (id, name, description, i18n, value, type, settable_type, settable_value, ad_group_name) VALUES (37, 'enable-font-smoothing', '텍스트 가장자리 렌더링 사용 여부', 'label.policy.desc.enable.font.smoothing', 'false', 'C', 'S', 'true,false', null);
INSERT INTO works.policy_list (id, name, description, i18n, value, type, settable_type, settable_value, ad_group_name) VALUES (38, 'enable-full-window-drag', '창을 이동할 때 창의 내용 표시 여부', 'label.policy.desc.enable.full.window.drag', 'false', 'C', 'S', 'true,false', null);
INSERT INTO works.policy_list (id, name, description, i18n, value, type, settable_type, settable_value, ad_group_name) VALUES (39, 'enable-desktop-composition', '투명한 창 및 그림자와 같은 그래픽 효과 사용 여부', 'label.policy.desc.enable.desktop.composition', 'false', 'C', 'S', 'true,false', null);
INSERT INTO works.policy_list (id, name, description, i18n, value, type, settable_type, settable_value, ad_group_name) VALUES (40, 'enable-menu-animations', '시작 메뉴 열기 및 닫기 애니메이션이 허용 여부', 'label.policy.desc.enable.menu.animations', 'false', 'C', 'S', 'true,false', null);