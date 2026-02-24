package random

type Random struct {
	current int
}

func New(seed int) Random {
	return Random{current: seed}
}

func (r *Random) Int() int {
	result := r.current

	q := r.current * r.current
	r.current = (q / 10) % 100

	return result
}

func (r *Random) GetCircle() []int {
	result := []int{r.Int()}

	for {
		for _, v := range result {
			if v == r.current {
				return result
			}
		}
		result = append(result, r.Int())
	}
}
