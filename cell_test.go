package gol

import "testing"

func TestCell_IsUnderpopulated(t *testing.T) {
	cell := Cell{
		State: Life,
	}
	neighbors := []Cell{
		{State: Life},
	}

	if !cell.IsUnderpopulated(neighbors) {
		t.Errorf("Expected cell to be underpopulated")
	}
}

func TestCell_IsNotUnderpopulated(t *testing.T) {
	cell := Cell{
		State: Life,
	}
	neighbors := []Cell{
		{State: Life},
		{State: Life},
	}

	if cell.IsUnderpopulated(neighbors) {
		t.Errorf("Expected cell not to be underpopulated")
	}
}

func TestCell_EvolveToNextGeneration_TwoNeighbors(t *testing.T) {
	cell := Cell{
		State: Life,
	}
	neighbors := []Cell{
		{State: Life},
		{State: Life},
	}

	if !cell.EvolveToNextGeneration(neighbors) {
		t.Errorf("Expected cell to evolve to next generation")
	}
}

func TestCell_EvolveToNextGeneration_ThreeNeighbors(t *testing.T) {
	cell := Cell{
		State: Life,
	}
	neighbors := []Cell{
		{State: Life},
		{State: Life},
		{State: Life},
	}

	if !cell.EvolveToNextGeneration(neighbors) {
		t.Errorf("Expected cell not to evolve to next generation")
	}
}

func TestCell_IsOverpopulated(t *testing.T) {
	cell := Cell{
		State: Life,
	}
	neighbors := []Cell{
		{State: Life},
		{State: Life},
		{State: Life},
		{State: Life},
	}

	if !cell.IsOverpopulated(neighbors) {
		t.Errorf("Expected cell to be overpopulated")
	}
}

func TestCell_IsNotOverpopulated(t *testing.T) {
	cell := Cell{
		State: Life,
	}
	neighbors := []Cell{
		{State: Life},
		{State: Life},
		{State: Life},
	}

	if cell.IsOverpopulated(neighbors) {
		t.Errorf("Expected cell not to be overpopulated")
	}
}

func TestCell_ShouldBecomeAlive(t *testing.T) {
	cell := Cell{
		State: Dead,
	}
	neighbors := []Cell{
		{State: Life},
		{State: Life},
		{State: Life},
	}

	if !cell.ShouldBecomeAlive(neighbors) {
		t.Errorf("Expected cell to become alive")
	}
}

func TestCell_ShouldNotBecomeAlive(t *testing.T) {
	cell := Cell{
		State: Dead,
	}
	neighbors := []Cell{
		{State: Life},
		{State: Life},
	}

	if cell.ShouldBecomeAlive(neighbors) {
		t.Errorf("Expected cell not to become alive")
	}
}
