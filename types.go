package bandiera

type HttpClient interface {
	GetUrlContent(url string, params Params) ([]byte, error)
}

type GroupName string
type FeatureName string
type Flags map[FeatureName]bool
type GroupFlags map[GroupName]Flags
type Params map[string]string

type AllResponse struct {
	Warning    string     `json:"warning,omitempty"`
	GroupFlags GroupFlags `json:"response"`
}

type GroupResponse struct {
	Warning string `json:"warning,omitempty"`
	Flags   Flags  `json:"response"`
}

type FeatureResponse struct {
	Warning string `json:"warning,omitempty"`
	Enabled bool   `json:"response"`
}
