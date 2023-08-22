package testconnection

import (
	"testing"

	"belajar-restapi/config"
)

func TestConnectionDB(t *testing.T) {
	config.Db_Mysql()
}
