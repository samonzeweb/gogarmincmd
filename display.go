package main

import (
  "fmt"
  "os"
  "math"
  "strconv"
  "github.com/samonzeweb/gogarmincmd/tcx"
)

func displayActivities(tcxData *tcx.TCXTrainingCenterDatabase, arguments *CmdArguments) {
  for _, activity := range((*tcxData).Activities) {
    fmt.Println("\n---")

    totalDistance, totalTime, avgFC, maxFC := summary(activity)
    fmt.Println("TOTAL :", formattedRun(totalDistance, totalTime, avgFC, maxFC, arguments))

    for index, lap := range(activity.Laps) {
      fmt.Fprintf(os.Stdout,"Partie %02d : %s", index,
                  formattedRun(lap.DistanceMeters, lap.TotalTimeSeconds, lap.AverageHeartRateBpm, lap.MaximumHeartRateBpm, arguments))
    }
  }
}

func summary(activity tcx.TCXActivity) (totalDistance float64, totalTime float64, avgFC int, maxFC int) {
  totalDistance = 0
  totalTime = 0
  var tmpSumFC float64 = 0
  var tmpMaxFC float64 = 0
  for _, lap := range(activity.Laps) {
    totalDistance += lap.DistanceMeters
    totalTime += lap.TotalTimeSeconds
    tmpSumFC += float64(lap.AverageHeartRateBpm) * lap.TotalTimeSeconds
    if float64(lap.MaximumHeartRateBpm) > tmpMaxFC {
      tmpMaxFC = float64(lap.MaximumHeartRateBpm)
    }
  }
  avgFC = (int)(tmpSumFC / totalTime)
  maxFC = (int)(tmpMaxFC)
  return
}

func formattedRun(distance float64, duration float64, avgFC int, maxFC int, arguments *CmdArguments) string {
  return fmt.Sprintf("%s - %s - %s - FC %s Max %s\n",
                     formattedDistance(distance), formattedDuration(duration) , formattedPace(distance, duration),
                     formattedFC(avgFC, arguments.fcMax), formattedFC(maxFC, arguments.fcMax))
}

func formattedDistance(distance float64) string {
  if distance >= 1000 {
    return fmt.Sprintf("%0.2f km", distance/1000)
  } else {
    return fmt.Sprintf(" %03.0f m ", distance)
  }
}

// the value is a int because no more precision is needed
func percentOfFC(fc int, fcMax int) string {
  percent := 100 * fc / fcMax
  return fmt.Sprintf("%02d%%", percent)
}

func formattedFC(fc int, fcMax int) string {
  formattedFC := strconv.Itoa(fc)
  if fcMax > 0 {
    formattedFC += "(" + percentOfFC(fc, fcMax) + ")"
  }
  return formattedFC
}

func formattedDuration(seconds float64) string {
  hours :=  math.Floor(seconds / 3600)
  remainingSeconds := seconds - 3600 * hours
  minutes := math.Floor(remainingSeconds / 60)
  remainingSeconds = math.Floor(remainingSeconds - 60 * minutes)
  if hours > 0 {
    return fmt.Sprintf("%02.0fH%02.0f'%02.0f", hours, minutes, remainingSeconds)
  } else {
    return fmt.Sprintf("%02.0f'%02.0f", minutes, remainingSeconds)
  }
}

func formattedPace(distance float64, seconds float64) string {
  pace := 1000 * seconds / distance
  return formattedDuration(pace)+"/km"
}
