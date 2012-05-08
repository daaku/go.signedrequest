// Package fbsr implements the standard Facebook signed_request structures.
package fbsr

import (
	"github.com/nshah/go.signedrequest"
	"time"
)

type Timestamp int64

type Page struct {
	ID    uint64 `json:"id,string"`
	Liked bool   `json:"liked"`
	Admin bool   `json:"admin"`
}

type Age struct {
	Min uint `json:"min,omitempty"`
}

type User struct {
	Country string `json:"country,omitempty"`
	Locale  string `json:"locale,omitempty"`
	Age     *Age   `json:"age,omitempty"`
}

type SignedRequest struct {
	Algorithm   string    `json:"algorithm"`
	ExpiresAt   Timestamp `json:"expires"`
	IssuedAt    Timestamp `json:"issued_at"`
	AccessToken string    `json:"oauth_token,omitempty"`
	Page        *Page     `json:"page,omitempty"`
	User        *User     `json:"user,omitempty"`
	UserID      uint64    `json:"user_id,string,omitempty"`
}

// Unmarshal a Facebook signed request.
func Unmarshal(data []byte, secret []byte) (*SignedRequest, error) {
	sr := &SignedRequest{}
	err := signedrequest.Unmarshal(data, secret, sr)
	if err != nil {
		return nil, err
	}
	return sr, err
}

func (t Timestamp) Time() time.Time {
	return time.Unix(int64(t), 0)
}
