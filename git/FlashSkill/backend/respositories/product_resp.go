// @Author zhangjiaozhu 2023/10/14 17:25:00
package respositories

import (
	"FlashSkill/backend/common"
	"FlashSkill/backend/models"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"strconv"
)

type IProduct interface {
	Conn() error
	Insert(product *models.Product)(int64,error)
	Delete(int642 int64)bool
	Update(product *models.Product) error
	SelectByKey(int64)(*models.Product,error)
	SelectAll()([]*models.Product,error)
}

type ProductManager struct {
	table string
	mysqlConn *sql.DB
}



func NewProductManager(table string)  IProduct{
	return &ProductManager{table: table,mysqlConn: common.DB}
}

func (m *ProductManager) Conn() (err error)  {
	if m.mysqlConn == nil {
		m.mysqlConn = common.DB
	}
	if m.table == ""{
		m.table = "product"
	}
	return
}

func (m *ProductManager) Insert(product *models.Product) (productId int64,err error)  {
	// 1 判断连接是否存在
	if err = m.Conn();err != nil {
		return
	}
	// 2 准备sql
	sql := "insert product set productName = ? , productNum = ? ,productImage = ? , productUrl = ?"
	stmt , errsql :=m.mysqlConn.Prepare(sql)
	if errsql != nil {
		return 0, errsql
	}
	// 3 sql传参
	result, err := stmt.Exec(product.ProductName, product.ProductNum, product.ProductImage, product.ProductUrl)
	if err != nil {
		return 0, err
	}
	return result.LastInsertId()
}
func (m *ProductManager) Update(product *models.Product) (err error)  {
	// 1 判断连接是否存在
	if err = m.Conn();err != nil {
		return
	}
	// 2 准备sql
	sql := "update product set productName = ? , productNum = ? ,productImage = ? , productUrl = ? where ID = ?"+strconv.FormatInt(product.ID,10)

	stmt , errsql :=m.mysqlConn.Prepare(sql)
	if errsql != nil {
		return  errsql
	}
	// 3 sql传参
	_, err = stmt.Exec(product.ProductName, product.ProductNum, product.ProductImage, product.ProductUrl)
	if err != nil {
		return  err
	}
	return nil
}
func (m *ProductManager) Delete(productID int64) bool  {
	// 1 判断连接是否存在
	if err := m.Conn();err != nil {
		return false
	}
	// 2 准备sql
	sql := "delete from product where ID = ?"
	stmt , errsql :=m.mysqlConn.Prepare(sql)
	if errsql != nil {
		return false
	}
	// 3 sql传参
	_, err := stmt.Exec(productID)
	if err != nil {
		return false
	}
	return true
}

func (m *ProductManager) SelectByKey(productID int64)(product *models.Product,err error)  {
	// 1 判断连接是否存在
	if err = m.Conn();err != nil {
		return nil,err
	}

	sql := "select * from " + m.table + "where ID = " +strconv.FormatInt(productID,10)
	row, err := m.mysqlConn.Query(sql)
	defer row.Close()
	if err != nil {
		return &models.Product{},err
	}
	result := common.GetResultRow(row)
	if len(result) == 0 {
		return &models.Product{},nil
	}
	common.DataToStructByTagSql(result,product)
	return
}

func (m *ProductManager) SelectAll() ([]*models.Product, error) {
	// 1 判断连接是否存在
	if err := m.Conn();err != nil {
		return nil,err
	}
	sql := "select * from "+m.table
	row , errRow := m.mysqlConn.Query(sql)
	defer row.Close()
	if errRow != nil {
		return nil,errRow
	}
	result := common.GetAllResultRows(row)
	if len(result) == 0 {
		return nil,nil
	}
	results :=make([]*models.Product,0)
	for _ , v := range result{
		product := &models.Product{}
		common.DataToStructByTagSql(v,product)
		log.Println(v)
		results = append(results,product)
	}
	return results,nil
}