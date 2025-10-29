package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"path/filepath"
	"strings"
	"syscall"
	"time"
)

func main() {
	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), "TaskHerald - A systemd service that monitors Taskwarrior tasks and sends ntfy.sh notifications\n\n")
		fmt.Fprintf(flag.CommandLine.Output(), "USAGE:\n")
		fmt.Fprintf(flag.CommandLine.Output(), "  taskherald [options]\n\n")
		fmt.Fprintf(flag.CommandLine.Output(), "OPTIONS:\n")
		flag.PrintDefaults()
		fmt.Fprintf(flag.CommandLine.Output(), "CONFIGURATION:\n")
		fmt.Fprintf(flag.CommandLine.Output(), "  Environment Variables:\n")
		fmt.Fprintf(flag.CommandLine.Output(), "    NTFY_SERVER        ntfy server URL (default: https://ntfy.sh)\n")
		fmt.Fprintf(flag.CommandLine.Output(), "    NTFY_TOPIC         default topic (default: taskherald, or taskherald-RANDOM for ntfy.sh)\n")
		fmt.Fprintf(flag.CommandLine.Output(), "    NTFY_TOPIC_FILE    path to file containing ntfy topic\n")
		fmt.Fprintf(flag.CommandLine.Output(), "    TASKHERALD_INTERVAL check interval in seconds (default: 60)\n")
		fmt.Fprintf(flag.CommandLine.Output(), "EXAMPLES:\n")
		fmt.Fprintf(flag.CommandLine.Output(), "  taskherald                    # Start the service\n")
		fmt.Fprintf(flag.CommandLine.Output(), "  taskherald --help             # Show this help\n")
		fmt.Fprintf(flag.CommandLine.Output(), "  NTFY_TOPIC=my-tasks taskherald # Start with custom topic\n")
	}

	help := flag.Bool("help", false, "Show help")
	flag.Parse()

	if *help {
		flag.Usage()
		return
	}

	config, err := LoadConfig()
	if err != nil {
		panic(err)
	}

	InitLogger()

	CheckTaskwarriorConfig()

	if config.NtfyServer != "https://ntfy.sh" {
		LogInfo("Using custom ntfy server: " + config.NtfyServer)
	}
	LogInfo("Using ntfy topic: " + config.NtfyTopic)
	LogInfo("Check interval: " + config.TaskHeraldInterval.String())

	LogInfo("TaskHerald starting")

	// Startup tasks
	HandleStartup(config)

	// Main loop - check every interval
	ticker := time.NewTicker(config.TaskHeraldInterval)
	defer ticker.Stop()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	for {
		select {
		case <-ticker.C:
			LogInfo("Ticker fired, checking for notifications")
			CheckAndNotify(config)
		case <-sigChan:
			LogInfo("Shutting down gracefully")
			return
		}
	}
}

func CheckTaskwarriorConfig() {
	taskrcPath := os.Getenv("TASKRC")
	if taskrcPath == "" {
		homeDir, err := os.UserHomeDir()
		if err != nil {
			LogInfo("WARNING: Could not get home directory for Taskwarrior config check")
			return
		}
		// Try common locations
		possiblePaths := []string{
			filepath.Join(homeDir, ".taskrc"),
			filepath.Join(homeDir, ".config", "task", "taskrc"),
			"/etc/taskrc",
		}
		for _, path := range possiblePaths {
			if _, err := os.Stat(path); err == nil {
				taskrcPath = path
				break
			}
		}
	}

	if taskrcPath == "" {
		LogInfo("WARNING: Could not find Taskwarrior config file (.taskrc). Please ensure UDAs are defined: uda.notification_date.type=date and uda.taskherald_notified.type=date")
		return
	}

	hasNotificationDate, hasTaskheraldNotified := scanConfigFile(taskrcPath, make(map[string]bool))

	if !hasNotificationDate {
		LogInfo("WARNING: uda.notification_date.type not found in Taskwarrior config. Please add: uda.notification_date.type=date")
	}
	if !hasTaskheraldNotified {
		LogInfo("WARNING: uda.taskherald_notified.type not found in Taskwarrior config. Please add: uda.taskherald_notified.type=date")
	}
}

func scanConfigFile(path string, visited map[string]bool) (bool, bool) {
	if visited[path] {
		return false, false
	}
	visited[path] = true

	file, err := os.Open(path)
	if err != nil {
		LogInfo(fmt.Sprintf("WARNING: Could not open Taskwarrior config at %s: %v", path, err))
		return false, false
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	hasNotificationDate := false
	hasTaskheraldNotified := false

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if strings.HasPrefix(line, "#") {
			continue
		}
		if strings.HasPrefix(line, "include ") {
			includePath := strings.TrimSpace(strings.TrimPrefix(line, "include "))
			if !filepath.IsAbs(includePath) {
				// Relative to the current file's directory
				includePath = filepath.Join(filepath.Dir(path), includePath)
			}
			incNotif, incNotified := scanConfigFile(includePath, visited)
			hasNotificationDate = hasNotificationDate || incNotif
			hasTaskheraldNotified = hasTaskheraldNotified || incNotified
		}
		if strings.Contains(line, "uda.notification_date.type=") {
			hasNotificationDate = true
		}
		if strings.Contains(line, "uda.taskherald_notified.type=") {
			hasTaskheraldNotified = true
		}
	}

	if err := scanner.Err(); err != nil {
		LogInfo(fmt.Sprintf("WARNING: Error reading Taskwarrior config at %s: %v", path, err))
	}

	return hasNotificationDate, hasTaskheraldNotified
}
