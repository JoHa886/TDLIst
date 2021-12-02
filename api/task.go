package api

import (
	"TDList/pkg/utils"
	"TDList/service"

	"github.com/gin-gonic/gin"
	logging "github.com/sirupsen/logrus"
)

func CreateTask(c *gin.Context) {
	var createTask service.CreateTaskService
	claim, _ := utils.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&createTask); err == nil {
		res := createTask.Create(claim.Id)
		c.JSON(200, res)
	} else {
		logging.Error(err)
		c.JSON(400, err)
	}

}

func TaskDetail(c *gin.Context) {
	var taskDetail service.TaskDetailService
	// claim, _ := utils.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&taskDetail); err == nil {
		res := taskDetail.Get(c.Param("id"))
		c.JSON(200, res)
	} else {
		logging.Error(err)
		c.JSON(400, err)
	}

}
