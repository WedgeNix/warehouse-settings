package wedgenix

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/WedgeNix/warehouse-settings/app"

	"os"
)

// Controller keeps data for Do.
type Controller struct {
	url  string
	user string
	pass string
}

func Settings(ptr app.Any) error {
	url, found := os.LookupEnv("SETTINGS_URL")
	if !found {
		return errors.New("missing SETTINGS_URL")
	}
	user, found := os.LookupEnv("SETTINGS_USER")
	if !found {
		return errors.New("missing SETTINGS_USER")
	}
	pass, found := os.LookupEnv("SETTINGS_PASS")
	if !found {
		return errors.New("missing SETTINGS_PASS")
	}
	ctr := Controller{
		url:  url,
		user: user,
		pass: pass,
	}
	return ctr.Do(ptr)
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
func (c *Controller) Do(ptr app.Any) error {
	q := "?appname="
	switch ptr.(type) {
	case *app.All:
		q += "all"
	case *app.Bananas:
		q += "bananas"
	case *app.D2s:
		q += "d2s"
	case *app.ScriptToRuleThemAll:
		q += "thescripttorulethemall"
	case *app.EmailGrabber:
		q += "EmailGrabber"
	default:
		return errors.New("ptr required")
	}
	req, err := http.NewRequest(http.MethodGet, c.url+q, nil)
	if err != nil {
		return err
	}
	req.Header.Add("User", c.user)
	req.Header.Add("Pass", c.pass)

	cl := http.Client{}
	resp, err := cl.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 300 {
		return errors.New("warehouse-settings: " + resp.Status)
	}

	return json.NewDecoder(resp.Body).Decode(ptr)
}
