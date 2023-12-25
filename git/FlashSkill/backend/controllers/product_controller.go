// @Author zhangjiaozhu 2023/10/14 20:10:00
package controllers

import (
	"FlashSkill/backend/models"
	"FlashSkill/backend/services"
	"github.com/gin-gonic/gin"
	"strconv"
)

//type ProductController struct {
//	Ctx gin.Context
//	ProductService services.IProductService
//}

	var ProductService = services.NewProductService()


func GetAllProduct(ctx *gin.Context)   {
	product, err := ProductService.GetAllProduct()
	if err != nil {
		ctx.JSON(200,"服务器繁忙")
		return
	}
	ctx.JSON(200,product)
}

func UpdateProduct(ctx *gin.Context)  {
	productDto := &models.ProductUpdateDto{}
	err := ctx.ShouldBind(productDto)
	if err != nil {
		ctx.JSON(200,"格式不匹配")
		return
	}
	if productDto == nil {
		ctx.JSON(200,"数据为空")
		return
	}

	err = ProductService.UpdateProduct(&models.Product{ProductImage: productDto.ProductImage, ProductName: productDto.ProductName, ProductUrl: productDto.ProductName, ProductNum: productDto.ProductNum})
	if err != nil {
		ctx.JSON(200,"服务器繁忙")
		return
	}
	ctx.JSON(200,"success")
}
func AddProduct(ctx *gin.Context)  {
	productDto := &models.ProductUpdateDto{}
	err := ctx.ShouldBind(productDto)
	if err != nil {
		ctx.JSON(200,"格式不匹配")
		return
	}
	//TODO 非空校验
	if productDto == nil {
		ctx.JSON(200,"数据为空")
		return
	}
	ProductService.InsertProduct(&models.Product{ProductImage: productDto.ProductImage, ProductName: productDto.ProductName, ProductUrl: productDto.ProductName, ProductNum: productDto.ProductNum})
	if err != nil {
		ctx.JSON(200,"服务器繁忙")
		return
	}
	ctx.JSON(200,"success")
}

func DelProduct(ctx *gin.Context)  {
	// TODO 非空校验    判断数据是否存在
	param := ctx.Param("id")
	id, _ := strconv.ParseInt(param,10,64)
	ok := ProductService.DeleteProductByID(id)
	if !ok {
		ctx.JSON(200,"服务器繁忙")
		return
	}
	ctx.JSON(200,"success")
}