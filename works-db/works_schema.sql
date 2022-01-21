CREATE USER works@'%' IDENTIFIED BY 'Ablecloud1!';
GRANT ALL ON *.* TO works@'%';
FLUSH PRIVILEGES;

-- auto-generated definition
create schema works collate utf8mb4_0900_ai_ci;

grant alter, alter routine, create, create routine, create temporary tables, create view, delete, drop, event, execute, index, insert, lock tables, references, select, show view, trigger, update on works.* to works;


create table works.async_job
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

create table works.configuration
(
    id            int auto_increment
        primary key,
    name          varchar(100)                    not null comment '설정 이름',
    category      varchar(100) default 'Advanced' not null comment '설정 카테고리',
    value         varchar(256)                    not null comment '설정의 값',
    default_value varchar(256)                    null comment '설정 기본 값',
    description   varchar(1024)                   not null comment '설정 설명',
    update_date   datetime                        null comment '설정 업데이트 일자'
)
    comment 'works API 설정값';

create table works.users
(
    id          int auto_increment
        primary key,
    uuid        varchar(40)             not null comment '유저 UUID',
    user_name   varchar(100)            not null comment '유저 계정 이름',
    create_date datetime                not null comment '유저 비밀번호',
    password    varchar(40) default '0' not null,
    removed     timestamp               null
);

create table works.vm_instances
(
    id               int auto_increment
        primary key,
    name             varchar(255)                    not null comment '가상머신 name',
    uuid             varchar(40)                     not null comment '가상머신 UUID',
    workspace_uuid   varchar(40)                     not null comment '워크스페이스 UUID',
    workspace_name   varchar(50)                     null comment '워크스페이스 이름',
    mold_uuid        varchar(40)                     null comment 'Mold 의 UUID',
    status           varchar(20) default 'Not Ready' not null comment '가상머신의 실행 여부',
    owner_account_id varchar(100)                    null comment '가상머신 할당된 사용자 계정',
    checked          tinyint     default 0           not null,
    connected        tinyint     default 0           not null comment '가상머신 접속여부',
    create_date      datetime                        null comment '가상머신 생성일',
    removed          datetime                        null comment '가상머신 삭제일',
    windows_uuid     varchar(40)                     null,
    checked_date     datetime                        null
)
    comment '데스크탑 및 앱용 가상머신 정보 테이블';

create table works.workspaces
(
    id                    int auto_increment comment '워크스페이스 ID'
        primary key,
    name                  varchar(255)                         not null comment '워크스페이스 이름',
    description           varchar(255)                         null comment '워크스페이스 설명',
    uuid                  varchar(40)                          not null comment '워크스페이스 UUID',
    state                 varchar(10) default 'Enable'         not null comment '워크스페이스 상태',
    workspace_type        varchar(10)                          not null comment '워크스페이스 타입',
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
    postfix               int         default 0                not null comment '워크스페이스 postfix',
    shared                tinyint     default 0                not null,
    create_date           datetime                             not null comment '워크스페이스 생성일',
    removed               datetime                             null comment '워크스페이스 삭제일.
해당 컬럼에 값이 있으면 삭제된 워크스페이스임'
)
    comment '워크스페이스 관련 테이블';

INSERT INTO works.configuration (id, name, category, value, default_value, description, update_date) VALUES (1, 'mold.default.domain.account', 'Advanced', 'works', null, 'Mold 의 domain 계정', null);
INSERT INTO works.configuration (id, name, category, value, default_value, description, update_date) VALUES (2, 'mold.default.domain.id', 'Advanced', 'd294bfc7-f86e-4f67-ac83-80220110b23f', null, 'Mold 의 domain ID', null);
INSERT INTO works.configuration (id, name, category, value, default_value, description, update_date) VALUES (3, 'mold.default.api.key', 'Advanced', 'WuiJ_rWUllvp5a9Wkj1ZzwE_VqShM-3eJr1O2Jvi4RmaIfsWgbJh-GKxMpJ78V47wVX8BWACi9KilsaRusjZ7w', null, 'Mold 의 도메인 관리자의 API Key 를 설정하는 항목입니다.', null);
INSERT INTO works.configuration (id, name, category, value, default_value, description, update_date) VALUES (4, 'mold.default.secret.key', 'Advanced', 'Y3sJNJuK6fHuMigdOYqaJHAGgVaIb9RDUTfTfAAeQzR7UAUPUfcOjhekIHStirtLIfgb6Nbre_x1Lz9S7c7HHQ', null, 'Mold 의 도메인 관리자의 Secret Key 를 설정하는 항목입니다.', null);
INSERT INTO works.configuration (id, name, category, value, default_value, description, update_date) VALUES (5, 'mold.default.protocol', 'Advanced', 'https://', null, 'Mold 의 프로토콜 타입', null);
INSERT INTO works.configuration (id, name, category, value, default_value, description, update_date) VALUES (6, 'mold.default.url', 'Advanced', 'mold.ablecloud.io', null, 'Mold 의 URL', null);
INSERT INTO works.configuration (id, name, category, value, default_value, description, update_date) VALUES (7, 'mold.default.port', 'Advanced', '80', null, 'Mold 의 Port', null);
INSERT INTO works.configuration (id, name, category, value, default_value, description, update_date) VALUES (8, 'mold.default.url.postfix', 'Advanced', '/client/api', null, 'Mold 의 URL 의 Postfix', null);
INSERT INTO works.configuration (id, name, category, value, default_value, description, update_date) VALUES (9, 'dc.default.protocol', 'Advanced', 'http://', null, 'Domain Controller 의 프로토콜 타입', null);
INSERT INTO works.configuration (id, name, category, value, default_value, description, update_date) VALUES (10, 'dc.default.url', 'Advanced', '10.10.1.99', null, 'Domain Controller 의 URL', null);
INSERT INTO works.configuration (id, name, category, value, default_value, description, update_date) VALUES (11, 'dc.default.port', 'Advanced', '8083', null, 'Domain Controller 의 Port', null);
INSERT INTO works.configuration (id, name, category, value, default_value, description, update_date) VALUES (12, 'dc.default.url.postfix', 'Advanced', '/api', null, 'Domain Controller 의 URL 의 Postfix', null);
INSERT INTO works.configuration (id, name, category, value, default_value, description, update_date) VALUES (13, 'mold.network.uuid', 'Advanced', '205', null, 'Works 에서 사용할 Mold network uuid', null);
INSERT INTO works.configuration (id, name, category, value, default_value, description, update_date) VALUES (14, 'mold.zone.id', 'Advanced', '82bfaf86-ea6b-4cbc-88c5-0a951fd57aa0', null, 'Works 가 생성된 Zone ID', null);
INSERT INTO works.configuration (id, name, category, value, default_value, description, update_date) VALUES (15, 'works.default.url', 'Advanced', '10.10.13.102', null, 'Works 의 URL', null);
INSERT INTO works.configuration (id, name, category, value, default_value, description, update_date) VALUES (16, 'works.default.port', 'Advanced', '8083', null, 'Works 의 Port', null);
INSERT INTO works.configuration (id, name, category, value, default_value, description, update_date) VALUES (17, 'samba.default.url', 'Advanced', '10.1.1.59', null, 'Samba 의 URL', null);
INSERT INTO works.configuration (id, name, category, value, default_value, description, update_date) VALUES (18, 'samba.default.domain', 'Advanced', 'testdomain', null, 'Samba 의 Domain', null);
INSERT INTO works.configuration (id, name, category, value, default_value, description, update_date) VALUES (19, 'guacamole.default.protocol', 'Advanced', 'http://', null, 'guacamole 의 protocol', null);
INSERT INTO works.configuration (id, name, category, value, default_value, description, update_date) VALUES (20, 'guacamole.default.url', 'Advanced', '10.10.1.100', null, 'guacamole 의 URL', null);
INSERT INTO works.configuration (id, name, category, value, default_value, description, update_date) VALUES (21, 'guacamole.default.port', 'Advanced', '8080', null, 'guacamole 의 Port', null);
INSERT INTO works.configuration (id, name, category, value, default_value, description, update_date) VALUES (22, 'guacamole.default.url.postfix', 'Advanced', '/guacamole/api', null, 'guacamole 의 URL 의 Postfix', null);
INSERT INTO works.configuration (id, name, category, value, default_value, description, update_date) VALUES (23, 'guacamole.default.username', 'Advanced', 'guacadmin', null, 'guacamole 의 기본 도메인', null);
INSERT INTO works.configuration (id, name, category, value, default_value, description, update_date) VALUES (24, 'cluster.default.name', 'Advanced', 'clusterTestName', null, 'Cluster 의 이름', null);
INSERT INTO works.configuration (id, name, category, value, default_value, description, update_date) VALUES (25, 'log.default.level', 'Advanced', 'logrus.DebugLevel', null, 'Works-API 의 Log Level', null);