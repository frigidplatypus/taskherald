package main

import (
	"encoding/json"
	"os/exec"
	"time"
)

type Task struct {
	UUID        string            `json:"uuid"`
	Description string            `json:"description"`
	Project     string            `json:"project,omitempty"`
	Due         string            `json:"due,omitempty"` // ISO datetime string
	Tags        []string          `json:"tags"`
	Priority    string            `json:"priority,omitempty"`
	UDAs        map[string]string `json:"udas"`
}

func GetTasksWithNotifications() ([]Task, error) {
	cmd := exec.Command("task", "notification_date.after:now", "export")
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
	cmd := exec.Command("task", uuid, "modify", "taskherald_notified:"+time.Now().Format(time.RFC3339))
	return cmd.Run()
}