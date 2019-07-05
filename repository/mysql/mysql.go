package mysql

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/strwrd/rest-scrapper/model"

	"github.com/go-sql-driver/mysql"
)

// Configuration loader
var (
	_MysqlUser     = "root"
	_MysqlPassword = "123sayangsemuax"
	_MysqlHost     = "scrapper.cxiwruko6bfi.ap-southeast-1.rds.amazonaws.com"
	_MysqlPort     = "3306"
	_MysqlDB       = "scrapper"
)

// Repository : repository mysql interface contract
type Repository interface {
	Close()
	GetAllArchieve(ctx context.Context) ([]*model.Archieve, error)
	GetAllJournal(ctx context.Context) ([]*model.Journal, error)
	GetArchieveByArchieveID(ctx context.Context, ID string) (*model.Archieve, error)
	GetArchieveByCode(ctx context.Context, code string) (*model.Archieve, error)
	GetJournalsByArchieveID(ctx context.Context, ID string) ([]*model.Journal, error)
	GetJournalByJournalID(ctx context.Context, ID string) (*model.Journal, error)
}

// Repository mysql object
type repository struct {
	//connection mysql server
	conn *sql.DB
}

// NewRepository : create repository mysql object
func NewRepository() (Repository, error) {

	// Mysql connection configuration
	config := &mysql.Config{
		User:                 _MysqlUser,
		Passwd:               _MysqlPassword,
		Addr:                 fmt.Sprintf("%s:%s", _MysqlHost, _MysqlPort),
		DBName:               _MysqlDB,
		Loc:                  time.UTC,
		ParseTime:            true,
		AllowNativePasswords: true,
		Net:                  "tcp",
	}

	// Do connecting mysql server
	conn, err := sql.Open("mysql", config.FormatDSN())
	if err != nil {
		return nil, err
	}

	// Check if mysql is available
	if err := conn.Ping(); err != nil {
		return nil, err
	}

	// return mysql object with connection & error value
	return &repository{conn}, nil
}

// Close : closing mysql connection
func (r *repository) Close() {
	r.conn.Close()
}
