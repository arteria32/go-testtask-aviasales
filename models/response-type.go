package models

import "encoding/xml"

type AirFareSearchResponse struct {
	XMLName           xml.Name `xml:"AirFareSearchResponse"`
	Text              string   `xml:",chardata"`
	RequestTime       string   `xml:"RequestTime,attr"`
	ResponseTime      string   `xml:"ResponseTime,attr"`
	RequestId         string   `xml:"RequestId"`
	PricedItineraries struct {
		Text    string                 `xml:",chardata"`
		Flights []PricedItinerariesXML `xml:"Flights"`
	} `xml:"PricedItineraries"`
}
type PricedItinerariesXML struct {
	OnwardPricedItinerary struct {
		Text    string `xml:",chardata"`
		Flights struct {
			Flight []FlightXML `xml:"Flight"`
		} `xml:"Flights"`
	} `xml:"OnwardPricedItinerary"`
	ReturnPricedItinerary struct {
		Text    string `xml:",chardata"`
		Flights struct {
			Flight []FlightXML `xml:"Flight"`
		} `xml:"Flights"`
	} `xml:"ReturnPricedItinerary"`
	Pricing PricingXML `xml:"Pricing"`
}
type PricingXML struct {
	Text           string `xml:",chardata"`
	Currency       string `xml:"currency,attr"`
	ServiceCharges []struct {
		Text       string `xml:",chardata"`
		Type       string `xml:"type,attr"`
		ChargeType string `xml:"ChargeType,attr"`
	} `xml:"ServiceCharges"`
}

type FlightXML struct {
	Carrier struct {
		Text string `xml:",chardata"`
		ID   string `xml:"id,attr"`
	} `xml:"Carrier"`
	FlightNumber       string `xml:"FlightNumber"`
	Source             string `xml:"Source"`
	Destination        string `xml:"Destination"`
	DepartureTimeStamp string `xml:"DepartureTimeStamp"`
	ArrivalTimeStamp   string `xml:"ArrivalTimeStamp"`
	Class              string `xml:"Class"`
	NumberOfStops      string `xml:"NumberOfStops"`
	FareBasis          string `xml:"FareBasis"`
	WarningText        string `xml:"WarningText"`
	TicketType         string `xml:"TicketType"`
}
