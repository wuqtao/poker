package poker

import (
	"testing"
)
//检查不同的扑克副是否冲突，一副牌里面是否有扑克牌冲突
func Test_Deck(t *testing.T){
	dec1 := CreateDeck()
	dec2 := CreateDeck()

	tempCard := dec1.GetCard(0)
	for i,_ := range dec1.ToPokerSet(){
		if i != 0{
			if dec1.GetCard(i).GetValue() == tempCard.GetValue() &&
				dec1.GetCard(i).GetCardName() == tempCard.GetCardName() &&
				dec1.GetCard(i).GetSuit() == tempCard.GetSuit(){
					t.Error("同一副牌中有重复的牌")
			}
		}
		if dec1.GetCard(i) == dec2.GetCard(i){
			t.Error("不同扑克牌地址相同")
		}

		if dec1.GetCard(i).GetValue() != dec2.GetCard(i).GetValue() ||
			dec1.GetCard(i).GetCardName() != dec2.GetCard(i).GetCardName() ||
			dec1.GetCard(i).GetSuit() != dec2.GetCard(i).GetSuit(){
				t.Error("同一索引扑克牌值不一致")
		}
	}
}

func TestPokerDeck_ToPokerSet(t *testing.T) {
	set := CreateDeck().ToPokerSet()
	if set[0].cardValue != CARD_VALUE_THREE || set[4].cardValue != CARD_VALUE_FOUR{
		t.Error("TestPokerDeck_ToPokerSet err")
	}
}

func TestPokerDeck_CountCards(t *testing.T) {
	if CreateDeck().CountCards() != 54{
		t.Error("TestPokerDeck_CountCards err")
	}
}
