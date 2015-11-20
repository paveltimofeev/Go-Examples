package vk_api

import "github.com/ptimofeev/vk_api"

var api vk_api.Api
var owner_id = "-userId"


Login()
GetVideos()
GetAudios()


func Login(){

  err := api.LoginAuth(
    "email/phone",
    "pass",
    "3087104", // client id
    "wall,video,offline", // scope (permissions)
  )

  if err != nil {
    panic(err)
  }
}

func GetVideos() string{

  params := make(map[string]string)
  params["owner_id"] = owner_id
  params["count"] = "100"
  params["offset"] = "0"

  return Invoke("video.get", nil)
}

func GetAudios() string{

  params := make(map[string]string)
  params["owner_id"] = owner_id
  params["count"] = "100"
  params["offset"] = "0"

  return Invoke("audio.get", nil)
}

func Invoke(req string, params map[string]string) string {

  strResp, err := api.Request(req, params)
    if err != nil {
        panic(err)
    }
  log.Println(strResp)
  return strResp
}
