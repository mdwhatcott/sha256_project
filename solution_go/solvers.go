package main

func (p *Problems) SolveProblem1(s *Solutions) {
	for _, problem := range p.Problem1 {
		s.Problem1 = append(s.Problem1, add32(problem[0], problem[1]))
	}
}

func add32(a, b int) int {
	A := uint32(a)
	B := uint32(b)
	return int(A + B)
}
func (p *Problems) SolveProblem2(solutions *Solutions)  {}
func (p *Problems) SolveProblem3(solutions *Solutions)  {}
func (p *Problems) SolveProblem4(solutions *Solutions)  {}
func (p *Problems) SolveProblem5(solutions *Solutions)  {}
func (p *Problems) SolveProblem6(solutions *Solutions)  {}
func (p *Problems) SolveProblem7(solutions *Solutions)  {}
func (p *Problems) SolveProblem8(solutions *Solutions)  {}
func (p *Problems) SolveProblem9(solutions *Solutions)  {}
func (p *Problems) SolveProblem10(solutions *Solutions) {}
func (p *Problems) SolveProblem11(solutions *Solutions) {}
func (p *Problems) SolveProblem12(solutions *Solutions) {}
func (p *Problems) SolveProblem13(solutions *Solutions) {}
func (p *Problems) SolveProblem14(solutions *Solutions) {}
func (p *Problems) SolveProblem15(solutions *Solutions) {}
func (p *Problems) SolveProblem16(solutions *Solutions) {}
