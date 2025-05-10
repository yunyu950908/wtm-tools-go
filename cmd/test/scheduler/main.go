package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
	"wtm-tools-go/core/scheduler"
	"wtm-tools-go/logger"
)

func main() {
	logger.Info().Msg("Starting scheduler test")

	timezone := flag.String("tz", "Asia/Shanghai", "Timezone for scheduler")
	flag.Parse()

	tz, err := time.LoadLocation(*timezone)
	if err != nil {
		logger.Error().Err(err).Msg("Failed to load timezone")
		os.Exit(1)
	}

	s := scheduler.NewScheduler(tz)
	logger.Info().Msg("Scheduler created")

	err = s.RegisterTask("test-task-1", "* * * * *", func() error {
		logger.Info().Msg("Task 1 executed - runs every minute")
		return nil
	})
	if err != nil {
		logger.Error().Err(err).Msg("Failed to register task 1")
		os.Exit(1)
	}

	err = s.RegisterTask("test-task-2", "*/2 * * * *", func() error {
		logger.Info().Msg("Task 2 executed - runs every 2 minutes")
		return nil
	})
	if err != nil {
		logger.Error().Err(err).Msg("Failed to register task 2")
		os.Exit(1)
	}

	// 测试错误任务
	err = s.RegisterTask("test-task-error", "* * * * *", func() error {
		logger.Info().Msg("Task with error is running")
		return fmt.Errorf("this is a test error")
	})
	if err != nil {
		logger.Error().Err(err).Msg("Failed to register error task")
		os.Exit(1)
	}

	// 测试panic任务
	err = s.RegisterTask("test-task-panic", "* * * * *", func() error {
		logger.Info().Msg("Task with panic is running")
		panic("this is a test panic")
	})
	if err != nil {
		logger.Error().Err(err).Msg("Failed to register panic task")
		os.Exit(1)
	}

	logger.Info().Int("count", s.TaskCount()).Msg("Registered tasks")

	// 启动调度器
	err = s.Start()
	if err != nil {
		logger.Error().Err(err).Msg("Failed to start scheduler")
		os.Exit(1)
	}
	logger.Info().Bool("running", s.IsRunning()).Msg("Scheduler status")

	// 测试重复启动
	err = s.Start()
	if err != nil {
		logger.Info().Err(err).Msg("Expected error when starting already running scheduler")
	}

	ctx, cancel := context.WithCancel(context.Background())

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		sig := <-sigCh
		logger.Info().Str("signal", sig.String()).Msg("Received signal")
		cancel()
	}()

	// 30秒后移除一个任务
	go func() {
		time.Sleep(30 * time.Second)
		if err := s.RemoveTask("test-task-1"); err != nil {
			logger.Error().Err(err).Msg("Failed to remove task")
		} else {
			logger.Info().Int("count", s.TaskCount()).Msg("Task removed, remaining tasks")
		}
	}()

	// 等待退出信号
	<-ctx.Done()

	// 停止调度器
	if err := s.Stop(); err != nil {
		logger.Error().Err(err).Msg("Failed to stop scheduler")
	}
	logger.Info().Bool("running", s.IsRunning()).Msg("Scheduler stopped")

	// 测试重复停止
	err = s.Stop()
	if err != nil {
		logger.Info().Err(err).Msg("Expected error when stopping already stopped scheduler")
	}

	logger.Info().Msg("Scheduler test completed")
}
