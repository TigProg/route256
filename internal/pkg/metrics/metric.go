package metrics

import (
	"strconv"
	"sync"
)

func NewMetric(name string) *Metric {
	return &Metric{
		Name:  name,
		value: 0,
		mu:    sync.RWMutex{},
	}
}

type Metric struct {
	Name  string
	value uint32
	mu    sync.RWMutex
}

func (m *Metric) Increment() {
	m.mu.Lock()
	defer m.mu.Unlock()

	m.value += 1
}

func (m *Metric) String() string {
	m.mu.RLock()
	defer m.mu.RUnlock()

	return strconv.FormatInt(int64(m.value), 10)
}
