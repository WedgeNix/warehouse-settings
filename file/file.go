package file

// Any is a type safe interface.
type Any interface {
	__()
}

// Ats is for the Script to rule them all part of the v2 settings file.
type Ats map[string]struct {
	FileCount   int
	TestMode    bool
	GoogleSheet bool
	SheetNumber int
}

// EmailDownlaod GAS scrupt uses this info.
type EmailDownlaod map[string]struct {
	Label string
	Query []QueryStruct
}

// QueryStruct holds search settings for emailDownlaod.
type QueryStruct map[string]struct {
	Ext   string
	Query string
	Time  bool
}

// SettingsFile is the overal data struct for whole v2 settings.
type SettingsFile map[string]struct {
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
	RecvBuffer    int
	ReorderWindow float64
}

// AppBananas struct for Bannanas app settings.
type AppBananas map[string]struct {
	Email         []string
	FileDownload  bool
	Location      string
	PONum         string
	ShareOffPrice bool
	WaitingPeriod int
	UseUPC        bool
	Monitor       bool
	RecvBuffer    int
	ReorderWindow float64
}

// AppD2s struct for D2s app settings.
type AppD2s map[string]struct {
	WeekdayBuffer int
	WeekendBuffer int
	WaitingPeriod int
}

// AppScriptToRuleThemAll struct for ScriptToRuleThemAll app settings.
type AppScriptToRuleThemAll map[string]map[string]struct {
	FileCount   int
	TestMode    bool
	GoogleSheet bool
	SheetNumber int
	Location    string
}

// AppEmailGrabber struct for EmailGrabber app settings.
type AppEmailGrabber map[string]struct {
	Label string
	Query []QueryStruct
}

func (_ Ats) __()                    {}
func (_ EmailDownlaod) __()          {}
func (_ QueryStruct) __()            {}
func (_ SettingsFile) __()           {}
func (_ AppBananas) __()             {}
func (_ AppD2s) __()                 {}
func (_ AppScriptToRuleThemAll) __() {}
func (_ AppEmailGrabber) __()        {}
