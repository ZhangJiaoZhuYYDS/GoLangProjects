// @Author zhangjiaozhu 2023/10/15 14:21:00
package respositories

import (
	"FlashSkill/backend/common"
	"FlashSkill/backend/models"
	"database/sql"
	"errors"
	"strconv"
)

type IUserResp interface {
	Conn() error
	Select(username string)(user *models.User,err error)
	Insert(user *models.User)(userID int64,err error)
}

type UserManagerResp struct {
	table string
	mysqlConn *sql.DB
}

func (u *UserManagerResp) Conn() error {
	if u.mysqlConn == nil {
		u.mysqlConn = common.DB
	}
	if u.table == ""{
		u.table = "user"
	}
	return nil
}

func (u *UserManagerResp) Select(username string) (user *models.User, err error) {
	if err = u.Conn();err != nil {
		return nil,err
	}
	sql := "select * from `user` where userName = ?"
	row , err := u.mysqlConn.Query(sql,username)
	defer row.Close()
	if err != nil {
		return nil, err
	}
	resultRow := common.GetResultRow(row)
	if len(resultRow) == 0 {
		return nil,errors.New("用户不存在")
	}
	common.DataToStructByTagSql(resultRow,user)
	return user,err
}

func (u *UserManagerResp) Insert(user *models.User) (userID int64, err error) {
	if err = u.Conn();err != nil {
		return
	}
	sql := "insert  `user` " + "set nickName = ? , userName = ? passWord = ?"
	stmt,err := u.mysqlConn.Prepare(sql)
	if err != nil {
		return
	}
	result, err := stmt.Exec(user.NickName, user.UserName, user.HashPassword)
	if err != nil {
		return
	}
	return result.LastInsertId()
}
func (u *UserManagerResp) SelectByID(userid int64)(*models.User,error)  {
	if err := u.Conn();err != nil {
		return nil, err
	}
	sql := "select * from `user` where ID = " + strconv.FormatInt(userid,10)
	row, err := u.mysqlConn.Query(sql)
	if err != nil {
		return nil, err
	}
	result := common.GetResultRow(row)
	if len(result) == 0 {
		return nil, errors.New("用户不存在")
	}
	user := &models.User{}
	common.DataToStructByTagSql(result,user)
	return user,nil
}

func NewUserResp(table string) IUserResp {
	return &UserManagerResp{table: table,mysqlConn: common.DB}
}
