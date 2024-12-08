package main

import (
	"fmt"
	"math/rand"
)

type Kilometrs uint
type Days uint
type MillionDollars uint
type KilometersSec uint

const DistanceToMars = 62_100_000
const SecInDay = 86_400

type Ticket struct {
	SpaceLineName string
	Duration      Days
	TripType      string
	Price         MillionDollars
}

var TripeType = [2]string{"RoundTripe", "OneWay"}

var SpaceStation = [5]string{
	"SpaceAdventures",
	"SpaceX",
	"VirginGalaxy",
	"Baykhonure",
	"DeadSpace"}

func SortStation() []int {
	var a []int = rand.Perm(5)
	var b []int = rand.Perm(5)
	for _, c := range b {
		a = append(a, c)
	}
	return a
}

func SpeedSet() KilometersSec {
	return KilometersSec(rand.Intn(15) + 16)
}

func PriceSet(a KilometersSec) MillionDollars {
	return MillionDollars(int(a) + 22)
}

func DaysSet() Days {
	return Days(DistanceToMars / uint(SpeedSet()*SecInDay))
}

func TripeTipeSet() string {
	return TripeType[rand.Intn(2)]
}

func (t *Ticket) TicketGeneration() {
	for i := 0; i < 10; i++ {
		(*t).SpaceLineName = SpaceStation[SortStation()[i]]
		(*t).TripType = TripeTipeSet()
		if t.TripType == TripeType[0] {
			(*t).Duration = DaysSet() * 2
			(*t).Price = PriceSet(SpeedSet()) * 2
		}
		if t.TripType == TripeType[1] {
			(*t).Duration = DaysSet()
			(*t).Price = PriceSet(SpeedSet())
		}
		fmt.Printf("%-15v %4v    %-10v   $%4v\n", t.SpaceLineName,
			t.Duration, t.TripType, t.Price)
	}
}

func PrintTicket() {
	var t Ticket
	fmt.Println()
	fmt.Println("Spaceline         Days  Trip type    Price")
	fmt.Println("==========================================")
	t.TicketGeneration()
	fmt.Println()
}

func main() {
	PrintTicket()
}
