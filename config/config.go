package config

import (
	"log"
	"os"
	"strconv"

	"github.com/spf13/viper"
)

var (
	JWT_KEY    string = ""
	AWS_REGION string = ""
	S3_KEY     string = ""
	S3_SECRET  string = ""
	AWS_BUCKET string = ""
)

type AppConfig struct {
	DBUser    string
	DBPass    string
	DBHost    string
	DBPort    int
	DBName    string
	JWTKEY    string
	AWSREGION string
	S3KEY     string
	S3SECRET  string
	AWSBUCKET string
}

func InitConfig() *AppConfig {
	return ReadEnv()
}

func ReadEnv() *AppConfig {
	app := AppConfig{}
	isRead := true

	// JWT
	if val, found := os.LookupEnv("JWT_KEY"); found {
		app.JWTKEY = val
		isRead = false
	}

	// AWS S3 Bucket
	if val, found := os.LookupEnv("AWS_REGION"); found {
		app.AWSREGION = val
		isRead = false
	}
	if val, found := os.LookupEnv("S3_KEY"); found {
		app.S3KEY = val
		isRead = false
	}
	if val, found := os.LookupEnv("S3_SECRET"); found {
		app.S3SECRET = val
		isRead = false
	}
	if val, found := os.LookupEnv("AWS_BUCKET"); found {
		app.AWSBUCKET = val
		isRead = false
	}

	// DATABASE
	if val, found := os.LookupEnv("DBUSER"); found {
		app.DBUser = val
		isRead = false
	}

	if val, found := os.LookupEnv("DBPASS"); found {
		app.DBPass = val
		isRead = false
	}
	if val, found := os.LookupEnv("DBHOST"); found {
		app.DBHost = val
		isRead = false
	}
	if val, found := os.LookupEnv("DBPORT"); found {
		cnv, _ := strconv.Atoi(val)
		app.DBPort = cnv
		isRead = false
	}
	if val, found := os.LookupEnv("DBNAME"); found {
		app.DBName = val
		isRead = false
	}
	if isRead {
		viper.AddConfigPath(".")
		viper.SetConfigName("local")
		viper.SetConfigType("env")

		err := viper.ReadInConfig()
		if err != nil {
			log.Println("error read config : ", err.Error())
			return nil
		}
		err = viper.Unmarshal(&app)
		if err != nil {
			log.Println("error parse config : ", err.Error())
			return nil
		}
	}

	JWT_KEY = app.JWTKEY
	AWS_REGION = app.AWSREGION
	S3_KEY = app.S3KEY
	S3_SECRET = app.S3SECRET
	AWS_BUCKET = app.AWSBUCKET
	return &app
}
