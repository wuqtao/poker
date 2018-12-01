package poker

import (
	"errors"
)
/**
	定义扑克集合
 */
type PokerSet []*PokerCard
//创建新的扑克集
func NewPokerSet() PokerSet{
	return PokerSet{}
}
//向扑克集中添加扑克
func (set PokerSet)AddPokers(cards PokerSet) PokerSet{
	for _,card := range cards{
		set = append(set,card)
	}
	//append元素可能需要重新分配空间，导致原切片引用发生变化，此处需要返回新的切片引用
	return set
}
//检查给定的索引是否存在
func (set PokerSet)checkIndex(indexs []int) error{
	setLen := set.CountCards()
	for _,index := range indexs{
		if index >= setLen{
			return errors.New("给定的索引超过扑克集的长度")
		}
	}
	return nil
}
//从扑克集中删除制定索引的扑克,删除后扑克集元素变少
func (set PokerSet)DelPokersByIndex(indexs []int) (PokerSet,error){

	err := set.checkIndex(indexs)
	if err != nil{
		return nil,err
	}

	for _,index := range indexs{
		set[index] = nil
	}
	newSet := NewPokerSet()

	for _,card := range set{
		if card != nil{
			newSet = append(newSet,card)
		}
	}
	return newSet,nil
}
//从扑克集中删除指定的子集
func (set PokerSet)DelPokers(pokers PokerSet) (PokerSet,error){
	indexs,err := set.GetPokerIndexs(pokers)
	if err != nil{
		return nil,err
	}
	newSet,err := set.DelPokersByIndex(indexs)
	return newSet,nil
}
//根据给定索引，从扑克集中获取指定扑克
func (set PokerSet)GetPokerByIndex(index int) (*PokerCard,error){
	err := set.checkIndex([]int{index})
	if err != nil{
		return nil,err
	}
	return set[index],nil
}
//根据给定索引，从扑克集中获取子扑克集
func (set PokerSet)GetPokersByIndexs(indexs []int) (PokerSet,error){

	err := set.checkIndex(indexs)
	if err != nil{
		return nil,err
	}

	newSet := NewPokerSet()

	for _,index := range indexs{
		newSet = append(newSet,set[index])
	}
	return newSet,nil
}
//将指定索引的扑克牌替换
func (set PokerSet)ReplacePoker(index int,card *PokerCard) error{
	err := set.checkIndex([]int{index})
	if err != nil{
		return err
	}
	set[index] = card
	return nil
}
//将扑克集中的各个扑克牌用于某个任务
func (set PokerSet) DoOnEachPokerCard(do func(index int,card *PokerCard)){
	for i,card := range set{
		do(i,card)
	}
}
//获取指定扑克牌在扑克集中的index
func (set PokerSet)GetPokerIndex(card *PokerCard) (int,error){
	for i,c := range set{
		if c == card{
			return i,nil
		}
	}
	return -1,errors.New("查找的扑克牌不在该扑克集中")
}
//获取指定扑克集中各扑克牌在扑克集中的index
func (set PokerSet)GetPokerIndexs(pokers PokerSet) ([]int,error){
	indexs := []int{}
	for _,card := range pokers{
		index,err := set.GetPokerIndex(card)
		if err != nil{
			return nil,err
		}
		indexs = append(indexs,index)
	}
	return indexs,nil
}
//计数扑克集中牌的数量
func (set PokerSet) CountCards() int{
	return len(set)
}
//对扑克集中的牌，根据value大小从小到大排序
func (set PokerSet)SortAsc(){
	BubbleSortCardsMin2Max(set, IsFirstCardValueBigger)
}
//对扑克集中的牌，根据value大小从大到小排序
func (set PokerSet)SortDesc(){
	BubbleSortCardsMax2Min(set, IsFirstCardValueBigger)
}
//检测是否有相同值的扑克牌
func (set PokerSet) HasSameValueCard(s PokerSet) bool{
	for _,card1 := range set{
		for _,card2 := range s{
			if card1.cardValue == card2.cardValue{
				return true
			}
		}
	}
	return false
}
//分析一组牌中，各值牌的数量,返回map[cardValue]num
func (set PokerSet) AnalyzeEachCardValueNum() map[int]int{
	cardMap := make(map[int]int)

	if len(set) == 0 {
		return nil
	}

	for _,card := range set {
		_,ok := cardMap[card.GetValue()]
		if ok {
			cardMap[card.GetValue()]++
		}else{
			cardMap[card.GetValue()] = 1
		}
	}

	return cardMap
}
//分析一组牌中，各花色牌的数量,返回map[CardSuit]num
func (set PokerSet) AnalyzeEachCardSuitNum() map[string]int{
	cardMap := make(map[string]int)

	if len(set) == 0 {
		return nil
	}

	for _,card := range set {
		_,ok := cardMap[card.GetSuit()]
		if ok {
			cardMap[card.GetSuit()]++
		}else{
			cardMap[card.GetSuit()] = 1
		}
	}

	return cardMap
}

//判断一组牌中，不同数字的数量是否相同
func (set PokerSet) IsUnsameCardNumSame() bool{
	numMap := make(map[int]int)
	for _,v := range set{
		_,ok := numMap[v.GetValue()]
		if ok {
			numMap[v.GetValue()]++
		}else{
			numMap[v.GetValue()] = 1
		}
	}
	temp := 0
	index := 1
	for _,v:= range numMap{
		if index == 1{
			temp = v
			index++
		}else{
			if temp != v{
				return false
			}
		}
	}
	return true
}

//判断一组牌中，给定索引的牌是否一样大小
func (set PokerSet)IsAllCardSame(cardIndexs []int) bool{
	temp := -1
	for i,v:= range cardIndexs{
		if i == 0{
			temp = set[v].GetValue()
		}else{
			if temp != set[v].GetValue(){
				return false
			}
		}
	}
	return true
}
