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
	Link  string
}

func (req *SFRequestRS3) BuildURI() *url.URL {
	var uri string
	if strings.HasPrefix(req.Param, "#") {
		uri = RS3Request + "/colors/" + url.QueryEscape(req.Param)
		req.Link = RS3Link + "/colors/" + url.QueryEscape(req.Param)
	} else {
		p := url.URL{Path: req.Param}
		uri = OSRSRequest + "/items/" + p.String()
		req.Link = OSRSLink + "/items/" + p.String()
	}

	params := url.Values{}
	if req.Slot != "" {
		params.Add("slot", string(req.Slot))
	}

	ret, _ := url.Parse(uri + "?" + params.Encode())
	return ret
}

func (req *SFRequestOSRS) BuildURI() *url.URL {
	var uri string
	if strings.HasPrefix(req.Param, "#") {
		uri = OSRSRequest + "/colors/" + url.QueryEscape(req.Param)
		req.Link = OSRSLink + "/colors/" + url.QueryEscape(req.Param)
	} else {
		p := url.URL{Path: req.Param}
		uri = OSRSRequest + "/items/" + p.String()
		req.Link = OSRSLink + "/items/" + p.String()
	}

	params := url.Values{}
	if req.Slot != "" {
		params.Add("slot", string(req.Slot))
	}

	ret, _ := url.Parse(uri + "?" + params.Encode())
	return ret
}

type SFResponse struct {
	Items []SFItem `json:"items"`
	Link  string
}

func SFSearch(req Requester) (SFResponse, error) {
	uri := req.BuildURI()

	fmt.Println(uri.String())

	body, err := makeRequest(uri)
	if err != nil {
		return SFResponse{}, err
	}

	var items SFResponse
	switch r := req.(type) {
	case *SFRequestRS3:
		items.Link = r.Link
	case *SFRequestOSRS:
		items.Link = r.Link
	}

	err = json.Unmarshal(body, &items)
	if err != nil {
		return SFResponse{}, err
	}

	return items, nil
}
