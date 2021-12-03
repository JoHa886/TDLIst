package service

import (
	"TDList/model"
	"TDList/serializer"
	"time"

	logging "github.com/sirupsen/logrus"
)

// 新增任务
type CreateTaskService struct {
	Title   string `json:"title" form:"title"`
	Content string `json:"content" form:"content"`
}

func (service *CreateTaskService) Create(id uint) serializer.Response {
	var user model.User
	model.GlobalDB.First(&user, id)
	task := model.Task{
		User:      user,
		Uid:       user.ID,
		Title:     service.Title,
		Status:    0,
		Content:   service.Content,
		StartTime: time.Now().Unix(),
		EndTime:   0,
	}
	err := model.GlobalDB.Create(&task).Error
	if err != nil {
		return serializer.Response{
			Status: 500,
			Msg:    "数据库创建任务失败",
		}
	}
	return serializer.Response{
		Status: 200,
		Msg:    "创建成功",
	}
}

// 查询任务
type TaskDetailService struct {
}

func (service *TaskDetailService) Get(tid string) serializer.Response {
	var task model.Task

	err := model.GlobalDB.First(&task, tid).Error
	if err != nil {
		return serializer.Response{
			Status: 400,
			Msg:    "查询失败",
		}
	}
	return serializer.Response{
		Status: 200,
		Msg:    "查询成功",
		Data:   serializer.BuildTask(task),
	}
}

// 查询任务列表
type TaskListService struct {
	PageNum  int `json:"page_num" form:"page_num"`
	PageSize int `json:"page_size" form:"page_size"`
}

func (service *TaskListService) List(uid uint) serializer.Response {
	var tasks []model.Task
	var total int64 = 0
	if service.PageSize == 0 {
		service.PageSize = 15
	}
	model.GlobalDB.Model(model.Task{}).Preload("User").Where("uid = ?", uid).Count(&total).
		Limit(service.PageSize).Offset((service.PageNum - 1) * service.PageSize).
		Find(&tasks)
	return serializer.BuildListResponse(serializer.BuildTasks(tasks), uint(total))
}

//删除任务的服务
type DeleteTaskService struct {
}

func (service *DeleteTaskService) Delete(id string) serializer.Response {
	var task model.Task
	code := 200
	err := model.GlobalDB.First(&task, id).Error
	if err != nil {
		logging.Info(err)
		code = 500
		return serializer.Response{
			Status: code,
			Msg:    "数据库错误",
			Error:  err.Error(),
		}
	}
	err = model.GlobalDB.Delete(&task).Error
	if err != nil {
		logging.Info(err)
		code = 500
		return serializer.Response{
			Status: code,
			Msg:    "数据库错误",
			Error:  err.Error(),
		}
	}
	return serializer.Response{
		Status: code,
		Msg:    "删除成功",
	}
}

//更新任务的服务
type UpdateTaskService struct {
	ID      uint   `form:"id" json:"id"`
	Title   string `form:"title" json:"title" binding:"required,min=2,max=100"`
	Content string `form:"content" json:"content" binding:"max=1000"`
	Status  int    `form:"status" json:"status"` //0 待办   1已完成
}

func (service *UpdateTaskService) Update(id string) serializer.Response {
	var task model.Task
	model.GlobalDB.Model(model.Task{}).Where("id = ?", id).First(&task)
	task.Content = service.Content
	task.Status = service.Status
	task.Title = service.Title
	code := 200
	err := model.GlobalDB.Save(&task).Error
	if err != nil {
		logging.Info(err)
		code = 500
		return serializer.Response{
			Status: code,
			Msg:    "数据库错误",
			Error:  err.Error(),
		}
	}
	return serializer.Response{
		Status: code,
		Msg:    "成功",
		Data:   "修改成功",
	}
}

//搜索任务的服务
type SearchTaskService struct {
	Info string `form:"info" json:"info"`
}

func (service *SearchTaskService) Search(uId uint) serializer.Response {
	var tasks []model.Task
	code := 200
	model.GlobalDB.Where("uid=?", uId).Preload("User").First(&tasks)
	err := model.GlobalDB.Model(&model.Task{}).Where("title LIKE ? OR content LIKE ?",
		"%"+service.Info+"%", "%"+service.Info+"%").Find(&tasks).Error
	if err != nil {
		logging.Info(err)
		code = 500
		return serializer.Response{
			Status: code,
			Msg:    "数据库错误",
			Error:  err.Error(),
		}
	}
	return serializer.Response{
		Status: code,
		Msg:    "成功",
		Data:   serializer.BuildTasks(tasks),
	}
}
