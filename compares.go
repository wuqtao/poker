package poker

//常规比较，按照值大小比较,如果第一个牌比第二个牌的值大，返回true
func IsFirstCardValueBigger(card1 *PokerCard,card2 *PokerCard) bool{
	if card1.GetValue() > card2.GetValue(){
		return true
	}else{
		return false
	}
}


