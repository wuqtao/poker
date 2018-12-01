package poker

//定义一副扑克牌
type PokerDeck struct {
	cards [54]PokerCard
}

//获取指定索引的扑克牌
func (deck PokerDeck) GetCard(index int) *PokerCard{
	return &deck.cards[index]
}
//计算扑克牌的数量
func (deck PokerDeck) CountCards() int{
	return len(deck.cards)
}
//获取该deck中的所有牌
func (deck PokerDeck) GetAllCards() []PokerCard{
	return deck.cards[:]
}
//将pokerdeck转换为pokerset
func (deck PokerDeck) ToPokerSet() PokerSet{
	set := NewPokerSet()
	for i,_ := range deck.cards{
		set = set.AddPokers(PokerSet{&deck.cards[i]})
	}
	return set
}

//原始的一副扑克牌
var originDeck PokerDeck

func init(){
	originDeck = createOriginDeck()
}

//创建原始扑克牌，后续只需要复制即可，不用再运算获得
func createOriginDeck() PokerDeck{
	deck := PokerDeck{}
	for i := 0;i<52;i++ {
		shang := i/4;
		yu := i%4;
		suit := ""
		switch(yu){
		case 0:
			suit = CARD_SUIT_CLUB
		case 1:
			suit = CARD_SUIT_DIAMOND
		case 2:
			suit = CARD_SUIT_HEART
		case 3:
			suit = CARD_SUIT_SPADE
		}
		pokerValue := 0
		pokerName := ""
		switch(shang){
		case 0:
			pokerValue = CARD_VALUE_THREE
			pokerName = CARD_SYMBOL_THREE
		case 1:
			pokerValue = CARD_VALUE_FOUR
			pokerName = CARD_SYMBOL_FOUR
		case 2:
			pokerValue = CARD_VALUE_FIVE
			pokerName = CARD_SYMBOL_FIVE
		case 3:
			pokerValue = CARD_VALUE_SIX
			pokerName = CARD_SYMBOL_SIX
		case 4:
			pokerValue = CARD_VALUE_SEVEN
			pokerName = CARD_SYMBOL_SEVEN
		case 5:
			pokerValue = CARD_VALUE_EIGHT
			pokerName = CARD_SYMBOL_EIGHT
		case 6:
			pokerValue = CARD_VALUE_NINE
			pokerName = CARD_SYMBOL_NINE
		case 7:
			pokerValue = CARD_VALUE_TEN
			pokerName = CARD_SYMBOL_TEN
		case 8:
			pokerValue = CARD_VALUE_JACK
			pokerName = CARD_SYMBOL_JACK
		case 9:
			pokerValue = CARD_VALUE_QUEEN
			pokerName = CARD_SYMBOL_QUEEN
		case 10:
			pokerValue = CARD_VALUE_KING
			pokerName = CARD_SYMBOL_KING
		case 11:
			pokerValue = CARD_VALUE_ACE
			pokerName = CARD_SYMBOL_ACE
		case 12:
			pokerValue = CARD_VALUE_TWO
			pokerName = CARD_SYMBOL_TWO
		}
		deck.cards[i] = PokerCard{
			pokerValue,
			suit,
			pokerName,
		}
	}
	deck.cards[52] = PokerCard{
		CARD_VALUE_BLACK_JOKER,
		CARD_SUIT_JOKER,
		CARD_SYMBOL_BLACK_JOKER,
	}

	deck.cards[53] = PokerCard{
		CARD_VALUE_RED_JOKER,
		CARD_SUIT_JOKER,
		CARD_SYMBOL_RED_JOKER,
	}
	return deck
}

//每个游戏桌子都有单独的扑克牌，防止洗牌等操作冲突
func CreateDeck() PokerDeck{
	copyDeck := originDeck
	return copyDeck
}
//根据输入的扑克副数，生成扑克集
func CreatePokerSetWithDeckNum(deckNum int) PokerSet{
	set := NewPokerSet()
	for i:=0;i<deckNum;i++{
		set = set.AddPokers(CreateDeck().ToPokerSet())
	}
	return set
}
