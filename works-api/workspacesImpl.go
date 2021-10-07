package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/sirupsen/logrus"
	"net/http"
	"os"
	"time"
)

type Workspace struct {
	Id                  int     `json:"id"`
	Name                string  `json:"name"`
	Description         string  `json:"description"`
	Uuid                string  `json:"uuid"`
	State               string  `json:"state"`
	WorkspaceType       string  `json:"workspace_type"`
	TemplateOkCheck     string  `json:"template_ok_check"`
	Quantity            int     `json:"quantity"`
	NetworkUuid         string  `json:"network_uuid"`
	ComputeOfferingUuid string  `json:"compute_offering_uuid"`
	TemplateUuid        string  `json:"template_uuid"`
	Postfix             int     `json:"postfix"`
	Shared              bool    `json:"shared"`
	CreateDate          string  `json:"create_date"`
	Removed             *string `json:"removed"`
}

type Instance struct {
	Id             int     `json:"id"`
	Name           string  `json:"name"`
	Uuid           string  `json:"uuid"`
	WorkspaceUuid  string  `json:"workspace_uuid"`
	WorkspaceName  string  `json:"workspace_name"`
	MoldUuid       string  `json:"mold_uuid"`
	Status         string  `json:"status"`
	OwnerAccountId *string `json:"owner_account_id,omitempty"`
	Checked        bool    `json:"checked"`
	Connected      bool    `json:"connected"`
	CreateDate     string  `json:"create_date"`
	CheckedDate    *string `json:"checked_date"`
	Removed        string  `json:"removed"`
	MoldStatus     string  `json:"mold_status"`
}

func selectWorkspaceList(workspaceUuid string) ([]Workspace, error) {
	log.WithFields(logrus.Fields{
		"workspaceImpl": "selectWorkspaceList",
	}).Infof("payload workspaceUuid [%v]", workspaceUuid)
	db, err := sql.Open(os.Getenv("MysqlType"), os.Getenv("DbInfo"))
	if err != nil {
		log.WithFields(logrus.Fields{
			"workspaceImpl": "selectWorkspaceList",
		}).Errorf("selectWorkspaceList DB Connect Error [%v]", err)
	}
	defer db.Close()
	log.WithFields(logrus.Fields{
		"workspaceImpl": "selectWorkspaceList",
	}).Infof("select WorkspaceList DB Connect success")
	queryString := "SELECT" +
		" id, name, description, uuid, state," +
		"workspace_type, template_ok_check, quantity, network_uuid, compute_offering_uuid," +
		"template_uuid, postfix, shared, create_date, removed" +
		" FROM workspaces" +
		" WHERE removed IS NULL"
	if workspaceUuid != ALL {
		queryString = queryString + " AND uuid = '" + workspaceUuid + "'"
	}
	queryString = queryString + " ORDER BY id DESC"
	log.WithFields(logrus.Fields{
		"workspaceImpl": "selectWorkspaceList",
	}).Infof("select WorkspaceList Query [%v]", queryString)
	rows, err := db.Query(queryString)
	if err != nil {
		log.WithFields(logrus.Fields{
			"workspaceImpl": "selectWorkspaceList",
		}).Errorf("Workspace Select Query FAILED [%v]", err)
	}
	var workspaceList []Workspace
	defer rows.Close()
	for rows.Next() {
		workspace := Workspace{}
		err = rows.Scan(
			&workspace.Id, &workspace.Name, &workspace.Description, &workspace.Uuid, &workspace.State,
			&workspace.WorkspaceType, &workspace.TemplateOkCheck, &workspace.Quantity, &workspace.NetworkUuid, &workspace.ComputeOfferingUuid,
			&workspace.TemplateUuid, &workspace.Postfix, &workspace.Shared, &workspace.CreateDate, &workspace.Removed)
		if err != nil {
			log.WithFields(logrus.Fields{
				"workspaceImpl": "selectWorkspaceList",
			}).Errorf("Workspace Select Query 이후 Scan 중 에러가 발생했습니다. [%v]", err)
		}

		workspaceList = append(workspaceList, workspace)
	}
	log.WithFields(logrus.Fields{
		"workspaceImpl": "selectWorkspaceList",
	}).Infof("selectWorkspaceList Query result [%v]", workspaceList)

	return workspaceList, err
}

func selectCountWorkspace() (int, error) {
	db, err := sql.Open(os.Getenv("MysqlType"), os.Getenv("DbInfo"))
	if err != nil {
		log.WithFields(logrus.Fields{
			"workspaceImpl": "selectCountWorkspace",
		}).Errorf("selectCountWorkspace DB Connect Error [%v]", err)
	}
	defer db.Close()
	log.WithFields(logrus.Fields{
		"workspaceImpl": "selectCountWorkspace",
	}).Info("selectCountWorkspace DB connect success")
	var workspaceCount int
	err = db.QueryRow("SELECT count(*) FROM workspaces where removed IS NULL").Scan(&workspaceCount)
	if err != nil {
		log.WithFields(logrus.Fields{
			"workspaceImpl": "selectCountWorkspace",
		}).Errorf("Workspace Count Select Query FAILED [%v]", err)
	}

	return workspaceCount, err
}

func selectInstanceList(uuid string, selectType string) ([]Instance, error) {
	log.WithFields(logrus.Fields{
		"workspaceImpl": "selectInstanceList",
	}).Infof("payload uuid [%v]", uuid)
	db, err := sql.Open(os.Getenv("MysqlType"), os.Getenv("DbInfo"))
	if err != nil {
		log.WithFields(logrus.Fields{
			"workspaceImpl": "selectInstanceList",
		}).Errorf("selectInstanceList DB connect error [%v]", err)
	}
	defer db.Close()
	log.WithFields(logrus.Fields{
		"workspaceImpl": "selectInstanceList",
	}).Info("selectInstanceList DB connect success")
	queryString := "SELECT" +
		" vi.id, vi.name, vi.uuid, vi.workspace_uuid, vi.mold_uuid, " +
		"IFNULL(vi.owner_account_id, '') as owner_account_id, vi.checked, vi.connected, vi.status, vi.create_date, " +
		"vi.checked_date, vi.workspace_name" +
		" FROM vm_instances AS vi" +
		" LEFT JOIN workspaces w on vi.workspace_uuid = w.uuid" +
		" WHERE vi.removed IS NULL"
	if selectType == WorkspaceString {
		if uuid != ALL {
			queryString = queryString + " AND vi.workspace_uuid = '" + uuid + "'"
		}
	} else if selectType == InstanceString {
		queryString = queryString + " AND vi.uuid = '" + uuid + "'"
	}
	queryString = queryString + " ORDER BY vi.id DESC"
	log.WithFields(logrus.Fields{
		"workspaceImpl": "selectInstanceList",
	}).Infof("selectInstanceList Query [%v]", queryString)
	rows, err := db.Query(queryString)
	if err != nil {
		log.WithFields(logrus.Fields{
			"workspaceImpl": "selectInstanceList",
		}).Errorf("selectInstanceList query FAILED [%v]", err)
	}
	defer rows.Close()
	var instanceList []Instance
	defer rows.Close()
	for rows.Next() {
		instance := Instance{}
		err = rows.Scan(
			&instance.Id, &instance.Name, &instance.Uuid, &instance.WorkspaceUuid, &instance.MoldUuid,
			&instance.OwnerAccountId, &instance.Checked, &instance.Connected, &instance.Status, &instance.CreateDate,
			&instance.CheckedDate, &instance.WorkspaceName)
		if err != nil {
			log.WithFields(logrus.Fields{
				"workspaceImpl": "selectInstanceList",
			}).Errorf("Instance Select Query 이후 Scan 중 에러가 발생했습니다. [%v]", err)
		}

		instanceList = append(instanceList, instance)
	}
	log.WithFields(logrus.Fields{
		"workspaceImpl": "selectInstanceList",
	}).Infof("selectWorkspaceList Query result [%v]", instanceList)

	return instanceList, err
}

func insertWorkspace(workspace Workspace) (map[string]interface{}, error) {
	db, err := sql.Open(os.Getenv("MysqlType"), os.Getenv("DbInfo"))
	resultData := map[string]interface{}{}
	if err != nil {
		log.WithFields(logrus.Fields{
			"workspaceImpl": "insertWorkspace",
		}).Errorf("insertWorkspace DB connect error [%v]", err)
		resultData["message"] = "DB connect error"
		resultData["status"] = BaseErrorCode
	}
	defer db.Close()

	result, err := db.Exec("INSERT INTO workspaces(name, description, uuid, workspace_type, network_uuid, compute_offering_uuid, template_uuid, create_date) VALUES (?, ?, ?, ?, ?, ?, ?,  NOW())",
		workspace.Name, workspace.Description, workspace.Uuid, workspace.WorkspaceType, workspace.NetworkUuid, workspace.ComputeOfferingUuid, workspace.TemplateUuid)
	if err != nil {
		log.WithFields(logrus.Fields{
			"workspaceImpl": "insertWorkspace",
		}).Errorf("UUID 생성 후 DB Insert 중 오류가 발생하였습니다. [%v]", err)
		resultData["message"] = "An error occurred while inserting the DB after generating the UUID."
		resultData["status"] = BaseErrorCode
	}
	n, err := result.RowsAffected()
	if n == 1 {
		log.Info("워크스페이스가 정상적으로 생성되었습니다.")
		resultData["message"] = "The workspace has been successfully created."
		resultData["status"] = http.StatusOK
	}

	return resultData, err
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

	result, err := db.Exec("INSERT INTO vm_instances(name, uuid, mold_uuid, workspace_uuid, workspace_name, create_date) VALUES (?, ?, ?, ?, ?, NOW())",
		instance.Name, instance.Uuid, instance.MoldUuid, instance.WorkspaceUuid, instance.WorkspaceName)
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

func updateWorkspaceQuantity(workspaceUuid string) map[string]interface{} {
	returnValue := map[string]interface{}{}
	db, err := sql.Open(os.Getenv("MysqlType"), os.Getenv("DbInfo"))
	if err != nil {
		fmt.Println("DB connect error")
		fmt.Println(err)
	}
	defer db.Close()
	result1, err := db.Exec("UPDATE workspaces set quantity=(select count(*) from vm_instances where removed is null and workspace_uuid = ?) where uuid=?", workspaceUuid, workspaceUuid)
	if err != nil {
		log.Error("워크스페이스 수량 업데이트 쿼리중 에러가 발생했습니다.")
		log.Error(err)
		returnValue["message"] = "An error occurred while querying the workspace Quantity update."
		returnValue["status"] = SQLQueryError
	}
	n, err := result1.RowsAffected()
	if n == 1 {
		log.Info("워크스페이스 수량이 업데이트 완료가 되었습니다.")
		returnValue["message"] = "Workspace Quantity update is complete."
		returnValue["status"] = http.StatusOK
	}
	return returnValue
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

func updateWorkspaceTemplateCheck(uuid string, typeString string) map[string]interface{} {
	db, err := sql.Open(os.Getenv("MysqlType"), os.Getenv("DbInfo"))
	resultReturn := map[string]interface{}{}
	if err != nil {
		log.WithFields(logrus.Fields{
			"workspaceImpl": "updateWorkspaceTemplateCheck",
		}).Errorf("DB Connect Error [%v]", err)
		resultReturn["message"] = MsgDBConnectError
		resultReturn["status"] = SQLConnectError
	}
	log.WithFields(logrus.Fields{
		"workspaceImpl": "updateWorkspaceTemplateCheck",
	}).Infof("uuid [%v]", uuid)
	defer db.Close()

	result, err := db.Exec("UPDATE workspaces set template_ok_check=?, state=? where uuid=?", typeString, Enable, uuid)
	if err != nil {
		log.Error(MsgDBConnectError)
		log.Error(err)
		resultReturn["message"] = MsgDBConnectError
		resultReturn["status"] = SQLQueryError
	}
	n1, _ := result.RowsAffected()
	log.Debugf("123123123 [%v]", n1)
	if n1 == 1 {
		resultReturn["message"] = "workspace template check OK"
		resultReturn["status"] = http.StatusOK
	}
	return resultReturn
}

func updateInstanceCheck(uuid string, loginInfo string, logoutInfo string) map[string]interface{} {
	db, err := sql.Open(os.Getenv("MysqlType"), os.Getenv("DbInfo"))
	resultReturn := map[string]interface{}{}
	resultReturn["status"] = http.StatusUnauthorized
	if err != nil {
		log.Error("DB connect error")
		log.Error(err)
		resultReturn["message"] = MsgDBConnectError
		resultReturn["status"] = BaseErrorCode
	}
	log.WithFields(logrus.Fields{
		"workspaceImpl": "updateInstanceCheck",
	}).Debugf("uuid [%v]", uuid)
	defer db.Close()

	loginInfoMap := map[string]interface{}{}
	logoutInfoMap := map[string]interface{}{}
	layout := "2006/01/02 15:04:05"
	connected := 0
	err1 := json.Unmarshal([]byte(loginInfo), &loginInfoMap)
	if err1 != nil {
		return nil
	}
	err2 := json.Unmarshal([]byte(logoutInfo), &logoutInfoMap)
	if err2 != nil {
		return nil
	}
	loginTime, _ := time.Parse(layout, loginInfoMap["time"].(string))
	logOutTime, _ := time.Parse(layout, logoutInfoMap["time"].(string))
	log.Debugf("loginInfoMap [%v], logoutInfoMap [%v]", loginInfoMap, logoutInfoMap)
	if logOutTime.Before(loginTime) {
		connected = 1
		log.Debugf("connected [%v]", logOutTime.Before(loginTime))
	}
	result, err := db.Exec("UPDATE vm_instances set checked=1, connected=?, checked_date=NOW(), status='Ready' where uuid=?", connected, uuid)
	if err != nil {
		log.Error(MsgDBConnectError)
		log.Error(err)
		resultReturn["message"] = MsgDBConnectError
		resultReturn["status"] = SQLQueryError
	}
	n1, _ := result.RowsAffected()
	if n1 == 1 {
		resultReturn["status"] = http.StatusOK
	}

	return resultReturn
}

func deleteInstance(instanceUuid string) map[string]interface{} {
	db, err := sql.Open(os.Getenv("MysqlType"), os.Getenv("DbInfo"))
	resultReturn := map[string]interface{}{}
	if err != nil {
		log.Error("DB connect error")
		log.Error(err)
		resultReturn["message"] = MsgDBConnectError
		resultReturn["status"] = BaseErrorCode
	}
	log.WithFields(logrus.Fields{
		"workspaceImpl": "deleteInstance",
	}).Infof("instanceUuid [%v]", instanceUuid)
	defer db.Close()

	result, err := db.Exec("UPDATE vm_instances set removed=NOW() where uuid=?", instanceUuid)
	if err != nil {
		log.Error(MsgDBConnectError)
		log.Error(err)
		resultReturn["message"] = MsgDBConnectError
		resultReturn["status"] = SQLQueryError
	}
	n1, _ := result.RowsAffected()
	if n1 == 1 {
		resultReturn["status"] = http.StatusOK
	}

	return resultReturn
}

func updateInstanceUser(instanceUuid string, userName string) map[string]interface{} {
	db, err := sql.Open(os.Getenv("MysqlType"), os.Getenv("DbInfo"))
	resultReturn := map[string]interface{}{}
	if err != nil {
		log.Error("DB connect error")
		log.Error(err)
		resultReturn["message"] = MsgDBConnectError
		resultReturn["status"] = BaseErrorCode
	}
	log.WithFields(logrus.Fields{
		"workspaceImpl": "deleteInstance",
	}).Infof("instanceUuid [%v], userName[%v]", instanceUuid, userName)
	defer db.Close()

	result, err := db.Exec("UPDATE vm_instances SET owner_account_id=? WHERE uuid=?", userName, instanceUuid)
	if err != nil {
		log.Error(MsgDBConnectError)
		log.Error(err)
		resultReturn["message"] = MsgDBConnectError
		resultReturn["status"] = SQLQueryError
	}
	n1, _ := result.RowsAffected()
	if n1 == 1 {
		resultReturn["status"] = http.StatusOK
	}

	return resultReturn
}

func updateInstanceChecked0() {
	db, err := sql.Open(os.Getenv("MysqlType"), os.Getenv("DbInfo"))
	resultReturn := map[string]interface{}{}
	if err != nil {
		log.Error("DB connect error")
		log.Error(err)
		resultReturn["message"] = MsgDBConnectError
		resultReturn["status"] = BaseErrorCode
	}
	log.WithFields(logrus.Fields{
		"workspaceImpl": "updateInstanceChecked0",
	}).Infof("updateInstanceChecked0 exec")
	defer db.Close()
	queryString := "UPDATE vm_instances" +
		" SET checked = 0," +
		" status = 'Not Ready'" +
		" WHERE uuid IN (" +
		"SELECT vmUuid.uuid FROM (" +
		"SELECT uuid, checked_date, connected, removed FROM vm_instances" +
		" WHERE removed IS NULL" +
		" AND checked = 1" +
		" AND checked_date < DATE_ADD(NOW(), INTERVAL -5 MINUTE)" +
		") vmUuid)"

	log.WithFields(logrus.Fields{
		"workspaceImpl": "updateInstanceChecked0",
	}).Debugf("updateInstanceChecked0 query [%v]", queryString)

	result, err := db.Exec(queryString)
	if err != nil {
		log.Error(MsgDBConnectError)
		log.Error(err)
		resultReturn["message"] = MsgDBConnectError
		resultReturn["status"] = SQLQueryError
	}
	n1, _ := result.RowsAffected()
	if n1 == 1 {
		resultReturn["status"] = http.StatusOK
	}

	time.Sleep(60 * time.Second)
}
