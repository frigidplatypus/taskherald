package main

import (
	"time"
)

func CheckAndNotify(config *Config) {
	tasks, err := GetTasksWithNotifications()
	if err != nil {
		LogError(err)
		return
	}

	now := time.Now()
	for _, task := range tasks {
		if shouldNotify(task, now) {
			err := SendNotification(config, task)
			if err != nil {
				LogError(err)
				continue
			}
			err = UpdateTaskNotified(task.UUID)
			if err != nil {
				LogError(err)
			}
		}
	}
}

func shouldNotify(task Task, now time.Time) bool {
	if task.UDAs["notification_date"] == "" {
		return false
	}
	notifyTime, err := time.Parse(time.RFC3339, task.UDAs["notification_date"])
	if err != nil {
		LogError(err)
		return false
	}
	return notifyTime.Before(now) && task.UDAs["taskherald_notified"] == ""
}