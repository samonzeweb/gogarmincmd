package main

import (
  "fmt"
  "github.com/samonzeweb/gogarmincmd/tcx"
)

func displayActivities(tcxData *tcx.TCXTrainingCenterDatabase, arguments *CmdArguments) {
  // TODO : produire un résultat plus propre, et gérer plusieurs langues
  for _, activity := range((*tcxData).Activities) {
    fmt.Println("\n--- Activité")
    for _, lap := range(activity.Laps) {
      line := fmt.Sprintf("Distance : %v, Temps : %v, FCmoyen : %v, FCMax : %v",
                          lap.DistanceMeters, lap.TotalTimeSeconds,
                          lap.AverageHeartRateBpm, lap.MaximumHeartRateBpm)
      if arguments.fcMax > 0 {
        line = line + "( moyenne " + percentOfFC((int)(lap.AverageHeartRateBpm), arguments.fcMax) + ")"
      }
      fmt.Println(line)
    }
  }
}

// the value is a int because no more precision is needed
func percentOfFC(fc int, fcMax int) string {
  percent := 100 * fc / fcMax
  return fmt.Sprintf("%02d%%", percent)
}
