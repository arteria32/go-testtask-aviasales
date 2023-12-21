package ticketsfinder

import (
	"encoding/xml"
	"io"
	"log"
	. "main/models"
	"os"
	"strconv"
	"time"
)

var LAYOUT_TIME = "2006-01-02T1504"

func getFullInfoFromRequest() (*AirFareSearchResponse, error) {
	reserchResponse := new(AirFareSearchResponse)
	// Working Directory
	wd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	xmlFile, err := os.Open(wd + "/mock/RS_Via-3.xml")

	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	defer xmlFile.Close()
	byteValue, err := io.ReadAll(xmlFile)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	err = xml.Unmarshal(byteValue, &reserchResponse)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return reserchResponse, nil
}

func GetFullTicketsInfo() ([]PricedItineraries, error) {
	response, err := getFullInfoFromRequest()
	if err != nil {
		return nil, err
	}
	var result []PricedItineraries
	for _, v := range response.PricedItineraries.Flights {
		onwardFlights, returnFlights, priceInfo := convertXMLPricedItineraries(v)
		newPricedItineraries := PricedItineraries{
			OnwardPricedItinerary: onwardFlights, ReturnPricedItinerary: returnFlights, Pricing: priceInfo}
		result = append(result, newPricedItineraries)
	}
	return result, nil
}

func convertXMLPricedItineraries(inpRes PricedItinerariesXML) ([]Flight, []Flight, Pricing) {
	onwardFlights := convertToFligts(inpRes.ReturnPricedItinerary.Flights.Flight)
	returnFlights := convertToFligts(inpRes.OnwardPricedItinerary.Flights.Flight)
	priceInfo := convertToPricing(inpRes.Pricing)
	return onwardFlights, returnFlights, priceInfo
}

/*.Формируем информацию по сервисным сборам  */
func convertToPricing(inpPricing PricingXML) Pricing {
	var tempServices []ServiceCharges
	for _, v := range inpPricing.ServiceCharges {
		newValue, _ := strconv.ParseFloat(v.Text, 64)
		newServiceCharge := ServiceCharges{
			Value: newValue, Type: v.Type, ChargeType: v.ChargeType}
		tempServices = append(tempServices, newServiceCharge)
	}
	priceInfo := Pricing{Currency: inpPricing.Currency, ServiceCharges: tempServices}
	return priceInfo
}

/*.Формируем информацию по билетам туда-обратно  */
func convertToFligts(flightInfo []FlightXML) []Flight {
	var flights []Flight
	for _, v := range flightInfo {
		carierInfo := Carrier{Text: v.Carrier.Text, ID: v.Carrier.ID}
		newFlight := Flight{
			Carrier:            carierInfo,
			FlightNumber:       v.FlightNumber,
			Source:             v.Source,
			Destination:        v.Destination,
			DepartureTimeStamp: v.DepartureTimeStamp,
			ArrivalTimeStamp:   v.ArrivalTimeStamp,
			Class:              v.Class,
			NumberOfStops:      v.NumberOfStops,
			FareBasis:          v.FareBasis,
			WarningText:        v.WarningText,
			TicketType:         v.TicketType,
		}
		flights = append(flights, newFlight)
	}
	return flights
}

func GetTicketsInfoByInterval(from, to string) ([]PricedItineraries, error) {
	res, err := GetFullTicketsInfo()
	if err != nil {
		return nil, err
	}
	fromDate, fromErr := time.Parse(LAYOUT_TIME, from)
	toDate, toErr := time.Parse(LAYOUT_TIME, to)
	if fromErr != nil {
		return nil, fromErr
	}
	if toErr != nil {
		return nil, toErr
	}
	var filteredRes []PricedItineraries

	for _, v := range res {
		startFlightV, _ := time.Parse(LAYOUT_TIME, v.OnwardPricedItinerary[0].DepartureTimeStamp)
		endFlightV, _ := time.Parse(LAYOUT_TIME, v.ReturnPricedItinerary[len(v.ReturnPricedItinerary)-1].ArrivalTimeStamp)
		if fromDate.Before(startFlightV) && toDate.After(endFlightV) {
			filteredRes = append(filteredRes, v)
		}
	}
	return filteredRes, nil
}
