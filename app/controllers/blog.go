package controllers

import (
	"github.com/revel/revel"
	// "letsgo/app/models"
)

type Blog struct {
	GorpController
}

func (c Blog) Index() revel.Result {
	return c.Render()
}

func (c Blog) Edit(id int) revel.Result {
	return c.Render()
}

func (c Blog) Detail(id int) revel.Result {
	return c.Render()
}

func (c Blog) Add() revel.Result {
	return c.Render()
}

func (c Blog) Review() revel.Result {
	return c.Render()
}

func (c Blog) List() revel.Result {
	return c.Render()
}

