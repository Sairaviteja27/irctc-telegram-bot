package model

type PassengerStruct struct {
	Name string `json:"name"`
	BookingStatus string `json:"booking_status"`
	CurrentStatus string `json:"current_status"`
}
type DepartureData struct{
	DepartureDate string `json:"departure_date"`
	DepartureTime string `json:"departure_time"`
	}

type ArrivalData struct{
	ArrivalDate string `json:"arrival_date"`
	ArrivalTime string `json:"arrival_time"`
}
type ResponseModel struct {
	Passenger []PassengerStruct
	BoardingStation string `json:"reservation_upto"`
	ReservationUpto string `json:"reservation_upto"`
	DepartureData
	ArrivalData
	Quota string `json:"quota"`
	Class string `json:"class"`
	ChartStatus string `json:"chart_status"`
	TrainName string `json:"train_name"`
	TrainNumber string `json:"train_number"`
}




