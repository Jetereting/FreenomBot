package config

import (
	"log"
	"os"
	"path/filepath"

	"github.com/BurntSushi/toml"
)

/** Tom file **/

// Config object at tom file
type Config struct {
	Accounts     []Account
	System       System
	WeChatNotify WeChatNotify
	LineNotify   LineNotify
}

// System for basic auth at tom file
type System struct {
	Account    string
	Password   string
	CronTiming string
	Lang       string
}

// LineNotify for send message at tom file
type LineNotify struct {
	Enable bool
	Daily  bool
	Token  string
}

// WeChatNotify for send message at tom file
type WeChatNotify struct {
	Enable     bool
	Daily      bool
	CorpID     string
	CorpSecret string
	AgentID    string
}

// Account struct at tom file
type Account struct {
	Username string
	Password string
	Domains  []Domain
}

// Domain data
type Domain struct {
	DomainName string
	Days       int
	ID         string
	RenewState int
	CheckTimes int
}

var configData *Config

func readConf(filename string) error {
	var (
		err error
	)
	// 先读取环境变量的配置文件
	envConfig := os.Getenv(filename)
	if len(envConfig) > 0 {
		if _, err = toml.Decode("FreenomBot.toml", &configData); err != nil {
			log.Fatal(err)
		}
		return err
	}
	// 没有则读文件
	filename, err = filepath.Abs(filename)
	if err != nil {
		log.Fatal(err)
		return err
	}
	if _, err = toml.DecodeFile(filename, &configData); err != nil {
		log.Fatal(err)
	}
	return err
}

// GetData get Config data
func GetData() *Config {
	return configData
}

func init() {
	if err := readConf("config.toml"); err != nil {
		log.Fatalln(err)
	}
}
