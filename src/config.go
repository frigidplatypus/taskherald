package main

import (
	"crypto/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

type Config struct {
	NtfyServer         string
	NtfyTopic          string
	TaskHeraldInterval time.Duration
}

func LoadConfig() (*Config, error) {
	config := &Config{
		NtfyServer: ensureProtocol(getEnv("NTFY_SERVER", "https://ntfy.sh")),
	}

	// Set topic: check file first, then env var, then generate random topic for default server
	topicFile := os.Getenv("NTFY_TOPIC_FILE")
	if topicFile != "" {
		topicBytes, err := os.ReadFile(topicFile)
		if err != nil {
			return nil, err
		}
		config.NtfyTopic = strings.TrimSpace(string(topicBytes))
	} else {
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

func ensureProtocol(url string) string {
	if strings.HasPrefix(url, "http://") || strings.HasPrefix(url, "https://") {
		return url
	}
	return "https://" + url
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
