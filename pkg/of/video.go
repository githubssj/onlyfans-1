package of

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// Video describes a video
type Video struct {
	Author struct {
		View string `json:"_view"`
		ID   int64  `json:"id"`
	} `json:"author"`
	CanComment         bool          `json:"canComment"`
	CanDelete          bool          `json:"canDelete"`
	CanEdit            bool          `json:"canEdit"`
	CanReport          bool          `json:"canReport"`
	CanToggleFavorite  bool          `json:"canToggleFavorite"`
	CanViewMedia       bool          `json:"canViewMedia"`
	CanVote            bool          `json:"canVote"`
	CommentsCount      int64         `json:"commentsCount"`
	ExpiredAt          interface{}   `json:"expiredAt"`
	FavoritesCount     int64         `json:"favoritesCount"`
	HasVoting          bool          `json:"hasVoting"`
	ID                 int64         `json:"id"`
	IsAddedToBookmarks bool          `json:"isAddedToBookmarks"`
	IsArchived         bool          `json:"isArchived"`
	IsDeleted          bool          `json:"isDeleted"`
	IsFavorite         bool          `json:"isFavorite"`
	IsMediaReady       bool          `json:"isMediaReady"`
	IsOpened           bool          `json:"isOpened"`
	IsPinned           bool          `json:"isPinned"`
	LinkedPosts        []interface{} `json:"linkedPosts"`
	LinkedUsers        []interface{} `json:"linkedUsers"`
	LockedText         bool          `json:"lockedText"`
	Media              []Media       `json:"media,omitempty"`
	MediaCount         int64         `json:"mediaCount"`
	MentionedUsers     []interface{} `json:"mentionedUsers"`
	PostedAt           string        `json:"postedAt"`
	PostedAtPrecise    string        `json:"postedAtPrecise"`
	Preview            []interface{} `json:"preview"`
	Price              interface{}   `json:"price"`
	RawText            string        `json:"rawText"`
	ResponseType       string        `json:"responseType"`
	StreamID           interface{}   `json:"streamId"`
	Text               string        `json:"text"`
}

// ListVideos from a user
func (c *Client) ListVideos(ctx context.Context, userID string) ([]*Video, error) {
	path := fmt.Sprintf("/users/%s/posts/videos?limit=1000&order=publish_date_desc&skip_users=all&skip_users_dups=1&app-token=%s", userID, c.Token)
	resp, err := c.Do(ctx, http.MethodGet, path, nil, http.StatusOK)
	if err != nil {
		return nil, err
	}

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	vs := []*Video{}
	err = json.Unmarshal(b, &vs)
	if err != nil {
		return nil, err
	}

	return vs, nil
}
