package poker

import (
	"testing"
)

var oringSet PokerSet

func init(){
	oringSet = CreateDeck().ToPokerSet()
}

func TestPokerSet_AddPokers(t *testing.T) {
	set1 := PokerSet{}
	for i := range oringSet{
		if set1.CountCards() != i{
			t.Error("addPokers长度不匹配")
		}
		set1 = set1.AddPokers(PokerSet{oringSet[i]})
	}
}


func TestPokerSet_DelPokers(t *testing.T) {
	set := PokerSet{oringSet[0],oringSet[1]}
	set,err := set.DelPokers(PokerSet{oringSet[0]})
	if err != nil{
		t.Error("TestPokerSet_DelPokers err")
	}else{
		if set.CountCards() != 1  || set[0].GetValue() != oringSet[1].GetValue(){
			t.Error("TestPokerSet_DelPokers err")
		}
	}
}

func TestPokerSet_DelPokersByIndex(t *testing.T) {
	set := PokerSet{oringSet[0],oringSet[1]}
	set,err := set.DelPokersByIndex([]int{0})
	if err != nil{
		t.Error("TestPokerSet_DelPokers err")
	}else{
		if set.CountCards() != 1  || set[0].GetValue() != oringSet[1].GetValue(){
			t.Error("TestPokerSet_DelPokers err")
		}
	}
}

func TestPokerSet_GetLength(t *testing.T) {
	set := PokerSet{oringSet[0],oringSet[1]}
	if len(set) != set.CountCards(){
		t.Error("TestPokerSet_GetLength err")
	}
}

func TestPokerSet_GetPokerByIndex(t *testing.T) {
	set := PokerSet{oringSet[0],oringSet[1]}
	card,err := set.GetPokerByIndex(1)
	if err != nil{
		t.Error("TestPokerSet_DelPokers err")
	}else{
		if card != oringSet[1]{
			t.Error("TestPokerSet_DelPokers err")
		}
	}
}

func TestPokerSet_GetPokerIndex(t *testing.T) {
	set := PokerSet{oringSet[0],oringSet[1]}
	index,err := set.GetPokerIndex(oringSet[1])
	if err != nil{
		t.Error("TestPokerSet_DelPokers err")
	}else{
		if index != 1 {
			t.Error("TestPokerSet_DelPokers err")
		}
	}
}

func TestPokerSet_GetPokerIndexs(t *testing.T) {
	set := PokerSet{oringSet[0],oringSet[1]}
	indexs,err := set.GetPokerIndexs(PokerSet{oringSet[0],oringSet[1]})
	if err != nil{
		t.Error("TestPokerSet_DelPokers err")
	}else{
		if indexs[0] != 0 || indexs[1] != 1 {
			t.Error("TestPokerSet_DelPokers err")
		}
	}
}

func TestPokerSet_GetPokersByIndexs(t *testing.T) {
	set := PokerSet{oringSet[0],oringSet[1]}
	cards,err := set.GetPokersByIndexs([]int{0,1})
	if err != nil{
		t.Error("TestPokerSet_DelPokers err")
	}else{
		if cards[0] != oringSet[0] || cards[1] != oringSet[1] {
			t.Error("TestPokerSet_DelPokers err")
		}
	}
}



func TestPokerSet_HasSamePoker(t *testing.T) {
	set1 := PokerSet{oringSet[0],oringSet[1]}
	if set1.HasSameValueCard(PokerSet{oringSet[4],oringSet[7]}){
		t.Error("TestPokerSet_HasSamePoker err1")
	}

	if !set1.HasSameValueCard(PokerSet{oringSet[0],oringSet[1]}){
		t.Error("TestPokerSet_HasSamePoker err1")
	}
}

func TestPokerSet_ReplacePoker(t *testing.T) {
	set := PokerSet{oringSet[36],oringSet[37]}
	err := set.ReplacePoker(0,oringSet[0])
	if err != nil{
		t.Error("1")
	}else{
		if set[0] != oringSet[0]{
			t.Error("2")
		}
	}
}

