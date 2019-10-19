package mysql

import (
	"database/sql"
	"fmt"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

type ConnConfig struct {
	Host     string
	Port     string
	Username string
	Password string
	DbName   string
}

var (
	configs = make(map[string]ConnConfig)
	ormers  = make(map[string]orm.Ormer)
	PgOrm   *sql.DB
)

func InitMysql() {
	//pgsql注册自定义配置
	PgOrm = GetPostgresSession()
}


func getDsn(config ConnConfig) string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8",
		config.Username, config.Password, config.Host, config.Port, config.DbName,
	)
}