package db

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"

	"github.com/longhaoteng/wineglass/config"
)

var (
	DB           *gorm.DB
	repositories []Repository
)

type Repository interface {
	TableName() string
	BeforeCreate(db *gorm.DB) error
	BeforeSave(db *gorm.DB) error
	AfterCreate(db *gorm.DB) error
	AfterSave(db *gorm.DB) error
}

func AddRepositories(dbRepositories ...Repository) {
	repositories = append(repositories, dbRepositories...)
}

func Init() error {
	var err error
	var dialect gorm.Dialector

	switch config.DB.Driver {
	case "mysql":
		dialect = mysql.New(
			mysql.Config{
				DSN: fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local",
					config.DB.User,
					config.DB.Password,
					config.DB.Host,
					config.DB.Port,
					config.DB.DBName,
				),
				DefaultStringSize: 255, // add default size for string fields, by default, will use db type `longtext` for fields without size, not a primary key, no index defined and don't have default values
			},
		)
	case "sqlite":
		dialect = sqlite.Open(fmt.Sprintf("%s.db", config.DB.DBName))
	}

	cfg := &gorm.Config{
		PrepareStmt: true,
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
		Logger: newLog(),
	}

	DB, err = gorm.Open(dialect, cfg)
	if err != nil {
		return err
	}

	sqldb, err := DB.DB()
	if err != nil {
		return err
	}
	sqldb.SetMaxOpenConns(config.DB.MaxOpenConns)
	sqldb.SetMaxIdleConns(config.DB.MaxIdleConns)

	for _, repository := range repositories {
		if err = DB.AutoMigrate(repository); err != nil {
			return err
		}
	}

	return nil
}

func IsRecordFound(e error) (found bool, err error) {
	switch e {
	case nil:
		found = true
	case gorm.ErrRecordNotFound:
		found = false
	default:
		err = e
	}
	return
}
