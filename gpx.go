/*
Package gpx provides xml En/Decodable structures for the GPX format as defined by
http://www.topografix.com/gpx/1/1/ gpx.xsd.
*/
package gpx

import (
	"encoding/xml"
	"time"
)

/*
<gpx
version="1.1 [1] ?"
creator="xsd:string [1] ?">
<metadata> metadataType </metadata> [0..1] ?
<wpt> wptType </wpt> [0..*] ?
<rte> rteType </rte> [0..*] ?
<trk> trkType </trk> [0..*] ?
<extensions> extensionsType </extensions> [0..1] ?
</gpx>
*/

type Gpx struct {
	XMLName  xml.Name      `xml:"gpx"`
	Version  string        `xml:"version,attr"`
	Creator  string        `xml:"creator,attr"`
	Metadata *MetadataType `xml:"metadata,omitempty"`
	Wpt      []*WptType    `xml:"wpt,omitempty"`
	Rte      []*RteType    `xml:"rte,omitempty"`
	Trk      []*TrkType    `xml:"trk,omitempty"`
}

/*
<name> xsd:string </name> [0..1] ?
<desc> xsd:string </desc> [0..1] ?
<author> personType </author> [0..1] ?
<copyright> copyrightType </copyright> [0..1] ?
<link> linkType </link> [0..*] ?
<time> xsd:dateTime </time> [0..1] ?
<keywords> xsd:string </keywords> [0..1] ?
<bounds> boundsType </bounds> [0..1] ?
<extensions> extensionsType </extensions> [0..1] ?
*/
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

/*
<name> xsd:string </name> [0..1] ?
<email> emailType </email> [0..1] ?
<link> linkType </link> [0..1] ?
*/
type PersonType struct {
	Name  *string    `xml:"name,omitempty"`
	Email *EmailType `xml:"email,omitempty"`
	Link  *LinkType  `xml:"link,omitempty"`
}

/*
<...
id="xsd:string [1] ?"
domain="xsd:string [1] ?"
/>
*/
type EmailType struct {
	Id     string `xml:"id,attr"`
	Domain string `xml:"domain,attr"`
}

/*
href="xsd:anyURI [1] ?">
<text> xsd:string </text> [0..1] ?
<type> xsd:string </type> [0..1] ?
*/
type LinkType struct {
	Href string  `xml:"href,attr"`
	Text *string `xml:"text,omitempty"`
	Type *string `xml:"type,omitempty"`
}

/*
author="xsd:string [1] ?">
<year> xsd:gYear </year> [0..1] ?
<license> xsd:anyURI </license> [0..1] ?
*/
type CopyrightType struct {
	Author  string  `xml:"author,attr"`
	Year    *int    `xml:"year,omitempty"`
	License *string `xml:"license,omitempty"`
}

/*
minlat="latitudeType [1] ?"
minlon="longitudeType [1] ?"
maxlat="latitudeType [1] ?"
maxlon="longitudeType [1] ?"/>
*/
type BoundsType struct {
	Minlat float64 `xml:"minlat,attr"`
	Minlon float64 `xml:"minlon,attr"`
	Maxlat float64 `xml:"maxlat,attr"`
	Maxlon float64 `xml:"maxlon,attr"`
}

/*
lat="latitudeType [1] ?"
lon="longitudeType [1] ?">
<ele> xsd:decimal </ele> [0..1] ?
<time> xsd:dateTime </time> [0..1] ?
<magvar> degreesType </magvar> [0..1] ?
<geoidheight> xsd:decimal </geoidheight> [0..1] ?
<name> xsd:string </name> [0..1] ?
<cmt> xsd:string </cmt> [0..1] ?
<desc> xsd:string </desc> [0..1] ?
<src> xsd:string </src> [0..1] ?
<link> linkType </link> [0..*] ?
<sym> xsd:string </sym> [0..1] ?
<type> xsd:string </type> [0..1] ?
<fix> fixType </fix> [0..1] ?
<sat> xsd:nonNegativeInteger </sat> [0..1] ?
<hdop> xsd:decimal </hdop> [0..1] ?
<vdop> xsd:decimal </vdop> [0..1] ?
<pdop> xsd:decimal </pdop> [0..1] ?
<ageofdgpsdata> xsd:decimal </ageofdgpsdata> [0..1] ?
<dgpsid> dgpsStationType </dgpsid> [0..1] ?
<extensions> extensionsType </extensions> [0..1] ?
*/
type WptType struct {
	Lat           float64     `xml:"lat,attr"`
	Lon           float64     `xml:"lon,attr"`
	Ele           *int        `xml:"ele,omitempty"`
	Time          *time.Time  `xml:"time,omitempty"`
	Magvar        *float64    `xml:"magvar,omitempty"`
	Geoidheight   *float64    `xml:"geoidheight,omitempty"`
	Name          *string     `xml:"name,omitempty"`
	Cmt           *string     `xml:"cmt,omitempty"`
	Desc          *string     `xml:"desc,omitempty"`
	Src           *string     `xml:"src,omitempty"`
	Link          []*LinkType `xml:"link,omitempty"`
	Sym           *string     `xml:"sym,omitempty"`
	Type          *string     `xml:"type,omitempty"`
	Fix           *string     `xml:"fix,omitempty"`
	Sat           *int        `xml:"sat,omitempty"`
	Hdop          *int        `xml:"hdop,omitempty"`
	Vdop          *int        `xml:"vdop,omitempty"`
	Pdop          *int        `xml:"pdop,omitempty"`
	Ageofdgpsdata *int        `xml:"ageofdgpsdata,omitempty"`
	Dgpsid        *string     `xml:"dgpsid,omitempty"`
}

/*
<name> xsd:string </name> [0..1] ?
<cmt> xsd:string </cmt> [0..1] ?
<desc> xsd:string </desc> [0..1] ?
<src> xsd:string </src> [0..1] ?
<link> linkType </link> [0..*] ?
<number> xsd:nonNegativeInteger </number> [0..1] ?
<type> xsd:string </type> [0..1] ?
<extensions> extensionsType </extensions> [0..1] ?
<rtept> wptType </rtept> [0..*] ?
*/
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

/*
<name> xsd:string </name> [0..1] ?
<cmt> xsd:string </cmt> [0..1] ?
<desc> xsd:string </desc> [0..1] ?
<src> xsd:string </src> [0..1] ?
<link> linkType </link> [0..*] ?
<number> xsd:nonNegativeInteger </number> [0..1] ?
<type> xsd:string </type> [0..1] ?
<extensions> extensionsType </extensions> [0..1] ?
<trkseg> trksegType </trkseg> [0..*] ?
*/
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

/*
<trkpt> wptType </trkpt> [0..*] ?
<extensions> extensionsType </extensions> [0..1] ?
*/
type TrksegType struct {
	Trkpt []*WptType `xml:"trkpt,omitempty"`
}
