package proc

import (
	"fmt"
	"math"

	"github.com/tormoder/fit"
)

const (
	METERS_TO_YARDS = 1.09361
)

func metersToYards(meters float64) float64 {
	dist := float64(meters) * METERS_TO_YARDS
	return math.Round(dist)
}

func formatPace(pace float64) string {
	minutes := int(pace) / 60
	seconds := int((pace - float64(minutes)*60))
	return fmt.Sprintf("%d:%02d ", minutes, seconds)
}

func ProcessOpenWaterSession(session fit.SessionMsg) {
	s := fmt.Sprintf("%.3f yards", metersToYards(session.GetTotalDistanceScaled()))
	fmt.Println("Total Distance Transformed:", s)
}

// TODO: Rename this function to capture that is for open Water really.
// TODO: Add numbers of minutes that you stopped to rest.
// TODO: Can be number of times that you stopped to rest.
func ProcessActivityLaps(activity fit.ActivityFile) {
	for _, lap := range activity.Laps {
		dist := metersToYards(lap.GetTotalDistanceScaled())
		lapTotalTime := formatPace(lap.GetTotalTimerTimeScaled())
		lapElapseTime := formatPace(lap.GetTotalElapsedTimeScaled())
		per_100 := formatPace(lap.GetTotalTimerTimeScaled() / (dist / 100))
		fmt.Println("Distance:", dist,
			"/ Avege HR:", lap.AvgHeartRate,
			"/ Max HR:", lap.MaxHeartRate,
			"/ time: ", lapTotalTime,
			"/ lapElapseTime: ", lapElapseTime,
			"/ pace: ", per_100)
	}
}
