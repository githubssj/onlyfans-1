package of

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// Photo describes a photo
type Photo struct {
	Author struct {
		View string `json:"_view,omitempty"`
		ID   int64  `json:"id,omitempty"`
	} `json:"author,omitempty"`
	CanComment         bool          `json:"canComment,omitempty"`
	CanDelete          bool          `json:"canDelete,omitempty"`
	CanEdit            bool          `json:"canEdit,omitempty"`
	CanReport          bool          `json:"canReport,omitempty"`
	CanToggleFavorite  bool          `json:"canToggleFavorite,omitempty"`
	CanViewMedia       bool          `json:"canViewMedia,omitempty"`
	CommentsCount      int64         `json:"commentsCount,omitempty"`
	ExpiredAt          interface{}   `json:"expiredAt,omitempty"`
	FavoritesCount     int64         `json:"favoritesCount,omitempty"`
	HasVoting          bool          `json:"hasVoting,omitempty"`
	ID                 int64         `json:"id,omitempty"`
	IsAddedToBookmarks bool          `json:"isAddedToBookmarks,omitempty"`
	IsArchived         bool          `json:"isArchived,omitempty"`
	IsDeleted          bool          `json:"isDeleted,omitempty"`
	IsFavorite         bool          `json:"isFavorite,omitempty"`
	IsMediaReady       bool          `json:"isMediaReady,omitempty"`
	IsOpened           bool          `json:"isOpened,omitempty"`
	IsPinned           bool          `json:"isPinned,omitempty"`
	LinkedPosts        []interface{} `json:"linkedPosts,omitempty"`
	LinkedUsers        []interface{} `json:"linkedUsers,omitempty"`
	LockedText         bool          `json:"lockedText,omitempty"`
	PhotoMedia         []PhotoMedia  `json:"media,omitempty"`
	MediaCount         int64         `json:"mediaCount,omitempty"`
	MentionedUsers     []interface{} `json:"mentionedUsers,omitempty"`
	PostedAt           string        `json:"postedAt,omitempty"`
	PostedAtPrecise    string        `json:"postedAtPrecise,omitempty"`
	Preview            []interface{} `json:"preview,omitempty"`
	Price              interface{}   `json:"price,omitempty"`
	RawText            string        `json:"rawText,omitempty"`
	ResponseType       string        `json:"responseType,omitempty"`
	StreamID           interface{}   `json:"streamId,omitempty"`
	Text               string        `json:"text,omitempty"`
}

// PhotoMedia describes photo media
type PhotoMedia struct {
	CanView          bool   `json:"canView,omitempty"`
	ConvertedToVideo bool   `json:"convertedToVideo,omitempty"`
	CreatedAt        string `json:"createdAt,omitempty"`
	Files            struct {
		Preview struct {
			URL string `json:"url,omitempty"`
		} `json:"preview,omitempty"`
	} `json:"files,omitempty"`
	Full     string `json:"full,omitempty"`
	HasError bool   `json:"hasError,omitempty"`
	ID       int64  `json:"id,omitempty"`
	Info     struct {
		Preview struct {
			Height int64 `json:"height,omitempty"`
			Size   int64 `json:"size,omitempty"`
			Width  int64 `json:"width,omitempty"`
		} `json:"preview,omitempty"`
		Source struct {
			Duration int64  `json:"duration,omitempty"`
			Height   int64  `json:"height,omitempty"`
			Size     int64  `json:"size,omitempty"`
			Source   string `json:"source,omitempty"`
			Width    int64  `json:"width,omitempty"`
		} `json:"source,omitempty"`
	} `json:"info,omitempty"`
	Preview string `json:"preview,omitempty"`
	Source  struct {
		Duration int64  `json:"duration,omitempty"`
		Height   int64  `json:"height,omitempty"`
		Size     int64  `json:"size,omitempty"`
		Source   string `json:"source,omitempty"`
		Width    int64  `json:"width,omitempty"`
	} `json:"source,omitempty"`
	SquarePreview string `json:"squarePreview,omitempty"`
	Thumb         string `json:"thumb,omitempty"`
	Type          string `json:"type,omitempty"`
	VideoSources  struct {
		Two40   interface{} `json:"240,omitempty"`
		Seven20 interface{} `json:"720,omitempty"`
	} `json:"videoSources,omitempty"`
}

// ListPhotos from a user
func (c *Client) ListPhotos(ctx context.Context, userID string) ([]*Photo, error) {
	path := fmt.Sprintf("/users/%s/posts/photos?limit=1000&order=publish_date_desc&skip_users=all&skip_users_dups=1&app-token=%s", userID, c.Token)
	resp, err := c.Do(ctx, http.MethodGet, path, nil, http.StatusOK)
	if err != nil {
		return nil, err
	}

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	ps := []*Photo{}
	err = json.Unmarshal(b, &ps)
	if err != nil {
		return nil, err
	}

	return ps, nil
}
