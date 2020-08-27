package endpoints

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strings"
)

const (
	OSRSRequest = "https://api.scape.fashion"
	OSRSLink    = "https://scape.fashion"

	RS3Request = "https://api.rune.scape.fashion"
	RS3Link    = "https://rune.scape.fashion"
)

type SFItem struct {
	Images struct {
		Detail string `json:"detail"`
	} `json:"images"`
	Name string `json:"name"`
	Slot string `json:"slot"`
	Wiki struct {
		Api    string `json:"api"`
		Link   string `json:"link"`
		PageID int    `json:"pageId"`
	} `json:"wiki"`
	Colors []string `json:"colors"`
	Match  float64  `json:"match"`
}

func (item *SFItem) String() string {
	jsonData, _ := json.MarshalIndent(item, "", "    ")
	return string(jsonData)
}

type SFSlot string

const (
	Ammunition SFSlot = "Ammunition"
	Body       SFSlot = "Body"
	Cape       SFSlot = "Cape"
	Feet       SFSlot = "Feet"
	Hand       SFSlot = "Hand"
	Head       SFSlot = "Head"
	Leg        SFSlot = "Leg"
	Neck       SFSlot = "Neck"
	Ring       SFSlot = "Ring"
	Shield     SFSlot = "Shield"
	Weapon     SFSlot = "Weapon"
)

type SFRequestRS3 struct {
	Slot  SFSlot // Optional. If included, will add an additional parameter for slot filters
	Param string // Will either be a hexcode (#00FF00) or an item name
	Link string // The link that goes along with the request
}

type SFRequestOSRS struct {
	Slot  SFSlot
	Param string
	Link string
} // TODO: Need an interface for URL generation based on Request. Maybe add game version as a struct field?

func (req *SFRequestRS3) BuildURI() (*url.URL, error) {
	// Find out if the param is an item or a color
	var uri string
	if strings.HasPrefix(req.Param, "#") {
		uri = RS3Request + "/colors/" + url.QueryEscape(req.Param)
		req.Link = RS3Link + "/colors/" + url.QueryEscape(req.Param)
	} else {
		uri = RS3Request + "/items/" + url.QueryEscape(req.Param)
		req.Link = RS3Link + "/items/" + url.QueryEscape(req.Param)
	}

	params := url.Values{}
	if req.Slot != "" {
		params.Add("slot", string(req.Slot))
	}

	return url.Parse(uri + "?" + params.Encode())
}

func (req *SFRequestOSRS) BuildURI() (*url.URL, error) {
	// Find out if the param is an item or a color
	var uri string
	if strings.HasPrefix(req.Param, "#") {
		uri = OSRSRequest + "/colors/" + url.QueryEscape(req.Param)
		req.Link = OSRSLink + "/colors/" + url.QueryEscape(req.Param)
	} else {
		uri = OSRSRequest + "/items/" + url.QueryEscape(req.Param)
		req.Link = OSRSLink + "/items/" + url.QueryEscape(req.Param)
	}

	params := url.Values{}
	if req.Slot != "" {
		params.Add("slot", string(req.Slot))
	}

	return url.Parse(uri + "?" + params.Encode())
}

type SFResponse struct {
	Items []SFItem `json:"items"`
	Link  string
}

// Color search -> WORKING
// Item search -> NOT WORKING // TODO
// Slot parameters -> NOT WORKING // TODO
func SFSearch(req SFRequestOSRS) (SFResponse, error) {
	uri, err := req.BuildURI()
	if err != nil {
		return SFResponse{}, err
	}

	fmt.Println(uri.String())

	body, err := makeRequest(uri)
	if err != nil {
		return SFResponse{}, err
	}

	items := SFResponse{
		Link: req.Link,
	}

	err = json.Unmarshal(body, &items)
	if err != nil {
		return SFResponse{}, err
	}

	return items, nil
}
