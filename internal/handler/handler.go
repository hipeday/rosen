package handler

import (
	"fmt"
	"sync"
)

// Type 处理程序类型
type Type string

type Handler interface {
	GetType() Type // 获取处理器类型
}

// Factory 工厂和缓存管理
type Factory struct {
	mu       sync.Mutex
	handlers sync.Map // 缓存已创建的 Handler 实例
}

const (
	Console Type = "console"
)

// GetHandler 根据类型获取 Handler，如果不存在则创建
func (f *Factory) GetHandler(t Type) (Handler, error) {
	// 尝试从缓存中获取
	if h, ok := f.handlers.Load(t); ok {
		return h.(Handler), nil
	}

	// 加锁确保创建线程安全
	f.mu.Lock()
	defer f.mu.Unlock()

	// 双重检查，避免重复创建
	if h, ok := f.handlers.Load(t); ok {
		return h.(Handler), nil
	}

	// 根据类型创建新的 Handler
	var handler Handler
	switch t {
	case Console:
		handler = NewConsoleHandler()
	default:
		return nil, fmt.Errorf("unknown handler type: %s", t)
	}

	// 缓存创建的 Handler
	f.handlers.Store(t, handler)
	return handler, nil
}

// NewHandlerFactory 创建一个新的 HandlerFactory
func NewHandlerFactory() *Factory {
	return &Factory{}
}
