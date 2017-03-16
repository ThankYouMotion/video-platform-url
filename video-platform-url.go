package videoPlatformURL

import (
	"net/url"
	"strings"
)

var (
	facebookVideoURL = "https://www.facebook.com/plugins/video.php?href="
)

//FindURL Create a well formated URL for different providers
func FindURL(str string, provider string) string {

	var url string
	switch provider {
	case "dailymotion":
		url = "http://www.dailymotion.com/video/" + string(str)
	case "youtube":
		url = "https://www.youtube.com/watch?v=" + string(str)
	case "proxy":
		url = analyzeDomain(str)
	case "digiteka":
		// if it's an URL
		if strings.Contains(str, "://") {
			url = string(str)
		} else {
			url = "https://www.ultimedia.com/deliver/generic/iframe/mdtk/01836272/src/" + string(str)
		}
	default:
		url = string(str)
	}

	return url
}

func analyzeDomain(str string) string {
	if len(str) > len(facebookVideoURL) && str[0:len(facebookVideoURL)] == facebookVideoURL {
		videoURL, err := url.QueryUnescape(str[len(facebookVideoURL):len(str)])
		if err != nil {
			return str[len(facebookVideoURL):len(str)]
		}
		return videoURL
	}
	return str
}
