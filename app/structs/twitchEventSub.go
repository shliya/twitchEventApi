package structs

type EventSubRespose struct {
	Subscription Subscription `json:"subscription"`
	Event        Event        `json:"event"`
	Challenge    string       `json:"challenge`
}

type SubscriptionRequset struct {
	Type      string    `json:"type"`
	Version   string    `json:"version"`
	Condition Condition `json:"condition"`
	Transport Transport `json:"transport"`
}

type SubscriptionList struct {
	Total int            `json:"total"`
	Data  []Subscription `json:"data"`
}

type Condition struct {
	BroadcasterUserId string `json:"broadcaster_user_id"`
	UserId            string `json:"_user_id"`
}

type Transport struct {
	Method   string `json:"method"`
	Callback string `json:"callback"`
	Secret   string `json:"secret"`
}

type Subscription struct {
	Id        string    `json:"id"`
	Status    string    `json:"status"`
	Type      string    `json:"type"`
	Version   string    `json:"version"`
	Cost      int       `json:"cost"`
	CreatedAt string    `json:"created_at"`
	Condition Condition `json:"condition"`
	Transport Transport `json:"transport"`
}

type Event struct {
	Id                   string `json:"id"`
	BroadcasterUserId    string `json:"broadcaster_user_id"`
	BroadcasterUserLogin string `json:"broadcaster_user_login"`
	Reward               Reward `json:"reward"`
	Cost                 int    `json:"cost"`
}

type Reward struct {
	Id     string `json:"id"`
	Title  string `json:"title"`
	Cost   int    `json:"cost"`
	Prompt string `json:"prompt"`
}

type TokenRes struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   string `json:"expires_in"`
	TokenType   string `json:"token_type"`
}
