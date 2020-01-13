package content

import (
	"playground/messaging"
	"playground/messaging/content/tabledata"
	"playground/messaging/content/unknown"
)

func GetContent(id messaging.ContentID) messaging.Content {
	switch id {
	case messaging.TableData:
		return tabledata.NewBuilder()
	default:
		return unknown.New()
	}
}
