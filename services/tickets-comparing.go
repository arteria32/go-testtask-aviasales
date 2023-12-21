package ticketsfinder

import (
	"errors"
	"log"
	. "main/models"
)

func updateByLowestPrice(curItinary PricedItineraries, newItinary PricedItineraries) PricedItineraries {

	var curPrice float64
	for _, v := range curItinary.Pricing.ServiceCharges {
		if v.ChargeType == "TotalAmount" {
			curPrice = v.Value
		}
	}
	var newPrice float64
	for _, v := range newItinary.Pricing.ServiceCharges {
		if v.ChargeType == "TotalAmount" {
			newPrice = v.Value
		}
	}
	if newPrice < curPrice || curPrice == 0 {
		return newItinary
	}
	return curItinary
}
func updateByHighestPrice(curItinary PricedItineraries, newItinary PricedItineraries) PricedItineraries {

	var curPrice float64
	for _, v := range curItinary.Pricing.ServiceCharges {
		if v.ChargeType == "TotalAmount" {
			curPrice = v.Value
		}
	}
	var newPrice float64
	for _, v := range newItinary.Pricing.ServiceCharges {
		if v.ChargeType == "TotalAmount" {
			newPrice = v.Value
		}
	}
	log.Println(newPrice, curPrice)
	if newPrice > curPrice || curPrice == 0 {
		return newItinary
	}
	return curItinary
}

type ComparingFunc func(curItinary PricedItineraries, newItinary PricedItineraries) PricedItineraries

func factoryComparingFactory(typeComp string) (ComparingFunc, error) {
	switch typeComp {
	case "lowestPrice":
		return updateByLowestPrice, nil
	case "highestPrice":
		return updateByHighestPrice, nil
	default:
		return nil, errors.New("Undefined Comparing Type")
	}

}
func GetTicketsInfoByIntervalAndType(from, to string, typeSearch string) (PricedItineraries, error) {
	filteredByIntervalData, err := GetTicketsInfoByInterval(from, to)
	var findedTicket PricedItineraries
	if err != nil {
		return findedTicket, err
	}
	curComparingFunc, funcError := factoryComparingFactory(typeSearch)
	if err != nil {
		return findedTicket, funcError
	}
	for _, v := range filteredByIntervalData {
		findedTicket = curComparingFunc(findedTicket, v)
	}

	return findedTicket, nil
}
