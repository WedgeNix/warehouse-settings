package app

import "github.com/WedgeNix/warehouse-settings/types"

// Any is a type safe
type Any interface {
	__()
}

type All map[string]types.All
type Bananas map[string]types.Bananas
type D2s map[string]types.D2s
type ScriptToRuleThemAll map[string]types.ScriptToRuleThemAll
type EmailGrabber map[string]types.EmailGrabber

func (_ SettingsFile) __()        {}
func (_ Bananas) __()             {}
func (_ D2s) __()                 {}
func (_ ScriptToRuleThemAll) __() {}
func (_ EmailGrabber) __()        {}
