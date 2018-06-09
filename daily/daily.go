// Package daily encapsulates work with daily images
package daily

// Daily fetch daily images
type Daily struct {
	dg DailyGetter
}

// NewDailyGetter creates new daily getter instance
func NewDailyGetter(dg DailyGetter) *Daily {
	return &Daily{
		dg: dg,
	}
}

// GetImage ImageGetter implementation
func (d *Daily) GetImage() ([]byte, error) {
	return d.dg.GetDailyImage()
}
