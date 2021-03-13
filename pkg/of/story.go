package of

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

// ListHighlightResponse is the response from listing highlights
type ListHighlightResponse struct {
	ID           int       `json:"id"`
	UserID       int       `json:"userId"`
	Title        string    `json:"title"`
	CoverStoryID int       `json:"coverStoryId"`
	Cover        string    `json:"cover"`
	StoriesCount int       `json:"storiesCount"`
	CreatedAt    time.Time `json:"createdAt"`
}

// Highlight describes a highlight
type Highlight struct {
	ID           int       `json:"id"`
	UserID       int       `json:"userId"`
	Title        string    `json:"title"`
	CoverStoryID int       `json:"coverStoryId"`
	Cover        string    `json:"cover"`
	StoriesCount int       `json:"storiesCount"`
	CreatedAt    time.Time `json:"createdAt"`
	Stories      []struct {
		ID                int           `json:"id"`
		UserID            int           `json:"userId"`
		CreatedAt         time.Time     `json:"createdAt"`
		ExpiredAt         time.Time     `json:"expiredAt"`
		IsReady           bool          `json:"isReady"`
		ViewersCount      int           `json:"viewersCount"`
		Viewers           []interface{} `json:"viewers"`
		CanLike           bool          `json:"canLike"`
		MediaCount        int           `json:"mediaCount"`
		IsWatched         bool          `json:"isWatched"`
		IsLiked           bool          `json:"isLiked"`
		CanDelete         bool          `json:"canDelete"`
		IsHighlightCover  bool          `json:"isHighlightCover"`
		IsLastInHighlight bool          `json:"isLastInHighlight"`
		Media             []Media       `json:"media"`
	} `json:"stories"`
}

// ListHighlights from a user
func (c *Onlyfans) ListHighlights(ctx context.Context, userID int) ([]*Highlight, error) {
	path := fmt.Sprintf("/users/%d/stories/highlights?unf=1&app-token=%s&limit=10000", userID, c.Token)
	b, err := c.Do(ctx, http.MethodGet, path, nil, http.StatusOK)
	if err != nil {
		log.Fatal(err)
	}

	lhr := []ListHighlightResponse{}
	err = json.Unmarshal(b, &lhr)
	if err != nil {
		return nil, err
	}

	hs := make([]*Highlight, 0)
	for _, r := range lhr {
		h, err := c.GetHighlight(ctx, userID, r.ID)
		if err != nil {
			return nil, err
		}

		hs = append(hs, h)
	}

	return hs, nil
}

// GetHighlight from a user
func (c *Onlyfans) GetHighlight(ctx context.Context, userID, storyID int) (*Highlight, error) {
	path := fmt.Sprintf("/stories/highlights/%d?unf=1&app-token=%s", storyID, c.Token)
	b, err := c.Do(ctx, http.MethodGet, path, nil, http.StatusOK)
	if err != nil {
		return nil, err
	}

	h := &Highlight{}
	err = json.Unmarshal(b, h)
	if err != nil {
		return nil, err
	}

	return h, nil
}
