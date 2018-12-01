package poker

import "testing"

func TestBubbleSortCardsMax2Min(t *testing.T) {
	currSet := PokerSet{set[0],set[10],set[5]}
	currSet.SortDesc()

	if currSet[0].GetValue() != set[10].GetValue() ||
		currSet[1].GetValue() != set[5].GetValue() ||
		currSet[2].GetValue() != set[0].GetValue(){
		t.Error("TestBubbleSortCardsMax2Min err")
	}
}

func TestBubbleSortCardsMin2Max(t *testing.T) {
	currSet := PokerSet{set[0],set[10],set[5]}
	currSet.SortAsc()

	if currSet[0].GetValue() != set[0].GetValue() ||
		currSet[1].GetValue() != set[5].GetValue() ||
		currSet[2].GetValue() != set[10].GetValue(){
		t.Error("TestBubbleSortCardsMin2Max err")
	}
}
