package main

import (
	"crypto/rand"
	"os"
	"strconv"
	"time"
)

type Config struct {
	NtfyServer         string
	NtfyTopic          string
	TaskHeraldInterval time.Duration
	TaskHeraldStateFile string
}

func LoadConfig() (*Config, error) {
	config := &Config{
		NtfyServer:         getEnv("NTFY_SERVER", "https://ntfy.sh"),
		TaskHeraldStateFile: getEnv("TASKHERALD_STATE_FILE", "/var/lib/taskherald/notifications.json"),
	}

	// Set topic: use env var if set, otherwise generate random topic for default server
	topicEnv := os.Getenv("NTFY_TOPIC")
	if topicEnv != "" {
		config.NtfyTopic = topicEnv
	} else if config.NtfyServer == "https://ntfy.sh" {
		randomSuffix, err := generateRandomAlphanumeric(8)
		if err != nil {
			return nil, err
		}
		config.NtfyTopic = "taskherald-" + randomSuffix
	} else {
		config.NtfyTopic = "taskherald"
	}

	intervalStr := getEnv("TASKHERALD_INTERVAL", "60")
	interval, err := strconv.Atoi(intervalStr)
	if err != nil {
		return nil, err
	}
	config.TaskHeraldInterval = time.Duration(interval) * time.Second

	return config, nil
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

func generateRandomAlphanumeric(length int) (string, error) {
	const charset = "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	bytes := make([]byte, length)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	for i, b := range bytes {
		bytes[i] = charset[b%byte(len(charset))]
	}
	return string(bytes), nil
}