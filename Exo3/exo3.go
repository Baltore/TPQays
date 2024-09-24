package exo

func Ft_non_overlap(intervals [][]int) int {

	if len(intervals) == 0 {
		return 0
	}

	for i := 0; i < len(intervals)-1; i++ {
		for j := i + 1; j < len(intervals); j++ {
			if intervals[i][1] > intervals[j][1] {
				intervals[i], intervals[j] = intervals[j], intervals[i]
			}
		}
	}

	count := 0
	end := intervals[0][1]

	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] < end {
			count++
		} else {
			end = intervals[i][1]
		}
	}

	return count
}
