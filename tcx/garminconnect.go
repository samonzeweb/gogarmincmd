package tcx

import (
  "fmt"
  "net/http"
  "io/ioutil"
)

const tmplURLActivity = "http://connect.garmin.com/proxy/activity-service-1.1/tcx/activity/"

func GetTCX(idActivity int64) (*TCXTrainingCenterDatabase, error) {

  tcxContent, err := fetchTCX(idActivity)
  if err != nil {
    return nil, err
  }

  tcxData, err  := parseTCX(tcxContent)
  if err != nil {
    return nil, err
  }

  return tcxData, err

}

func fetchTCX(idActivity int64) ([]byte, error) {
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
