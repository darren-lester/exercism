package react

type reactor struct {
	cells      []*Cell
	dirtyCells map[*cell]int
}

// New creates a new reactor
func New() Reactor {
	r := reactor{dirtyCells: make(map[*cell]int)}
	return &r
}

// CreateInput creates an input cell
func (r *reactor) CreateInput(value int) InputCell {
	inputCell := cell{reactor: r, value: value}
	return &inputCell
}

// CreateCompute1 creates a compute cell dependent on an input cell
func (r *reactor) CreateCompute1(input Cell, valueFn func(int) int) ComputeCell {
	return r.createCompute(func(inputValues ...int) int {
		return valueFn(inputValues[0])
	}, input)
}

// CreateCompute2 creates a compute cell dependent on two input cells
func (r *reactor) CreateCompute2(input1, input2 Cell, valueFn func(int, int) int) ComputeCell {
	return r.createCompute(func(inputValues ...int) int {
		return valueFn(inputValues[0], inputValues[1])
	}, input1, input2)
}

func (r *reactor) createCompute(valueFn func(...int) int, inputCells ...Cell) ComputeCell {
	computeCell := cell{reactor: r, valueFn: func() int {
		inputValues := make([]int, len(inputCells))
		for i, inputCell := range inputCells {
			inputValues[i] = inputCell.Value()
		}
		return valueFn(inputValues...)
	}, callbacks: make(map[*func(int)]bool)}

	for _, inputCell := range inputCells {
		in := inputCell.(*cell)
		in.dependants = append(in.dependants, &computeCell)
	}

	computeCell.SetValue(computeCell.valueFn())
	return &computeCell
}

func (r *reactor) update(updatedCell *cell, prevValue int, runCallbacks bool) {
	if _, ok := r.dirtyCells[updatedCell]; !ok {
		if updatedCell.value != prevValue {
			r.dirtyCells[updatedCell] = prevValue
		}
	}

	for _, dependant := range updatedCell.dependants {
		dependant.setValue(dependant.valueFn(), false)
	}

	if runCallbacks {
		r.runCallbacks()
	}
}

func (r *reactor) runCallbacks() {
	for cell, prevValue := range r.dirtyCells {
		value := cell.Value()
		if value != prevValue {
			for callback, _ := range cell.callbacks {
				(*callback)(value)
			}
		}
		delete(r.dirtyCells, cell)
	}
}
