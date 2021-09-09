package main

import (
	"database/sql"
	"fmt"
	"github.com/sirupsen/logrus"
	"net/http"
	"os"
	"strconv"
	"time"
)

const (
	VMStart  = "VMStart"
	VMsStart = "VMsStart"
	VMStop   = "VMStop"
	VMsStop  = "VMsStop"
	//VMDeploy = "VMDeploy"
	VMsDeploy  = "VMsDeploy"
	VMDestroy  = "VMDestroy"
	VMsDestroy = "VMsDestroy"
)

type AsyncJob struct {
	Id, Name, ExecUuid, Parameter string
	CreateDate                    time.Time
	Ready                         int
}

func asyncJobMonitoring() {
	db, err := sql.Open(os.Getenv("MysqlType"), os.Getenv("DbInfo"))
	if err != nil {
		fmt.Println("DB connect error")
		fmt.Println(err)
	}
	defer db.Close()
	for {
		var count int
		err = db.QueryRow("SELECT count(*) FROM async_job where ready = 1").Scan(&count)
		if err != nil {
			log.Error("async_job 테이블 조회중 에러가 발생했습니다.")
			log.Error(err)
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
	if asyncJob.Name == VMDestroy {
		log.Info("Async Job VM Destroy Execution.")
		params := []MoldParams{
			{"id": asyncJob.ExecUuid},
		}
		result := getDestroyVirtualMachine(params)
		log.Info(result["destroyvirtualmachineresponse"].(map[string]interface{})["jobid"].(string))
		if result["destroyvirtualmachineresponse"].(map[string]interface{})["jobid"].(string) != "" {
			deleteAsyncJob(asyncJob.Id)
		}
	} else if asyncJob.Name == VMsDeploy {
		log.Info("Async Job VMs Deploy Execution.")
		deployQuantity, _ := strconv.Atoi(asyncJob.Parameter)
		resultQuantity := 0
		resultFailQuantity := 0
		for i := 1; i <= deployQuantity; i++ {
			resultSelectWorkspaceDetail := selectWorkspaceDetail(asyncJob.ExecUuid)
			log.WithFields(logrus.Fields{"asyncJob.go": "resultSelectWorkspaceDetail"}).Infof("%v", resultSelectWorkspaceDetail)
			for y := 1; y <= 3; y++ {
				resultGetDeployVirtualMachine := getDeployVirtualMachine(resultSelectWorkspaceDetail.Uuid)
				log.WithFields(logrus.Fields{"asyncJob.go": "resultGetDeployVirtualMachine"}).Infof("%v", resultGetDeployVirtualMachine)
				time.Sleep(time.Second * 10)
				if resultGetDeployVirtualMachine["deployvirtualmachineresponse"].(map[string]interface{})["jobid"] == nil {
					log.WithFields(logrus.Fields{"asyncJob.go": "VMsDeploy 실패"}).Infof("%v", resultGetDeployVirtualMachine)
					resultFailQuantity = resultFailQuantity + 1
				} else if resultGetDeployVirtualMachine["deployvirtualmachineresponse"].(map[string]interface{})["jobid"].(string) != "" {
					resultPostfixFill := postfixFill(resultSelectWorkspaceDetail.Postfix)
					instance := Instance{}
					instance.Uuid = resultGetDeployVirtualMachine["deployvirtualmachineresponse"].(map[string]interface{})["id"].(string)
					instance.Name = resultSelectWorkspaceDetail.Name + "-" + resultPostfixFill
					instance.WorkspaceUuid = resultSelectWorkspaceDetail.Uuid
					resultInsertInstance := insertInstance(instance)
					params := []MoldParams{
						{"resourceids": instance.Uuid},
						{"tags[0].key": ServiceDaaS},
						{"tags[0].value": AblecloudWorks},
						{"tags[1].key": WorkspaceName},
						{"tags[1].value": resultSelectWorkspaceDetail.Name},
					}
					aaa := getCreateTags(params)
					log.Info("00000000000000000000000000000000000")
					log.Info(aaa)
					log.Info("00000000000000000000000000000000000")

					updateWorkspacePostfix(resultSelectWorkspaceDetail.Uuid, resultSelectWorkspaceDetail.Postfix)

					log.Info("The virtual machine has been successfully created.")
					log.Info(resultInsertInstance)

					break
				}
			}
			if resultFailQuantity == 3 {
				updateWorkspacePostfix(resultSelectWorkspaceDetail.Uuid, resultSelectWorkspaceDetail.Postfix)
			}
		}
		log.Infof("%v개의 가상머신 async_job 등록이 완료 되었습니다..\n", resultQuantity)
		deleteAsyncJob(asyncJob.Id)
	}
	log.Info("Async Job Execution end")
}

func selectAsyncJob(asyncJobUuid string) (AsyncJob, error) {
	db, err := sql.Open(os.Getenv("MysqlType"), os.Getenv("DbInfo"))
	if err != nil {
		fmt.Println("DB connect error")
		fmt.Println(err)
	}
	defer db.Close()
	asyncJob := AsyncJob{}
	err = db.QueryRow("SELECT id, name, exec_uuid, ready, parameter, create_date FROM async_job where id = ?", asyncJobUuid).Scan(&asyncJob.Id, &asyncJob.Name, &asyncJob.ExecUuid, &asyncJob.Ready, &asyncJob.Parameter, &asyncJob.CreateDate)
	log.Info("asyncJob")
	log.Info(asyncJob)

	return asyncJob, err
}

func insertAsyncJob(asyncJob AsyncJob) map[string]interface{} {
	asyncJob.Id = getUuid()
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
		log.Errorf("Async Job 등록중 에러가 발생했습니다.\n")
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
		returnValue["message"] = "DB connect error"
		returnValue["status"] = BaseErrorCode
	}
	defer db.Close()
	result, err := db.Exec("DELETE FROM async_job where id =?", asyncJobUuid)
	if err != nil {
		log.Error("Async 데이터 삭제에 실패했습니다.")
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
