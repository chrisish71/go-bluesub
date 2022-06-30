package util

import (
	"github.com/chrisish71/go-bluesub"
	"github.com/chrisish71/go-bluesub/blue/model"
	"math"
	"strconv"
	"strings"
	"time"
)

const (
	VTT    string = "vtt"
	STL    string = "stl"
	TTML   string = "ttml"
	LEFT   string = "LEFT"
	RIGHT  string = "RIGHT"
	CENTER string = "CENTER"
)

func BlueModelToAsticodeModel(blueSubtitles []model.Subtitle, format string) *astisub.Subtitles {
	var items []*astisub.Item
	asticodeSubtitles := astisub.NewSubtitles()
	var vttRegions = make(map[string]*astisub.Region)
	for _, blueSubtitle := range blueSubtitles {
		var item = new(astisub.Item)
		item.StartAt = durationOf(blueSubtitle.StartTime)
		item.EndAt = durationOf(blueSubtitle.EndTime)
		item.InlineStyle = subtitleStyle(blueSubtitle, format)
		for _, blueLine := range blueSubtitle.Lines {
			item.Lines = append(item.Lines, astisub.Line{
				Items: []astisub.LineItem{
					{
						Text:        formatText(blueLine, format),
						InlineStyle: lineStyle(blueLine, format),
					},
				},
			})
		}
		if format == VTT {
			id := "sub_" + strconv.Itoa(blueSubtitle.VerticalAlign) + "_" + strconv.Itoa(len(blueSubtitle.Lines))
			var vttStyleAttributes = new(astisub.StyleAttributes)
			vttStyleAttributes.WebVTTWidth = "100%"
			vttStyleAttributes.WebVTTLines = len(blueSubtitle.Lines)
			vttStyleAttributes.WebVTTRegionAnchor = "50%,0%"
			vttStyleAttributes.WebVTTViewportAnchor = "50%," + vttPosition(blueSubtitle.VerticalAlign) + "%"
			vttStyleAttributes.WebVTTScroll = "up"
			var vttRegion = new(astisub.Region)
			vttRegion.ID = id
			vttRegion.InlineStyle = vttStyleAttributes
			vttRegions[id] = vttRegion
			item.Region = vttRegion
		}
		items = append(items, item)
	}
	asticodeSubtitles.Regions = vttRegions
	asticodeSubtitles.Items = items
	return asticodeSubtitles
}

func durationOf(millisecond int64) time.Duration {
	return time.Duration(millisecond * 1e6)
}

func lineStyle(blueLine model.Line, format string) *astisub.StyleAttributes {
	var styleAttributes = new(astisub.StyleAttributes)
	switch format {
	case VTT:
		styleAttributes.WebVTTItalics = blueLine.Italic
		styleAttributes.WebVTTColor = blueLine.Color
		styleAttributes.WebVTTBold = blueLine.Bold
		styleAttributes.WebVTTUnderline = blueLine.Underline
		styleAttributes.WebVTTBackgroundColor = blueLine.BoxingColor
	case STL:
		styleAttributes.STLItalics = &blueLine.Italic
		styleAttributes.STLColor = &blueLine.Color
		styleAttributes.STLBoxing = boxing(blueLine)
		styleAttributes.STLUnderline = &blueLine.Underline
	case TTML:
		styleAttributes.TTMLColor = &blueLine.Color
	}
	return styleAttributes
}

func subtitleStyle(blueSubtitle model.Subtitle, format string) *astisub.StyleAttributes {
	var styleAttributes = new(astisub.StyleAttributes)
	switch format {
	case VTT:
		styleAttributes.WebVTTAlign = strings.ToLower(blueSubtitle.HorizontalAlign)
	case STL:
		styleAttributes.STLPosition = stlPosition(blueSubtitle)
		styleAttributes.STLJustification = &astisub.JustificationCentered
	case TTML:
	}
	return styleAttributes
}

func vttPosition(align int) string {
	return strconv.Itoa((align * 100) / 23)
}

func stlPosition(blueSubtitle model.Subtitle) *astisub.STLPosition {
	var stlPosition = new(astisub.STLPosition)
	stlPosition.VerticalPosition = blueSubtitle.VerticalAlign
	return stlPosition
}

func formatText(line model.Line, format string) string {
	if format == STL {
		length := len(line.Text)
		switch line.Justification {
		case LEFT:
			return line.Text + strings.Repeat(" ", int(math.Max(float64(36-length), 0)))
		case CENTER:
			if length%2 != 0 {
				length++
			}
			return strings.Repeat(" ", int(math.Max(float64((36-length)/2), 0))) + line.Text + strings.Repeat(" ", int(math.Max(float64((36-length)/2), 0)))
		case RIGHT:
			return strings.Repeat(" ", int(math.Max(float64(36-length), 0))) + line.Text
		default:
			return line.Text
		}
	}
	return line.Text
}

func boxing(line model.Line) *bool {
	var boxing = new(bool)
	*boxing = len(line.BoxingColor) > 0
	return boxing
}
