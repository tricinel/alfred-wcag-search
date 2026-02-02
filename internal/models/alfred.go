package models

type AlfredResponse struct {
	Items []AlfredItem `json:"items"`
}

type AlfredItem struct {
	Title    string      `json:"title"`
	Subtitle string      `json:"subtitle"`
	Arg      string      `json:"arg"`
	Match    string      `json:"match"`
	Valid    bool        `json:"valid"`
	Icon     *AlfredIcon `json:"icon,omitempty"`
}

type AlfredIcon struct {
	Path string `json:"path"`
}
