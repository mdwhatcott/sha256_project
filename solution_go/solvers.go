package main

func (p *Problems) SolveProblem1(s *Solutions) {
	for _, problem := range p.Problem1 {
		s.Problem1 = append(s.Problem1, problem[0]+problem[1])
	}
}
func (p *Problems) SolveProblem2(s *Solutions) {
	for _, problem := range p.Problem2 {
		s.Problem2 = append(s.Problem2, rightRotate32(problem[0], int(problem[1])))
	}
}
func (p *Problems) SolveProblem3(s *Solutions) {
	s.Problem3 = littleSigma0(p.Problem3)
}
func (p *Problems) SolveProblem4(s *Solutions) {
	s.Problem4 = littleSigma1(p.Problem4)
}
func (p *Problems) SolveProblem5(s *Solutions) {
	s.Problem5 = messageSchedule(p.Problem5)
}

func (p *Problems) SolveProblem6(s *Solutions)  {}
func (p *Problems) SolveProblem7(s *Solutions)  {}
func (p *Problems) SolveProblem8(s *Solutions)  {}
func (p *Problems) SolveProblem9(s *Solutions)  {}
func (p *Problems) SolveProblem10(s *Solutions) {}
func (p *Problems) SolveProblem11(s *Solutions) {}
func (p *Problems) SolveProblem12(s *Solutions) {}
func (p *Problems) SolveProblem13(s *Solutions) {}
func (p *Problems) SolveProblem14(s *Solutions) {}
func (p *Problems) SolveProblem15(s *Solutions) {}
func (p *Problems) SolveProblem16(s *Solutions) {}
