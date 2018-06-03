/*
  cutgpx reads in a gpx file and writes out copies with only the waypoints and the separate tracks.

*/
package main

import (
	"encoding/xml"
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	gpx "github.com/lvdlvd/go-encoding-xml-gpx"
)

func main() {

	flag.Parse()

	f := os.Stdin
	var err error
	if len(flag.Args()) > 0 {
		f, err = os.Open(flag.Arg(0))
		if err != nil {
			log.Fatalln(err)
		}
	}

	log.Println("reading from ", f.Name())

	fname := filepath.Base(f.Name())
	ext := filepath.Ext(fname)
	fname = strings.TrimSuffix(fname, ext)

	var gpx gpx.GPX
	if err := xml.NewDecoder(f).Decode(&gpx); err != nil {
		log.Fatalln(err)
	}

	tracks := gpx.Trk
	gpx.Trk = nil

	{
		of, err := os.Create(fmt.Sprintf("%s-WPT%s", fname, ext))
		if err != nil {
			log.Fatalln(err)
		}
		log.Printf("Writing %s", of.Name())

		xml.NewEncoder(of).Encode(gpx)
		of.Close()
	}

	gpx.Wpt = nil

	for i := range tracks {
		of, err := os.Create(fmt.Sprintf("%s-%d%s", fname, i, ext))
		if err != nil {
			log.Fatalln(err)
		}
		log.Printf("Writing %s", of.Name())

		gpx.Trk = tracks[i : i+1]

		xml.NewEncoder(of).Encode(gpx)
		of.Close()
	}

}
