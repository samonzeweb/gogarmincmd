package tcx

import (
  "fmt"
  "errors"
  "net/http"
  "io/ioutil"
)

const tmplURLActivity = "http://connect.garmin.com/proxy/activity-service-1.1/tcx/activity/"

func GetTCX(idActivity string) (*TCXTrainingCenterDatabase, error) {

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

func fetchTCX(idActivity string) ([]byte, error) {
  url := fmt.Sprintf("%s%s", tmplURLActivity, idActivity)
  resp, err := http.Get(url)
  if err != nil {
    return nil, err
  }
  if resp.StatusCode != http.StatusOK {
    return nil,  errors.New("Status "+resp.Status)
  }
  defer resp.Body.Close()
  xml, err := ioutil.ReadAll(resp.Body)
  if err != nil {
    return nil, err
  }
  return xml, nil
}
