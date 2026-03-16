package internal

import "time"

type User struct {
	StreetID      string `json:"streetId"`
	Nick          string `json:"nick"`
	Tier          string `json:"tier"`
	Role          string `json:"role"`
	Level         string `json:"level"`
	PlanSlots     int    `json:"planSlots"`
	PlanUsed      int    `json:"planUsed"`
	Penalized     bool   `json:"penalized"`
	PenaltyReason string `json:"penaltyReason,omitempty"`
}

type StudentPublic struct {
	StreetID string `json:"streetId"`
	Nick     string `json:"nick"`
	Tier     string `json:"tier"`
	Role     string `json:"role"`
	Level    string `json:"level"`
}

type GrowthRoute struct {
	ID        string  `json:"id"`
	StudentID string  `json:"studentId"`
	Character string  `json:"character"`
	Question  string  `json:"question"`
	VideoID   string  `json:"videoId"`
	Tier      string  `json:"tier"`
	Replies   []Reply `json:"replies"`
}

type Reply struct {
	ID        string    `json:"id"`
	Author    string    `json:"author"` // admin or student
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"createdAt"`
}

type Banner struct {
	ID    string `json:"id"`
	Title string `json:"title"`
	Image string `json:"image"`
	Link  string `json:"link"`
}

type Participant struct {
	StreetID string `json:"streetId"`
	Nick     string `json:"nick"`
	Tier     string `json:"tier"`
}

type Event struct {
	ID           string        `json:"id"`
	Title        string        `json:"title"`
	Announce     string        `json:"announce"`
	Groups       []string      `json:"groups,omitempty"`
	Awards       []string      `json:"awards,omitempty"`
	Participants []Participant `json:"participants,omitempty"`
}

type Stats struct {
	Students int `json:"students"`
	Routes   int `json:"routes"`
	Events   int `json:"events"`
	Banners  int `json:"banners"`
}
