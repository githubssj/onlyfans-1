package of

import (
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

// Client describes an onlyfans client
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
func (c *Client) Do(ctx context.Context, method, path string, body io.Reader, expectedStatus int) ([]byte, error) {
	req, err := http.NewRequestWithContext(ctx, method, fmt.Sprintf("%s/%s", c.BaseURL, path), body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Accept", fmt.Sprintf("%s, %s, %s", "application/json", "text/plain", "*/*"))
	req.Header.Add("User-Agent", c.UserAgent)
	req.Header.Add("Access-Token", c.Session)
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
	defer resp.Body.Close()

	if resp.StatusCode != expectedStatus {
		return nil, fmt.Errorf("unexpected status: %d", resp.StatusCode)
	}

	return ioutil.ReadAll(resp.Body)
}

// DownloadContent downloads a content
func (c *Client) DownloadContent(ctx context.Context, media []Media, name, saveDir string) {
	dir := strings.ReplaceAll(name, " ", "")
	for _, m := range media {
		source := getSource(m)
		if source == "" {
			return
		}

		req, err := http.NewRequest(http.MethodGet, source, nil)
		if err != nil {
			log.Println(err)
		}

		resp, err := c.Client.Do(req)
		if err != nil {
			log.Println(err)
		}
		defer resp.Body.Close()

		f := fmt.Sprintf("%d.%s", m.ID, getExtensionFromURL(source))
		err = SaveFile(saveDir, dir, f, resp.Body)
		if err != nil {
			log.Println(err)
		}
	}
}

func getExtensionFromURL(url string) string {
	u := url[:strings.Index(url, "?")+1]
	return u[strings.LastIndex(u, ".")+1 : strings.Index(u, "?")]
}

func getSource(m Media) string {
	if m.Source != nil && m.Source.FileSource != "" {
		return m.Source.FileSource
	}

	if m.Files != nil && m.Files.Source != nil && m.Files.Source.URL != "" {
		return m.Files.Source.URL
	}

	return ""
}
