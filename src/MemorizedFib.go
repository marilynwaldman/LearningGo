package src

func Memoize(mf Memoized) Memoized {
	cache := make(map[int]int)
	return func(key int) int {
		if val, found := cache[key]; found {
			return val
		}
		temp := mf(key)
		cache[key] = temp
		return temp
	}
}

type Memoized func(int) int

var memFib Memoized

func  MemorizedFib(n int) int {
	n += 1
	memFib = Memoize(func(n int) int{
		if n == 0 || n == 1 {
			return n
		}
		return memFib(n-1) + memFib(n-2)
	})
	return memFib(n)
}

