package twitter

import (
	"fmt"
	"net/url"
	"strconv"
)

func Tweet(text string) {
	api := getTwitterAPI()

	tweet, err := api.PostTweet(text, nil)

	if err != nil {
		fmt.Printf("failed to tweet: %v\n", err)
	}

	fmt.Println(tweet.Text)
}

func GetSearch(word string) {
	api := getTwitterAPI()

	searchResult, err := api.GetSearch(word, nil)

	if err != nil {
		fmt.Printf("failed to get search: %v\n", err)
	}

	for _, tweet := range searchResult.Statuses {
		fmt.Println(tweet.FullText)
	}
}

func GetHomeTimeline(quantity int) {
	api := getTwitterAPI()

	v := url.Values{}
	s := strconv.Itoa(quantity)

	v.Set("count", s)
	tweets, err := api.GetHomeTimeline(v)

	if err != nil {
		fmt.Printf("failed to get home timeline: %v\n", err)
	}

	for i, tweet := range tweets {
		fmt.Printf("%d, tweetname:%v\n tweet: %v\n", i, tweet.User.Name, tweet.FullText)
	}
}
