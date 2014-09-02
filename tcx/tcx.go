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
  TotalTimeSeconds    float32   `xml:"TotalTimeSeconds"`
  DistanceMeters      float32   `xml:"DistanceMeters"`
  AverageHeartRateBpm float32   `xml:"AverageHeartRateBpm>Value"`
  MaximumHeartRateBpm float32   `xml:"MaximumHeartRateBpm>Value"`
}

func ParseTCX(xmlTCX []byte) (*TCXTrainingCenterDatabase, error) {
  var tcx TCXTrainingCenterDatabase
  err := xml.Unmarshal(xmlTCX, &tcx)
  if err != nil {
    return nil, err
  }
  return &tcx, nil
}
