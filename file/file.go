package file

import "github.com/WedgeNix/warehouse-settings/types"

// Any is a type safe
type Any interface {
	__()
}
type EmailDownlaod map[string]types.EmailDownlaod
type QueryStruct map[string]types.QueryStruct
type SettingsFile map[string]types.SettingsFile
type AppBananas map[string]types.AppBananas
type AppD2s map[string]types.AppD2s
type AppScriptToRuleThemAll map[string]types.AppScriptToRuleThemAll
type AppEmailGrabber map[string]types.AppEmailGrabber

func (_ EmailDownlaod) __()          {}
func (_ QueryStruct) __()            {}
func (_ SettingsFile) __()           {}
func (_ AppBananas) __()             {}
func (_ AppD2s) __()                 {}
func (_ AppScriptToRuleThemAll) __() {}
func (_ AppEmailGrabber) __()        {}
