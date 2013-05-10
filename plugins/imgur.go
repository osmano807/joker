package plugins

import . "github.com/osmano807/joker/interfaces"

func Imgur(il *InputLine) (ol *OutputLine) {
	ol = &OutputLine{}

	ol.ChannelId = il.ChannelId

	if il.URL.Host != "i.imgur.com" {
		ol.Result = NO_CHANGE
		return
	}

	ol.Result = NEW_STOREID
	ol.StoreId = "http://" + "Joker/" + il.URL.Host + "/"

	return
}
