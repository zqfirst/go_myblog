package mg_user

import (
	"database/sql"
	"fmt"
	"myblog_go/models"
	"time"
)

type MgUser struct {
	UserId     int       `xorm:"not null pk autoincr INT(11)"`
	Username   string    `xorm:"not null comment('用户名') VARCHAR(20)"`
	Password   string    `xorm:"not null default '' VARCHAR(60)"`
	CreateTime time.Time `xorm:"not null default 'CURRENT_TIMESTAMP' TIMESTAMP"`
	UpdateTime time.Time `xorm:"not null default 'CURRENT_TIMESTAMP' TIMESTAMP"`
}

func GetByUsername(username string) (m MgUser) {
	sql := "select * from mg_user where username=?"
	Scan(&m, models.Db.QueryRow(sql, username))
	return m;
}

func CheckPwd(pwd string, m MgUser) (b bool) {
	return m.Password == pwd
}

//将获取结果传入
func Scan(m *MgUser, row *sql.Row) {
	err := row.Scan(&m.UserId, &m.Username, &m.Password, &m.CreateTime, &m.UpdateTime)
	if err != nil {
		fmt.Println(err)
	}
}
