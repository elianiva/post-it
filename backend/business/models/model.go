package models

type Model interface {
	GetSchema() (string, error)
	Drop() string
}
