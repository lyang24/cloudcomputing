package main

import "fmt"

type Features struct{
  Id    string `json:"id"`
  Description string `json:"description"`
  url     string `json:"url"`
  IamgeUrl  string `json:"image_url"`
}

type Action String

const(
  CLICK_ACTION Action="click"
  VIEW_ACTION Action="view"
)

type InteractionRequest struct {
  Id String `json:"id"`
  Action Action `json:"Action"`
}

type Interaction struct {
  View uint64 `json:"View"`
  Click uint64 `json:"click"`
}


func main() {
  fmt.Println()
}
