package main

import (
	"database/sql"
	"github.com/sirupsen/logrus"
	"os"
)

func selectUserDesktopWorkspaceList(userName string) ([]string, error) {
	log.WithFields(logrus.Fields{
		"userDesktopDao": "selectUserDesktopList",
	}).Infof("payload none")
	db, err := sql.Open(os.Getenv("MysqlType"), os.Getenv("DbInfo"))
	if err != nil {
		log.WithFields(logrus.Fields{
			"userDesktopDao": "selectUserDesktopList",
		}).Errorf("selectUserDesktopList DB Connect Error [%v]", err)
	}
	defer db.Close()
	log.WithFields(logrus.Fields{
		"userDesktopDao": "selectUserDesktopList",
	}).Infof("select UserDesktopList DB Connect success")
	queryString := "SELECT" +
		" DISTINCT workspace_uuid" +
		" FROM vm_instances" +
		" WHERE removed IS NULL" +
		" AND owner_account_id = '" + userName + "'"
	log.WithFields(logrus.Fields{
		"configurationDao": "selectUserDesktopList",
	}).Infof("select selectConfigurationList Query [%v]", queryString)
	rows, err := db.Query(queryString)
	if err != nil {
		log.WithFields(logrus.Fields{
			"configurationDao": "selectConfigurationList",
		}).Errorf("Configuration Select Query FAILED [%v]", err)
	}
	var workspaceList []string
	defer rows.Close()
	for rows.Next() {
		workspaceUuid := ""
		err = rows.Scan(&workspaceUuid)
		if err != nil {
			log.WithFields(logrus.Fields{
				"configurationDao": "selectUserDesktopList",
			}).Errorf("configuratio Select Query 이후 Scan 중 에러가 발생했습니다. [%v]", err)
		}

		workspaceList = append(workspaceList, workspaceUuid)
	}
	log.WithFields(logrus.Fields{
		"workspaceImpl": "selectUserDesktopList",
	}).Infof("selectUserDesktopList Query result [%v]", workspaceList)

	return workspaceList, err
}

func selectUserDesktopListForWorkspace(workspaceUuid string, userName string) ([]Instance, error) {
	log.WithFields(logrus.Fields{
		"userDesktopDao": "selectUserDesktopListForWorkspace",
	}).Infof("payload workspaceUuid [%v], userName [%v]", workspaceUuid, userName)
	db, err := sql.Open(os.Getenv("MysqlType"), os.Getenv("DbInfo"))
	if err != nil {
		log.WithFields(logrus.Fields{
			"userDesktopDao": "selectUserDesktopListForWorkspace",
		}).Errorf("selectInstanceList DB connect error [%v]", err)
	}
	defer db.Close()
	log.WithFields(logrus.Fields{
		"userDesktopDao": "selectUserDesktopListForWorkspace",
	}).Info("selectUserDesktopListForWorkspace DB connect success")
	queryString := "SELECT" +
		" vi.id, vi.name, vi.uuid, vi.workspace_uuid, vi.mold_uuid," +
		" IFNULL(vi.owner_account_id, '') as owner_account_id, vi.checked, vi.connected, vi.status, vi.create_date," +
		" vi.checked_date, vi.workspace_name, vi.handshake_status, vi.ipaddress, w.workspace_type" +
		" FROM vm_instances AS vi" +
		" LEFT JOIN workspaces w on vi.workspace_uuid = w.uuid" +
		" WHERE vi.removed IS NULL" +
		" AND vi.workspace_uuid = '" + workspaceUuid + "'" +
		" AND vi.owner_account_id = '" + userName + "'"
	queryString = queryString + " ORDER BY vi.id DESC"
	log.WithFields(logrus.Fields{
		"userDesktopDao": "selectUserDesktopListForWorkspace",
	}).Infof("selectUserDesktopListForWorkspace Query [%v]", queryString)
	rows, err := db.Query(queryString)
	if err != nil {
		log.WithFields(logrus.Fields{
			"userDesktopDao": "selectUserDesktopListForWorkspace",
		}).Errorf("selectUserDesktopListForWorkspace query FAILED [%v]", err)
	}
	defer rows.Close()
	var instanceList []Instance
	defer rows.Close()
	for rows.Next() {
		instance := Instance{}
		err = rows.Scan(
			&instance.Id, &instance.Name, &instance.Uuid, &instance.WorkspaceUuid, &instance.MoldUuid,
			&instance.OwnerAccountId, &instance.Checked, &instance.Connected, &instance.Status, &instance.CreateDate,
			&instance.CheckedDate, &instance.WorkspaceName, &instance.HandshakeStatus, &instance.Ipaddress, &instance.WorkspaceType)
		if err != nil {
			log.WithFields(logrus.Fields{
				"userDesktopDao": "selectUserDesktopListForWorkspace",
			}).Errorf("Instance Select Query 이후 Scan 중 에러가 발생했습니다. [%v]", err)
		}

		instanceList = append(instanceList, instance)
	}
	log.WithFields(logrus.Fields{
		"userDesktopDao": "selectUserDesktopListForWorkspace",
	}).Infof("selectUserDesktopListForWorkspace Query result [%v]", instanceList)

	return instanceList, err
}

func selectUserPassword(userName string) (string, error) {
	log.WithFields(logrus.Fields{
		"userDesktopDao": "selectUserPassword",
	}).Infof("payload none")
	db, err := sql.Open(os.Getenv("MysqlType"), os.Getenv("DbInfo"))
	if err != nil {
		log.WithFields(logrus.Fields{
			"userDesktopDao": "selectUserPassword",
		}).Errorf("selectUserPassword DB Connect Error [%v]", err)
	}
	defer db.Close()
	log.WithFields(logrus.Fields{
		"userDesktopDao": "selectUserPassword",
	}).Infof("select selectUserPassword DB Connect success")
	queryString := "SELECT" +
		" password" +
		" FROM users" +
		" WHERE removed IS NULL" +
		" AND user_name = '" + userName + "'"
	log.WithFields(logrus.Fields{
		"configurationDao": "selectUserPassword",
	}).Infof("select selectConfigurationList Query [%v]", queryString)
	var userPassword string
	err = db.QueryRow(queryString).Scan(&userPassword)
	//rows, err := db.Query(queryString)
	if err != nil {
		log.WithFields(logrus.Fields{
			"configurationDao": "selectConfigurationList",
		}).Errorf("Configuration Select Query FAILED [%v]", err)
	}
	//var workspaceList []string
	//defer rows.Close()
	//for rows.Next() {
	//	workspaceUuid := ""
	//	err = rows.Scan(&workspaceUuid)
	//	if err != nil {
	//		log.WithFields(logrus.Fields{
	//			"configurationDao": "selectUserDesktopList",
	//		}).Errorf("configuratio Select Query 이후 Scan 중 에러가 발생했습니다. [%v]", err)
	//	}
	//
	//	workspaceList = append(workspaceList, workspaceUuid)
	//}
	//log.WithFields(logrus.Fields{
	//	"workspaceImpl": "selectUserDesktopList",
	//}).Infof("selectUserDesktopList Query result [%v]", workspaceList)

	return userPassword, err
}

func selectConnectionRdpDesktop(instanceUuid string, userName string) (Instance, error) {
	log.WithFields(logrus.Fields{
		"userDesktopDao": "selectConnectionRdpDesktop",
	}).Infof("payload instanceUuid [%v], userName [%v]", instanceUuid, userName)
	db, err := sql.Open(os.Getenv("MysqlType"), os.Getenv("DbInfo"))
	if err != nil {
		log.WithFields(logrus.Fields{
			"userDesktopDao": "selectConnectionRdpDesktop",
		}).Errorf("selectInstanceList DB connect error [%v]", err)
	}
	defer db.Close()
	log.WithFields(logrus.Fields{
		"userDesktopDao": "selectConnectionRdpDesktop",
	}).Info("selectConnectionRdpDesktop DB connect success")
	queryString := "SELECT" +
		" vi.id, vi.name, vi.uuid, vi.workspace_uuid, vi.mold_uuid," +
		" IFNULL(vi.owner_account_id, '') as owner_account_id, vi.checked, vi.connected, vi.status, vi.create_date," +
		" vi.checked_date, vi.workspace_name, vi.handshake_status, vi.ipaddress, w.workspace_type," +
		" vi.hash" +
		" FROM vm_instances AS vi" +
		" LEFT JOIN workspaces w on vi.workspace_uuid = w.uuid" +
		" WHERE vi.removed IS NULL" +
		" AND vi.uuid = '" + instanceUuid + "'" +
		" AND vi.owner_account_id = '" + userName + "'"
	queryString = queryString + " ORDER BY vi.id DESC"
	log.WithFields(logrus.Fields{
		"userDesktopDao": "selectConnectionRdpDesktop",
	}).Infof("selectUserDesktopListForWorkspace Query [%v]", queryString)
	rows, err := db.Query(queryString)
	if err != nil {
		log.WithFields(logrus.Fields{
			"userDesktopDao": "selectConnectionRdpDesktop",
		}).Errorf("selectUserDesktopListForWorkspace query FAILED [%v]", err)
	}
	defer rows.Close()
	var instanceList []Instance
	for rows.Next() {
		instance := Instance{}
		err = rows.Scan(
			&instance.Id, &instance.Name, &instance.Uuid, &instance.WorkspaceUuid, &instance.MoldUuid,
			&instance.OwnerAccountId, &instance.Checked, &instance.Connected, &instance.Status, &instance.CreateDate,
			&instance.CheckedDate, &instance.WorkspaceName, &instance.HandshakeStatus, &instance.Ipaddress, &instance.WorkspaceType,
			&instance.Hash)
		if err != nil {
			log.WithFields(logrus.Fields{
				"userDesktopDao": "selectUserDesktopListForWorkspace",
			}).Errorf("Instance Select Query 이후 Scan 중 에러가 발생했습니다. [%v]", err)
		}

		instanceList = append(instanceList, instance)
	}
	log.WithFields(logrus.Fields{
		"userDesktopDao": "selectConnectionRdpDesktop",
	}).Infof("selectUserDesktopListForWorkspace Query result [%v]", instanceList)

	return instanceList[0], err
}
