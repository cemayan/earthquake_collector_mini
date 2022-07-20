package service

import "github.com/cemayan/earthquake_collector_mini/src/types/earthquake"

type XMLService interface {
	Fetch(url string) []byte
	Parse(data []byte) []earthquake.Earthquake
}
