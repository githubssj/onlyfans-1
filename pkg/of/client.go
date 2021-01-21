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

// DownloadContent downloads a content
func (c *Client) DownloadContent(ctx context.Context, req []Media, name, saveDir string) error {
	dir := strings.ReplaceAll(name, " ", "")
	for _, m := range req {
		req, err := http.NewRequest(http.MethodGet, m.Source.Source, nil)
		if err != nil {
			return err
		}

		resp, err := c.Client.Do(req)
		if err != nil {
			return err
		}

		f := fmt.Sprintf("%d.%s", m.ID, getExtensionFromURL(m.Source.Source))
		err = SaveFile(saveDir, dir, f, resp.Body)
		if err != nil {
			return err

		}
	}

	return nil
}

func getExtensionFromURL(url string) string {
	u := url[:strings.Index(url, "?")+1]
	return u[strings.LastIndex(u, ".")+1 : strings.Index(u, "?")]
}
