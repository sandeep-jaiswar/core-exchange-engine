package config

import (
    "github.com/spf13/viper"
    "log"
)

func InitConfig() {
    viper.SetConfigName("config") // Name of config file (without extension)
    viper.SetConfigType("yaml")   // REQUIRED if the config file does not have the extension in the name
    viper.AddConfigPath("./config") // Path to look for the config file in

    if err := viper.ReadInConfig(); err != nil {
        log.Fatalf("Error reading config file: %v", err)
    }
}

func GetMongoURI() string {
    return viper.GetString("mongo.uri")
}

func GetMongoDatabase() string {
    return viper.GetString("mongo.database")
}