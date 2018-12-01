package poker

/**
	定义扑克牌花色、显示牌型、值以及扑克牌
*/
//定义扑克牌值
const (
	CARD_VALUE_THREE = iota
	CARD_VALUE_FOUR
	CARD_VALUE_FIVE
	CARD_VALUE_SIX
	CARD_VALUE_SEVEN
	CARD_VALUE_EIGHT
	CARD_VALUE_NINE
	CARD_VALUE_TEN
	CARD_VALUE_JACK
	CARD_VALUE_QUEEN
	CARD_VALUE_KING
	CARD_VALUE_ACE
	CARD_VALUE_TWO
	CARD_VALUE_BLACK_JOKER
	CARD_VALUE_RED_JOKER
)
//定义扑克牌符号
const (
	CARD_SYMBOL_THREE= "3"
	CARD_SYMBOL_FOUR = "4"
	CARD_SYMBOL_FIVE = "5"
	CARD_SYMBOL_SIX = "6"
	CARD_SYMBOL_SEVEN = "7"
	CARD_SYMBOL_EIGHT = "8"
	CARD_SYMBOL_NINE = "9"
	CARD_SYMBOL_TEN = "10"
	CARD_SYMBOL_JACK = "J"
	CARD_SYMBOL_QUEEN = "Q"
	CARD_SYMBOL_KING = "K"
	CARD_SYMBOL_ACE = "A"
	CARD_SYMBOL_TWO = "2"
	CARD_SYMBOL_BLACK_JOKER = "BlackJoker"
	CARD_SYMBOL_RED_JOKER = "RedJoker"
)
//定义扑克牌花色
const(
	CARD_SUIT_DIAMOND = "Diamond" //方片
	CARD_SUIT_HEART   = "Heart"   //红桃
	CARD_SUIT_SPADE   = "Spade"   //黑桃
	CARD_SUIT_CLUB    = "Club"    //梅花
	CARD_SUIT_JOKER   = "Joker"   //大小王无花色
)
//定义扑克牌
type PokerCard struct {
	cardValue int			//card值用于排序比较
	cardSuit string			//card花色
	cardName string			//card显示的字符
}
//获取扑克牌的值
func (card PokerCard)GetValue() int{
	return card.cardValue
}
//获取扑克牌的花色
func (card PokerCard)GetSuit() string{
	return card.cardSuit
}
//获取扑克牌的名字
func (card PokerCard)GetCardName() string{
	return card.cardName
}








