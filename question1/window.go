package main

type NumberWindow struct {
	size   int
	values []int
}

func NewNumberWindow(size int) *NumberWindow {
	return &NumberWindow{size: size}
}

func (w *NumberWindow) Update(newValues []int) {
	for _, val := range newValues {
		w.values = append(w.values, val)
		if len(w.values) > w.size {
			w.values = w.values[1:]
		}
	}
}

func (w *NumberWindow) CurrentState() []int {
	// Return a copy to avoid mutation
	state := make([]int, len(w.values))
	copy(state, w.values)
	return state
}

func (w *NumberWindow) Average() float64 {
	if len(w.values) == 0 {
		return 0
	}
	sum := 0
	for _, val := range w.values {
		sum += val
	}
	return float64(sum) / float64(len(w.values))
}
