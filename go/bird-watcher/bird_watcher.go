package birdwatcher

const (
	daysInWeek = 7
)

func splitDataIntoWeeks(Data []int) (Weeks [][]int) {
	for i := 0; i < len(Data); i += daysInWeek {
		end := i + daysInWeek

		// necessary check to avoid slicing beyond
		// slice capacity
		if end > len(Data) {
			end = len(Data)
		}

		Weeks = append(Weeks, Data[i:end])
	}

	return Weeks
}

func getWeekData(birdsPerDay []int, week int) (weekData []int) {
	Weeks := splitDataIntoWeeks(birdsPerDay)
	return Weeks[week-1]
}

// TotalBirdCount return the total bird count by summing
// the individual day's counts.
func TotalBirdCount(birdsPerDay []int) (total int) {
	for _, v := range birdsPerDay {
		total += v
	}
	return total
}

// BirdsInWeek returns the total bird count by summing
// only the items belonging to the given week.
func BirdsInWeek(birdsPerDay []int, week int) int {
	weekData := getWeekData(birdsPerDay, week)
	return TotalBirdCount(weekData)
}

// FixBirdCountLog returns the bird counts after correcting
// the bird counts for alternate days.
func FixBirdCountLog(birdsPerDay []int) (updatedCount []int) {
	for index := range birdsPerDay {
		if index % 2 == 0 {
			birdsPerDay[index] = birdsPerDay[index] + 1
		} 
	}
	return birdsPerDay
}
