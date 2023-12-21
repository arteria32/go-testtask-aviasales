package models

type Flight struct {
	Carrier            Carrier `xml:"Carrier"`
	FlightNumber       string  `xml:"FlightNumber"`
	Source             string  `xml:"Source"`
	Destination        string  `xml:"Destination"`
	DepartureTimeStamp string  `xml:"DepartureTimeStamp"`
	ArrivalTimeStamp   string  `xml:"ArrivalTimeStamp"`
	Class              string  `xml:"Class"`
	NumberOfStops      string  `xml:"NumberOfStops"`
	FareBasis          string  `xml:"FareBasis"`
	WarningText        string  `xml:"WarningText"`
	TicketType         string  `xml:"TicketType"`
}
type Carrier struct {
	Text string `xml:",chardata"`
	ID   string `xml:"id,attr"`
}
type ServiceCharges struct {
	Value      float64
	Type       string
	ChargeType string
}
type Pricing struct {
	Currency       string
	ServiceCharges []ServiceCharges
}

type PricedItineraries struct {
	OnwardPricedItinerary []Flight
	ReturnPricedItinerary []Flight
	Pricing               Pricing
}
