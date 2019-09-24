package controllers

import (
	"fmt"

	"github.com/astaxie/beego"
)

// ValidacionController operations for Validacion
type ValidacionController struct {
	beego.Controller
}

// URLMapping ...
func (c *ValidacionController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
}

// Post ...
// @Title Create
// @Description create Validacion
// @Param	docID	query	string	true		"ID del documento"
// @Success 200 {}
// @Failure 403 body is empty
// @router / [post]
func (c *ValidacionController) Post() {
	DocumentID := c.GetString("docID")
	fmt.Println(DocumentID)
}

// GetOne ...
// @Title GetOne
// @Description get Validacion by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Validacion
// @Failure 403 :id is empty
// @router /:id [get]
func (c *ValidacionController) GetOne() {

}

// GetAll ...
// @Title GetAll
// @Description get Validacion
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.Validacion
// @Failure 403
// @router / [get]
func (c *ValidacionController) GetAll() {

}

// Put ...
// @Title Put
// @Description update the Validacion
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.Validacion	true		"body for Validacion content"
// @Success 200 {object} models.Validacion
// @Failure 403 :id is not int
// @router /:id [put]
func (c *ValidacionController) Put() {

}
