package internal

import "testing"

func TestName(t *testing.T) {
	country, err := NewCountry()
	if err == nil {
		print(err)
	}

	vertical, err := NewVertical()
	if err == nil {
		print(err)
	}
	dosomething(country, vertical)
}

func TestName1(t *testing.T) {
	ResolveLeftTwo(NewCountry1(), NewVertical1(), func(country int, vertical int) {
		dosomething(country, vertical)
	})
	ResolveRightTwo(NewCountry1(), NewVertical1(), func(country error, vertical error) {
		print(country, vertical)
	})
}

func NewCountry() (int, error) {
	return 1, nil
}

func NewVertical() (int, error) {
	return 1, nil
}

func dosomething(country int, vertical int) {

}

func NewVertical1() Either[int, error] {
	return NewLeft[int, error](2)
}

func NewCountry1() Either[int, error] {
	return NewLeft[int, error](2)
}
