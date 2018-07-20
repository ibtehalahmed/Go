package slicewithminimalaverage

import (
	"time"
	"log"
)

func SliceWithMinimalAverage(a []int) int {
	defer timeTrack(time.Now())
	n := len(a)
	minAvg := float64(a[0] +a [1]) / 2
	startingPosition := 0

	if n > 2 {
		avg := float64(a[0] + a[1] + a[2]) / 3
		if avg < minAvg {
			minAvg = avg
		}
	}
	for i := 1; i < n-2; i++ {
		avg := float64(a[i] + a[i+1]) / 2
		if avg < minAvg {
			minAvg = avg
			startingPosition = i
		}
		avg = float64(a[i] + a[i+1]  +a[i+2]) / 3
		if avg < minAvg {
			minAvg = avg
			startingPosition = i
		}
	}
	if n > 2 {
		avg := float64(a[n-2] + a[n-1]) / 2
		if avg < minAvg {
			minAvg = avg
			startingPosition = n - 2
		}
	}
	return startingPosition
}
func timeTrack(start time.Time) {
	elapsed := time.Since(start)
	log.Printf("Time spent: %s", elapsed)
}