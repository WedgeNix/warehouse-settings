package warehouse

// Vendor holds vendor-specific settings.
type Vendor struct {
	WeekendBuffer int
	WeekdayBuffer int
	Location      string
	Email         string
	PONum         string
	ShareOffPrice bool // Send warning off price if true.
	WaitingPeriod int  // Maximum days a vendor not provide info.
	FileDownload  bool // If vendor needs to download file for orders CSV file.
	UseUPC        bool // If vendor only accepts UPC for email orders.
}

// Settings is the JSON settings file format.
type Settings map[string]Vendor
