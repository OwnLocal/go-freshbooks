package freshbooks

import (
	"encoding/xml"
	"time"
)

const (
	DATE_FORMAT = "2006-01-02 15:04:05"
)

type Date struct {
	time.Time
}

func (date *Date) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	// Encoding tokens manually to inject attributes and comments
	e.EncodeElement([]byte(date.Format(DATE_FORMAT)), start)
	return nil
}

func (date *Date) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var v string
	d.DecodeElement(&v, &start)
	parse, err := time.Parse(DATE_FORMAT, v)
	if err != nil {
		return nil
	}
	*date = Date{parse}
	return nil
}
