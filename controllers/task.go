package controllers

import (
	"log"
	"net/http"
	"test-v2/models"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type CreateTaskInput struct {
	AssignedTo string `json: "assignedTo"`
	Task       string `json: "test"`
	Deadline   string `json: "deadline"`
}

type UpdateTaskInput struct {
	AssignedTo string `json: "assignedTo"`
	Task       string `json: "test"`
	Deadline   string `json: "deadline"`
}

// POST/Tasks
// CreateTask
func CreateTask(c *gin.Context) {
	log.Println("c =>", c)
	var input CreateTaskInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	log.Println("input =>", input)

	date := "2006-01-02"
	deadline, _ := time.Parse(date, input.Deadline)

	task := models.Task{
		AssignedTo: input.AssignedTo,
		Task:       input.Task,
		Deadline:   deadline,
	}

	db := c.MustGet("db").(*gorm.DB)
	db.Create(&task)
	c.JSON(http.StatusOK, gin.H{
		"data": task,
	})
}

// GET /tasks
// Get all task
func FindTasks(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var task []models.Task
	db.Find(&task)

	c.JSON(http.StatusOK, gin.H{
		"data": task,
	})
}

// GET /task/:id
// Find a task
func FindTask(c *gin.Context) {
	var task models.Task

	db := c.MustGet("db").(*gorm.DB)

	if err := db.Where("id = ?", c.Param("id")).First(&task).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Record not fonud",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": task,
	})
}

func UpdateTask(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var task models.Task

	if err := db.Where("id = ?", c.Param("id")).First(&task).Error; err != nil {
		c.JSON(http.StatusBadGateway, gin.H{
			"error": "Record not found",
		})
		return
	}

	var input UpdateTaskInput
	log.Println("input =>", input)
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	log.Println("input =>", input)
	log.Println("input.Deadline =>", input.Deadline)

	date := "2006-01-02"
	deadline, _ := time.Parse(date, input.Deadline)

	var updateInput models.Task
	updateInput.AssignedTo = input.AssignedTo
	updateInput.Deadline = deadline
	// updateInput.Deadline = deadline
	updateInput.Task = input.Task

	log.Println("updateInput => ", updateInput)
	log.Println("updateInput.Deadline => ", updateInput.Deadline)

	db.Model(&task).Updates(updateInput)

	c.JSON(http.StatusOK, gin.H{
		"data": updateInput,
	})
}

// Delete /task/:id
// Delete a task
func DeleteTask(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var book models.Task
	if err := db.Where("id = ?", c.Param("id")).First(&book).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Record not found",
		})
	}

	db.Delete(&book)
	c.JSON(http.StatusOK, gin.H{
		"data": true,
	})
}
