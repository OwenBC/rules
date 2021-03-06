package rules

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSoloRulesetInterface(t *testing.T) {
	var _ Ruleset = (*SoloRuleset)(nil)
}

func TestSoloCreateNextBoardStateSanity(t *testing.T) {
	boardState := &BoardState{}
	r := SoloRuleset{}
	_, err := r.CreateNextBoardState(boardState, []SnakeMove{})
	require.NoError(t, err)
}

func TestSoloIsGameOver(t *testing.T) {
	tests := []struct {
		Snakes   []Snake
		Expected bool
	}{
		{[]Snake{}, true},
		{[]Snake{{}}, false},
		{[]Snake{{}, {}, {}}, false},
		{[]Snake{{EliminatedCause: EliminatedByOutOfBounds}}, true},
		{
			[]Snake{
				{EliminatedCause: EliminatedByOutOfBounds},
				{EliminatedCause: EliminatedByOutOfBounds},
				{EliminatedCause: EliminatedByOutOfBounds},
			},
			true,
		},
	}

	r := SoloRuleset{}
	for _, test := range tests {
		b := &BoardState{
			Height: 11,
			Width:  11,
			Snakes: test.Snakes,
			Food:   []Point{},
		}

		actual, err := r.IsGameOver(b)
		require.NoError(t, err)
		require.Equal(t, test.Expected, actual)
	}
}
