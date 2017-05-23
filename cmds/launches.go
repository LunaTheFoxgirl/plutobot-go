package cmds

type Failure struct {
	Status  string `json:"status"`
	Message string `json:"msg"`
}

type Agency struct {
	ID           int      `json:"id"`
	Name         string   `json:"name"`
	Abbriviation string   `json:"abbrev"`
	Type         int      `json:"type"`
	CountryCode  string   `json:"countryCode"`
	WikiURL      string   `json:"wikiURL"`
	InfoURL      string   `json:"infoURL"`
	InfoURLs     []string `json:"infoURLs"`
}

type AgencyType struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type Calendar struct {
	Offset int `json:"offset"`
	Count  int `json:"count"`
	Total  int `json:"total"`
}

type EventType struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type Launch struct {
	Offset   int `json:"offset"`
	Count    int `json:"count"`
	Total    int `json:"total"`
	Launches []struct {
		ID          int       `json:"id"`
		Name        string    `json:"name"`
		Windowstart string    `json:"windowstart"`
		Windowend   string    `json:"windowend"`
		Net         string    `json:"net"`
		Wsstamp     int       `json:"wsstamp"`
		Westamp     int       `json:"westamp"`
		Netstamp    int       `json:"netstamp"`
		Isostart    string    `json:"isostart"`
		Isoend      string    `json:"isoend"`
		Isonet      string    `json:"isonet"`
		Status      int       `json:"status"`
		Inhold      int       `json:"inhold"`
		Tbdtime     int       `json:"tbdtime"`
		VidURLs     []string  `json:"vidURLs"` //depricated??
		VidURL      string    `json:"vidURL"`
		InfoURLs    []string  `json:"infoURLs"`
		InfoURL     string    `json:"infoURL"`
		Holdreason  string    `json:"holdreason"`
		Failreason  string    `json:"failreason"` //depricated??
		Tbddate     int       `json:"tbddate"`
		Probability int       `json:"probability"`
		Hashtag     string    `json:"hashtag"`
		Location    Location  `json:"location"`
		Rocket      Rocket    `json:"rocket"`
		Missions    []Mission `json:"missions"`
	} `json:"launches"`
}

type LaunchEvent struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	RelativeTime int    `json:"relativeTime"`
	Type         int    `json:"type"`
	Duration     int    `json:"duration"`
	Description  string `json:"description"`
	ParentID     int    `json:"parentid"`
}

type LaunchStatus struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type Location struct {
	ID          int      `json:"id"`
	Name        string   `json:"name"`
	CountryCode string   `json:"countryCode"`
	WikiURL     string   `json:"wikiURL"`
	InfoURL     string   `json:"infoURL"`
	InfoURLs    []string `json:"infoURLs"`
}

type Rocket struct {
	ID            int      `json:"id"`
	Name          string   `json:"name"`
	Configuration string   `json:"configuration"`
	Familyname    string   `json:"familyname"`
	Agencies      []Agency `json:"agencies"`
	WikiURL       string   `json:"wikiURL"`
	InfoURLs      []string `json:"infoURLs"` //depricated??
	ImageURL      string   `json:"imageURL"`
	ImageSizes    []int    `json:"imageSizes"`
}

type Pad []struct {
	ID        int      `json:"id"`
	Name      string   `json:"name"`
	InfoURL   string   `json:"infoURL"` //depricated??
	WikiURL   string   `json:"wikiURL"`
	MapURL    string   `json:"mapURL"`
	Latitude  float64  `json:"latitude"`
	Longitude float64  `json:"longitude"`
	Agencies  []Agency `json:"agencies"`
}

type Mission struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type MissionEvent struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	RelativeTime int    `json:"relativeTime"`
	Type         int    `json:"type"`
	Duration     int    `json:"duration"`
	Description  string `json:"description"`
	ParentID     int    `json:"parentid"`
}

type MissionType struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}
