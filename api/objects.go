package api

type BaseIndicator struct {
	AccessReason string `json:"access_reason"`
	AccessType   string `json:"access_type"`
	Content      string `json:"content"`
	Description  string `json:"description"`
	Id           int    `json:"id"`
	Indicator    string `json:"indicator"`
	Title        string `json:"title"`
	Type         string `json:"type"`
}

type PulseInfo struct {
	Count      int      `json:"count"`
	Pulses     []Pulse  `json:"pulses"`
	References []string `json:"references"`
}

type PulseAuthor struct {
	AvatarUrl    string `json:"avatar_url"`
	Id           string `json:"id"`
	IsFollowing  bool   `json:"is_following"`
	IsSubscribed bool   `json:"is_subscribed"`
	Username     string `json:"username"`
}

type PulseGroup struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type IndicatorTypeCounts struct {
	IPv4     int `json:"ipv4"`
	SHA256   int `json:"sha256"`
	Domain   int `json:"domain"`
	Hostname int `json:"hostname"`
}

type Observation struct {
	AccessReason       string `json:"access_reason"`
	AccessType         string `json:"access_type"`
	Content            string `json:"content"`
	Description        string `json:"description"`
	Expiration         string `json:"expiration"`
	IsActive           bool   `json:"is_active"`
	ObservationCreated string `json:"observation_created"`
	PulseKey           string `json:"pulse_key"`
	Role               string `json:"role"`
	Title              string `json:"title"`
}

type Pulse struct {
	TLP                 string              `json:"tlp"`
	Adversary           string              `json:"adversay"`
	Author              PulseAuthor         `json:"author"`
	ClonedFrom          interface{}         `json:"cloned_from"`
	CommentCount        int                 `json:"comment_count"`
	Created             string              `json:"created"`
	Description         string              `json:"description"`
	DownvotesCount      int                 `json:"downvote_count"`
	ExportCount         int                 `json:"export_count"`
	FollwerCount        int                 `json:"follower_count"`
	Groups              []PulseGroup        `json:"groups"`
	Id                  string              `json:"id"`
	InGroup             bool                `json:"in_group"`
	IndicatorCount      int                 `json:"indicator_count"`
	IndicatorTypeCounts IndicatorTypeCounts `json:"indicator_type_counts"`
	Industries          []string            `json:"industries"`
	IsAuthor            bool                `json:"is_author"`
	IsFollowing         bool                `json:"is_following"`
	IsModified          bool                `json:"is_modified"`
	IsSubscribing       bool                `json:"is_subscribing"`
	Locked              bool                `json:"locked"`
	Modified            string              `json:"modified"`
	ModifiedText        string              `json:"modified_text"`
	Name                string              `json:"name"`
	Observation         Observation         `json:"observation"`
	Public              bool                `json:"public"`
	PulseSource         string              `json:"pulse_source"`
	Reference           []string            `json:"reference"`
	SubscriberCount     int                 `json:"subscriber_count"`
	Tags                []string            `json:"tags"`
	TargetCountries     []string            `json:"target_countries"`
	UpvoteCount         int                 `json:"upvote_count"`
	ValidatorCount      int                 `json:"validator_count"`
	Vote                int                 `json:"vote"`
	VoteCount           int                 `json:"vote_count"`
}

type Validation struct {
	Message string `json:"message"`
	Name    string `json:"name"`
	Source  string `json:"source"`
}

type DomainGeneral struct {
	Alexa         string        `json:"alexa"`
	BaseIndicator BaseIndicator `json:"base_indicator"`
	Indicator     string        `json:"indicator"`
	PulseInfo     PulseInfo     `json:"pulse_info"`
	Sections      []string      `json:"sections"`
	Type          string        `json:"type"`
	Validation    []Validation  `json:"validation"`
	Whois         string        "json:`whois`"
}

type IPv4General struct {
	AreaCode      int           `json:"area_code"`
	Asn           string        `json:"asn"`
	BaseIndicator BaseIndicator `json:"base_indicator"`
	Charset       int           `json:"charset"`
	City          string        `json:"city"`
	CityData      bool          `json:"city_data"`
	ContinentCode string        `json:"continet_code"`
	CountryCode   string        `json:"country_code"`
	CountryCode3  string        `json:"country_code3"`
	CountryName   string        `json:"country_name"`
	DmaCode       int           `json:"dma_code"`
	FlagTitle     string        `json:"flag_title"`
	FlagUrl       string        `json:"flag_url"`
	Latitude      float32       `json:"latitude"`
	Longitude     float32       `json:"longitude"`
	PostalCode    string        `json:"postal_code"`
	PulseInfo     PulseInfo     `json:"pulse_info"`
	Region        string        `json:"region"`
	Sections      []string      `json:"sections"`
	Type          string        `json:"type"`
	TypeTitle     string        `json:"type_title"`
	Validation    []Validation  `json:"validation"`
	Whois         string        "json:`whois`"
}

type IPv4Reputation struct {
}

type Malware struct {
	Count    int           `json:"count"`
	Data     []MalwareData `json:"data"`
	Next     string        `json:"next"`
	Previous string        `json:"previous"`
	Size     int           `json:"size"`
}

type IPv4UrlList struct {
	ActualSize int        `json:"actual_size"`
	FullSize   int        `json:"full_size"`
	HasNext    bool       `json:"has_next"`
	Limit      int        `json:limit"`
	PageNum    int        `json:page_num"`
	Paged      bool       `json:"paged"`
	UrlList    []UrlEntry `json:"url_list"`
}

type GenPassiveDns struct {
	Count      int          `json:"count"`
	PassiveDns []PassiveDns `json:"passive_dns"`
}

type Geo struct {
	AreaCode      int     `json:"area_code"`
	Asn           string  `json:"asn"`
	Charset       int     `json:"charset"`
	City          string  `json:"city"`
	CityData      bool    `json:"city_data"`
	ContinentCode string  `json:"continet_code"`
	CountryCode   string  `json:"country_code"`
	CountryCode3  string  `json:"country_code3"`
	CountryName   string  `json:"country_name"`
	DmaCode       int     `json:"dma_code"`
	FlagTitle     string  `json:"flag_title"`
	FlagUrl       string  `json:"flag_url"`
	Latitude      float32 `json:"latitude"`
	Longitude     float32 `json:"longitude"`
	PostalCode    string  `json:"postal_code"`
	Region        string  `json:"region"`
}

type MalwareData struct {
	DatetimeInt int    `json:"datetime_int"`
	Hash        string `json:"hash"`
}

type UrlEntry struct {
	Date     string      `json:"data"`
	Domain   string      `json:"domain"`
	Encoded  string      `json:"encoded"`
	Hostname string      `json:hostname"`
	Httpcode string      `json.httpcode"`
	Result   interface{} `json:"result"`
	Url      string      `json:"url"`
}

type DomainUrlList struct {
	ActualSize int        `json:"actual_size"`
	FullSize   int        `json:"full_size"`
	HasNext    bool       `json:"has_next"`
	Limit      int        `json:limit"`
	PageNum    int        `json:page_num"`
	Paged      bool       `json:"paged"`
	UrlList    []UrlEntry `json:"url_list"`
}

type PassiveDns struct {
	Address       string `json:"address"`
	AssetType     string `json:"asset_type"`
	First         string `json:"first"`
	FlagTitle     string `json:"flag_title"`
	FlagUrl       string `json:"flag_url"`
	Hostname      string `json:"hostname"`
	IndicatorLink string `json:"indicator_link"`
	Last          string `json:"last"`
}

type DomainFull struct {
	PassiveDns GenPassiveDns
	UrlList    DomainUrlList
	Malware    Malware
}

type IPv6General struct {
	BaseIndicator BaseIndicator `json:"base_indicator"`
	Indicator     string        `json:"indicator"`
	PulseInfo     PulseInfo     `json:"pulse_info"`
	Reputatioin   int           `json:"reputation"`
	Sections      []string      `json:"sections"`
	Type          string        `json:"type"`
	TypeTitle     string        `json:"type_title"`
	Validation    []Validation  `json:"validation"`
	Whois         string        "json:`whois`"
}

type IPv6Reputation struct {
}

type IPv6UrlList struct {
	ActualSize int        `json:"actual_size"`
	FullSize   int        `json:"full_size"`
	HasNext    bool       `json:"has_next"`
	Limit      int        `json:limit"`
	PageNum    int        `json:page_num"`
	Paged      bool       `json:"paged"`
	UrlList    []UrlEntry `json:"url_list"`
}

type HostnameGeneral struct {
	Alexa         string        `json:"alexa"`
	BaseIndicator BaseIndicator `json:"base_indicator"`
	Indicator     string        `json:"indicator"`
	PulseInfo     PulseInfo     `json:"pulse_info"`
	Sections      []string      `json:"sections"`
	Type          string        `json:"type"`
	TypeTitle     string        `json:"type_title"`
	Validation    []Validation  `json:"validation"`
	Whois         string        "json:`whois`"
}

type HostnameUrlList struct {
	ActualSize int        `json:"actual_size"`
	FullSize   int        `json:"full_size"`
	HasNext    bool       `json:"has_next"`
	Limit      int        `json:limit"`
	PageNum    int        `json:page_num"`
	Paged      bool       `json:"paged"`
	UrlList    []UrlEntry `json:"url_list"`
}

type FileGeneral struct {
	BaseIndicator BaseIndicator `json:"base_indicator"`
	Indicator     string        `json:"indicator"`
	PulseInfo     PulseInfo     `json:"pulse_info"`
	Sections      []string      `json:"sections"`
	Type          string        `json:"type"`
	TypeTitle     string        `json:"type_title"`
	Validation    []Validation  `json:"validation"`
}

type FileAnalysis struct {
	Analysis map[string]interface{} `json:"analysis"`
	Malware  map[string]interface{} `json:"malware"`
	PageType string                 `json:"page_type"`
}

type UrlGeneral struct {
	Alexa         string        `json:"alexa"`
	BaseIndicator BaseIndicator `json:"base_indicator"`
	Indicator     string        `json:"indicator"`
	PulseInfo     PulseInfo     `json:"pulse_info"`
	Sections      []string      `json:"sections"`
	Type          string        `json:"type"`
	TypeTitle     string        `json:"type_title"`
	Validation    []Validation  `json:"validation"`
	Whois         string        "json:`whois`"
}

type UrlUrlList struct {
	AreaCode      int        `json:"area_code"`
	Asn           string     `json:"asn"`
	Charset       int        `json:"charset"`
	City          string     `json:"city"`
	CityData      bool       `json:"city_data"`
	ContinentCode string     `json:"continet_code"`
	CountryCode   string     `json:"country_code"`
	CountryCode3  string     `json:"country_code3"`
	CountryName   string     `json:"country_name"`
	DmaCode       int        `json:"dma_code"`
	FlagTitle     string     `json:"flag_title"`
	FlagUrl       string     `json:"flag_url"`
	Latitude      float32    `json:"latitude"`
	Longitude     float32    `json:"longitude"`
	PostalCode    string     `json:"postal_code"`
	Region        string     `json:"region"`
	UrlList       []UrlEntry `json:"url_list"`
}

type Reference struct {
	ExternalSource string `json:"external_source"`
	Href           string `json:"href"`
	Name           string `json:"name"`
	Type           string `json:"type"`
}

type CveGeneral struct {
	BaseIndicator BaseIndicator     `json:"base_indicator"`
	Cve           string            `json:"cve"`
	Cvss          map[string]string `json:"cvss"`
	Cwe           string            `json:"cwe"`
	DateCreated   string            `json:"date_created"`
	DateModified  string            `json:"date_modified"`
	Description   string            `json:"description"`
	Indicator     string            `json:"indicator"`
	MitreUrl      string            `json:"mitre_url"`
	NvdUrl        string            `json:"nvd_url"`
	Products      []string          `json:"products"`
	PulseInfo     PulseInfo         `json:"pulse_info"`
	References    []Reference       `json:"references"`
	Sections      []string          `json:"sections"`
}

type NidsGeneral struct {
	BaseIndicator BaseIndicator `json:"base_indicator"`
	Category      string        `json:"category"`
	Cve           string        `json:"cve"`
	EventActivity interface{}   `json:"event_activity"`
	Indicator     string        `json:"indicator"`
	MalwareName   []string      `json:"malware_name"`
	Name          string        `json:"name"`
	PulseInfo     PulseInfo     `json:"pulse_info"`
	Sections      []string      `json:"sections"`
	Subcategory   string        `json:"subcategory"`
}

type CorrelationRuleGeneral struct {
	BaseIndicator BaseIndicator `json:"base_indicator"`
	Cve           string        `json:"cve"`
	Indicator     string        `json:"indicator"`
	PulseInfo     PulseInfo     `json:"pulse_info"`
	Sections      []string      `json:"sections"`
}
