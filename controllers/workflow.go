package controllers

import (
	"github.com/astaxie/beego"
	"github.com/udistrital/nuxeo_mid/models"
)

// WorkflowController operations for Workflow
type WorkflowController struct {
	beego.Controller
}

// URLMapping ...
func (c *WorkflowController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
}

// Post ...
// @Title Create
// @Description create Workflow
// @Param	docID	query	string	true		"ID del documento"
// @Success 200 {}
// @Failure 403 body is empty
// @router / [post]
func (c *WorkflowController) Post() {
	var alertErr models.Alert
	// alertas := append([]interface{}{"Response:"})

	DocumentID := c.GetString("docID")
	alertErr.Type = "OK"
	alertErr.Code = "200"
	alertErr.Body = DocumentID
	c.Data["json"] = alertErr
	c.ServeJSON()
}

// GetAll ...
// @Title GetAll
// @Description get Workflow
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.Workflow
// @Failure 403
// @router / [get]
func (c *WorkflowController) GetAll() {

}

// Put ...
// @Title Put
// @Description update the Workflow
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.Workflow	true		"body for Workflow content"
// @Success 200 {object} models.Workflow
// @Failure 403 :id is not int
// @router /:id [put]
func (c *WorkflowController) Put() {

}