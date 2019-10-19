package mysql
import (
	"database/sql"
	"errors"
	"fmt"
	"time"
	_ "github.com/lib/pq"
	"github.com/astaxie/beego/utils"
	"github.com/go-ini/ini"
	"strconv"
)

var (
	host     string = ""
	port     int
	user     string = ""
	password string = ""
	dbname   string = ""
	max_conn  int = 40
	idle_conn int = 10
	postgreConn *sql.DB //全局sql连接,已实现连接池,所以可以只创建一个实例
	DB_CONN_ERROR error
)

func init() {
	//使用beego方法读取配置
	//host = beego.AppConfig.String("postgres_host")
	//port, _ = beego.AppConfig.Int("postgres_port")
	//user = beego.AppConfig.String("postgres_user")
	//password = beego.AppConfig.String("postgres_password")
	//dbname = beego.AppConfig.String("postgres_dbname")
	//max_conn = beego.AppConfig.DefaultInt("postgres_max_conn", 50)
	//idle_conn = beego.AppConfig.DefaultInt("postgres_idle_conn", 10)

	//使用ini包读取配置
	dbIniPath:=getPgsqlConf()
	b := utils.FileExists(dbIniPath)
	if !b {
		panic("PgsqlDatabase config not exists")
	}
	cfg, err := ini.Load(dbIniPath)
	if err != nil {
		panic(fmt.Sprintf("Config File Read Error!, the err is %v", err))
	}
	sections := cfg.Sections()
	for _, section := range sections {
		name := section.Name()
		if name == "DEFAULT" {
			continue
		}
		sec := cfg.Section(name)
		host, _ := sec.GetKey("HOST")
		port, _ := sec.GetKey("PORT")
		username, _ := sec.GetKey("USERNAME")
		password, _ := sec.GetKey("PASSWORD")
		dbname, _ := sec.GetKey("DBNAME")
		configs[section.Name()] = ConnConfig{
			Host:     host.String(),
			Port:     port.String(),
			Username: username.String(),
			Password: password.String(),
			DbName:   dbname.String(),
		}
	}

	for key,val :=range configs{
		switch key {
		case "master":
			host = val.Host
			port,_  =strconv.Atoi(val.Port)
			user = val.Username
			password = val.Password
			dbname = val.DbName
			break
		default:
			break
		}
	}
	DB_CONN_ERROR = errors.New("数据库连接失败")
}

func getPgsqlConf()  string{
	return "conf/databases/pgsql.ini"
}

func GetPostgresSession() *sql.DB {
	if postgreConn == nil {
		psqlInfo := fmt.Sprintf(`host=%s port=%d user=%s password=%s dbname=%s sslmode=disable`,
			host, port, user, password, dbname)
		db, err := sql.Open("postgres", psqlInfo)
		if err != nil {
			return nil
		}

		db.SetConnMaxLifetime(30 * time.Minute)
		db.SetMaxOpenConns(max_conn)
		db.SetMaxIdleConns(idle_conn)

		err = db.Ping()
		if err != nil {
			fmt.Println("pgsql connect is file!",DB_CONN_ERROR)
			return nil
		}
		postgreConn = db
	}
	return postgreConn
}
