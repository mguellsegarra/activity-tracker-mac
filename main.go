package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/getlantern/systray"
	"github.com/lextoumbourou/idle"
)

const idleThreshold = 1 * time.Minute // Set the idle time threshold here

func main() {
	systray.Run(onReady, onExit)
}

func onReady() {
	systray.SetTitle("Starting...")
	systray.SetTooltip("Activity Tracker")

	mQuit := systray.AddMenuItem("Quit", "Quit the application")

	ticker := time.NewTicker(1 * time.Second)
	totalActiveTime, lastReset := readActiveTime()
	lastActiveTime := time.Now()

	go func() {
		for {
			select {
			case <-ticker.C:
				idleTime, err := idle.Get()
				if err != nil {
					log.Printf("Error getting idle time: %v\n", err)
					continue
				}
				log.Printf("Idle time: %s\n", idleTime)

				if time.Now().Day() != lastReset.Day() {
					log.Println("Day change detected, resetting active time")
					totalActiveTime = 0
					lastReset = time.Now()
					writeActiveTime(totalActiveTime, lastReset)
				}

				if idleTime < idleThreshold {
					newActiveTime := time.Now().Sub(lastActiveTime)
					totalActiveTime += newActiveTime
					log.Printf("Active: %s, Total active: %s\n", newActiveTime, totalActiveTime)
					writeActiveTime(totalActiveTime, lastReset)
					lastActiveTime = time.Now()
				} else {
					lastActiveTime = time.Now()
					log.Println("System idle")
				}

				systray.SetTitle(formatDuration(totalActiveTime))
			case <-mQuit.ClickedCh:
				log.Println("Exiting, writing final active time.")
				writeActiveTime(totalActiveTime, lastReset)
				systray.Quit()
				return
			}
		}
	}()
}

func onExit() {
	log.Println("Clean up before exit")
}

func writeActiveTime(d time.Duration, lastReset time.Time) {
	configPath := filepath.Join(os.Getenv("HOME"), ".config", "activity_tracker")
	os.MkdirAll(configPath, os.ModePerm)

	dataPath := filepath.Join(configPath, "active_time")
	data := fmt.Sprintf("%d,%d", int64(d), lastReset.Unix())
	err := ioutil.WriteFile(dataPath, []byte(data), 0644)
	if err != nil {
		log.Printf("Error writing active time to file: %v", err)
	}
}

func readActiveTime() (time.Duration, time.Time) {
	configPath := filepath.Join(os.Getenv("HOME"), ".config", "activity_tracker")
	dataPath := filepath.Join(configPath, "active_time")
	data, err := ioutil.ReadFile(dataPath)
	if err != nil {
		if os.IsNotExist(err) {
			log.Println("Active time file not found, initializing new file...")
			now := time.Now()
			writeActiveTime(0, now)
			return 0, now
		}
		log.Printf("Error reading active time from file: %v", err)
		return 0, time.Now()
	}
	parts := strings.Split(string(data), ",")
	if len(parts) != 2 {
		log.Println("File data format incorrect, reinitializing...")
		now := time.Now()
		writeActiveTime(0, now)
		return 0, now
	}
	totalTime, err := strconv.ParseInt(parts[0], 10, 64)
	if err != nil {
		log.Printf("Error parsing active time from file: %v", err)
		return 0, time.Now()
	}
	lastReset, err := strconv.ParseInt(parts[1], 10, 64)
	if err != nil {
		log.Printf("Error parsing last reset time from file: %v", err)
		return 0, time.Now()
	}
	return time.Duration(totalTime), time.Unix(lastReset, 0)
}

func formatDuration(d time.Duration) string {
	hours := int(d.Hours())
	minutes := int(d.Minutes()) % 60
	if hours > 0 {
		return fmt.Sprintf("%dh %dm", hours, minutes)
	}
	return fmt.Sprintf("%dm", minutes)
}
