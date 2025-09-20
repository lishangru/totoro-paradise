package types

type Point struct {
	Longitude string `json:"longitude"`
	Latitude  string `json:"latitude"`
}

type RunPoint struct {
	TaskID    string  `json:"taskId"`
	PointID   string  `json:"pointId"`
	PointName string  `json:"pointName"`
	Longitude string  `json:"longitude"`
	Latitude  string  `json:"latitude"`
	PointList []Point `json:"pointList"`
}
