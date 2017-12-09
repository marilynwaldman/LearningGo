package src

import (

)

type StrFunc func(string)  string

func Compose(f StrFunc, g StrFunc) StrFunc {
	return func (s string) string {
		return g(f(s))
	}
}



