package main

import (
	"database/sql"
	"fmt"
	"os"
	"time"
)

type AsyncJob struct {
	Id, Name, ExecUuid, Parameter string
	CreateDate                    time.Time
	Ready                         int
}

func asyncJobMonitoring() {
	db, err := sql.Open(os.Getenv("MsqlType"), os.Getenv("DbInfo"))
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
	db, err := sql.Open(os.Getenv("MsqlType"), os.Getenv("DbInfo"))
	if err != nil {
		fmt.Println("DB connect error")
		fmt.Println(err)
	}
	defer db.Close()
	asyncJob := AsyncJob{}
	err = db.QueryRow("SELECT id, name, exec_uuid, ready, parameter, create_date FROM async_job where ready = 1 order by create_date limit 1").Scan(&asyncJob.Id, &asyncJob.Name, &asyncJob.ExecUuid, &asyncJob.Ready, &asyncJob.Parameter, &asyncJob.CreateDate)
	log.Info("asyncJob")
	log.Info(asyncJob)
	if asyncJob.Name == "VMDestroy" {
		log.Info("Async Job VM Destroy Execution.")
		params := []MoldParams{
			{"id": asyncJob.ExecUuid},
		}
		result := getDestroyVirtualMachine(params)
		log.Info("123123123123123123123")
		log.Info(result["destroyvirtualmachineresponse"].(map[string]interface{})["jobid"].(string))
		if result["destroyvirtualmachineresponse"].(map[string]interface{})["jobid"].(string) != "" {
			result, err := db.Exec("DELETE FROM async_job where id =?", asyncJob.Id)
			if err != nil {
				log.Error("Async 데이터 삭제에 실패했습니다.")
			}
			n, _ := result.RowsAffected()
			if n > 1 {
				log.Info("Async 데이터를 삭제했습니다.")
			}
		}
	}
	log.Info("Async Job Execution end")
}

//상태값 : VMStart, VMStop, VMDeploy, VMDestroy
func insertAsyncJob(asyncJob AsyncJob) map[string]interface{} {
	asyncJob.Id = getUuid()
	db, err := sql.Open(os.Getenv("MsqlType"), os.Getenv("DbInfo"))
	returnValue := map[string]interface{}{}
	if err != nil {
		log.Error("DB connect error")
		log.Error(err)
		returnValue["message"] = "DB connect error"
	}
	defer db.Close()

	result, err := db.Exec("INSERT INTO async_job(id, name, exec_uuid, ready, parameter, create_date) VALUES (?, ?, ?, ?, ?, now())",
		asyncJob.Id, asyncJob.Name, asyncJob.ExecUuid, asyncJob.Ready, asyncJob.Parameter)
	if err != nil {
		log.Error("Async Job 등록중 에러가 발생했습니다.")
		log.Error(err)
		returnValue["message"] = "An error occurred while registering Async Job."
	}
	n, err := result.RowsAffected()
	if n == 1 {
		log.Info(result.LastInsertId())
		log.Info("async job이 정상적으로 등록되었습니다.")
		returnValue["code"] = 200
		returnValue["jobid"] = asyncJob.Id
	}

	return returnValue
}
