package _1_singleton

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetSingleton(t *testing.T) {
	assert.Equal(t, GetSingleton(), GetSingleton())
}

func BenchmarkGetInstanceParallel(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			if GetSingleton() != GetSingleton() {
				b.Errorf("test fail")
			}
		}
	})
}

func TestGetLazySingleton(t *testing.T) {
	assert.Equal(t, GetLazySingleton(), GetLazySingleton())
}

func BenchmarkGetLazySingletonParallel(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			if GetLazySingleton() != GetSingleton() {
				b.Errorf("test fail")
			}
		}
	})
}
