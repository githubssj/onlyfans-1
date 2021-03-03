package of

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

// Message describes a message
type Message struct {
	ResponseType string        `json:"responseType,omitempty"`
	Text         string        `json:"text,omitempty"`
	LockedText   bool          `json:"lockedText,omitempty"`
	IsFree       bool          `json:"isFree,omitempty"`
	Price        interface{}   `json:"price,omitempty"` //this fields type seems to switch occasionally..?
	IsMediaReady bool          `json:"isMediaReady,omitempty"`
	MediaCount   int           `json:"mediaCount,omitempty"`
	Media        []Media       `json:"media,omitempty"`
	Previews     []interface{} `json:"previews,omitempty"`
	IsTip        bool          `json:"isTip,omitempty"`
	FromUser     struct {
		ID   int    `json:"id,omitempty"`
		View string `json:"_view,omitempty"`
	} `json:"fromUser,omitempty"`
	IsFromQueue        bool      `json:"isFromQueue,omitempty"`
	QueueID            int       `json:"queueId,omitempty"`
	CanUnsendQueue     bool      `json:"canUnsendQueue,omitempty"`
	UnsendSecondsQueue int       `json:"unsendSecondsQueue,omitempty"`
	ID                 int64     `json:"id,omitempty"`
	IsOpened           bool      `json:"isOpened,omitempty"`
	IsNew              bool      `json:"isNew,omitempty"`
	CreatedAt          time.Time `json:"createdAt,omitempty"`
	ChangedAt          time.Time `json:"changedAt,omitempty"`
	CancelSeconds      int       `json:"cancelSeconds,omitempty"`
	IsLiked            bool      `json:"isLiked,omitempty"`
	CanPurchase        bool      `json:"canPurchase,omitempty"`
	CanPurchaseReason  string    `json:"canPurchaseReason,omitempty"`
}

type listMessagesResponse struct {
	Messages []Message `json:"list,omitempty"`
	HasMore  bool      `json:"hasMore,omitempty"`
}

// ListMessages lists messages from a user
func (c *Client) ListMessages(ctx context.Context, userID int) ([]Message, error) {
	hasMore := true
	offset := 0
	ms := make([]Message, 0)

	for hasMore {
		path := fmt.Sprintf("/chats/%d/messages?limit=1000&offset=%d&order=desc&skip_users=all&skip_users_dups=1&app-token=%s", userID, offset, c.Token)
		b, err := c.Do(ctx, http.MethodGet, path, nil, http.StatusOK)
		if err != nil {
			return nil, err
		}

		r := listMessagesResponse{}
		err = json.Unmarshal(b, &r)
		if err != nil {
			return nil, err
		}

		for _, m := range r.Messages {
			ms = append(ms, m)
		}

		b, err = json.Marshal(&ms)
		if err != nil {
			log.Fatal(err)
		}

		hasMore = r.HasMore
		offset++
	}

	return ms, nil
}
