package main

import (
	"fmt"

	"heckel.io/ntfy/client"
)

func SendNotification(config *Config, task Task) error {
	c := client.New(&client.Config{
		DefaultHost: config.NtfyServer,
	})

	topic := config.NtfyTopic
	if task.UDAs["ntfy_topic"] != "" {
		topic = task.UDAs["ntfy_topic"]
	}

	msg := formatMessage(task)

	_, err := c.Publish(topic, msg, client.WithTags(task.Tags), client.WithPriority(mapPriority(task.Priority)))
	return err
}

func formatMessage(task Task) string {
	project := task.Project
	if project == "" {
		project = ""
	}
	due := task.Due
	if due == "" {
		due = ""
	}
	return fmt.Sprintf("Project:%s %s Due:%s", project, task.Description, due)
}

func mapPriority(p string) string {
	switch p {
	case "H":
		return "5"
	case "M":
		return "4"
	case "L":
		return "1"
	default:
		return "3"
	}
}