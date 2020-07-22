package main

import (
	"github.com/kataras/golog"
)

// Environment is struct that simulate reading environment 
type Environment struct {
	Logger *golog.Logger
}

// NewEnvironment return new instance Environemnt 
func NewEnvironment(logger *golog.Logger) *Environment {
	logger.Debug("Prentending to read heavily from environment");
	logger.Debug("New Controller Again");
	return &Environment{
		Logger: logger,
	}
}