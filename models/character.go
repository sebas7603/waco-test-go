package models

import (
	"time"
)

type Character struct {
	ID       int               `json:"id"`
	Name     string            `json:"name"`
	Status   string            `json:"status"`
	Species  string            `json:"species"`
	Type     string            `json:"type"`
	Gender   string            `json:"gender"`
	Origin   CharacterOrigin   `json:"origin"`
	Location CharacterLocation `json:"location"`
	Image    string            `json:"image"`
	Episode  []string          `json:"episode"`
	URL      string            `json:"url"`
	Created  time.Time         `json:"created"`
}

type CharacterOrigin struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type CharacterLocation struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}
