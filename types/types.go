package types

// Ats is for the Script to rule them all part of the v2 settings file.
type Ats struct {
	FileCount   int
	TestMode    bool
	GoogleSheet bool
	SheetNumber int
}

// EmailDownlaod GAS scrupt uses this info.
type EmailDownlaod struct {
	Label string
	Query []QueryStruct
}

// QueryStruct holds search settings for emailDownlaod.
type QueryStruct struct {
	Ext   string
	Query string
	Time  bool
}

// All is the overal data struct for whole v2 settings.
type All struct {
	Email         []string
	FileDownload  bool
	Location      string
	PONum         string
	ShareOffPrice bool
	WaitingPeriod int
	WeekdayBuffer int
	WeekendBuffer int
	UseUPC        bool
	ATS           Ats
	EmaiDownload  EmailDownlaod
	Monitor       bool
	Hybrid        bool
	Regex         string
	ReordPtAdd    int
}

// Bananas struct for Bannanas app settings.
type Bananas struct {
	Email         []string
	FileDownload  bool
	Location      string
	PONum         string
	ShareOffPrice bool
	WaitingPeriod int
	UseUPC        bool
	Monitor       bool
	Regex         string
	ReordPtAdd    int
	Hybrid        bool
}

func (b Bananas) Copy() Bananas {
	db := b
	db.Email = make([]string, len(b.Email))
	copy(db.Email, b.Email)
	return db
}

// D2s struct for D2s app settings.
type D2s struct {
	WeekdayBuffer int
	WeekendBuffer int
	WaitingPeriod int
}

// ScriptToRuleThemAll struct for ScriptToRuleThemAll app settings.
type ScriptToRuleThemAll struct {
	FileCount   int
	TestMode    bool
	GoogleSheet bool
	SheetNumber int
	Location    string
}

// EmailGrabber struct for EmailGrabber app settings.
type EmailGrabber struct {
	Label string
	Query []QueryStruct
}
