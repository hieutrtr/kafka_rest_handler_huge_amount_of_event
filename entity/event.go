package entity

type object struct {
	Name string `json:"name"`
	Age int `json:"age"`
}
type Event struct {
	Type string `json:"type"`
	Action string `json:"action"`
	Object object `json:"object"`
	CreatedTime int64 `json:"createdTime"`
}
