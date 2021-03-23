package oauth

import "time"

type AccessToken struct {
	AccessToken string `json:"access_token"`
	UserId      int    `json:"user_id"`
	Expires     int    `json:"expires"`
}

func (at *AccessToken) IsExpired() bool {
	return time.Unix(int64(at.Expires), 0).UTC().Before(time.Now().UTC())
}
