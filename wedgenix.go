package wedgenix

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"os"

	"github.com/WedgeNix/warehouse-settings/file"
)

// Controller keeps data for Do.
type Controller struct {
	url  string
	user string
	pass string
}

// New new controller, and gets credentials from environment variables
func New() *Controller {
	settingsURL := os.Getenv("SETTINGS_URL")
	settingsUser := os.Getenv("SETTINGS_USER")
	settingsPass := os.Getenv("SETTINGS_PASS")

	return &Controller{
		url:  settingsURL,
		user: settingsUser,
		pass: settingsPass,
	}
}

// Do sends request to settings server for the particular settings file.
func (c *Controller) Do(f file.Any) []error {
	errs := []error{}
	q := "?appname="
	switch f.(type) {
	case file.AppBananas:
		q += "bananas"
	case file.AppD2s:
		q += "d2s"
	case file.AppScriptToRuleThemAll:
		q += "thescripttorulethemall"
	case file.AppEmailGrabber:
		q += "EmailGrabber"
	}
	req, err := http.NewRequest(http.MethodGet, c.url+q, nil)
	if err != nil {
		return append(errs, err)
	}
	req.Header.Add("User", c.user)
	req.Header.Add("Pass", c.pass)

	cl := http.Client{}
	resp, err := cl.Do(req)
	if err != nil {
		return append(errs, err)
	}
	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return append(errs, err)
	}

	err = json.Unmarshal(b, f)
	if err != nil {
		return append(errs, err)
	}
	return nil
}