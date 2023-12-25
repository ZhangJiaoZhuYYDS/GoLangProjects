// @Author zhangjiaozhu 2023/10/15 12:46:00
package respositories

import (
	"FlashSkill/backend/common"
	"FlashSkill/backend/models"
	"database/sql"
	"strconv"
)

type IOrderResp interface {
	Conn()error
	Insert(*models.Order)(int64,error)
	Delete(int64)bool
	Update(*models.Order)error
	SelectByKey (int64)(*models.Order,error)
	SelectAll ()([]*models.Order,error)
	SelectAllWithInfo()(map[int]map[string]string,error)
}
type OrderManagerResp struct {
	table string
	mysqlConn *sql.DB
}

func (o *OrderManagerResp) Conn() (err error) {
	if o.mysqlConn == nil {
		o.mysqlConn = common.DB
	}
	if o.table == ""{
		o.table = "order"
	}
	return
}

func (o *OrderManagerResp) Insert(order *models.Order) (int64, error) {
	if err := o.Conn();err != nil {
		return 0, err
	}
	sql :=  "insert `order` set userID = ? , productID = ? ,orderStatus = ? "
	stmt, err := o.mysqlConn.Prepare(sql)
	if err != nil {
		return 0, err
	}
	result, err := stmt.Exec(order.UserId, order.ProductId, order.OrderStatus)
	if err != nil {
		return 0, err
	}
	return result.LastInsertId()
}

func (o *OrderManagerResp) Delete(i int64) bool {
	// 1 判断连接是否存在
	if err := o.Conn();err != nil {
		return false
	}
	// 2 准备sql
	sql := "delete from `order` where ID = ?"
	stmt , errsql :=o.mysqlConn.Prepare(sql)
	if errsql != nil {
		return false
	}
	// 3 sql传参
	_, err := stmt.Exec(i)
	if err != nil {
		return false
	}
	return true
}

func (o *OrderManagerResp) Update(order *models.Order) error {
	// 1 判断连接是否存在
	if err := o.Conn();err != nil {
		return err
	}
	// 2 准备sql
	sql := "update `order` set userID = ? , productID = ? ,orderStatus = ? where ID = ?"+strconv.FormatInt(order.ID,10)

	stmt , errsql :=o.mysqlConn.Prepare(sql)
	if errsql != nil {
		return  errsql
	}
	// 3 sql传参
	_, err := stmt.Exec(order.UserId, order.ProductId, order.OrderStatus)
	if err != nil {
		return  err
	}
	return nil
}

func (o *OrderManagerResp) SelectByKey(i int64) (*models.Order, error) {
	// 1 判断连接是否存在
	if err := o.Conn();err != nil {
		return nil,err
	}

	sql := "select * from `order` where ID = " +strconv.FormatInt(i,10)
	row,err := o.mysqlConn.Query(sql)
	defer row.Close()
	if err != nil {
		return nil, err
	}
	result := common.GetResultRow(row)
	if len(result) == 0 {
		return &models.Order{},nil
	}
	order := &models.Order{}
	// 反射结构
	common.DataToStructByTagSql(result,order)
	return order,nil
}

func (o *OrderManagerResp) SelectAll() ([]*models.Order, error) {
	// 1 判断连接是否存在
	if err := o.Conn();err != nil {
		return nil,err
	}
	sql := "select * from `order` "
	row , errRow := o.mysqlConn.Query(sql)
	defer row.Close()
	if errRow != nil {
		return nil,errRow
	}
	result := common.GetAllResultRows(row)
	if len(result) == 0 {
		return nil,nil
	}
	results :=make([]*models.Order,0)
	for _ , v := range result{
		order := &models.Order{}
		common.DataToStructByTagSql(v,order)
		results = append(results,order)
	}
	return results,nil
}

func (o *OrderManagerResp) SelectAllWithInfo() (map[int]map[string]string, error) {
	if errConn := o.Conn();errConn != nil {
		return nil, errConn
	}
	sql := "select o.ID , p.productName , o.orderStatus from `order` as o left join product as p on o.productID=p.ID"
	rows, err := o.mysqlConn.Query(sql)
	if err != nil {
		return nil, err
	}
	return common.GetAllResultRows(rows),nil
}

func NewOrderManagerResp(table string) IOrderResp  {
	return &OrderManagerResp{table: table,mysqlConn:common.DB}
}