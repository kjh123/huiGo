package cmd

import (
    "github.com/jinzhu/configor"
    "github.com/jinzhu/gorm"
    "github.com/kjh123/huiGo/config"
    "github.com/kjh123/huiGo/database"
)

func initConfigDB(configFile string) (*config.Config, *gorm.DB, error) {
    conf := &config.Config{}
    if err := configor.Load(conf, configFile); err != nil {
        return conf, nil, err
    }
    db, err := database.NewDatabase(conf)
    if err != nil {
        return conf, nil, err
    }
    return conf, db, nil
}
