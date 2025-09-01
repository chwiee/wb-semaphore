package config

import (
	"github.com/spf13/viper"
)

const (
	DefaultLogLevel        = "debug"
	DefaultTimestampFormat = "2006-01-02 15:04:05"
	DefaultLogoutput       = "stdout"

	DefaultDatabasePort         = 5432
	DefaultDatabaseName         = "app_db"
	DefaultDatabaseUser         = "postgres"
	DefaultDatabaseHost         = "localhost"
	DefaultDatabaseSSLMode      = "disable"
	DefaultDatabasePassword     = "password"
	DefaultDatabaseMaxIdleConns = 10
	DefaultDatabaseMaxOpenConns = 100

	DefaultSemaphoreURL = "http://localhost:3000/api/"
	DefaultBearerToken  = "s9iu1iwfobl4k4k7ylfvcudjeln2r8-lyw3o8r8k_b4="
)

var Viper *viper.Viper

func init() {
	v := viper.New()
	v.SetEnvPrefix("wb")
	v.AutomaticEnv()
	Viper = v
}

func GetLogLevel() string {
	level := Viper.GetString("log.level")
	if len(level) > 0 {
		return level
	}
	return DefaultLogLevel
}

func GetTimestampFormat() string {
	format := Viper.GetString("log.timestamp_format")
	if len(format) > 0 {
		return format
	}
	return DefaultTimestampFormat
}

func GetLogOutput() string {
	output := Viper.GetString("log.output")
	if len(output) > 0 {
		return output
	}
	return DefaultLogoutput
}

func GetDatabasePort() int {
	port := Viper.GetInt("database.port")
	if port > 0 {
		return port
	}
	return DefaultDatabasePort
}

func GetDatabaseName() string {
	name := Viper.GetString("database.name")
	if len(name) > 0 {
		return name
	}
	return DefaultDatabaseName
}

func GetDatabaseUser() string {
	user := Viper.GetString("database.user")
	if len(user) > 0 {
		return user
	}
	return DefaultDatabaseUser
}

func GetDatabaseHost() string {
	host := Viper.GetString("database.host")
	if len(host) > 0 {
		return host
	}
	return DefaultDatabaseHost
}

func GetDatabaseSSLMode() string {
	sslMode := Viper.GetString("database.sslmode")
	if len(sslMode) > 0 {
		return sslMode
	}
	return DefaultDatabaseSSLMode
}

func GetDatabasePassword() string {
	password := Viper.GetString("database.password")
	if len(password) > 0 {
		return password
	}
	return DefaultDatabasePassword
}

func GetDatabaseMaxIdleConns() int {
	maxIdleConns := Viper.GetInt("database.max_idle_conns")
	if maxIdleConns > 0 {
		return maxIdleConns
	}
	return DefaultDatabaseMaxIdleConns
}

func GetDatabaseMaxOpenConns() int {
	maxOpenConns := Viper.GetInt("database.max_open_conns")
	if maxOpenConns > 0 {
		return maxOpenConns
	}
	return DefaultDatabaseMaxOpenConns
}

func GetSemaphoreURL() string {
	url := Viper.GetString("semaphore.url")
	if len(url) > 0 {
		return url
	}
	return DefaultSemaphoreURL
}

func GetBearerToken() string {
	token := Viper.GetString("semaphore.token")
	if len(token) > 0 {
		return token
	}
	return DefaultBearerToken
}
