package tcx

import (
    "encoding/xml"
)

type TCXTrainingCenterDatabase struct {
  XMLName     xml.Name       `xml:"TrainingCenterDatabase"`
  Activities  []TCXActivity  `xml:"Activities>Activity"`
}

 type TCXActivity struct {
   XMLName   xml.Name  `xml:"Activity"`
   Laps      []TCXLap  `xml:"Lap"`
}

type TCXLap struct {
  XMLName             xml.Name  `xml:"Lap"`
  TotalTimeSeconds    float64   `xml:"TotalTimeSeconds"`
  DistanceMeters      float64   `xml:"DistanceMeters"`
  AverageHeartRateBpm int       `xml:"AverageHeartRateBpm>Value"`
  MaximumHeartRateBpm int       `xml:"MaximumHeartRateBpm>Value"`
}

func parseTCX(xmlTCX []byte) (*TCXTrainingCenterDatabase, error) {
  var tcx TCXTrainingCenterDatabase
  err := xml.Unmarshal(xmlTCX, &tcx)
  if err != nil {
    return nil, err
  }
  return &tcx, nil
}
