package warehouse

// Vendor holds vendor-specific settings.
type Vendor struct {
	WeekendBuffer  int
	WeekdayBuffer  int
	Location       string
	Email          string
	PoNum          string
	ShareOffPrice  bool // Send warning off price if true.
	WarningTrigger int  // Maximum days a vendor not provide info.
}
