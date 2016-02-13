package widgets

import (
	"time"
	"github.com/lietu/godashery/godashery"
)

type ClockData struct {
	Type string    `json:"type"`
	Id   string    `json:"id"`
	Time string    `json:"time"`
}

func NewClockData() *ClockData {
	cd := ClockData{
		"widgetData",
		"clock",
		time.Now().Format("15:04:05"),
	}

	return &cd
}

type ClockWidget struct {
	godashery.BaseWidget
}

func (c *ClockWidget) CheckUpdate() {
	if (c.CheckUpdateTime()) {
		c.Updating = true  // Prevents multiple simultaneous updates
		c.Update()
		c.Updating = false
	}
}

func (c *ClockWidget) Update() {
	c.SetValue(NewClockData())
}

func NewClock(id string) godashery.Widget {
	c := ClockWidget{
		godashery.NewBaseWidget(id),
	}

	return &c
}

func init() {
	godashery.RegisterWidget("ClockWidget", NewClock)
}
