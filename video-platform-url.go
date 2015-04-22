package videoPlatformUrl

type videoUrl string

func FindUrl(id string, provider string) videoUrl {

	var url string
	switch provider {
	case "dailymotion":
		url = "http://www.dailymotion.com/video/" + string(id)
	case "youtube":
		url = "https://www.youtube.com/watch?v=" + string(id)
	case "generic":
		url = id
	}

	return videoUrl(url)
}
