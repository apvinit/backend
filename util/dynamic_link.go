package util

import (
	"backend/model"
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

const key = "AIzaSyD5nKQMrrZnRN089ynNPQ6z0AJPxe1j-hA"

const dynamicKeyURL = "https://firebasedynamiclinks.googleapis.com/v1/shortLinks?key=" + key

// DynamicLinkParameter object
type DynamicLinkParameter struct {
	DynamicLinkInfo `json:"dynamicLinkInfo"`
	Suffix          `json:"suffix"`
}

// DynamicLinkInfo object for dynamic link object
type DynamicLinkInfo struct {
	DomainURIPrefix   string `json:"domainUriPrefix"`
	Link              string `json:"link"`
	AndroidInfo       `json:"androidInfo"`
	SocialMetaTagInfo `json:"socialMetaTagInfo"`
}

// Suffix for deciding if the link is SHORT or UNGUESSABLE
type Suffix struct {
	Option string `json:"option"`
}

// AndroidInfo for android info object
type AndroidInfo struct {
	AndroidPackageName           string `json:"androidPackageName,omitempty"`
	AndroidFallbackLink          string `json:"androidFallbackLink,omitempty"`
	AndroidMinPackageVersionCode string `json:"androidMinPackageVersionCode,omitempty"`
}

// SocialMetaTagInfo object
type SocialMetaTagInfo struct {
	SocialTitle       string `json:"socialTitle,omitempty"`
	SocialDescription string `json:"socialDescription,omitempty"`
	SocialImageLink   string `json:"socialImageLink,omitempty"`
}

type shortLinkResp struct {
	ShortLink string `json:"shortLink"`
}

// CreateDynamicLink creates short dynamic link
func CreateDynamicLink(p *model.Post) (string, error) {
	a := AndroidInfo{
		AndroidPackageName:           "xyz.codingabc.jobadda.debug",
		AndroidMinPackageVersionCode: "7",
	}

	s := SocialMetaTagInfo{
		SocialTitle:       p.Title,
		SocialDescription: p.Info,
		SocialImageLink:   p.ImageLink,
	}

	d := DynamicLinkInfo{
		DomainURIPrefix:   "https://sjobadda.page.link",
		Link:              "https://sarkarijobadda.in/" + p.ID.Hex(),
		AndroidInfo:       a,
		SocialMetaTagInfo: s,
	}

	su := Suffix{
		Option: "SHORT",
	}

	param := DynamicLinkParameter{
		DynamicLinkInfo: d,
		Suffix:          su,
	}

	output, _ := json.Marshal(param)

	resp, err := http.Post(dynamicKeyURL, "application/json", bytes.NewBuffer(output))
	if err != nil {
		print(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		print(err)
	}
	var sl shortLinkResp
	err = json.Unmarshal(body, &sl)
	if err != nil {
		return "", err
	}
	return sl.ShortLink, nil
}
