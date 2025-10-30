package main

import (
	"fmt"
	"strings"

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
	var parts []string

	if task.Project != "" {
		parts = append(parts, fmt.Sprintf("Project:%s", task.Project))
	}

	parts = append(parts, task.Description)

	if task.Due != "" {
		parts = append(parts, fmt.Sprintf("Due:%s", task.Due))
	}

	return strings.Join(parts, " ")
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