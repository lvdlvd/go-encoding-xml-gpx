/*
Package gpx provides xml En/Decodable structures for the GPX format as defined by
http://www.topografix.com/gpx/1/1/ gpx.xsd.

GPX schema version 1.1

For more information on GPX and this schema, visit http://www.topografix.com/gpx.asp

GPX uses the following conventions: all coordinates are relative to the WGS84 datum. All measurements are in metric units.
*/
package gpx

import (
	"encoding/xml"
	"time"
)

// GPX is the root element in the XML file.
//
// GPX documents contain a metadata header, followed by waypoints, routes, and tracks.
type GPX struct {
	XMLName  xml.Name      `xml:"gpx"`
	Version  string        `xml:"version,attr"`
	Creator  string        `xml:"creator,attr"`
	Metadata *MetadataType `xml:"metadata,omitempty"`
	Wpt      []*WptType    `xml:"wpt,omitempty"`
	Rte      []*RteType    `xml:"rte,omitempty"`
	Trk      []*TrkType    `xml:"trk,omitempty"`
}

// Information about the GPX file, author, and copyright restrictions goes in the metadata section.
// Providing rich, meaningful information about your GPX files allows others to search for and use your GPS data.
type MetadataType struct {
	Name      *string        `xml:"name,omitempty"`
	Desc      *string        `xml:"desc,omitempty"`
	Author    *PersonType    `xml:"author,omitempty"`
	Copyright *CopyrightType `xml:"copyright,omitempty"`
	Link      []*LinkType    `xml:"link,omitempty"`
	Time      *time.Time     `xml:"time,omitempty"`
	Keywords  *string        `xml:"keywords,omitempty"`
	Bounds    *BoundsType    `xml:"bounds,omitempty"`
}

// A person or organization.
type PersonType struct {
	Name  *string    `xml:"name,omitempty"`
	Email *EmailType `xml:"email,omitempty"`
	Link  *LinkType  `xml:"link,omitempty"`
}

// An email address. Broken into two parts (id and domain) to help prevent email harvesting.
type EmailType struct {
	Id     string `xml:"id,attr"`
	Domain string `xml:"domain,attr"`
}

// A link to an external resource (Web page, digital photo, video clip, etc) with additional information.
type LinkType struct {
	Href string  `xml:"href,attr"`
	Text *string `xml:"text,omitempty"`
	Type *string `xml:"type,omitempty"`
}

// Information about the copyright holder and any license governing use of this file.
// By linking to an appropriate license, you may place your data into the public domain or grant additional usage rights.
type CopyrightType struct {
	Author  string  `xml:"author,attr"`
	Year    *int    `xml:"year,omitempty"`
	License *string `xml:"license,omitempty"`
}

// Two lat/lon pairs defining the extent of an element.
type BoundsType struct {
	Minlat float64 `xml:"minlat,attr"`
	Minlon float64 `xml:"minlon,attr"`
	Maxlat float64 `xml:"maxlat,attr"`
	Maxlon float64 `xml:"maxlon,attr"`
}

// wpt represents a waypoint, point of interest, or named feature on a map.
type WptType struct {
	Lat           float64     `xml:"lat,attr"`
	Lon           float64     `xml:"lon,attr"`
	Ele           *int        `xml:"ele,omitempty"`           // Elevation (in meters) of the point.
	Time          *time.Time  `xml:"time,omitempty"`          // Creation/modification timestamp for element. Date and time in are in Univeral Coordinated Time (UTC), not local time! Conforms to ISO 8601 specification for date/time representation. Fractional seconds are allowed for millisecond timing in tracklogs.
	Magvar        *float64    `xml:"magvar,omitempty"`        // Magnetic variation (in degrees) at the point
	Geoidheight   *float64    `xml:"geoidheight,omitempty"`   // Height (in meters) of geoid (mean sea level) above WGS84 earth ellipsoid.  As defined in NMEA GGA message.
	Name          *string     `xml:"name,omitempty"`          // The GPS name of the waypoint. This field will be transferred to and from the GPS. GPX does not place restrictions on the length of this field or the characters contained in it. It is up to the receiving application to validate the field before sending it to the GPS.
	Cmt           *string     `xml:"cmt,omitempty"`           // GPS waypoint comment. Sent to GPS as comment.
	Desc          *string     `xml:"desc,omitempty"`          // A text description of the element. Holds additional information about the element intended for the user, not the GPS.
	Src           *string     `xml:"src,omitempty"`           // Source of data. Included to give user some idea of reliability and accuracy of data.  "Garmin eTrex", "USGS quad Boston North", e.g.
	Link          []*LinkType `xml:"link,omitempty"`          // Link to additional information about the waypoint.
	Sym           *string     `xml:"sym,omitempty"`           // Text of GPS symbol name. For interchange with other programs, use the exact spelling of the symbol as displayed on the GPS.  If the GPS abbreviates words, spell them out.
	Type          *string     `xml:"type,omitempty"`          // Type (classification) of the waypoint.
	Fix           *string     `xml:"fix,omitempty"`           // Type of GPX fix. ('none'|'2d'|'3d'|'dgps'|'pps')
	Sat           *int        `xml:"sat,omitempty"`           // Number of satellites used to calculate the GPX fix.
	Hdop          *int        `xml:"hdop,omitempty"`          // Horizontal dilution of precision.
	Vdop          *int        `xml:"vdop,omitempty"`          // Vertical dilution of precision.
	Pdop          *int        `xml:"pdop,omitempty"`          // Position dilution of precision.
	Ageofdgpsdata *int        `xml:"ageofdgpsdata,omitempty"` // Number of seconds since last DGPS update.
	Dgpsid        *string     `xml:"dgpsid,omitempty"`        // ID of DGPS station used in differential correction.
}

// rte represents route - an ordered list of waypoints representing a series of turn points leading to a destination.
type RteType struct {
	Name   *string     `xml:"name,omitempty"`
	Cmt    *string     `xml:"cmt,omitempty"`
	Desc   *string     `xml:"desc,omitempty"`
	Src    *string     `xml:"src,omitempty"`
	Link   []*LinkType `xml:"link,omitempty"`
	Number *int        `xml:"number,omitempty"`
	Type   *string     `xml:"type,omitempty"`
	Rtept  []*WptType  `xml:"rtept,omitempty"`
}

// trk represents a track - an ordered list of points describing a path.
type TrkType struct {
	Name   *string       `xml:"name,omitempty"`
	Cmt    *string       `xml:"cmt,omitempty"`
	Desc   *string       `xml:"desc,omitempty"`
	Src    *string       `xml:"src,omitempty"`
	Link   []*LinkType   `xml:"link,omitempty"`
	Number *int          `xml:"number,omitempty"`
	Type   *string       `xml:"type,omitempty"`
	Trkseg []*TrksegType `xml:"trkseg,omitempty"`
}

// A Track Segment holds a list of Track Points which are logically connected in order.
// To represent a single GPS track where GPS reception was lost, or the GPS receiver was
// turned off, start a new Track Segment for each continuous span of track data.
type TrksegType struct {
	Trkpt []*WptType `xml:"trkpt,omitempty"`
}
