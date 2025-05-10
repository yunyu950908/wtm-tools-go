package scheduler

import (
	"fmt"
	"sync"
	"time"
	"wtm-tools-go/logger"

	"github.com/go-co-op/gocron/v2"
)

// TaskFunc 定义任务函数类型
type TaskFunc func() error

// Scheduler 任务调度器结构体
type Scheduler struct {
	scheduler gocron.Scheduler
	tasks     map[string]gocron.Job // 存储任务，键为任务名称
	mutex     sync.RWMutex          // 保护任务列表的并发访问
	isRunning bool                  // 调度器运行状态
}

// NewScheduler 创建一个新的调度器实例
func NewScheduler(timezone *time.Location) *Scheduler {
	if timezone == nil {
		tz, err := time.LoadLocation("Asia/Shanghai")
		if err != nil {
			panic(err)
		}
		timezone = tz
	}
	scheduler, err := gocron.NewScheduler(gocron.WithLocation(timezone))
	if err != nil {
		panic(err)
	}
	return &Scheduler{
		scheduler: scheduler,
		tasks:     make(map[string]gocron.Job),
		isRunning: false,
	}
}

func safeJobWrapper(jobName string, fn func() error) func() {
	return func() {
		defer func() {
			if r := recover(); r != nil {
				logger.Error().Str("task", jobName).Any("panic", r).Msg("Task panic")
			}
		}()
		if err := fn(); err != nil {
			logger.Error().Str("task", jobName).Err(err).Msg("Task error")
		}
	}
}

// RegisterTask 注册一个新任务
// name: 任务名称，唯一标识
// cron: cron 表达式
// taskFunc: 任务执行的函数
func (s *Scheduler) RegisterTask(name, cron string, taskFunc TaskFunc) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	// 检查任务名称是否已存在
	if _, exists := s.tasks[name]; exists {
		return fmt.Errorf("task with name '%s' already exists", name)
	}

	// 创建任务
	job, err := s.scheduler.NewJob(
		gocron.CronJob(cron, false),
		gocron.NewTask(
			safeJobWrapper(name, taskFunc),
		),
		gocron.WithSingletonMode(gocron.LimitModeReschedule),
	)
	if err != nil {
		return fmt.Errorf("failed to create task '%s': %w", name, err)
	}

	// 存储任务
	s.tasks[name] = job
	logger.Info().Str("task", name).Str("cron", cron).Msg("Task registered successfully")
	return nil
}

// RemoveTask 删除一个任务
func (s *Scheduler) RemoveTask(name string) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	if job, exists := s.tasks[name]; exists {
		if err := s.scheduler.RemoveJob(job.ID()); err != nil {
			return fmt.Errorf("failed to remove task '%s': %w", name, err)
		}
		delete(s.tasks, name)
		logger.Info().Str("task", name).Msg("Task removed successfully")
		return nil
	}
	return fmt.Errorf("task with name '%s' not found", name)
}

// TaskCount 查询当前任务数量
func (s *Scheduler) TaskCount() int {
	s.mutex.RLock()
	defer s.mutex.RUnlock()
	return len(s.tasks)
}

// Start 启动调度器（非阻塞）
func (s *Scheduler) Start() error {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	if s.isRunning {
		return fmt.Errorf("scheduler is already running")
	}

	s.scheduler.Start()
	s.isRunning = true
	logger.Info().Msg("Scheduler started")
	return nil
}

// Stop 停止调度器
func (s *Scheduler) Stop() error {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	if !s.isRunning {
		return fmt.Errorf("scheduler is not running")
	}

	if err := s.scheduler.StopJobs(); err != nil {
		return fmt.Errorf("failed to stop scheduler: %w", err)
	}
	s.isRunning = false
	logger.Info().Msg("Scheduler stopped")
	return nil
}

// IsRunning 检查调度器是否正在运行
func (s *Scheduler) IsRunning() bool {
	s.mutex.RLock()
	defer s.mutex.RUnlock()
	return s.isRunning
}
