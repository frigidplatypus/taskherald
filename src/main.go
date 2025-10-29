package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
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

	if config.NtfyServer != "https://ntfy.sh" {
		LogInfo("Using custom ntfy server: " + config.NtfyServer)
	}
	LogInfo("Using ntfy topic: " + config.NtfyTopic)
	LogInfo("Check interval: " + config.TaskHeraldInterval.String())

	LogInfo("TaskHerald starting")

	// Startup tasks
	HandleStartup(config)

	// Main loop
	ticker := time.NewTicker(config.TaskHeraldInterval)
	defer ticker.Stop()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	for {
		select {
		case <-ticker.C:
			CheckAndNotify(config)
		case <-sigChan:
			LogInfo("Shutting down gracefully")
			return
		}
	}
}