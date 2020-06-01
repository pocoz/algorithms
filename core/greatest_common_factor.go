package core

func Find(m, n int) int {
	for {
		if m < n {
			mOld := m
			m = n
			n = mOld
		}
		r := m % n
		if r == 0 {
			return n
		}
		m = n
		n = r
	}
}
