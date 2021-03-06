/*
 * Author: zheng-ji.info
 */

package logic

import (
	"errors"
	"io/ioutil"
	"launchpad.net/goyaml"
)

type AppConfig struct {
	Redis      RedisConf `yaml:"redis"`
	Logconf    string    `yaml:"log_config"`
	Listen     string    `yaml:"listen"`
	Port       string    `yaml:"port"`
	Ttl        int       `yaml:"ttl"`
	Dnsservers []string  `yaml:"dnsservers"`
}

type RedisConf struct {
	Host string `yaml:"host"`
	Db   string `yaml:"db"`
}

func (rc *RedisConf) isValid() bool {
	return len(rc.Host) > 0 && len(rc.Db) > 0
}

var appConfig AppConfig

func (ac *AppConfig) isValid() bool {
	return ac.Redis.isValid() &&
		len(ac.Listen) > 0 &&
		len(ac.Port) > 0
}

func parseConfigFile(filepath string) error {
	if config, err := ioutil.ReadFile(filepath); err == nil {
		if err = goyaml.Unmarshal(config, &appConfig); err != nil {
			return err
		}
		if appConfig.Ttl == 0 {
			appConfig.Ttl = 10 * 60
		}
		if !appConfig.isValid() {
			return errors.New("Invalid configuration")
		}
	} else {
		return err
	}
	return nil
}
