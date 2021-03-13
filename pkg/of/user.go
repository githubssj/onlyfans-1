package of

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

// User describes an of user
type User struct {
	View         string `json:"view"`
	Avatar       string `json:"avatar"`
	AvatarThumbs struct {
		C50  string `json:"c50"`
		C144 string `json:"c144"`
	} `json:"avatarThumbs"`
	Header     string `json:"header"`
	HeaderSize struct {
		Width  int `json:"width"`
		Height int `json:"height"`
	} `json:"headerSize"`
	HeaderThumbs struct {
		W480 string `json:"w480"`
		W760 string `json:"w760"`
	} `json:"headerThumbs"`
	ID                   int     `json:"id"`
	Name                 string  `json:"name"`
	Username             string  `json:"username"`
	CanLookStory         bool    `json:"canLookStory"`
	CanCommentStory      bool    `json:"canCommentStory"`
	HasNotViewedStory    bool    `json:"hasNotViewedStory"`
	IsVerified           bool    `json:"isVerified"`
	CanPayInternal       bool    `json:"canPayInternal"`
	HasScheduledStream   bool    `json:"hasScheduledStream"`
	HasStream            bool    `json:"hasStream"`
	HasStories           bool    `json:"hasStories"`
	TipsEnabled          bool    `json:"tipsEnabled"`
	TipsTextEnabled      bool    `json:"tipsTextEnabled"`
	TipsMin              int     `json:"tipsMin"`
	TipsMax              int     `json:"tipsMax"`
	Bookmarked           bool    `json:"bookmarked"`
	CanBeBookmarked      bool    `json:"canBeBookmarked"`
	CanEarn              bool    `json:"canEarn"`
	CanAddSubscriber     bool    `json:"canAddSubscriber"`
	SubscribePrice       float64 `json:"subscribePrice"`
	IsPaywallRestriction bool    `json:"isPaywallRestriction"`
	ListsStates          []struct {
		ID         int    `json:"id"`
		Type       string `json:"type"`
		Name       string `json:"name"`
		HasUser    bool   `json:"hasUser"`
		CanAddUser bool   `json:"canAddUser"`
	} `json:"listsStates"`
	IsMuted                 bool        `json:"isMuted"`
	IsRestricted            bool        `json:"isRestricted"`
	CanRestrict             bool        `json:"canRestrict"`
	SubscribedBy            bool        `json:"subscribedBy"`
	SubscribedByExpire      bool        `json:"subscribedByExpire"`
	SubscribedByExpireDate  time.Time   `json:"subscribedByExpireDate"`
	SubscribedByAutoprolong bool        `json:"subscribedByAutoprolong"`
	SubscribedIsExpiredNow  bool        `json:"subscribedIsExpiredNow"`
	CurrentSubscribePrice   float64     `json:"currentSubscribePrice"`
	SubscribedOn            bool        `json:"subscribedOn"`
	SubscribedOnExpiredNow  interface{} `json:"subscribedOnExpiredNow"`
	SubscribedOnDuration    interface{} `json:"subscribedOnDuration"`
	JoinDate                time.Time   `json:"joinDate"`
	IsReferrerAllowed       bool        `json:"isReferrerAllowed"`
	About                   string      `json:"about"`
	RawAbout                string      `json:"rawAbout"`
	Website                 string      `json:"website"`
	Wishlist                string      `json:"wishlist"`
	Location                string      `json:"location"`
	PostsCount              int         `json:"postsCount"`
	ArchivedPostsCount      int         `json:"archivedPostsCount"`
	PhotosCount             int         `json:"photosCount"`
	VideosCount             int         `json:"videosCount"`
	AudiosCount             int         `json:"audiosCount"`
	MediasCount             int         `json:"mediasCount"`
	LastSeen                time.Time   `json:"lastSeen"`
	FavoritesCount          int         `json:"favoritesCount"`
	FavoritedCount          int         `json:"favoritedCount"`
	ShowPostsInFeed         bool        `json:"showPostsInFeed"`
	CanReceiveChatMessage   bool        `json:"canReceiveChatMessage"`
	IsPerformer             bool        `json:"isPerformer"`
	IsRealPerformer         bool        `json:"isRealPerformer"`
	IsSpotifyConnected      bool        `json:"isSpotifyConnected"`
	SubscribersCount        interface{} `json:"subscribersCount"`
	HasPinnedPosts          bool        `json:"hasPinnedPosts"`
	CanChat                 bool        `json:"canChat"`
	CallPrice               float64     `json:"callPrice"`
	IsPrivateRestriction    bool        `json:"isPrivateRestriction"`
	ShowSubscribersCount    bool        `json:"showSubscribersCount"`
	ShowMediaCount          bool        `json:"showMediaCount"`
	SubscribedByData        struct {
		Price              float64     `json:"price"`
		NewPrice           float64     `json:"newPrice"`
		RegularPrice       float64     `json:"regularPrice"`
		SubscribePrice     float64     `json:"subscribePrice"`
		DiscountPercent    int         `json:"discountPercent"`
		DiscountPeriod     int         `json:"discountPeriod"`
		SubscribeAt        time.Time   `json:"subscribeAt"`
		ExpiredAt          time.Time   `json:"expiredAt"`
		RenewedAt          time.Time   `json:"renewedAt"`
		DiscountFinishedAt interface{} `json:"discountFinishedAt"`
		DiscountStartedAt  interface{} `json:"discountStartedAt"`
		Status             interface{} `json:"status"`
		IsMuted            bool        `json:"isMuted"`
		UnsubscribeReason  string      `json:"unsubscribeReason"`
		Duration           string      `json:"duration"`
		ShowPostsInFeed    bool        `json:"showPostsInFeed"`
		Subscribes         []struct {
			ID           int         `json:"id"`
			UserID       int         `json:"userId"`
			SubscriberID int         `json:"subscriberId"`
			Date         time.Time   `json:"date"`
			Duration     int         `json:"duration"`
			StartDate    time.Time   `json:"startDate"`
			ExpireDate   time.Time   `json:"expireDate"`
			CancelDate   interface{} `json:"cancelDate"`
			Price        float64     `json:"price"`
			RegularPrice float64     `json:"regularPrice"`
			Discount     int         `json:"discount"`
			EarningID    int         `json:"earningId"`
			Action       string      `json:"action"`
			Type         string      `json:"type"`
			OfferStart   interface{} `json:"offerStart"`
			OfferEnd     interface{} `json:"offerEnd"`
			IsCurrent    bool        `json:"isCurrent"`
		} `json:"subscribes"`
	} `json:"subscribedByData"`
	SubscribedOnData   interface{} `json:"subscribedOnData"`
	CanPromotion       bool        `json:"canPromotion"`
	CanCreatePromotion bool        `json:"canCreatePromotion"`
	CanCreateTrial     bool        `json:"canCreateTrial"`
	IsAdultContent     bool        `json:"isAdultContent"`
	CanTrialSend       bool        `json:"canTrialSend"`
	HasFriends         bool        `json:"hasFriends"`
	HasLinks           bool        `json:"hasLinks"`
	Promotion          interface{} `json:"promotion"`
	PromotionExpired   interface{} `json:"promotionExpired"`
	IsFriend           bool        `json:"isFriend"`
	IsBlocked          bool        `json:"isBlocked"`
}

// GetUser looks up an of user by their username
func (c *Client) GetUser(ctx context.Context, username string) (*User, error) {
	path := fmt.Sprintf("/users/%s?app-token=%s", username, c.Token)
	b, err := c.Do(ctx, http.MethodGet, path, nil, http.StatusOK)
	if err != nil {
		return nil, err
	}

	user := &User{}
	err = json.Unmarshal(b, user)
	if err != nil {
		return nil, err
	}

	return user, nil
}
