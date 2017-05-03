package engin

import (
	"fmt"
	"os"
	"sync"

	"github.com/flourish-ship/work-attendance/consts"

	"github.com/go-xorm/core"
	"github.com/go-xorm/xorm"
)

var ormEngin *xorm.Engine
var once sync.Once

type DataBaseInfo struct {
	Database string
	User     string
	Password string
	Host     string
}

var dbInfo string

func init() {
	dbInfo = databaseLink()
}

func prepare(engin *xorm.Engine) {
	engin.SetMapper(core.GonicMapper{})
	engin.ShowSQL(false)
}

func NewOrmEngin() *xorm.Engine {
	var err error
	once.Do(func() {
		ormEngin, err = xorm.NewEngine("mysql", dbInfo)
		if err != nil {
			errInfo := fmt.Sprintf("create mysql database engin failed, error info is %s", err.Error())
			fmt.Println(errInfo)
			panic(errInfo)
		}
		fmt.Printf("has connect to database, database info is %s \n", dbInfo)
		prepare(ormEngin)
	})

	return ormEngin
}

func databaseLink() string {
	user := os.Getenv("MYSQL_USER")
	if len(user) == 0 {
		user = consts.DATABASE_DEFAULT_USER
	}
	password := os.Getenv("MYSQL_ROOT_PASSWORD")
	if len(password) == 0 {
		password = consts.DATABASE_DEFAULT_PASSWORD
	}
	host := os.Getenv("MYSQL_HOST")
	if len(host) == 0 {
		host = consts.DATABASE_DEFAULT_HOST
	}
	port := os.Getenv("MYSQL_PORT")
	if len(port) == 0 {
		port = consts.DATABASE_DEFAULT_PORT
	}
	database := os.Getenv("MYSQL_DATABASE")
	if len(database) == 0 {
		database = consts.DATABASE_DEFAULT
	}

	return fmt.Sprintf(consts.DATA_SOURCE, user, password, host, port, database)
}
