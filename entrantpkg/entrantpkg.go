package entrantpkg

func CalculateEntry(entrants []TEntrant) float64 {
	SE := countEntrants(entrants)
	sum := SE.Adult*Prices.Adult + SE.Child*Prices.Child + SE.Senior*Prices.Senior
	return sum
}

func countEntrants(entrants []TEntrant) TSortedEntry {
	var sortedEntr TSortedEntry
	for _, v := range entrants {
		switch {
		case v.Age < 18:
			sortedEntr.Child++
		case v.Age < 50:
			sortedEntr.Adult++
		default:
			sortedEntr.Senior++
		}
	}
	return sortedEntr
}
