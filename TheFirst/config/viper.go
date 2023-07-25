package config

import (
	"encoding/json"
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"github.com/spf13/viper"
	"golang.org/x/text/language"
)

var Viperconfig Configuration

func init() {
	runtimeViper := viper.New()
	runtimeViper.AddConfigPath(".")
	runtimeViper.SetConfigName("config")
	runtimeViper.SetConfigType("json")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal err config file : %s\n", err))
	}
	runtimeViper.Unmarshal(&Viperconfig)
	//初始化
	bundle := i18n.NewBundle(language.English)
	bundle.RegisterUnmarshalFunc("json", json.Unmarshal)
	bundle.MustLoadMessageFile(Viperconfig.App.Locale + "/active.en.json")
	bundle.MustLoadMessageFile(Viperconfig.App.Locale + "/active." + Viperconfig.App.Language + ".json")
	Viperconfig.LocaleBundle = bundle

	runtimeViper.WatchConfig()
	runtimeViper.OnConfigChange(func(e fsnotify.Event) {
		runtimeViper.Unmarshal(&Viperconfig)
		Viperconfig.LocaleBundle.MustLoadMessageFile(Viperconfig.App.Locale + "/active." + Viperconfig.App.Language + ".json")
	})
}
