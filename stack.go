package quicksort

type Stack struct{
	arr []int
	size int
}

func (s *Stack) Push(a int){
	s.arr = append(s.arr, a)
	s.size ++
}

func (s *Stack) Pop()int{
	temp := s.arr[s.size-1]
	s.size --
	s.arr = s.arr[:s.size]
	return temp
}

func (s Stack) Top()int{
	return s.arr[s.size-1]
}

func (s Stack) IsEmpty()bool{
	return s.size == 0
}

