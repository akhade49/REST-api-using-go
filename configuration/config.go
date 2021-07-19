package configuration

import (
	"os"
	"mutlicontainer/constants"
	"encoding/json"
	"github.com/gin-gonic/gin/binding"
)

type Config interface {
	LoadConfig(fileName string) (ProductConfig, error)
}

type ProductConfig interface {
	GetDBUser() string
	GetDBPassword() string
	GetDBName() string
	GetDBHost() string
	GetDBPort() int
}

type configLoader struct {
}

func NewConfigLoader() Config {
	return configLoader{}
}

type productConfig struct {
	DBUser     *string 
	DBPassword *string
	DBName     *string
	DBHost     *string
	DBPort     *int
}


func (configLoader configLoader) LoadEnvironmentVariable(configuration *productConfig) {

	configuration.DBUser = configLoader.getEnvironmentValue(constants.DBUser)
	configuration.DBPassword = configLoader.getEnvironmentValue(constants.DBPassword)
	configuration.DBHost = configLoader.getEnvironmentValue(constants.DBHost)
	configuration.DBName = configLoader.getEnvironmentValue(constants.DBName)
	//configuration.DBPort = configLoader.getEnvironmentValue(constants.DBPort)

}

func (configLoader configLoader) getEnvironmentValue(key string) *string {
	val, keyPresent := os.LookupEnv(key)
	if !keyPresent {
		return nil
	}
	return &val
}


func (configLoader configLoader) LoadConfig(fileName string) (ProductConfig, error) {
	configuration := productConfig{}
	file, err := os.Open(fileName)
	configLoader.LoadEnvironmentVariable(&configuration)
	defer file.Close()
	if err != nil {
		return configuration, err
	}
	decoder := json.NewDecoder(file)

	err = decoder.Decode(&configuration)
	if err != nil {
		return configuration, err
	}

	err = binding.Validator.ValidateStruct(configuration)
	if err != nil {
		//logger := logging.GetLogger(context.TODO())
		//logger.Error("Config load failed with error", err)
		return configuration, err
	}

	return configuration, err
}


// func (configLoader configLoader) LoadConfig() ProductConfig {
// 	configuration := productConfig{}
// 	configLoader.LoadEnvironmentVariable(&configuration)

// 	return configuration
// }




func (c productConfig) GetDBUser() string {
	return *c.DBUser
}

func (c productConfig) GetDBPassword() string {
	return *c.DBPassword
}
func (c productConfig) GetDBName() string {
	return *c.DBName
}

func (c productConfig) GetDBHost() string {
	return *c.DBHost
}

func (c productConfig) GetDBPort() int {
	return 5432
}