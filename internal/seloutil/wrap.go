package seloutil

import "github.com/kon3gor/selo"

type ConfigFactory[T any, C any] func(c C) T

func Wrap[T any, C any](c C, f ConfigFactory[T, C]) selo.Factory[T] {
	return func() T {
		return f(c)
	}
}
