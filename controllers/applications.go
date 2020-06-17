package controllers

import (
	"beego_project_test/models"
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"
	"strconv"
	"strings"

	"github.com/google/uuid"
	"github.com/labstack/gommon/log"
)

const applicationsPath = "static/applications/"

// ApplicationsController operations for Applications
type ApplicationsController struct {
	BaseController
}

type ApplicationData struct {
	AimID       int    `json:"aim_id"`
	FileName    string `json:"file_name"`
	File        []byte `json:"file"`
	Description string `json:"description"`
}

// URLMapping ...
func (c *ApplicationsController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

// Post ...
// @Title Post
// @Description create Applications
// @Param	body		body 	models.Applications	true		"body for Applications content"
// @Success 201 {int} models.Applications
// @Failure 403 body is empty
// @router / [post]
func (c *ApplicationsController) Post() {
	var ApplicationData ApplicationData
	var err error
	var aim *models.Aims
	if err = json.Unmarshal(c.Ctx.Input.RequestBody, &ApplicationData); err == nil {

		if aim, err = models.GetAimsById(ApplicationData.AimID); err != nil {
			c.Response(400, nil, errors.New("No such Aim"))
		}

		var uniqFilename uuid.UUID

		if uniqFilename, err = uuid.NewUUID(); err != nil {
			c.Response(500, nil, err)
		}
		filenameSlice := strings.Split(ApplicationData.FileName, ".")
		imgFileName := uniqFilename.String() + "." + filenameSlice[len(filenameSlice)-1]

		url := applicationsPath + imgFileName

		if err = ioutil.WriteFile(url, ApplicationData.File, 0644); err != nil {
			c.Response(400, nil, err)
		}

		applic := &models.Applications{

			Aim:         aim,
			UniqId:      uniqFilename.String(),
			Filename:    ApplicationData.FileName,
			Url:         url,
			Description: ApplicationData.Description,
			Active:      true,
		}

		if _, err := models.AddApplications(applic); err == nil {
			c.Response(201, ApplicationData, nil)
		} else {
			c.Response(400, nil, err)
		}
	} else {
		c.Response(400, nil, err)
	}

}

// GetOne ...
// @Title Get One
// @Description get Applications by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Applications
// @Failure 403 :id is empty
// @router /:id [get]
func (c *ApplicationsController) GetOne() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	application, err := models.GetApplicationsById(id)
	if err != nil {
		c.Response(400, nil, err)
	} else {
		c.Response(200, application, nil)
	}

}

// GetAll ...
// @Title Get All
// @Description get Applications
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.Applications
// @Failure 403
// @router / [get]
func (c *ApplicationsController) GetAll() {
	var fields []string
	var sortby []string
	var order []string
	var query = make(map[string]string)
	var limit int64 = 10
	var offset int64

	// fields: col1,col2,entity.col3
	if v := c.GetString("fields"); v != "" {
		fields = strings.Split(v, ",")
	}
	// limit: 10 (default is 10)
	if v, err := c.GetInt64("limit"); err == nil {
		limit = v
	}
	// offset: 0 (default is 0)
	if v, err := c.GetInt64("offset"); err == nil {
		offset = v
	}
	// sortby: col1,col2
	if v := c.GetString("sortby"); v != "" {
		sortby = strings.Split(v, ",")
	}
	// order: desc,asc
	if v := c.GetString("order"); v != "" {
		order = strings.Split(v, ",")
	}
	// query: k:v,k:v
	if v := c.GetString("query"); v != "" {
		for _, cond := range strings.Split(v, ",") {
			kv := strings.SplitN(cond, ":", 2)
			if len(kv) != 2 {
				c.Response(400, nil, errors.New("Error: invalid query key/value pair"))
				//c.Data["json"] = errors.New("Error: invalid query key/value pair")
				//c.ServeJSON()
				return
			}
			k, v := kv[0], kv[1]
			query[k] = v
		}
	}

	allApplications, err := models.GetAllApplications(query, fields, sortby, order, offset, limit)
	if err != nil {
		c.Response(400, nil, err)
	} else {
		c.Response(200, allApplications, nil)
	}
}

// Put ...
// @Title Put
// @Description update the Applications
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.Applications	true		"body for Applications content"
// @Success 200 {object} models.Applications
// @Failure 403 :id is not int
// @router /:id [put]
func (c *ApplicationsController) Put() {
	var err error
	var id int
	var aim *models.Aims

	idStr := c.Ctx.Input.Param(":id")
	if id, err = strconv.Atoi(idStr); err != nil {
		c.Response(400, nil, err)
	}
	var v *models.Applications

	if v, err = models.GetApplicationsById(id); err != nil {
		c.Response(400, nil, err)
	}
	var appData ApplicationData

	if err = json.Unmarshal(c.Ctx.Input.RequestBody, &appData); err != nil {
		c.Response(400, nil, err)
	}
	if aim, err = models.GetAimsById(appData.AimID); err != nil {
		c.Response(400, nil, err)
	}
	v.Aim = aim
	v.Description = appData.Description

	if appData.File != nil {
		log.Info(v.Url)
		if err = os.Remove(v.Url); err != nil {
			log.Error(err)
			c.Response(500, nil, err)
		}

		var uniqFilename uuid.UUID

		if uniqFilename, err = uuid.NewUUID(); err != nil {
			c.Response(500, nil, err)
		}
		filenameSlice := strings.Split(appData.FileName, ".")
		imgFileName := uniqFilename.String() + "." + filenameSlice[len(filenameSlice)-1]

		url := applicationsPath + imgFileName

		if err = ioutil.WriteFile(url, appData.File, 0644); err != nil {
			log.Error(err)
			c.Response(400, nil, err)
		}
		v.Url = url
		v.Filename = appData.FileName
		v.UniqId = uniqFilename.String()
	}

	if err = models.UpdateApplicationsById(v); err != nil {
		c.Response(500, nil, err)
	}
	c.Response(200, v, nil)

}

// Delete ...
// @Title Delete
// @Description delete the Applications
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *ApplicationsController) Delete() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	if err := models.DeleteApplications(id); err == nil {
		c.Data["json"] = "OK"
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}
