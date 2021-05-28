package Controllers

import (
	"vahidid/20-project-go/ApiHelpers"
	"vahidid/20-project-go/Models"

	"github.com/gin-gonic/gin"
)

func ListTask(c *gin.Context) {
	var task []Models.Task

	err := Models.GetAllTasks(&task)
	if err != nil {
		ApiHelpers.RespondJSON(c, 404, nil)
	} else {
		ApiHelpers.RespondJSON(c, 200, task)
	}
}

func GetTask(c *gin.Context) {
	id := c.Params.ByName("id")
	var task Models.Task
	err := Models.GetOneTask(&task, id)
	if err != nil {
		ApiHelpers.RespondJSON(c, 404, nil)
	} else {
		ApiHelpers.RespondJSON(c, 200, task)
	}
}

func CreateTask(c *gin.Context) {
	var task Models.Task

	if err := c.ShouldBindJSON(&task); err != nil {
		ApiHelpers.RespondJSON(c, 400, err.Error())
		return
	}

	c.BindJSON(&task)
	err := Models.SaveTask(&task)

	if err != nil {
		ApiHelpers.RespondJSON(c, 400, nil)
	} else {
		ApiHelpers.RespondJSON(c, 201, task)
	}
}

func UpdateTask(c *gin.Context) {
	var task Models.Task
	id := c.Params.ByName("id")
	err := Models.GetOneTask(&task, id)
	if err != nil {
		ApiHelpers.RespondJSON(c, 404, nil)
	}

	c.BindJSON(&task)
	err = Models.PutOneTask(&task, id)
	if err != nil {
		ApiHelpers.RespondJSON(c, 400, nil)
	} else {
		ApiHelpers.RespondJSON(c, 200, task)
	}
}

func DeleteTask(c *gin.Context) {
	id := c.Params.ByName("id")
	var task Models.Task
	err := Models.DeleteTask(&task, id)
	if err != nil {
		ApiHelpers.RespondJSON(c, 404, nil)
	} else {
		ApiHelpers.RespondJSON(c, 200, task)
	}
}
