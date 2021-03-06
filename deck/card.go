package deck

import (
	"strconv"
)

type CardType struct {
	name   string
	symbol rune
}

type SerializableCardType struct {
	Name   string
	Symbol rune
}

type Card struct {
	value     int
	cardType  *CardType
	IsVisible bool
}

type SerializableCard struct {
	Value     int
	CardType  *SerializableCardType
	IsVisible bool
}

func (c *Card) GetBlackjackValue() int {
	if c.value > 10 {
		return 10
	}
	return c.value
}

func (c *Card) GetSymbol() string {
	return string(c.cardType.symbol)
}

func (c *Card) GetDisplayingValue() string {
	switch c.value {
	case 1:
		return "A"
	case 12:
		return "J"
	case 13:
		return "Q"
	case 14:
		return "K"
	default:
		return strconv.Itoa(c.value)
	}
}

func (c *Card) GetSerializable() *SerializableCard {
	cardTypeJson := SerializableCardType{
		Name:   c.cardType.name,
		Symbol: c.cardType.symbol,
	}
	serializableCard := SerializableCard{
		Value:     c.value,
		CardType:  &cardTypeJson,
		IsVisible: c.IsVisible,
	}
	return &serializableCard
}

func (serializableCard *SerializableCard) DeserializeCard() *Card {
	return &Card{
		value: serializableCard.Value,
		cardType: &CardType{
			name:   serializableCard.CardType.Name,
			symbol: serializableCard.CardType.Symbol,
		},
		IsVisible: serializableCard.IsVisible,
	}
}

//help with testing
func (c *Card) SetCard(value int, name string, symbol rune) {
	c.value = value
	c.cardType = &CardType{name, symbol}
}
