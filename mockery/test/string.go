package test

type Stringer interface {
	String() string
}

type SendFunc func(data string) (int, error)
