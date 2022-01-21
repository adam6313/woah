package conv

// TOPIC -
type TOPIC string

// String -
func (t TOPIC) String() string {
	return string(t)
}

const (
	// EVENT -
	EVENT TOPIC = "topic.event"
)
