package migrations

import (
	"gofr.dev/pkg/datastore"
	"gofr.dev/pkg/log"
)

type K20231119174053 struct {
}

const createTable = "create table users ( name varchar(50), age int, email varchar(50));"
const dropTable = "drop table users"

func (k K20231119174053) Up(d *datastore.DataStore, logger log.Logger) error {
	_, err := d.DB().Exec(createTable)
	if err != nil {
		return err
	}

	return nil
}

func (k K20231119174053) Down(d *datastore.DataStore, logger log.Logger) error {
	_, err := d.DB().Exec(dropTable)
	if err != nil {
		return err
	}

	return nil
}
