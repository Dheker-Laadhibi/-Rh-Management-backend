package notifications

type PriorityType int

const (
	High PriorityType = iota
	Medium
	Low
)

type Notification struct {
	Idnotif        int          `json:"idnotifcation" gorm:"primaryKey;autoIncrement"`
	Typenotif      string       `json:"typenotif"`
	Subject        string       `json:"subject"`
	Content        string       `json:"content"`
	Priority       PriorityType `json:"priority"`
	Deliverystatus string       `json:"deliverystatus"`
}
