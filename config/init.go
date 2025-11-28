package config

import (
	"github.com/spf13/viper"
)

func Init() error {
	// (drk изменения 2025.11) Использование абсолютного пути внутри контейнера
	// Скорее всего не нужно, но система работает, оставлю
	viper.AddConfigPath("/root/config")
	//viper.AddConfigPath("./config") //# drk изменения 2025.11
	viper.SetConfigName("config")

	return viper.ReadInConfig()
}
