package main

type SomeInterface interface {
	Dit()
	Dah()
}


type SomeType struct {}

func (s *SomeType) Dit() {}

func main() {
	var i SomeInterface
	t := SomeType{}

	i = t
}
