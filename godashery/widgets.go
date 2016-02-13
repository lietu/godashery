package godashery

import (
	"log"
	"time"
	"bytes"
	"encoding/json"
)

var availableWidgets map[string]WidgetConstructor = map[string]WidgetConstructor{};

var widgets map[string]Widget = map[string]Widget{}

type Widget interface {
	CheckUpdate()
	GetTemplate() []byte
	GetLastValue() []byte
}

type WidgetConstructor func(id string) Widget

type BaseWidget struct {
	Id             string
	IntervalMillis time.Duration
	NextUpdate     time.Time
	Template       string
	LastValue      []byte
	Updating	   bool
}

func (b *BaseWidget) GetTemplate() []byte {
	return []byte("template")
}

func (b *BaseWidget) GetLastValue() []byte {
	return b.LastValue;
}

func (b *BaseWidget) SetValue(v interface{}) {
	result, err := json.Marshal(&v)

	if err != nil {
		log.Fatalf("error: %v", err)
	}

	// No need to bother clients with identical "updates"
	if bytes.Equal(result, b.LastValue) {
		return
	}

	log.Printf("Widget %s value updated", b.Id)

	b.LastValue = result

	for _, c := range connections {
		c.Send(result)
	}
}

func (b *BaseWidget) CheckUpdateTime() bool {
	if b.Updating {
		return false
	}

	now := time.Now()

	if now.After(b.NextUpdate) {
		b.NextUpdate = now.Add(b.IntervalMillis * time.Millisecond)
		return true
	}

	return false
}

func NewBaseWidget(id string) BaseWidget {
	b := BaseWidget{
		id,
		100,
		time.Now(),
		"",
		[]byte(""),
		false,
	}

	return b
}

func RegisterWidget(name string, constructor WidgetConstructor) {
	availableWidgets[name] = constructor
}

func LoadWidgets() {
	constructor := availableWidgets["ClockWidget"]

	id := "clock"
	widget := constructor(id)
	widgets[id] = widget
}

func RunWidgets() {
	for {
		for _, w := range widgets {
			go w.CheckUpdate()
		}
	}
}
