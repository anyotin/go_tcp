package Entities

import "time"

type Request struct {
	ContentType string        `json:"content_type"`
	RequestDate time.Time     `json:"request_date"`
	StatusCode  StatusCode    `json:"status_code"`
	Body		[]RequestBody `json:"body"`
}
