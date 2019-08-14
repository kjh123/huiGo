package cmd

import (
    "fmt"
    "github.com/kjh123/huiGo/util/migrations"
)

func Migrate(configFile string) error {
    _, db, err := initConfigDB(configFile)
    if err != nil {
        return fmt.Errorf("migrate error: %s", err)
    }
    defer db.Close()

    if err := migrations.BootStrap(db); err != nil {
        return err
    }

    if err := migrations.MigrateAll(db); err != nil {
        return err
    }

    return nil
}
