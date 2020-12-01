// JPEG toy
// https://d33wubrfki0l68.cloudfront.net/bdc1363abbd5744200ec5283d4154e55143df86c/8c624/images/decoding_jpeg/jpegrgb_dissected.png
// https://yasoob.me/posts/understanding-and-writing-jpeg-decoder-in-python/
package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
)

type Segment struct {
	Name   string
	Marker SegmentMarker
	// offset and length are measured in bytes
	Offset int64
	Length int64
}

type SegmentMarker struct {
	Length int64
}

func main() {
	var (
		jpegMarker = SegmentMarker{2}
		filePtr    = flag.String("file", "", "jpeg file to use")
	)

	var (
		headerSegment    = Segment{"Header", jpegMarker, 0, 18}
		luminanceSegment = Segment{"Luminance", jpegMarker, jpegMarker.Length + headerSegment.Offset + headerSegment.Length, 132}
		frameSegment     = Segment{"Frame", jpegMarker, jpegMarker.Length + luminanceSegment.Offset + luminanceSegment.Length, 17}
	)

	flag.Parse()
	if *filePtr == "" || (filepath.Ext(*filePtr) != ".jpeg" && filepath.Ext(*filePtr) != ".jpg") {
		panic("You must define a jpeg file to use")
	}

	file, err := os.Open(*filePtr)
	if err != nil {
		panic(err)
	}

	// FileInfo structure give us basic data about the file
	fileinfo, err := file.Stat()
	if err != nil {
		panic(err)
	}
	//size := fileinfo.Size()
	fmt.Println(fileinfo.Name())

	printSegment(file, headerSegment)
	printSegment(file, luminanceSegment)
	printSegment(file, frameSegment)

}

func readSegment(file *os.File, segment Segment) (string, error) {
	data := make([]byte, segment.Length)
	// Why not ioutil.ReadFile?
	// ReadFile loads EVERYTHING in memory, much like PHP's file_get_contents()!
	_, err := file.ReadAt(data, segment.Offset)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%q", data), nil
}

func printSegment(file *os.File, segment Segment) {
	data, err := readSegment(file, segment)
	if err != nil {
		panic(err)
	}
	fmt.Println(data)
}
