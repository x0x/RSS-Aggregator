package main

import (
	"time"

	"RSS-Aggregator/internal/database"
	"github.com/google/uuid"
)

type User struct {
	Id         uuid.UUID `json:"id"`
	Created_at time.Time `json:"created_at"`
	Updated_at time.Time `json:"updated_at"`
	Name       string    `json:"name"`
	ApiKey     string    `json:"api_key"`
}

type Feed struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `json:"name"`
	Url       string    `json:"url"`
	UserID    uuid.UUID `json:"user_id"`
}

type FeedFollow struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	FeedId    uuid.UUID `json:"feed_id"`
	UserID    uuid.UUID `json:"user_id"`
}

func dbUsertoUser(dbUser database.User) User {
	return User{
		Id:         dbUser.ID,
		Created_at: dbUser.CreatedAt,
		Updated_at: dbUser.UpdatedAt,
		Name:       dbUser.Name,
		ApiKey:     dbUser.ApiKey,
	}
}

func dbFeedtoFeed(dbFeed database.Feed) Feed {
	return Feed{
		ID:        dbFeed.ID,
		CreatedAt: dbFeed.CreatedAt,
		UpdatedAt: dbFeed.UpdatedAt,
		Name:      dbFeed.Name,
		Url:       dbFeed.Url,
		UserID:    dbFeed.UserID,
	}
}

func dbFeedstoFeeds(dbFeeds []database.Feed) []Feed {
	feeds := []Feed{}
	for _, dbFeed := range dbFeeds {
		feeds = append(feeds, dbFeedtoFeed(dbFeed))
	}
	return feeds
}

func dbFeedFollowtoFeedFollow(dbFeedFollow database.FeedFollow) FeedFollow {
	return FeedFollow{
		ID:        dbFeedFollow.ID,
		CreatedAt: dbFeedFollow.CreatedAt,
		UpdatedAt: dbFeedFollow.UpdatedAt,
		FeedId:    dbFeedFollow.FeedID,
		UserID:    dbFeedFollow.UserID,
	}
}

func dbFeedFollowstoFeedFollows(dbFeedFollows []database.FeedFollow) []FeedFollow {
	feedFollows := []FeedFollow{}

	for _, dbFeedFollow := range dbFeedFollows {
		feedFollows = append(feedFollows, dbFeedFollowtoFeedFollow(dbFeedFollow))
	}
	return feedFollows
}
