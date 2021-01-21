package of

import (
	"context"
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

const (
	// RFC2822 time format
	RFC2822 = "Mon, 02 Jan 2006 15:04:05 -0700"
)

// Client describes an of client
type Client struct {
	Client    *http.Client
	Token     string
	Session   string
	UserAgent string
	AuthID    string
	BaseURL   string
}

// NewClient returns a new client
func NewClient(token, session, userAgent, authID string) *Client {
	return &Client{
		Client:    http.DefaultClient,
		Token:     token,
		Session:   session,
		UserAgent: userAgent,
		AuthID:    authID,
		BaseURL:   "https://onlyfans.com/api2/v2/",
	}
}

// Do makes an api call
func (c *Client) Do(ctx context.Context, method, path string, body io.Reader, expectedStatus int) (*http.Response, error) {
	req, err := http.NewRequestWithContext(ctx, method, fmt.Sprintf("%s/%s", c.BaseURL, path), body)
	if err != nil {
		return nil, err
	}

	t := fmt.Sprintf("%d", time.Now().UTC().Unix()*1000-301000)
	a := []byte(strings.Join([]string{c.Session, t, path, c.UserAgent, "onlyfans"}, "\n"))
	h := sha1.New()
	h.Write(a)
	sha := hex.EncodeToString(h.Sum(nil))
	req.Header.Add("Accept", fmt.Sprintf("%s, %s, %s", "application/json", "text/plain", "*/*"))
	req.Header.Add("User-Agent", c.UserAgent)
	req.Header.Add("Access-Token", c.Session)
	req.Header.Add("Time", t)
	req.Header.Add("Sign", sha)
	req.AddCookie(&http.Cookie{
		Name:  "sess",
		Value: c.Session,
	})
	req.AddCookie(&http.Cookie{
		Name:  "auth_id",
		Value: c.AuthID,
	})

	resp, err := c.Client.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != expectedStatus {
		return nil, fmt.Errorf("unexpected status: %d", resp.StatusCode)
	}

	return resp, nil
}

// DownloadFile downloads a file
func (c *Client) DownloadFile(ctx context.Context, Media interface{}, name, base string) error {
	url := ""
	id := 0
	switch Media.(type) {
	case PostMedia:
		p, ok := Media.(PostMedia)
		if !ok {
			return fmt.Errorf("invalid post media")
		}

		url = p.Source.Source
		id = int(p.ID)
	case PhotoMedia:
		p, ok := Media.(PhotoMedia)
		if !ok {
			return fmt.Errorf("invalid photo media")
		}

		url = p.Source.Source
		id = int(p.ID)
	case VideoMedia:
		p, ok := Media.(VideoMedia)
		if !ok {
			return fmt.Errorf("invalid video media")
		}

		url = p.Source.Source
		id = int(p.ID)
	default:
		return fmt.Errorf("unable to retrieve media url")
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return err
	}

	resp, err := c.Client.Do(req)
	if err != nil {
		return err
	}

	dir := strings.ReplaceAll(name, " ", "")
	f := fmt.Sprintf("%d.%s", id, getExtensionFromURL(url))
	err = SaveFile(base, dir, f, resp.Body)
	if err != nil {
		return err
	}

	return nil
}

func getExtensionFromURL(url string) string {
	u := url[:strings.Index(url, "?")+1]
	return u[strings.LastIndex(u, ".")+1 : strings.Index(u, "?")]
}
