package migrations

import (
	"gofr.dev/pkg/datastore"
	"gofr.dev/pkg/log"
)

const alterData = "ALTER TABLE users MODIFY COLUMN age VARCHAR(255);"

const revertData = "ALTER TABLE users MODIFY COLUMN age int;"

type K20231119175509 struct {
}

func (k K20231119175509) Up(d *datastore.DataStore, logger log.Logger) error {
	_, err := d.DB().Exec(alterData)
	if err != nil {
		return err
	}

	return nil
}

func (k K20231119175509) Down(d *datastore.DataStore, logger log.Logger) error {
	_, err := d.DB().Exec(revertData)
	if err != nil {
		return err
	}

	return nil
}
