package migrations

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/kjh123/huiGo/log"
	"github.com/kjh123/huiGo/models"
)

type MigrationStage struct {
	Model models.ModelInterface
	Function func(db *gorm.DB, m models.ModelInterface) error
}

func Migrate(db *gorm.DB, migrations []MigrationStage) error {
	for _, m := range migrations {
		if MigrationExists(db, m.Model.TableName()) {
			continue
		}

		if err := m.Function(db, m.Model); err != nil {
			return err
		}

		if err := SaveMigration(db, m.Model.TableName()); err != nil {
			return err
		}

	}
	return nil
}

func MigrationExists(db *gorm.DB, migrationName string) bool {
	migration := new(Migration)

	found := !db.Where("name = ?", migrationName).First(migration).RecordNotFound()
	if found {
		log.INFO.Printf("Skipping %s migration", migrationName)
	} else {
		log.INFO.Printf("Running %s migration", migrationName)
	}

	return found
}

func SaveMigration(db *gorm.DB, migrationName string) error {
	migration := new(Migration)
	migration.Name = migrationName
	if err := db.Create(migration).Error; err != nil {
		log.ERROR.Printf("Error saving record to migrations table: %s", err)
		return fmt.Errorf("Error saving record to migrations table: %s", err)
	}

	return nil
}

