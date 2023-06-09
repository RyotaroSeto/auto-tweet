package main

import "auto-tweet/twitter"

func main() {
	twitter.Tweet("こんにちは")
	twitter.GetSearch("こんにちは")
	twitter.GetHomeTimeline(10)
}
