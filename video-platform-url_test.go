package videoPlatformURL

import "testing"

func TestAnalyzeDomain(t *testing.T) {
	facebookFullURL := "https://www.facebook.com/plugins/video.php?href=https%3A%2F%2Fwww.facebook.com%2FMTV%2Fvideos%2F10153999880236701%2F&show_text=0&width=560"
	expected := "https://www.facebook.com/MTV/videos/10153999880236701/&show_text=0&width=560"
	actual := analyzeDomain(facebookFullURL)
	if actual != expected {
		t.Errorf("Test failed, expected: '%s', got:  '%s'", expected, actual)
	}
}

func TestFindURL(t *testing.T) {

	var expected string
	var actual string

	youtubeID := "T8GRwkZDCBU"
	expected = "https://www.youtube.com/watch?v=T8GRwkZDCBU"
	actual = FindURL(youtubeID, "youtube")
	if actual != expected {
		t.Errorf("Test failed, expected: '%s', got:  '%s'", expected, actual)
	}

	dailymotionID := "x4qfa0s"
	expected = "http://www.dailymotion.com/video/x4qfa0s"
	actual = FindURL(dailymotionID, "dailymotion")
	if actual != expected {
		t.Errorf("Test failed, expected: '%s', got:  '%s'", expected, actual)
	}

	digitekaID := "8mruqx"
	expected = "https://www.ultimedia.com/deliver/generic/iframe/mdtk/01836272/src/8mruqx"
	actual = FindURL(digitekaID, "digiteka")
	if actual != expected {
		t.Errorf("Test failed, expected: '%s', got:  '%s'", expected, actual)
	}

	proxyURL := "http://www.20minutes.fr/insolite/1914323-20160828-harcele-trolls-zoo-cincinnati-decide-fermer-compte-twitter"
	expected = "http://www.20minutes.fr/insolite/1914323-20160828-harcele-trolls-zoo-cincinnati-decide-fermer-compte-twitter"
	actual = FindURL(proxyURL, "proxy")
	if actual != expected {
		t.Errorf("Test failed, expected: '%s', got:  '%s'", expected, actual)
	}

	proxyShortURL := "http://www.20minutes.fr/1914323"
	expected = "http://www.20minutes.fr/1914323"
	actual = FindURL(proxyShortURL, "proxy")
	if actual != expected {
		t.Errorf("Test failed, expected: '%s', got:  '%s'", expected, actual)
	}
}
