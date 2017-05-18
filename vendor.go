package warehouse

// Vendor holds vendor-specific quantity settings.
type Vendor struct {
	WeekendBuffer int
	WeekdayBuffer int
	Location      string
	Email         string
	PoNum         string
}
