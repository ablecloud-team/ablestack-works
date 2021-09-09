package main

import (
	"database/sql"
	"fmt"
	"github.com/sirupsen/logrus"
	"net/http"
	"os"
	"strconv"
)

type Workspace struct {
	Id                                             int
	Quantity, Postfix                              int
	Name, Description, Uuid, State, Type           string
	TemplateOkCheck                                string
	Shared                                         bool
	NetworkUuid, ComputeOfferingUuid, TemplateUuid string
	CreateDate, Removed                            []uint8
}

type Instance struct {
	Id                                                     int
	Name, Uuid, WorkspaceUuid, WindowsUuid, ownerAccountId string
	Checked, Connected                                     bool
	CreateDate, Removed                                    []uint8
}

func selectWorkspaceList() string {
	////pageInt, _ := strconv.Atoi(payload["page"].(string))
	////pageSizeInt, _ := strconv.Atoi(payload["pagesize"].(string))
	//pageInt = (pageInt - 1) * pageSizeInt

	//log.Info("page=["+strconv.Itoa(pageInt)+"]pageSize=["+strconv.Itoa(pageSizeInt)+"]")
	db, err := sql.Open(os.Getenv("MysqlType"), os.Getenv("DbInfo"))
	if err != nil {
		log.Error("DB connect error")
		log.Error(err)
	}
	defer db.Close()
	log.Info("DB connect success")

	rows, err := db.Query("SELECT * FROM workspaces ORDER BY id")
	if err != nil {
		log.Error("workspace select query FAILED")
		log.Error(err.Error())
	}
	defer rows.Close()

	result, err := rowsToString(rows)
	if err != nil {
		fmt.Println("err")
		fmt.Println(err)
	}
	return result
}

func selectCountWorkspace() string {
	db, err := sql.Open(os.Getenv("MysqlType"), os.Getenv("DbInfo"))
	if err != nil {
		log.Error("DB connect error")
		log.Error(err)
	}
	defer db.Close()
	log.Info("DB connect success")
	var workspaceCount int
	err = db.QueryRow("SELECT count(*) FROM workspaces where removed IS NULL").Scan(&workspaceCount)
	if err != nil {
		log.Error("workspace count select query FAILED")
		log.Error(err.Error())
	}

	return strconv.Itoa(workspaceCount)
}

func selectWorkspaceDetail(workspaceUuid string) Workspace {
	db, err := sql.Open(os.Getenv("MysqlType"), os.Getenv("DbInfo"))
	if err != nil {
		log.Error("DB connect error")
		log.Error(err)
	}
	defer db.Close()
	log.Info("DB connect success")
	log.Info(workspaceUuid)
	workspace := Workspace{}
	err = db.QueryRow("SELECT id, name, description, uuid, state, type, template_ok_check, quantity, shared, postfix, network_uuid, compute_offering_uuid, template_uuid, create_date FROM workspaces where uuid = ?", workspaceUuid).Scan(
		&workspace.Id, &workspace.Name, &workspace.Description,
		&workspace.Uuid, &workspace.State, &workspace.Type,
		&workspace.TemplateOkCheck, &workspace.Quantity, &workspace.Shared, &workspace.Postfix,
		&workspace.NetworkUuid, &workspace.ComputeOfferingUuid, &workspace.TemplateUuid,
		&workspace.CreateDate)
	if err != nil {
		log.Error("workspace Detail select query FAILED")
		log.Error(err.Error())
	}
	log.Debugln(workspace)

	return workspace
}

func selectInstanceDetail(instanceUuid string) Instance {
	db, err := sql.Open(os.Getenv("MysqlType"), os.Getenv("DbInfo"))
	if err != nil {
		log.Error("DB connect error")
		log.Error(err)
	}
	defer db.Close()
	log.Info(MsgDBConnectOK)
	log.Info(instanceUuid)
	instance := Instance{}
	err = db.QueryRow("SELECT id, name, uuid, workspace_uuid, windows_uuid, owner_account_id, checked, connected, create_date FROM vm_instances where removed IS NULL AND uuid = ?", instanceUuid).Scan(
		&instance.Id, &instance.Name, &instance.WorkspaceUuid, &instance.WindowsUuid,
		&instance.ownerAccountId, &instance.Checked, &instance.Connected, &instance.CreateDate)
	if err != nil {
		log.Error("instance select query FAILED")
		log.Error(err.Error())
	}
	log.Debugln(instance)
	if err != nil {
		fmt.Println("err")
		fmt.Println(err)
	}
	return instance
}

func insertWorkspace(workspace Workspace) map[string]interface{} {
	db, err := sql.Open(os.Getenv("MysqlType"), os.Getenv("DbInfo"))
	resultData := map[string]interface{}{}
	if err != nil {
		log.Error("DB connect error")
		log.Error(err)
		resultData["message"] = "DB connect error"
		resultData["status"] = BaseErrorCode
	}
	defer db.Close()

	result, err := db.Exec("INSERT INTO workspaces(name, description, uuid, type, network_uuid, compute_offering_uuid, template_uuid, create_date) VALUES (?, ?, ?, ?, ?, ?, ?,  now())",
		workspace.Name, workspace.Description, workspace.Uuid, workspace.Type, workspace.NetworkUuid, workspace.ComputeOfferingUuid, workspace.TemplateUuid)
	if err != nil {
		log.Error("UUID 생성 후 DB Insert 중 오류가 발생하였습니다.")
		log.Error(err)
		resultData["message"] = "An error occurred while inserting the DB after generating the UUID."
		resultData["status"] = BaseErrorCode
	}
	n, err := result.RowsAffected()
	if n == 1 {
		log.Info("워크스페이스가 정상적으로 생성되었습니다.")
		resultData["message"] = "The workspace has been successfully created."
		resultData["status"] = http.StatusOK
	}

	return resultData
}

func insertInstance(instance Instance) map[string]interface{} {
	db, err := sql.Open(os.Getenv("MysqlType"), os.Getenv("DbInfo"))
	resultData := map[string]interface{}{}
	if err != nil {
		log.Error("DB connect error")
		log.Error(err)
		resultData["message"] = "DB connect error"
		resultData["status"] = BaseErrorCode
	}
	defer db.Close()

	result, err := db.Exec("INSERT INTO vm_instances(name, uuid, workspace_uuid, create_date) VALUES (?, ?, ?, now())",
		instance.Name, instance.Uuid, instance.WorkspaceUuid)
	if err != nil {
		log.Error("가상머신 DB Insert 중 오류가 발생하였습니다.")
		log.Error(err)
		resultData["message"] = "An error occurred while inserting the virtual machine DB."
		resultData["status"] = SQLQueryError
	}
	n, err := result.RowsAffected()
	if n == 1 {
		log.Info("가상머신이 정상적으로 등록되였습니다.")
		resultData["message"] = "The virtual machine has been successfully registered."
		resultData["status"] = http.StatusOK
	}

	return resultData
}

func updateWorkspacePostfix(workspaceUuid string, postfixInt int) map[string]interface{} {
	resultReturn := map[string]interface{}{}
	db, err := sql.Open(os.Getenv("MysqlType"), os.Getenv("DbInfo"))
	if err != nil {
		fmt.Println("DB connect error")
		fmt.Println(err)
	}
	defer db.Close()
	updatePostfixInt := postfixInt + 1
	log.Infof("워크스페이스에 업데이트 될 postfix 은 [%v] 입니다.", updatePostfixInt)
	result1, err := db.Exec("UPDATE workspaces set postfix=? where uuid=?", updatePostfixInt, workspaceUuid)
	if err != nil {
		log.Error("워크스페이스 postfix 업데이트 쿼리중 에러가 발생했습니다.")
		log.Error(err)
		resultReturn["message"] = "An error occurred while querying the workspace postfix update."
		resultReturn["status"] = SQLQueryError
	}
	n, err := result1.RowsAffected()
	if n == 1 {
		log.Info("워크스페이스 postfix 업데이트 완료가 되었습니다.")
		resultReturn["message"] = "Workspace postfix update is complete."
		resultReturn["status"] = http.StatusOK
	}

	return resultReturn
}

func selectNetworkDetail() string {
	db, err := sql.Open(os.Getenv("MysqlType"), os.Getenv("DbInfo"))
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
	db, err := sql.Open(os.Getenv("MysqlType"), os.Getenv("DbInfo"))
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

func updateWorkspaceAgent(uuid string, asyncJobUuid string, typeString string) map[string]interface{} {
	db, err := sql.Open(os.Getenv("MysqlType"), os.Getenv("DbInfo"))
	resultReturn := map[string]interface{}{}
	if err != nil {
		log.Error("DB connect error")
		log.Error(err)
		resultReturn["message"] = MsgDBConnectError
		resultReturn["status"] = BaseErrorCode
	}
	log.WithFields(logrus.Fields{
		"workspaceImpl": "updateWorkspaceAgent",
	}).Infof("uuid=[%v], asyncJobUuid=[%v], sypeString=[%v]", uuid, asyncJobUuid, typeString)
	defer db.Close()
	if typeString == WorkspaceString {
		result, err := db.Exec("UPDATE workspaces set template_ok_check=?, state=? where uuid=?", AgentOK, Enable, uuid)
		if err != nil {
			log.Error(MsgDBConnectError)
			log.Error(err)
			resultReturn["message"] = MsgDBConnectError
			resultReturn["status"] = SQLQueryError
		}
		n1, _ := result.RowsAffected()
		if n1 == 1 {
			updateAsyncJobRead(asyncJobUuid, 1)
		}
	} else if typeString == InstanceString {
		result, err := db.Exec("UPDATE vm_instances set checked=? where uuid=?", 1, uuid)
		if err != nil {
			log.Error(MsgDBConnectError)
			log.Error(err)
			resultReturn["message"] = MsgDBConnectError
			resultReturn["status"] = SQLQueryError
		}
		n1, _ := result.RowsAffected()
		if n1 == 1 {
			updateAsyncJobRead(asyncJobUuid, 1)
		}
	}

	return resultReturn
}
