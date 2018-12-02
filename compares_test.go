package poker

import (
	"testing"
)

var set PokerSet

func init(){
	set = CreateDeck().ToPokerSet()
}


func TestGetBigCard(t *testing.T) {
	if set[4].CardValue != GetBigCard(set[4],set[0]).CardValue {
		t.Error("TestGetBigCard err")
	}

	if set[4].CardValue != GetBigCard(set[0],set[4]).CardValue {
		t.Error("TestGetBigCard err")
	}

}

func TestGetSmallCard(t *testing.T) {
	if set[0].CardValue != GetSmallCard(set[4],set[0]).CardValue {
		t.Error("TestGetSmallCard err")
	}

	if set[0].CardValue != GetSmallCard(set[0],set[4]).CardValue {
		t.Error("TestGetSmallCard err")
	}
}

func TestIsFirstCardValueBigger(t *testing.T) {

	if IsFirstCardValueBigger(set[0],set[4]){
		t.Error("TestIsFirstCardValueBigger1 err")
	}

	if !IsFirstCardValueBigger(set[4],set[0]){
		t.Error("TestIsFirstCardValueBigger2 err")
	}
}
