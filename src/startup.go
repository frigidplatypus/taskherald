package main

import (
	"fmt"
	"time"
)

func HandleStartup(config *Config) {
	tasks, err := GetTasksWithNotifications()
	if err != nil {
		LogError(err)
		return
	}

	now := time.Now()
	// Log up to 10 tasks with future notification dates
	count := 0
	for _, task := range tasks {
		if count >= 10 {
			break
		}
		if task.NotificationDate != "" {
			notifyTime, err := time.Parse("20060102T150405Z", task.NotificationDate)
			if err != nil {
				continue
			}
			if notifyTime.After(now) {
				localTime := notifyTime.Local()
				notifyDate := localTime.Format("2006-01-02 15:04:05")
				LogInfo(fmt.Sprintf("Task: %s - %s (notify: %s)", task.UUID, task.Description, notifyDate))
				count++
			}
		}
	}
	if count == 0 {
		LogInfo("No tasks with future notification dates found")
	}
}
