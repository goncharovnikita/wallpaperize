package daily

// DailyGetter daily getter interface
type DailyGetter interface {
	GetDailyImage() ([]byte, error)
}
