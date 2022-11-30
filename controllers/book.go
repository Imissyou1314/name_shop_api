package controllers

import (
	"name_shop_api/models"
)

// Operations about Book
type BookController struct {
	BaseController
}

// @Title Get
// @Description get all book data
// @Success 200 {[]Book} models.Book
// @Failure 403 :name is empty
// @router / [get]
func (o *BookController) Get() {
	resultList, err := models.GetAllBooks()
	o.SetResult(resultList, err)
}
