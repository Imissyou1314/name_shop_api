package controllers

import (
	"name_shop_api/models"
	"strconv"
)

// Operations about Entity
type EntityController struct {
	BaseController
}

// @Title Get
// @Description find Entity by name
// @Param	name		path 	string	true		"the file name you want to get"
// @Success 200 {[]Entity} models.Entity
// @Failure 403 :name is empty
// @router /:name [get]
func (o *EntityController) GetAll() {
	name := o.Ctx.Input.Param(":name")
	if name != "" {
		resultList, err := models.GetEntityByName(name)
		o.SetResult(resultList, err)
	} else {
		o.SetResult("名字不能为空", nil)
	}
}

// @Title Get
// @Description find Entity by name
// @Param	name		path 	string	true		"the file name you want to get"
// @Success 200 {[]Entity} models.Entity
// @Failure 403 :name is empty
// @router /:name/:index [get]
func (o *EntityController) GetOne() {
	name := o.Ctx.Input.Param(":name")
	indexStr := o.Ctx.Input.Param(":index")
	indexValue, _ := strconv.Atoi(indexStr)
	if name != "" {
		resultEntity, err := models.GetEntityByNameAndIndex(name, indexValue)
		o.SetResult(resultEntity, err)
	} else {
		o.SetResult("名字不能为空", nil)
	}
}
