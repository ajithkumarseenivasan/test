package model

import "os"

type status string

const (
	Success status = "success"
	Failed  status = "failed"
)

type MasterUiResponse struct {
	Status  bool        `json:"status"`
	Content interface{} `json:"content"`
	Message status      `json:"message"`
}

type TestingData struct {
	Data1 string              `jason:"data1"`
	Data2 int                 `json:"data2"`
	Data3 bool                `json:"data3"`
	Data4 float64             `json:"data4"`
	Data5 map[int]string      `json:"data5"`
	Data6 []int               `json:"data6"`
	Data7 MasterUiResponse    `json:"data7"`
	Data8 [10]MasterUiRequest `json:"data8"`
	Data9 os.Signal           `json:"osSignal"`
}
