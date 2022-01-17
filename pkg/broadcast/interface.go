package broadcast

// Broadcast -
type Broadcast interface {
	// Sub - 監聽
	Sub(topic string, cb Handler)

	// Pub - 推播
	Pub(topic string, v interface{}) error

	// CloseAll -
	CloseAll()
}
