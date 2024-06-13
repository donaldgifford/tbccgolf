/*
	package loggy

loggy.go

This file contains a log function and its configuration. This just makes pretty logs in colors so that
it easier to see and configure log output.

Copyright © 2023-2024 CG Forge LLC contact@cgforge.io
*/
package loggy

import (
	"os"

	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/log"
)

func Loggy() *log.Logger {
	styles := log.DefaultStyles()
	// set debug level styles
	styles.Levels[log.DebugLevel] = lipgloss.NewStyle().
		SetString("Debug").
		Padding(0, 1, 0, 1).
		Background(lipgloss.Color("#38d1d6")).
		Foreground(lipgloss.Color("0"))
	// set info level styles
	styles.Levels[log.InfoLevel] = lipgloss.NewStyle().
		SetString("Info").
		Padding(0, 1, 0, 1).
		Background(lipgloss.Color("#42f5b9")).
		Foreground(lipgloss.Color("0"))
	// set warn level styles
	styles.Levels[log.WarnLevel] = lipgloss.NewStyle().
		SetString("WARN").
		Padding(0, 1, 0, 1).
		Background(lipgloss.Color("#e69a30")).
		Foreground(lipgloss.Color("0"))
	// set error level styles
	styles.Levels[log.ErrorLevel] = lipgloss.NewStyle().
		SetString("🔥ERROR!!🔥").
		Padding(0, 1, 0, 1).
		Background(lipgloss.Color("204")).
		Foreground(lipgloss.Color("0"))
	// set fatal level styles
	styles.Levels[log.FatalLevel] = lipgloss.NewStyle().
		SetString("💀 FATAL!! 💀").
		Padding(0, 1, 0, 1).
		Background(lipgloss.Color("#f558f0")).
		Foreground(lipgloss.Color("0"))
	// set styles for 'err' keys and values
	styles.Keys["err"] = lipgloss.NewStyle().Foreground(lipgloss.Color("204"))
	styles.Values["err"] = lipgloss.NewStyle().Bold(true)
	// set styles for 'fatal' keys and values
	styles.Keys["fatal"] = lipgloss.NewStyle().Foreground(lipgloss.Color("#f558f0"))
	styles.Values["fatal"] = lipgloss.NewStyle().Bold(true)
	// set styles for 'warn' keys
	styles.Keys["warn"] = lipgloss.NewStyle().Background(lipgloss.Color("#e69a30")).Foreground(lipgloss.Color("0"))
	// configure timestamp styles
	styles.Timestamp = lipgloss.NewStyle()
	// create new logger with styles
	logger := log.New(os.Stderr)
	logger.SetStyles(styles)
	logger.SetReportTimestamp(true)
	logger.SetTimeFormat("01/02/2006 15:04:05")
	logger.SetLevel(log.Level(0))
	// logger.SetLevel(log.Level(viper.GetInt32("logs.level")))
	// return logger object to use
	return logger
}
