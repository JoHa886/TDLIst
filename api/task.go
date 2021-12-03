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

// 任务列表
func TaskList(c *gin.Context) {
	var taskList service.TaskListService
	claim, _ := utils.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&taskList); err == nil {
		res := taskList.List(claim.Id)
		c.JSON(200, res)
	} else {
		logging.Error(err)
		c.JSON(400, err)
	}

}

// 删除任务
func DeleteTask(c *gin.Context) {
	deleteTaskService := service.DeleteTaskService{}
	res := deleteTaskService.Delete(c.Param("id"))
	c.JSON(200, res)
}

// 修改任务
func UpdateTask(c *gin.Context) {
	updateTaskService := service.UpdateTaskService{}
	if err := c.ShouldBind(&updateTaskService); err == nil {
		res := updateTaskService.Update(c.Param("id"))
		c.JSON(200, res)
	} else {
		c.JSON(200, err)
		logging.Info(err)
	}
}

// 查询任务
func SearchTasks(c *gin.Context) {
	searchTaskService := service.SearchTaskService{}
	claim, _ := utils.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&searchTaskService); err == nil {
		res := searchTaskService.Search(claim.Id)
		c.JSON(200, res)
	} else {
		c.JSON(200, err)
		logging.Info(err)
	}
}
