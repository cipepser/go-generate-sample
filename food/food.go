package food

//go:generate stringer -type=Fruit
type Fruit int

const (
	Apple Fruit = iota
	Orange
)
