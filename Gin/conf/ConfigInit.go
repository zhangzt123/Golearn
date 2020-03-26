package conf

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"log"
)

var Conf *viper.Viper

func init() {
	viper.GetViper()
	viper.Debug()
	viper.SetConfigName("config")                                                // name of config file (without extension)
	viper.SetConfigType("yaml")                                                  // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath("/home/zhangzt/github.com/zhangzt123/Golearn/Gin/conf/") // optionally look for config in the working directory
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// Config file not found; ignore error if desired
			log.Println("no such config file")
		} else {
			// Config file was found but another error was produced
			log.Println("read config error")
		}
		log.Fatal(err)
	}

	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("Config file changed:", e.Name)
	})
	for _, str := range viper.AllKeys() {
		log.Println(str)
	}
	log.Println("all config load success!")
	Conf = viper.GetViper()

}
