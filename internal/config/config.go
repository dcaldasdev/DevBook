package config

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

var (
	ConnectionString string
	ApiPort          uint16
)

// Load initialize environment variables
func Load() {

	switch {
	case os.Getenv("PROFILE") == "dev":
		viper.SetConfigName("application")
	case os.Getenv("PROFILE") == "stg":
		viper.SetConfigName("application-staging")
	case os.Getenv("PROFILE") == "prd":
		viper.SetConfigName("application-production")
	default:
		panic("Could not find environment to load")

	}

	// Set the path to look for the configurations file
	viper.AddConfigPath("../internal/config/")

	// Enable VIPER to read Environment Variables
	viper.AutomaticEnv()

	viper.SetConfigType("yml")

	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("Error reading config file, %s", err)
	}

	dbUser, _ := viper.Get("db.mysql.user").(string)

	dbPass, _ := viper.Get("db.mysql.pass").(string)

	dbName, _ := viper.Get("db.mysql.name").(string)

	ConnectionString = fmt.Sprintf("%s:%s@/%s?charset=utf8&parseTime=True&loc=Local",
		dbUser,
		dbPass,
		dbName,
	)

	ApiPort, _ = viper.Get("api-port").(uint16)

}
