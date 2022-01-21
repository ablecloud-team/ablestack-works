package main

import (
	"database/sql"
	"github.com/sirupsen/logrus"
	"net/http"
	"os"
)

func insertApplicationDao(application Application) map[string]interface{} {
	db, err := sql.Open(os.Getenv("MysqlType"), os.Getenv("DbInfo"))
	resultData := map[string]interface{}{}
	if err != nil {
		log.WithFields(logrus.Fields{
			"applicationDao": "insertApplicationDao",
		}).Errorf("DB connect error. error [%v]", err)
		resultData["message"] = "DB connect error"
		resultData["status"] = BaseErrorCode
	}
	defer db.Close()

	result, err := db.Exec("INSERT INTO application(uuid, workspace_uuid, application_server, path, description, working_directory, create_date) VALUES (?, ?, ?, ?, ?, ?, NOW())",
		application.Uuid, application.WorkspaceUuid, application.RegApplicationServer, application.Path, application.Description, application.WorkingDirectory)
	if err != nil {
		log.WithFields(logrus.Fields{
			"applicationDao": "insertApplicationDao",
		}).Errorf("An error occurred while registering the application. [%v]", err)
		resultData["message"] = "An error occurred while registering the application."
		resultData["status"] = SQLQueryError
	}
	n, err := result.RowsAffected()
	if n == 1 {
		log.WithFields(logrus.Fields{
			"applicationDao": "insertApplicationDao",
		}).Infof("An error occurred while registering the application")
		resultData["message"] = "An error occurred while registering the application"
		resultData["status"] = http.StatusOK
	}

	return resultData
}

func selectApplicationDao(workspaceUuid string) ([]Application, error) {
	log.WithFields(logrus.Fields{
		"applicationDao": "selectApplicationDao",
	}).Infof("payload workspaceUuid [%v]", workspaceUuid)
	db, err := sql.Open(os.Getenv("MysqlType"), os.Getenv("DbInfo"))
	if err != nil {
		log.WithFields(logrus.Fields{
			"applicationDao": "selectApplicationDao",
		}).Errorf("selectWorkspaceList DB Connect Error [%v]", err)
	}
	defer db.Close()
	log.WithFields(logrus.Fields{
		"applicationDao": "selectApplicationDao",
	}).Infof("select WorkspaceList DB Connect success")
	queryString := "SELECT" +
		" id, uuid, workspace_uuid, application_server, path," +
		" description, working_directory, create_date" +
		" FROM application" +
		" WHERE removed IS NULL"
	if workspaceUuid != ALL {
		queryString = queryString + " AND workspace_uuid = '" + workspaceUuid + "'"
	}
	queryString = queryString + " ORDER BY id DESC"
	log.WithFields(logrus.Fields{
		"applicationDao": "selectApplicationDao",
	}).Infof("select WorkspaceList Query [%v]", queryString)
	rows, err := db.Query(queryString)
	if err != nil {
		log.WithFields(logrus.Fields{
			"applicationDao": "selectApplicationDao",
		}).Errorf("Workspace Select Query FAILED [%v]", err)
	}
	var applicationList []Application
	defer rows.Close()
	for rows.Next() {
		application := Application{}
		err = rows.Scan(
			&application.id, &application.Uuid, &application.WorkspaceUuid, &application.RegApplicationServer, &application.Path,
			&application.Description, &application.WorkingDirectory, &application.CreateDate)
		if err != nil {
			log.WithFields(logrus.Fields{
				"applicationDao": "selectApplicationDao",
			}).Errorf("An error occurred during Scan after Application Select Query. [%v]", err)
		}

		applicationList = append(applicationList, application)
	}
	log.WithFields(logrus.Fields{
		"applicationDao": "selectApplicationDao",
	}).Infof("selectWorkspaceList Query result [%v]", applicationList)

	return applicationList, err
}

func deleteApplicationDao(application Application) map[string]interface{} {
	db, err := sql.Open(os.Getenv("MysqlType"), os.Getenv("DbInfo"))
	resultReturn := map[string]interface{}{}
	if err != nil {
		log.Error("DB connect error")
		log.Error(err)
		resultReturn["message"] = MsgDBConnectError
		resultReturn["status"] = BaseErrorCode
	}
	log.WithFields(logrus.Fields{
		"applicationDao": "deleteApplicationDao",
	}).Infof("application [%v]", application)
	defer db.Close()

	result, err := db.Exec("UPDATE application set removed=NOW() where uuid=?", application.Uuid)
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
