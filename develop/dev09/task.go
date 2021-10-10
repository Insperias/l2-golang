package main

import (
	"errors"
	"io"
	"log"
	"math"
	"mime"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

func main() {
	var err error
	if len(os.Args) < 2 {
		log.Printf("usage: %s url [optioanl filename]", os.Args[0])
	}
	if len(os.Args) == 3 {
		err = wget(os.Args[1], os.Args[2])
	} else {
		err = wget(os.Args[1], "")
	}
	if err != nil {
		log.Printf("Error: %+v\n", err)
		os.Exit(1)
	}

}
func tidyFilename(filename, defaultFilename string) string {
	if filename == "" || filename == "/" || filename == "\\" || filename == "." {
		filename = defaultFilename
	}
	return filename
}

func wget(link, outputFilename string) error {
	if !strings.Contains(link, ":") {
		link = "http://" + link
	}

	startTime := time.Now()
	req, err := http.NewRequest("GET", link, nil)

	if err != nil {
		return err
	}

	filename := ""
	if outputFilename != "" {
		filename = outputFilename
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	log.Printf("Http response status: %s\n", resp.Status)

	lenS := resp.Header.Get("Content-Length")
	length := int64(-1)
	if lenS != "" {
		length, err = strconv.ParseInt(lenS, 10, 32)
		if err != nil {
			return err
		}
	}

	typ := resp.Header.Get("Content-Type")
	log.Printf("Content-Length: %v Content-Type: %s\n", lenS, typ)

	if filename == "" {
		filename, err = getFilename(req, resp)
		if err != nil {
			return err
		}
	}

	var out io.Writer
	var outFile *os.File

	log.Printf("Saving to: '%v'\n\n", filename)
	outFile, err = os.OpenFile(filename, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0660)
	if err != nil {
		return err
	}
	defer outFile.Close()
	out = outFile

	buf := make([]byte, 4068)
	tot := int64(0)
	i := 0

	for {
		//read chunk
		n, err := resp.Body.Read(buf)
		if err != nil && err != io.EOF {
			return err
		}
		if n == 0 {
			break
		}
		tot += int64(n)

		//write chunk
		if _, err := out.Write(buf[:n]); err != nil {
			return err
		}
		i++
		if length > -1 {
			if length < 1 {
				log.Printf("\r	[ <=>                                  ] %d\t-.--KB/s eta ?s             ", tot)
			} else {
				perc := (100 * tot) / length
				prog := progress(perc)
				nowTime := time.Now()
				totTime := nowTime.Sub(startTime)
				spd := float64(tot/1000) / totTime.Seconds()
				remKb := float64(length-tot) / float64(1000)
				eta := remKb / spd
				log.Printf("\r%3d%% [%s] %d\t%0.2fKB/s eta %0.1fs             ", perc, prog, tot, spd, eta)

			}
		} else {
			if math.Mod(float64(i), 20) == 0 {
				log.Printf(".")
			}
		}
	}

	nowTime := time.Now()
	totTime := nowTime.Sub(startTime)
	spd := float64(tot/1000) / totTime.Seconds()
	if length < 1 {
		log.Printf("\r     [ <=>                                  ] %d\t-.--KB/s in %0.1fs             ", tot, totTime.Seconds())
		log.Printf("\n (%0.2fKB/s) - '%v' saved [%v]\n", spd, filename, tot)
	} else {
		perc := (100 * tot) / length
		prog := progress(perc)
		log.Printf("\r%3d%% [%s] %d\t%0.2fKB/s in %0.1fs             ", perc, prog, tot, spd, totTime.Seconds())
		log.Printf("\n '%v' saved [%v/%v]\n", filename, tot, length)
	}
	if err != nil {
		return err
	}
	if outFile != nil {
		err = outFile.Close()
	}
	return err
}

func progress(perc int64) string {
	equalses := perc * 38 / 100
	if equalses < 0 {
		equalses = 0
	}

	spaces := 38 - equalses
	if spaces < 0 {
		spaces = 0
	}
	prog := strings.Repeat("=", int(equalses)) + ">" + strings.Repeat(" ", int(spaces))
	return prog
}

func getFilename(req *http.Request, resp *http.Response) (string, error) {
	filename := filepath.Base(req.URL.Path)

	if !strings.Contains(filename, ".") {
		filename = filepath.Base(resp.Request.URL.Path)
	}
	filename = tidyFilename(filename, "index.html")

	if !strings.Contains(filename, ".") {
		ct := resp.Header.Get("Content-Type")
		ext := "htm"
		mediatype, _, err := mime.ParseMediaType(ct)
		if err != nil {
			log.Printf("mime error: %v\n", err)
		} else {
			log.Printf("mime type: %v (from Content-Type %v)\n", mediatype, ct)
			slash := strings.Index(mediatype, "/")
			if slash != -1 {
				_, sub := mediatype[:slash], mediatype[slash+1:]
				if sub != "" {
					ext = sub
				}
			}
		}
		filename = filename + "." + ext
	}
	_, err := os.Stat(filename)
	if err != nil {
		if os.IsNotExist(err) {
			return filename, nil
		}
		return "", err

	}
	num := 1

	fnSlice := strings.Split(filename, ".")
	for num < 100 {
		filenameNew := fnSlice[0] + "(" + strconv.Itoa(num) + ")" + "." + fnSlice[1]
		_, err := os.Stat(filenameNew)
		if err != nil {
			if os.IsNotExist(err) {
				return filenameNew, nil
			}
			return "", err

		}
		num++
	}
	return filename, errors.New("stopping after trying 100 filename variants")
}
