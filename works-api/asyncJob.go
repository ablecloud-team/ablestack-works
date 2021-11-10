package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/sirupsen/logrus"
	"net/http"
	"os"
	"strconv"
	"time"
)

type AsyncJob struct {
	Id         string `json:"id"`
	Name       string `json:"name"`
	ExecUuid   string `json:"exec_uuid"`
	Parameter  string `json:"parameter"`
	CreateDate string `json:"create_date"`
	Ready      int    `json:"ready"`
}

func asyncJobMonitoring() {
	db, err := sql.Open(os.Getenv("MysqlType"), os.Getenv("DbInfo"))
	log.WithFields(logrus.Fields{
		"asyncJob.go": "asyncJobMonitoring",
	}).Infof("Exec")
	if err != nil {
		log.WithFields(logrus.Fields{
			"asyncJob.go": "asyncJobMonitoring",
		}).Errorf("DB connect error")
	}
	defer db.Close()
	for {
		log.Debugf("asyncjob exec")
		var count int
		err = db.QueryRow("SELECT count(*) FROM async_job where ready = 1").Scan(&count)
		if err != nil {
			log.WithFields(logrus.Fields{
				"asyncJob.go": "asyncJobMonitoring",
			}).Errorf("async_job 테이블 조회중 에러가 발생했습니다. [%v]", err)
		}
		if count > 0 {
			log.Info("Async Job Execution")
			asyncJobExec()
		}
		time.Sleep(10 * time.Second)
	}
}

func asyncJobExec() {
	log.Info("Async Job Execution start")
	db, err := sql.Open(os.Getenv("MysqlType"), os.Getenv("DbInfo"))
	if err != nil {
		fmt.Println("DB connect error")
		fmt.Println(err)
	}
	defer db.Close()
	asyncJob := AsyncJob{}
	err = db.QueryRow("SELECT id, name, exec_uuid, ready, parameter, create_date FROM async_job where ready = 1 order by create_date limit 1").Scan(&asyncJob.Id, &asyncJob.Name, &asyncJob.ExecUuid, &asyncJob.Ready, &asyncJob.Parameter, &asyncJob.CreateDate)
	log.WithFields(logrus.Fields{"asyncJob.go": "asyncJob 정보"}).Infof("%v", asyncJob)

	if asyncJob.Name == VMDestroy { //instance 삭제
		log.Info("Async Job VM Destroy Execution.")
		instanceList, _ := selectInstanceList(asyncJob.ExecUuid, InstanceString)
		if instanceList == nil {
			return
		} else {
			instanceInfo := instanceList[0]
			params := []MoldParams{
				{"id": instanceInfo.MoldUuid},
			}
			result := getDestroyVirtualMachine(params)
			log.Info(result["destroyvirtualmachineresponse"].(map[string]interface{})["jobid"])
			if result["destroyvirtualmachineresponse"].(map[string]interface{})["jobid"] != nil {
				workspaceList, _ := selectWorkspaceList(instanceInfo.WorkspaceUuid)
				workspaceInfo := workspaceList[0]
				deleteAsyncJob(asyncJob.Id)
				deleteInstance(asyncJob.ExecUuid)
				updateWorkspaceQuantity(workspaceInfo.Uuid)
			} else {
				deleteAsyncJob(asyncJob.Id)

			}
		}

	} else if asyncJob.Name == VMsDeploy { //instance 추가
		log.Info("Async Job VMs Deploy Execution.")
		deployQuantity, _ := strconv.Atoi(asyncJob.Parameter)
		resultQuantity := 0
		resultFailQuantity := 0
		workspaceList, _ := selectWorkspaceList(asyncJob.ExecUuid)
		workspaceInfo := workspaceList[0]
		for i := 1; i <= deployQuantity; i++ {
			log.Infof("resultWorkspaceInfo [%v]", workspaceInfo)
			instanceUuid := getUuid()
			resultGetDeployVirtualMachine := getDeployVirtualMachine(workspaceInfo.Uuid, instanceUuid, InstanceString)
			log.WithFields(logrus.Fields{"asyncJob.go": "resultGetDeployVirtualMachine"}).Infof("%v", resultGetDeployVirtualMachine)
			time.Sleep(time.Second * 10)
			if resultGetDeployVirtualMachine["deployvirtualmachineresponse"].(map[string]interface{})["jobid"] == nil {
				log.WithFields(logrus.Fields{"asyncJob.go": "VMsDeploy 실패"}).Infof("%v", resultGetDeployVirtualMachine)
				resultFailQuantity = resultFailQuantity + 1
			} else if resultGetDeployVirtualMachine["deployvirtualmachineresponse"].(map[string]interface{})["jobid"].(string) != "" {
				updateWorkspacePostfix(workspaceInfo.Uuid, workspaceInfo.Postfix)
				paramsMold := []MoldParams{
					{"id": resultGetDeployVirtualMachine["deployvirtualmachineresponse"].(map[string]interface{})["id"].(string)},
				}
				resultMoldInstanceInfo := getListVirtualMachinesMetrics(paramsMold)
				listVirtualMachinesMetrics := ListVirtualMachinesMetrics{}
				virtualMachineInfo, _ := json.Marshal(resultMoldInstanceInfo["listvirtualmachinesmetricsresponse"])
				json.Unmarshal([]byte(virtualMachineInfo), &listVirtualMachinesMetrics)

				instance := Instance{}
				instance.Uuid = instanceUuid
				instance.MoldUuid = listVirtualMachinesMetrics.Virtualmachine[0].Id
				instance.Name = listVirtualMachinesMetrics.Virtualmachine[0].Displayname
				instance.WorkspaceUuid = workspaceInfo.Uuid
				instance.WorkspaceName = workspaceInfo.Name
				instance.Ipaddress = listVirtualMachinesMetrics.Virtualmachine[0].Ipaddress
				resultInsertInstance := insertInstance(instance)
				params := []MoldParams{
					{"resourceids": instance.MoldUuid},
					{"resourcetype": UserVm},
					{"tags[0].key": ServiceDaaS},
					{"tags[0].value": AblecloudWorks},
					{"tags[1].key": WorkspaceName},
					{"tags[1].value": workspaceInfo.Name},
					{"tags[2].key": ClusterName},
					{"tags[2].value": os.Getenv("ClutsterName")},
				}
				resultGetCreateTags := getCreateTags(params)
				log.Infof("Create Tag Result [%v], params [%v]",resultGetCreateTags, params)

				log.Info("The virtual machine has been successfully created.")
				log.Info(resultInsertInstance)
				go handshakeVdi(instance, InstanceString)
			}
			updateWorkspaceQuantity(workspaceInfo.Uuid)
		}
		log.Infof("%v개의 가상머신 async_job 등록이 완료 되었습니다..\n", resultQuantity)
		deleteAsyncJob(asyncJob.Id)
	} else if asyncJob.Name == VMStop { //instance stop
		log.Info("Async Job VM Stop Execution.")
		instanceList, _ := selectInstanceList(asyncJob.ExecUuid, InstanceString)
		instanceInfo := instanceList[0]
		params := []MoldParams{
			{"id": instanceInfo.MoldUuid},
		}
		resultdata := getStopVirtualMachine(params)
		log.Info(resultdata)
		//stopvirtualmachineresponse:map[jobid:ff6e9d7a-63ea-4165-b843-5514b3327fdb]] (/Users/smlee/github/ablestack-desktop/works-api/asyncJ
		if resultdata["stopvirtualmachineresponse"].(map[string]interface{})["jobid"].(string) != "" {
			deleteAsyncJob(asyncJob.Id)
		}
	} else if asyncJob.Name == VMStart { //instance start
		log.Info("Async Job VM Start Execution.")
		instanceList, _ := selectInstanceList(asyncJob.ExecUuid, InstanceString)
		instanceInfo := instanceList[0]
		params := []MoldParams{
			{"id": instanceInfo.MoldUuid},
		}
		resultdata := getStartVirtualMachine(params)
		log.Info(resultdata)
		if resultdata["startvirtualmachineresponse"].(map[string]interface{})["jobid"].(string) != "" {
			deleteAsyncJob(asyncJob.Id)
		}
	}
	log.Info("Async Job Execution end")
}

func selectAsyncJob(execUuid string) (AsyncJob, error) {
	db, err := sql.Open(os.Getenv("MysqlType"), os.Getenv("DbInfo"))
	if err != nil {
		fmt.Println("DB connect error")
		fmt.Println(err)
	}
	defer db.Close()
	asyncJob := AsyncJob{}
	err = db.QueryRow("SELECT id, name, exec_uuid, ready, parameter, create_date FROM async_job where exec_uuid = ?", execUuid).Scan(&asyncJob.Id, &asyncJob.Name, &asyncJob.ExecUuid, &asyncJob.Ready, &asyncJob.Parameter, &asyncJob.CreateDate)
	log.Info("asyncJob")
	log.Info(asyncJob)

	return asyncJob, err
}

func insertAsyncJob(asyncJob AsyncJob) map[string]interface{} {
	db, err := sql.Open(os.Getenv("MysqlType"), os.Getenv("DbInfo"))
	resultReturn := map[string]interface{}{}
	if err != nil {
		log.Error("DB connect error")
		log.Error(err)
		resultReturn["message"] = "DB connect error"
	}
	defer db.Close()
	result, err := db.Exec("INSERT INTO async_job(id, name, exec_uuid, ready, parameter, create_date) VALUES (?, ?, ?, ?, ?, now())",
		asyncJob.Id, asyncJob.Name, asyncJob.ExecUuid, asyncJob.Ready, asyncJob.Parameter)
	if err != nil {
		log.Error("Async Job 등록중 에러가 발생했습니다.")
		log.Error(err)
		resultReturn["message"] = "An error occurred while registering Async Job."
	}
	n, _ := result.RowsAffected()
	if n == 1 {
		log.Info(result.LastInsertId())
		log.Info("async job 이 정상적으로 등록되었습니다.")
		resultReturn["status"] = http.StatusOK
		resultReturn["jobid"] = asyncJob.Id
		resultReturn["message"] = "The async job has been successfully registered."
	} else {
		resultReturn["status"] = BaseErrorCode
		resultReturn["message"] = "Failed to register async job."
	}

	return resultReturn
}

func updateAsyncJobRead(asyncJobUuid string, readyValue int) map[string]interface{} {
	db, err := sql.Open(os.Getenv("MysqlType"), os.Getenv("DbInfo"))
	returnValue := map[string]interface{}{}
	if err != nil {
		log.Errorf("%v\n", "DB connect error")
		log.Errorf("%v\n", err)
		returnValue["status"] = BaseErrorCode
		returnValue["message"] = "DB connect error"
	}
	defer db.Close()
	result, err := db.Exec("UPDATE async_job set ready=? where id=?", readyValue, asyncJobUuid)
	if err != nil {
		log.Errorf("Async Job 업데이트중 에러가 발생했습니다.\n")
		log.Errorf("%v\n", err)
		returnValue["status"] = BaseErrorCode
		returnValue["message"] = "An error occurred while registering Async Job."
	}
	n, _ := result.RowsAffected()
	if n == 1 {
		log.Infof("%v\n", result)
		log.Infof("async job이 정상적으로 업데이트 되였습니다.\n")
		returnValue["status"] = http.StatusOK
		returnValue["message"] = "async job has been updated normally."
	}

	return returnValue
}

func deleteAsyncJob(asyncJobUuid string) map[string]interface{} {
	db, err := sql.Open(os.Getenv("MysqlType"), os.Getenv("DbInfo"))
	returnValue := map[string]interface{}{}
	if err != nil {
		log.Errorf("%v\n", "DB connect error")
		log.Errorf("%v\n", err)
		returnValue["message"] = MsgDBConnectError
		returnValue["status"] = SQLConnectError
	}
	defer db.Close()
	result, err := db.Exec("DELETE FROM async_job where id =?", asyncJobUuid)
	if err != nil {
		log.Error("Async data deletion failed.")
		returnValue["message"] = "Async data deletion failed."
		returnValue["status"] = BaseErrorCode
	}
	n, _ := result.RowsAffected()
	if n > 1 {
		log.Info("Async 데이터를 삭제했습니다.")
		returnValue["message"] = "Deleted Async data."
		returnValue["status"] = http.StatusOK
	}

	return returnValue
}
