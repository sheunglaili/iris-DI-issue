package main

import "github.com/kataras/golog"

// Controller is controller for /api
type Controller struct {
	Logger *golog.Logger
	Db	   *Database
}

// Get is get endpoint for /api/
func (c *Controller) Get() {
	c.Logger.Info("Get endpoint !")
	
}

