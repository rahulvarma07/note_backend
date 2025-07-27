package config

import (
	"log"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
)

type Storage struct {
	DataBaseAName  string `yaml:"dbName" env_required:"true"`
	DataBaseString string `yaml:"dbString" env_required:"true"`
	DataBaseSecret string `yaml:"dbSecret" env_required:"true"`
}

type HttpServer struct {
	Host    string `yaml:"host" env_required:"true"`
	Port    int    `yaml:"port" env_required:"true"`
	BaseUrl string `yaml:"baseUrl" env_required:"true"`
}

type Mail struct {
	MailPort     int    `yaml:"mailPort" env_required:"true"`
	SenderMailID string `yaml:"mailID" env_required:"true"`
	MailHost     string `yaml:"mailHost" env_required:"true"`
	MailPassword string `yaml:"mailPassword" env_required:"true"`
}

type Config struct {
	Dev        string `yaml:"mode" env_required:"true"`
	Storage    `yaml:"storage" env_required:"true"`
	HttpServer `yaml:"http_server" env_required:"true"`
	Mail       `yaml:"mail_attributes" env_required:"true"`
}

func MustLoad() *Config {

	// Steps
	// load the env file
	// check whether t
	// the file exsists or not..
	// create a config #(as per as the config.yaml model)
	// load or match the data from env to Config model
	// use cleanenv package to do the above step

	if err := godotenv.Load(); err != nil {
		log.Fatal("there is an error in loading env File", err)
	}

	yaml_path := os.Getenv("YAML_PATH")
	_, err := os.Stat(yaml_path)

	if os.IsNotExist(err) {
		log.Fatalf("the file path does not exsist %s", yaml_path)
	}

	var cnf Config
	err = cleanenv.ReadConfig(yaml_path, &cnf)

	if err != nil {
		log.Fatalf("failed to load yaml path %s", err)
	}

	return &cnf
}
