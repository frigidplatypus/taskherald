package main

import (
	"encoding/json"
	"os/exec"
)

type Task struct {
	UUID               string            `json:"uuid"`
	Description        string            `json:"description"`
	Project            string            `json:"project,omitempty"`
	Due                string            `json:"due,omitempty"` // ISO datetime string
	Tags               []string          `json:"tags"`
	Priority           string            `json:"priority,omitempty"`
	NotificationDate   string            `json:"notification_date,omitempty"`
	UDAs               map[string]string `json:"udas"`
	TaskheraldNotified string            `json:"taskherald_notified,omitempty"`
}

func GetTasksWithNotifications() ([]Task, error) {
	cmd := exec.Command("task", "notification_date.any:", "export")
	output, err := cmd.Output()
	if err != nil {
		return nil, err
	}

	var tasks []Task
	err = json.Unmarshal(output, &tasks)
	if err != nil {
		return nil, err
	}

	return tasks, nil
}

func UpdateTaskNotified(uuid string) error {
	// Update the UDA via task command
	cmd := exec.Command("task", uuid, "modify", "taskherald_notified:now")
	return cmd.Run()
}
