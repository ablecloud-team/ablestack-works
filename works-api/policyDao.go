package main

import (
	"database/sql"
	"github.com/sirupsen/logrus"
	"net/http"
	"os"
)

type Policy struct {
	Id             int    `json:"id"`
	WorkspacesUuid string `json:"workspaces_uuid"`
	Name           string `json:"name"`
	Description    string `json:"description"`
	I18n           string `json:"i18N"`
	Value          string `json:"value"`
	Type           string `json:"type"`
	SettableValue  string `json:"settable_value"`
	SettableType   string `json:"settable_type"`
	AdGroupName    string `json:"ad_group_name"`
}

type PolicyName struct {
	Id                       int    `json:"id"`
	WorkspaceUuid            string `json:"workspace_uuid"`
	RdpPort                  int    `json:"rdp_port"`
	RdpAccessAllow           int    `json:"rdp_access_allow"`
	ClipboardRedirection     int    `json:"clipboard_redirection"`
	WindowsInstaller         int    `json:"windows_installer"`
	NetworkFileShare         int    `json:"network_file_share"`
	NetworkPrinterShare      int    `json:"network_printer_share"`
	LocalPrinter             int    `json:"local_printer"`
	WindowsAutoUpdate        int    `json:"windows_auto_update"`
	RemovableStorage         int    `json:"removable_storage"`
	CmdUse                   int    `json:"cmd_use"`
	SettingsUse              int    `json:"settings_use"`
	PcPowerUse               int    `json:"pc_power_use"`
	Remotefx                 int    `json:"remotefx"`
	Console                  int    `json:"console"`
	InitialProgram           int    `json:"initial_program"`
	ServerLayout             int    `json:"server_layout"`
	ColorDepth               int    `json:"color_depth"`
	Width                    int    `json:"width"`
	Height                   int    `json:"height"`
	Dpi                      int    `json:"dpi"`
	ResizeMethod             string `json:"resize_method"`
	ForceLossless            int    `json:"force_lossless"`
	DisableAudio             int    `json:"disable_audio"`
	EnableAudioInput         int    `json:"enable_audio_input"`
	EnablePrinting           int    `json:"enable_printing"`
	PrinterName              string `json:"printer_name"`
	EnableDrive              int    `json:"enable_drive"`
	DisableDownload          int    `json:"disable_download"`
	DisableUpload            int    `json:"disable_upload"`
	DrivePath                string `json:"drive_path"`
	CreateDrivePath          int    `json:"create_drive_path"`
	ConsoleAudio             int    `json:"console_audio"`
	EnableWallpaper          int    `json:"enable_wallpaper"`
	EnableTheming            int    `json:"enable_theming"`
	EnableFontSmoothing      int    `json:"enable_font_smoothing"`
	EnableFullWindowDrag     int    `json:"enable_full_window_drag"`
	EnableDesktopComposition int    `json:"enable_desktop_composition"`
	EnableMenuAnimations     int    `json:"enable_menu_animations"`
}

func selectDefaultPolicyList() ([]Policy, error) {
	log.WithFields(logrus.Fields{
		"policyDao": "selectDefaultPolicyList",
	}).Infof("payload")
	db, err := sql.Open(os.Getenv("MysqlType"), os.Getenv("DbInfo"))
	if err != nil {
		log.WithFields(logrus.Fields{
			"policyDao": "selectDefaultPolicyList",
		}).Errorf("selectDefaultPolicyList DB Connect Error [%v]", err)
	}
	defer db.Close()
	log.WithFields(logrus.Fields{
		"policyDao": "selectDefaultPolicyList",
	}).Infof("select selectDefaultPolicyList DB Connect success")

	policyString := " id, name, description, i18n, IFNULL(value,'')," +
		" type, settable_type, IFNULL(settable_value,''), IFNULL(ad_group_name,'')"

	queryString := "SELECT" +
		policyString +
		" FROM policy_list" +
		" ORDER BY id DESC"

	log.WithFields(logrus.Fields{
		"policyDao": "selectDefaultPolicyList",
	}).Infof("select selectDefaultPolicyList Query [%v]", queryString)
	rows, err := db.Query(queryString)
	if err != nil {
		log.WithFields(logrus.Fields{
			"policyDao": "selectDefaultPolicyList",
		}).Errorf("selectDefaultPolicyList Select Query FAILED [%v]", err)
	}
	var policyList []Policy
	defer rows.Close()
	for rows.Next() {
		policy := Policy{}
		err = rows.Scan(
			&policy.Id, &policy.Name, &policy.Description, &policy.I18n, &policy.Value,
			&policy.Type, &policy.SettableType, &policy.SettableValue, &policy.AdGroupName)
		if err != nil {
			log.WithFields(logrus.Fields{
				"policyDao": "selectDefaultPolicyList",
			}).Errorf("WorkspacePolicy Select Query 이후 Scan 중 에러가 발생했습니다. [%v]", err)
		}

		policyList = append(policyList, policy)
	}
	log.WithFields(logrus.Fields{
		"policyDao": "selectDefaultPolicyList",
	}).Infof("selectDefaultPolicyList Query result [%v]", policyList)

	return policyList, err
}

func selectDefaultPolicyInfo(policyName string) (Policy, error) {
	log.WithFields(logrus.Fields{
		"policyDao": "selectDefaultPolicyList",
	}).Infof("payload")
	db, err := sql.Open(os.Getenv("MysqlType"), os.Getenv("DbInfo"))
	if err != nil {
		log.WithFields(logrus.Fields{
			"policyDao": "selectDefaultPolicyList",
		}).Errorf("selectDefaultPolicyList DB Connect Error [%v]", err)
	}
	defer db.Close()
	log.WithFields(logrus.Fields{
		"policyDao": "selectDefaultPolicyList",
	}).Infof("select selectDefaultPolicyList DB Connect success")

	policyString := " id, name, description, i18n, value," +
		" type, settable_type, settable_value, ad_group_name"

	queryString := "SELECT" +
		policyString +
		" FROM policy_list" +
		" WHERE name = ?" +
		" ORDER BY id DESC"

	log.WithFields(logrus.Fields{
		"policyDao": "selectDefaultPolicyList",
	}).Infof("select selectDefaultPolicyList Query [%v]", queryString)

	var policy Policy

	err = db.QueryRow(queryString, policyName).
		Scan(&policy.Id, &policy.Name, &policy.Description, &policy.I18n, &policy.Value,
			&policy.Type, &policy.SettableType, &policy.SettableValue, &policy.AdGroupName)

	rows, err := db.Query(queryString, policyName)

	if err != nil {
		log.WithFields(logrus.Fields{
			"policyDao": "selectDefaultPolicyList",
		}).Errorf("selectDefaultPolicyList Select Query FAILED [%v]", err)
	}
	defer rows.Close()
	//var policyList []Policy
	//for rows.Next() {
	//	policy := Policy{}
	//	err = rows.Scan(
	//		&policy.Id, &policy.Name, &policy.Description, &policy.I18n, &policy.Value,
	//		&policy.Type, &policy.SettableType, &policy.SettableValue, &policy.AdGroupName)
	//	if err != nil {
	//		log.WithFields(logrus.Fields{
	//			"policyDao": "selectDefaultPolicyList",
	//		}).Errorf("WorkspacePolicy Select Query 이후 Scan 중 에러가 발생했습니다. [%v]", err)
	//	}
	//
	//	policyList = append(policyList, policy)
	//}
	log.WithFields(logrus.Fields{
		"policyDao": "selectDefaultPolicyList",
	}).Infof("selectDefaultPolicyList Query result [%v]", policyList)

	return policy, err
}

func selectWorkspacesPolicyList(workspaceUuid string) ([]Policy, error) {
	log.WithFields(logrus.Fields{
		"policyDao": "selectWorkspacesPolicyList",
	}).Infof("payload")
	db, err := sql.Open(os.Getenv("MysqlType"), os.Getenv("DbInfo"))
	if err != nil {
		log.WithFields(logrus.Fields{
			"policyDao": "selectWorkspacesPolicyList",
		}).Errorf("selectDefaultPolicyList DB Connect Error [%v]", err)
	}
	defer db.Close()
	log.WithFields(logrus.Fields{
		"policyDao": "selectWorkspacesPolicyList",
	}).Infof("select selectDefaultPolicyList DB Connect success")

	policyString := " id, name, workspaces_uuid, value"

	queryString := "SELECT" +
		policyString +
		" FROM workspaces_policy" +
		" WHERE workspaces_uuid = ?" +
		" ORDER BY id DESC"
	log.WithFields(logrus.Fields{
		"policyDao": "selectWorkspacesPolicyList",
	}).Infof("select selectWorkspacesPolicyList Query [%v]", queryString)
	rows, err := db.Query(queryString, workspaceUuid)
	if err != nil {
		log.WithFields(logrus.Fields{
			"policyDao": "selectWorkspacesPolicyList",
		}).Errorf("selectWorkspacesPolicyList Select Query FAILED [%v]", err)
	}
	var policyList []Policy
	defer rows.Close()
	for rows.Next() {
		policy := Policy{}
		err = rows.Scan(
			&policy.Id, &policy.Name, &policy.WorkspacesUuid, &policy.Value)
		if err != nil {
			log.WithFields(logrus.Fields{
				"policyDao": "selectWorkspacesPolicyList",
			}).Errorf("WorkspacePolicy Select Query 이후 Scan 중 에러가 발생했습니다. [%v]", err)
		}

		policyList = append(policyList, policy)
	}
	log.WithFields(logrus.Fields{
		"policyDao": "selectWorkspacesPolicyList",
	}).Infof("selectWorkspacesPolicyList Query result [%v]", policyList)

	return policyList, err
}

func insertPolicy(policy Policy) (map[string]interface{}, error) {
	log.WithFields(logrus.Fields{
		"policyDao": "insertPolicy",
	}).Infof("payload")
	db, err := sql.Open(os.Getenv("MysqlType"), os.Getenv("DbInfo"))
	resultData := map[string]interface{}{}
	if err != nil {
		log.WithFields(logrus.Fields{
			"workspaceImpl": "insertPolicy",
		}).Errorf("insertWorkspace DB connect error [%v]", err)
		resultData["message"] = "DB connect error"
		resultData["status"] = BaseErrorCode
	}
	defer db.Close()

	result, err := db.Exec("INSERT INTO workspaces_policy(name,workspaces_uuid, value) VALUES (?, ?, ?)",
		policy.Name, policy.WorkspacesUuid, policy.Value)
	if err != nil {
		log.WithFields(logrus.Fields{
			"workspaceImpl": "insertPolicy",
		}).Errorf("변경된 policy 저장중 에러가 발생 했습니다. [%v]", err)
		resultData["message"] = "An error occurred while inserting the DB after generating the UUID."
		resultData["status"] = BaseErrorCode
	}
	n, err := result.RowsAffected()
	if n == 1 {
		log.Info("변경된 policy 저장이 정상적으로 진행되었습니다.")
		resultData["message"] = "The workspace has been successfully created."
		resultData["status"] = http.StatusOK
	}

	return resultData, err
}

func deleteWorkspacesPolicy(policy Policy) map[string]interface{} {
	db, err := sql.Open(os.Getenv("MysqlType"), os.Getenv("DbInfo"))
	returnValue := map[string]interface{}{}
	if err != nil {
		log.Errorf("%v\n", "DB connect error")
		log.Errorf("%v\n", err)
		returnValue["message"] = MsgDBConnectError
		returnValue["status"] = SQLConnectError
	}
	defer db.Close()
	result, err := db.Exec("DELETE FROM workspaces_policy WHERE name =? AND workspaces_uuid = ?", policy.Name, policy.WorkspacesUuid)
	if err != nil {
		log.Error("deleteWorkspacesPolicy delete failed.")
		returnValue["message"] = "deleteWorkspacesPolicy delete failed."
		returnValue["status"] = BaseErrorCode
	}
	n, _ := result.RowsAffected()
	if n > 1 {
		log.Info("워크스페이스 정책 데이터를 삭제했습니다.")
		returnValue["message"] = "Deleted deleteWorkspacesPolicy"
		returnValue["status"] = http.StatusOK
	}

	return returnValue
}
