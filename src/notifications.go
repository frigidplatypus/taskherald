package main

import (
	"fmt"
	"time"
)

func CheckAndNotify(config *Config) {
	tasks, err := GetTasksWithNotifications(config.TaskBinary)
	if err != nil {
		LogError(err)
		return
	}

	now := time.Now()
	LogInfo(fmt.Sprintf("Checking %d tasks for notifications at %s", len(tasks), now.Format("15:04:05")))
	for _, task := range tasks {
		if shouldNotify(task, now) {
			LogInfo(fmt.Sprintf("Sending notification for task: %s", task.Description))
			err := SendNotification(config, task)
			if err != nil {
				LogError(fmt.Errorf("failed to send notification for task %s: %w", task.Description, err))
				continue
			}
			LogInfo(fmt.Sprintf("Notification sent, marking task %s as notified", task.Description))
			err = UpdateTaskNotified(config.TaskBinary, task.UUID)
			if err != nil {
				LogError(fmt.Errorf("failed to mark task %s as notified: %w", task.Description, err))
			}
		}
	}
}

func shouldNotify(task Task, now time.Time) bool {
	if task.NotificationDate == "" {
		return false
	}
	notifyTime, err := time.Parse("20060102T150405Z", task.NotificationDate)
	if err != nil {
		LogError(err)
		return false
	}
	// Only notify for tasks due today that are past due and not yet notified
	today := now.Truncate(24 * time.Hour)
	notifyDay := notifyTime.In(time.Local).Truncate(24 * time.Hour)
	if notifyDay != today {
		return false
	}
	notifyTimeLocal := notifyTime.In(time.Local)
	result := notifyTimeLocal.Before(now) && task.TaskheraldNotified == ""
	LogInfo(fmt.Sprintf("Task %s: notifyTime=%s, now=%s, before=%v, notified=%s, shouldNotify=%v",
		task.Description, notifyTimeLocal.Format("15:04:05"),
		now.Format("15:04:05"), notifyTimeLocal.Before(now), task.TaskheraldNotified, result))
	return result
}
