package main

import (
  "os"
  "fmt"
  "strconv"
  "net/http"
  "io/ioutil"
  "github.com/samonzeweb/gogarmincmd/tcx"
)


func main() {
  idActivity := activityInArgs()

  tcxContent, err := getTCX(idActivity)
  if err != nil {
    panic(err)
  }
  tcxData, err  := tcx.ParseTCX(tcxContent)
  if err != nil {
    panic(err)
  }

  for _, activity := range((*tcxData).Activities) {
    fmt.Println("\n--- Activité")
    for _, lap := range(activity.Laps) {
      fmt.Fprintf(os.Stdout, "Distance : %v, Temps : %v, FCmoyen : %v, FCMax : %v\n",
        lap.DistanceMeters, lap.TotalTimeSeconds, lap.AverageHeartRateBpm, lap.MaximumHeartRateBpm)
    }
  }
}

// ---

func activityInArgs() int64 {
  if len(os.Args) != 2 {
    printUsageAndQuit()
  }
  idActivity, err := strconv.ParseInt(os.Args[1], 10, 64)
  if err != nil {
    printUsageAndQuit()
  }
  return idActivity
}


func printUsageAndQuit() {
  // TODO utiliser basename du programme (pas tout le chemin)
  fmt.Fprintf(os.Stderr, "usage : %v id_activity\n", os.Args[0])
  os.Exit(1)
}

//---

const tmplURLActivity = "http://connect.garmin.com/proxy/activity-service-1.1/tcx/activity/"

func getTCX(idActivity int64) ([]byte, error) {
  url := fmt.Sprintf("%s%d", tmplURLActivity, idActivity)
  resp, err := http.Get(url)
  if err != nil {
    return nil, err
  }
  defer resp.Body.Close()
  xml, err := ioutil.ReadAll(resp.Body)
  if err != nil {
    return nil, err
  }
  return xml, nil
}
