package controller

import (
	"assigment-2/database"
	helpers "assigment-2/helper"
	"assigment-2/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

var (
	appJSON = "application/json"
)

func CreateOrders(c *gin.Context) {
	db := database.GetDB()
	contentType := helpers.GetContentType(c)
	Orders := models.Orders{}

	if contentType == appJSON {
		err := c.ShouldBindJSON(&Orders)
		if err != nil {
			panic(err.Error())
			return
		}
	} else {
		err := c.ShouldBind(&Orders)
		if err != nil {
			panic(err.Error())
			return
		}
	}

	err := db.Debug().Create(&Orders).Error

	if err != nil {
		fmt.Println("error", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "Bad Request",
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"status":  true,
		"message": "Success get data",
		"data":    Orders,
	})

}
func GetOrders(c *gin.Context) {
	db := database.GetDB()
	contentType := helpers.GetContentType(c)
	var Orders []models.Orders
	//var Item models.Item
	//var Order models.Item

	if contentType == appJSON {
		err := c.ShouldBindJSON(&Orders)
		if err != nil {
			panic(err.Error())
			return
		}
	} else {
		err := c.ShouldBind(&Orders)
		if err != nil {
			panic(err.Error())
			return
		}
	}

	//err := db.Debug().Model(models.Orders{}).Select("orders.customer_name, items.item_code") .Joins("left join items on items.order_id = orders.order.id") .Error
	err := db.Debug().Find(&Orders).Error
	//err := db.Joins("orders").Joins("items").Find(&Orders) .Error
	//err:=db.Joins("orders", db.Where(&Item.ItemId==&Order.OrderId)).Find(&Orders).Error

	if err != nil {
		fmt.Println("error", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "Bad Request",
			"message": err.Error(),
		})
		return
	}
	if len(Orders) == 0 {
		c.JSON(http.StatusCreated, gin.H{
			"status":  true,
			"message": "Data kosong",
			"data":    Orders,
		})
	} else {
		c.JSON(http.StatusCreated, gin.H{
			"status":  true,
			"message": "Success get data",
			"data":    Orders,
		})
	}

}

func UpdateOrders(c *gin.Context) {
	db := database.GetDB()
	contentType := helpers.GetContentType(c)
	Orders := models.Orders{}
	orderId, _ := strconv.Atoi(c.Param("orderId"))

	if contentType == appJSON {
		err := c.ShouldBindJSON(&Orders)
		if err != nil {
			panic(err.Error())
			return
		}
	} else {
		err := c.ShouldBind(&Orders)
		if err != nil {
			panic(err.Error())
			return
		}
	}
	//var dataItem models.Item
	err := db.Debug().Model(&Orders).Where("order_id = ?", orderId).Updates(models.Orders{
		CustomerName: Orders.CustomerName,
		OrderedAt:    Orders.OrderedAt,
		//Items: []models.Item{{
		//	ItemCode:    dataItem.ItemCode,
		//	Description: dataItem.Description,
		//	Quantity:    dataItem.Quantity,
		//}},
	}).Error

	if err != nil {
		fmt.Println("error", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "Bad Request",
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"status":  true,
		"message": "Success get data",
		"data":    Orders,
	})

}

func DeleteOrder(c *gin.Context) {
	db := database.GetDB()
	contentType := helpers.GetContentType(c)
	Order := models.Orders{}
	OrderId, _ := strconv.Atoi(c.Param("orderId"))

	if contentType == appJSON {
		err := c.ShouldBindJSON(&Order)
		if err != nil {
			panic(err.Error())
			return
		}
	} else {
		err := c.ShouldBind(&Order)
		if err != nil {
			panic(err.Error())
			return
		}
	}
	Order.OrderId = uint(OrderId)

	err := db.Debug().Model(&Order).Where("order_id = ?", OrderId).Delete(&Order, OrderId).Error

	if err != nil {
		fmt.Println("error", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "Bad Request",
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"status":  true,
		"message": "Success delete data",
	})

}
