package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/hipeday/rosen/internal/repository"
)

type ConsoleHandler struct {
	repo repository.Repository[int64]
}

// Login 后台管理系统登录
func (c *ConsoleHandler) Login(ctx *gin.Context) {
	ctx.JSON(200, gin.H{"ping": "pong", "action": "login"})
}

// Logout 后台管理系统登录
func (c *ConsoleHandler) Logout(ctx *gin.Context) {
	ctx.JSON(200, gin.H{"ping": "pong", "action": "logout"})
}

// Current 后台管理系统登录
func (c *ConsoleHandler) Current(ctx *gin.Context) {
	ctx.JSON(200, gin.H{"ping": "pong", "action": "logout"})
}

func (c *ConsoleHandler) GetType() Type {
	return Console
}

func (c *ConsoleHandler) GetRepository() repository.Repository[int64] {
	return c.repo
}

func NewConsoleHandler() *ConsoleHandler {
	return &ConsoleHandler{}
}
