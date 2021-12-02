package service

import (
	"TDList/model"
	"TDList/serializer"
	"time"
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
