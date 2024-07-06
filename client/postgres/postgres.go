package postgres

import (
	"database/sql"
	"fmt"
	"go-boilerplate/config"

	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *PostgreSQLClient

type PostgreSQLClient struct {
	Master *gorm.DB
	Slave  *gorm.DB
}

func InitPostgres() (*PostgreSQLClient, error) {
	dbConfig := config.Config.Database

	masterDsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		dbConfig.Master.Host,
		dbConfig.Master.Port,
		dbConfig.Master.Username,
		dbConfig.Master.Password,
		dbConfig.Master.DBName,
	)

	slaveDsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		dbConfig.Slave.Host,
		dbConfig.Slave.Port,
		dbConfig.Slave.Username,
		dbConfig.Slave.Password,
		dbConfig.Slave.DBName,
	)

	// Open connection
	masterConn, err := sql.Open("postgres", masterDsn)
	if err != nil {
		return nil, err
	}
	slaveConn, err := sql.Open("postgres", slaveDsn)
	if err != nil {
		return nil, err
	}

	// Test ping connection
	err = masterConn.Ping()
	if err != nil {
		return nil, err
	}
	err = slaveConn.Ping()
	if err != nil {
		return nil, err
	}

	// Connect
	dbMaster, err := gorm.Open(postgres.New(postgres.Config{
		Conn: masterConn,
	}))
	if err != nil {
		return nil, err
	}

	dbSlave, err := gorm.Open(postgres.New(postgres.Config{
		Conn: slaveConn,
	}))
	if err != nil {
		return nil, err
	}

	DB = &PostgreSQLClient{
		Master: dbMaster,
		Slave:  dbSlave,
	}

	return DB, nil
}
