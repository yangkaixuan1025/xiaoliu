package main

import "sort"

type TicketSecond struct {
	Src    string
	Dst    string
	Second int
}

func findItinerary(tickets [][]string) []string {
	var (
		targets = map[string][]string{}
	)

	for _, ticket := range tickets {
		targets[ticket[0]] = append(targets[ticket[0]], ticket[1])
	}
	for key := range targets {
		sort.Strings(targets[key])
	}

	return resS

}

func findItineraryBackTracking(targets map[string][]string, ticketNum int, start string) bool {
	if len(resS) == ticketNum+1 {
		return true
	}

	for _, target := range targets[start] {
		resS = append(resS, target)

	}
	return false

}
