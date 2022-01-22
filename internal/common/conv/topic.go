package conv

// TOPIC -
type TOPIC string

// String -
func (t TOPIC) String() string {
	return string(t)
}

const (
	// EVENT - 過去式
	EVENT TOPIC = "topic.event"
)
