package of

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

// Post describes a post
type Post struct {
	ResponseType    string      `json:"responseType,omitempty"`
	ID              int         `json:"id,omitempty"`
	PostedAt        time.Time   `json:"postedAt,omitempty"`
	PostedAtPrecise string      `json:"postedAtPrecise,omitempty"`
	ExpiredAt       interface{} `json:"expiredAt,omitempty"`
	Author          struct {
		View         string `json:"view,omitempty"`
		Avatar       string `json:"avatar,omitempty"`
		AvatarThumbs struct {
			C50  string `json:"c50,omitempty"`
			C144 string `json:"c144,omitempty"`
		} `json:"avatarThumbs,omitempty"`
		Header     string `json:"header,omitempty"`
		HeaderSize struct {
			Width  int `json:"width,omitempty"`
			Height int `json:"height,omitempty"`
		} `json:"headerSize,omitempty"`
		HeaderThumbs struct {
			W480 string `json:"w480,omitempty"`
			W760 string `json:"w760,omitempty"`
		} `json:"headerThumbs,omitempty"`
		ID                   int     `json:"id,omitempty"`
		Name                 string  `json:"name,omitempty"`
		Username             string  `json:"username,omitempty"`
		CanLookStory         bool    `json:"canLookStory,omitempty"`
		CanCommentStory      bool    `json:"canCommentStory,omitempty"`
		HasNotViewedStory    bool    `json:"hasNotViewedStory,omitempty"`
		IsVerified           bool    `json:"isVerified,omitempty"`
		CanPayInternal       bool    `json:"canPayInternal,omitempty"`
		HasScheduledStream   bool    `json:"hasScheduledStream,omitempty"`
		HasStream            bool    `json:"hasStream,omitempty"`
		HasStories           bool    `json:"hasStories,omitempty"`
		TipsEnabled          bool    `json:"tipsEnabled,omitempty"`
		TipsTextEnabled      bool    `json:"tipsTextEnabled,omitempty"`
		TipsMin              int     `json:"tipsMin,omitempty"`
		TipsMax              int     `json:"tipsMax,omitempty"`
		Bookmarked           bool    `json:"bookmarked,omitempty"`
		CanBeBookmarked      bool    `json:"canBeBookmarked,omitempty"`
		CanEarn              bool    `json:"canEarn,omitempty"`
		CanAddSubscriber     bool    `json:"canAddSubscriber,omitempty"`
		SubscribePrice       float64 `json:"subscribePrice,omitempty"`
		IsPaywallRestriction bool    `json:"isPaywallRestriction,omitempty"`
		ListsStates          []struct {
			ID         int    `json:"id,omitempty"`
			Type       string `json:"type,omitempty"`
			Name       string `json:"name,omitempty"`
			HasUser    bool   `json:"hasUser,omitempty"`
			CanAddUser bool   `json:"canAddUser,omitempty"`
		} `json:"listsStates,omitempty"`
		IsMuted                 bool        `json:"isMuted,omitempty"`
		IsRestricted            bool        `json:"isRestricted,omitempty"`
		CanRestrict             bool        `json:"canRestrict,omitempty"`
		SubscribedBy            bool        `json:"subscribedBy,omitempty"`
		SubscribedByExpire      bool        `json:"subscribedByExpire,omitempty"`
		SubscribedByExpireDate  time.Time   `json:"subscribedByExpireDate,omitempty"`
		SubscribedByAutoprolong bool        `json:"subscribedByAutoprolong,omitempty"`
		SubscribedIsExpiredNow  bool        `json:"subscribedIsExpiredNow,omitempty"`
		CurrentSubscribePrice   float64     `json:"currentSubscribePrice,omitempty"`
		SubscribedOn            bool        `json:"subscribedOn,omitempty"`
		SubscribedOnExpiredNow  interface{} `json:"subscribedOnExpiredNow,omitempty"`
		SubscribedOnDuration    interface{} `json:"subscribedOnDuration,omitempty"`
		ShowPostsInFeed         bool        `json:"showPostsInFeed,omitempty"`
		CanTrialSend            bool        `json:"canTrialSend,omitempty"`
	} `json:"author,omitempty"`
	Text               string        `json:"text,omitempty"`
	RawText            string        `json:"rawText,omitempty"`
	LockedText         bool          `json:"lockedText,omitempty"`
	IsFavorite         bool          `json:"isFavorite,omitempty"`
	CanReport          bool          `json:"canReport,omitempty"`
	CanDelete          bool          `json:"canDelete,omitempty"`
	CanComment         bool          `json:"canComment,omitempty"`
	CanEdit            bool          `json:"canEdit,omitempty"`
	IsPinned           bool          `json:"isPinned,omitempty"`
	FavoritesCount     int           `json:"favoritesCount,omitempty"`
	MediaCount         int           `json:"mediaCount,omitempty"`
	IsMediaReady       bool          `json:"isMediaReady,omitempty"`
	IsOpened           bool          `json:"isOpened,omitempty"`
	CanToggleFavorite  bool          `json:"canToggleFavorite,omitempty"`
	StreamID           interface{}   `json:"streamId,omitempty"`
	Price              interface{}   `json:"price,omitempty"`
	HasVoting          bool          `json:"hasVoting,omitempty"`
	IsAddedToBookmarks bool          `json:"isAddedToBookmarks,omitempty"`
	IsArchived         bool          `json:"isArchived,omitempty"`
	IsDeleted          bool          `json:"isDeleted,omitempty"`
	CommentsCount      int           `json:"commentsCount,omitempty"`
	MentionedUsers     []interface{} `json:"mentionedUsers,omitempty"`
	LinkedUsers        []interface{} `json:"linkedUsers,omitempty"`
	LinkedPosts        []interface{} `json:"linkedPosts,omitempty"`
	CanViewMedia       bool          `json:"canViewMedia,omitempty"`
	Media              []Media       `json:"media,omitempty"`
}

// ListPosts from a user
func (c *Client) ListPosts(ctx context.Context, userID string) ([]*Post, error) {
	path := fmt.Sprintf("/users/%s/posts?app-token=%s&limit=1000", userID, c.Token)
	resp, err := c.Do(ctx, http.MethodGet, path, nil, http.StatusOK)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	ps := make([]*Post, 0)
	err = json.Unmarshal(b, &ps)
	if err != nil {
		return nil, err
	}

	return ps, nil
}

// ListArchivedPosts list archived posts from a user
func (c *Client) ListArchivedPosts(ctx context.Context, userID int) ([]*Post, error) {
	path := fmt.Sprintf("/users/%d/posts/archived?app-token=%s", userID, c.Token)
	resp, err := c.Do(ctx, http.MethodGet, path, nil, http.StatusOK)
	if err != nil {
		return nil, err
	}

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	ps := make([]*Post, 0)
	err = json.Unmarshal(b, &ps)
	if err != nil {
		return nil, err
	}

	return ps, nil
}
