package rules

type SoloRuleset struct {
	StandardRuleset
}

func (r *SoloRuleset) IsGameOver(b *BoardState) (bool, error) {
	for i := 0; i < len(b.Snakes); i++ {
		if b.Snakes[i].EliminatedCause == NotEliminated {
			return false, nil
		}
	}
	return true, nil
}
