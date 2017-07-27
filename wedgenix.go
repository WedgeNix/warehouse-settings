package wedgenix

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"os"

	"github.com/WedgeNix/warehouse-settings/file"
)

type Controller struct {
	url  string
	user string
	pass string
}

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

func (c *Controller) Do(f file.Any) []error {
	errs := []error{}
	req, err := http.NewRequest(http.MethodGet, c.url+f.q, nil)
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

	settings := map[string]f{}
	err = json.Unmarshal(b, &settings)
	if err != nil {
		return append(errs, err)
	}

}
