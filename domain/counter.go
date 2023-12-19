package domain

func (s *Store) GetCounter() int {
	counter := 0
	s.sc.Mt.Lock()
	s.sc.Count += 1
	counter = s.sc.Count
	s.sc.Mt.Unlock()
	return counter
}
