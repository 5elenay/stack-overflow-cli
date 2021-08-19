package main

type StackOverflowItem struct {
	Link  string   `json:"link"`
	Title string   `json:"title"`
	Tags  []string `json:"tags"`
}

type StackOverflowItemList struct {
	Items []StackOverflowItem `json:"items"`
}
