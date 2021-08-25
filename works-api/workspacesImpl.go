package main

import (
	"database/sql"
	"fmt"
	"os"
	"time"
)

type Workspace struct {
	Id, Quantity                             int
	Name, Description, Uuid, State, Type     string
	NetworkName, NetworkType, NetworkUuid    string
	ComputeOfferingName, ComputeOfferingUuid string
	TemplateName, TemplateUuid               string
	DiskOfferingName, DiskOfferingUuid       string
	CreateDate, Removed                      time.Time
}

func insertWorkspace(workspace Workspace) string {
	db, err := sql.Open(os.Getenv("MsqlType"), os.Getenv("DbInfo"))
	message := ""
	if err != nil {
		log.Error("DB connect error")
		log.Error(err)
		message = "DB connect error"
	}
	defer db.Close()

	result, err := db.Exec("INSERT INTO workspaces(name, description, uuid, type, network_uuid, compute_offering_uuid, template_uuid, disk_offering_uuid, create_date) VALUES (?, ?, ?, ?, ?, ?, ?, ?, now())",
		workspace.Name, workspace.Description, workspace.Uuid, workspace.Type, workspace.NetworkUuid, workspace.ComputeOfferingUuid, workspace.TemplateUuid, workspace.DiskOfferingUuid)
	if err != nil {
		log.Error("UUID 생성 후 DB Insert 중 오류가 발생하였습니다.")
		log.Error(err)
		message = "An error occurred while inserting the DB after generating the UUID."
	}
	n, err := result.RowsAffected()
	if n == 1 {
		log.Info("워크스페이스가 정상적으로 생성되었습니다.")
		message = "The workspace has been successfully created."
	}

	return message
}

func selectNetworkDetail() string {
	db, err := sql.Open(os.Getenv("MsqlType"), os.Getenv("DbInfo"))
	if err != nil {
		fmt.Println("DB connect error")
		fmt.Println(err)
	}
	defer db.Close()
	var networkUuid string
	err = db.QueryRow("SELECT value as networkUuid FROM configuration where name = 'mold.network.uuid'").Scan(&networkUuid)
	if err != nil {
		log.Error("Network UUID 조회중 에러가 발생하였습니다.")
		log.Error(err)
	}

	return networkUuid
}

func selectZoneId() string {
	db, err := sql.Open(os.Getenv("MsqlType"), os.Getenv("DbInfo"))
	if err != nil {
		fmt.Println("DB connect error")
		fmt.Println(err)
	}
	defer db.Close()
	var zoneId string
	err = db.QueryRow("SELECT value as zoneId FROM configuration where name = 'mold.zone.id'").Scan(&zoneId)
	if err != nil {
		log.Error("Network UUID 조회중 에러가 발생하였습니다.")
		log.Error(err)
	}

	return zoneId
}

func insertWrokspaceAgent(workspaceUuid string, asyncJobUuid string) string {
	db, err := sql.Open(os.Getenv("MsqlType"), os.Getenv("DbInfo"))
	message := ""
	if err != nil {
		log.Error("DB connect error")
		log.Error(err)
		message = "DB connect error"
	}
	defer db.Close()

	result1, err := db.Exec("UPDATE workspaces set template_ok_check='AgentOK' where uuid=?", workspaceUuid)
	if err != nil {
		log.Error("Agent 체크 쿼리중 에러가 발생했습니다.")
		log.Error(err)
		message = "An error occurred during agent check query."
	}
	n1, err := result1.RowsAffected()
	if n1 == 1 {
		result2, err := db.Exec("UPDATE async_job set ready=1 where id=?", asyncJobUuid)
		if err != nil {
			log.Error("Async Job 쿼리중 에러가 발생했습니다.")
			log.Error(err)
		}
		n2, err := result2.RowsAffected()
		if n2 == 1 {
			log.Info("템플릿 Agent 체크가 완료되었습니다.")
			message = "Template Agent check is completed."
		} else {
			log.Error("요청하신 async_job UUID 가 없습니다.")
		}
	} else {
		log.Error("요청하신 workspace UUID 가 없습니다.")
	}

	return message
}
