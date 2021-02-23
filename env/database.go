package env

import (
	"entgo.io/ent/examples/start/ent"
	"github.com/spf13/viper"
	"log"
	"path/filepath"
	"fmt"
)

type Database struct {
	Type     string
	User     string
	Host     string
	Name     string
	Password string
}

var DBEnv Database
var DB *ent.Client

func loadEnv() error {
	rootDir, _ := filepath.Abs("./")
	viper.SetConfigFile(rootDir+"/config/config.json")
	var err error
	if err = viper.ReadInConfig(); err != nil {
		log.Fatal("Error read config file")
	}
	database := viper.Sub("database")

	err = database.Unmarshal(&DBEnv)

	if err != nil {
		log.Fatalf("unable to decode , %v", err)
		return err
	}
	return nil
}

func SetupEnv() error {
	if err := loadEnv(); err != nil {
		return err
	}

	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s)/%s",
		DBEnv.User,
		DBEnv.Password,
		DBEnv.Host,
		DBEnv.Name,
	)
	var err error
	DB, err = ent.Open("mysql", dataSourceName)

	if err != nil {
		return err
	}

	return nil
}