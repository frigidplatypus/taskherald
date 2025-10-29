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

	// Log up to 10 tasks
	count := len(tasks)
	if count > 10 {
		count = 10
	}
	for i := 0; i < count; i++ {
		LogInfo(fmt.Sprintf("Task: %s - %s", tasks[i].UUID, tasks[i].Description))
	}
	if len(tasks) == 0 {
		LogInfo("No tasks with notification_date found")
	}

	// Send summary for past 7 days unsent
	SendStartupSummary(config, tasks)
}

func SendStartupSummary(config *Config, tasks []Task) {
	sevenDaysAgo := time.Now().Add(-7 * 24 * time.Hour)
	var summaryTasks []Task
	for _, task := range tasks {
		if task.UDAs["notification_date"] != "" {
			notifyTime, err := time.Parse(time.RFC3339, task.UDAs["notification_date"])
			if err != nil {
				continue
			}
			if notifyTime.After(sevenDaysAgo) && task.UDAs["taskherald_notified"] == "" {
				summaryTasks = append(summaryTasks, task)
			}
		}
	}

	if len(summaryTasks) > 0 {
		summaryMsg := "Missed notifications:\n"
		for _, t := range summaryTasks {
			summaryMsg += fmt.Sprintf("- %s\n", t.Description)
		}
		// Send summary via ntfy
		summaryTask := Task{Description: summaryMsg}
		err := SendNotification(config, summaryTask)
		if err != nil {
			LogError(err)
		}
	}
}