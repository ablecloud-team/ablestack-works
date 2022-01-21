package main

import (
	"database/sql"
	"github.com/sirupsen/logrus"
	"os"
)

type Configuration struct {
	Id           string  `json:"id"`
	Name         string  `json:"name"`
	Category     string  `json:"category"`
	Value        string  `json:"value"`
	DefaultValue *string `json:"default_value"`
	Description  string  `json:"description"`
	UpdateDate   *string `json:"update_date"`
	InputType    string  `json:"input_type"`
	InputValue   *string `json:"input_value"`
}

func selectConfigurationList() ([]Configuration, error) {
	log.WithFields(logrus.Fields{
		"configurationDao": "selectConfigurationList",
	}).Infof("payload none")
	db, err := sql.Open(os.Getenv("MysqlType"), os.Getenv("DbInfo"))
	if err != nil {
		log.WithFields(logrus.Fields{
			"configurationDao": "selectConfigurationList",
		}).Errorf("selectConfigurationList DB Connect Error [%v]", err)
	}
	defer db.Close()
	log.WithFields(logrus.Fields{
		"configurationDao": "selectConfigurationList",
	}).Infof("select ConfigurationList DB Connect success")
	queryString := "SELECT" +
		" id, name, category, value, IFNULL(default_value, '')," +
		" description, IFNULL(update_date,''), input_type, IFNULL(input_value,'')" +
		" FROM configuration" +
		" ORDER BY id DESC"
	log.WithFields(logrus.Fields{
		"configurationDao": "selectConfigurationList",
	}).Infof("select selectConfigurationList Query [%v]", queryString)
	rows, err := db.Query(queryString)
	if err != nil {
		log.WithFields(logrus.Fields{
			"configurationDao": "selectConfigurationList",
		}).Errorf("Configuration Select Query FAILED [%v]", err)
	}
	var configurationList []Configuration
	defer rows.Close()
	for rows.Next() {
		configuration := Configuration{}
		err = rows.Scan(
			&configuration.Id, &configuration.Name, &configuration.Category, &configuration.Value, &configuration.DefaultValue,
			&configuration.Description, &configuration.UpdateDate, &configuration.InputType, &configuration.InputValue)
		if err != nil {
			log.WithFields(logrus.Fields{
				"configurationDao": "selectConfigurationList",
			}).Errorf("configuratio Select Query 이후 Scan 중 에러가 발생했습니다. [%v]", err)
		}

		configurationList = append(configurationList, configuration)
	}
	log.WithFields(logrus.Fields{
		"workspaceImpl": "selectWorkspaceList",
	}).Infof("selectWorkspaceList Query result [%v]", configurationList)

	return configurationList, err
}

func updateConfiguration(configuration Configuration) (int64, error) {
	log.WithFields(logrus.Fields{
		"configurationDao": "updateConfiguration",
	}).Infof("payload none")
	db, err := sql.Open(os.Getenv("MysqlType"), os.Getenv("DbInfo"))
	if err != nil {
		log.WithFields(logrus.Fields{
			"configurationDao": "updateConfiguration",
		}).Errorf("updateConfiguration DB Connect Error [%v]", err)
	}
	defer db.Close()
	log.WithFields(logrus.Fields{
		"configurationDao": "updateConfiguration",
	}).Infof("select updateConfiguration DB Connect success")
	result1, err := db.Exec("UPDATE configuration set value =? where id=?", configuration.Value, configuration.Id)
	if err != nil {
		log.WithFields(logrus.Fields{
			"configurationDao": "updateConfiguration",
		}).Errorf("An error occurred while changing the configuration. [%v]", err)
	}
	n, err := result1.RowsAffected()
	if n == 1 {
		log.WithFields(logrus.Fields{
			"configurationDao": "updateConfiguration",
		}).Infof("Configuration change is complete.")
	}

	return n, err
}
