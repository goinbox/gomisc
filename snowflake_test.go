package gomisc

import (
	"fmt"
	"sync"
	"testing"
)

func TestSnowflakeGenerateID(t *testing.T) {
	m := make(map[int64]bool)
	s := NewSnowflake(10)
	for i := 0; i < 1000000; i++ {
		id := s.GenerateID()
		t.Log(id, fmt.Sprintf("%064b", id))

		_, ok := m[id]
		if ok {
			t.Fatalf("duplicated id %d", id)
		}
		m[id] = true
	}
}

func BenchmarkSnowflakeGenerateID(b *testing.B) {
	s := NewSnowflake(10)
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			s.GenerateID()
		}
	})
}

func TestDuplicateSnowflakeGenerateID(t *testing.T) {
	m := new(sync.Map)
	s := NewSnowflake(10)
	wg := new(sync.WaitGroup)
	for i := 0; i < 1000000; i++ {
		go func() {
			defer wg.Done()
			wg.Add(1)

			for i := 0; i < 100; i++ {
				id := s.GenerateID()
				_, ok := m.Load(id)
				if ok {
					t.Errorf("duplicated id %d", id)
				} else {
					m.Store(id, true)
				}
			}
		}()
	}
	wg.Wait()
}
