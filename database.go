package main

import "github.com/kataras/golog"

// Database is a stub struct which pretent to be a datasource
type Database struct {
	Logger *golog.Logger
	Env    *Environment
}

// NewDatabse return new instance of Database
func NewDatabase(logger *golog.Logger, env *Environment) *Database{
	logger.Debug("Prentend to connecting to db");
	logger.Debug("New Datbase again");
	return &Database{
		Logger : logger,
		Env: env,
	}
}