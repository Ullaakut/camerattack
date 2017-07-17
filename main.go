package main

import (
	"fmt"
	"os"

	curl "github.com/andelf/go-curl"
)

func rtspOptions() {
	fmt.Println("Beginning preparation of OPTIONS request...")
	easy := curl.EasyInit()
	defer easy.Cleanup()
	if easy != nil {
		fmt.Println("Initialization of CURL OK")
		easy.Setopt(curl.OPT_NOBODY, 1)
		easy.Setopt(curl.OPT_URL, "rtsp://localhost:8554/live.sdp")
		easy.Setopt(curl.OPT_RTSP_STREAM_URI, "rtsp://localhost:8554/live.sdp")
		easy.Setopt(curl.OPT_FOLLOWLOCATION, 0)
		easy.Setopt(curl.OPT_HEADER, 0)
		easy.Setopt(curl.OPT_VERBOSE, 1)
		easy.Setopt(curl.OPT_RTSP_REQUEST, 1)
		fmt.Println("Performing request...")
		easy.Perform()
		fmt.Println("OK")
	}
}

func rtspDescribe() {
	fmt.Println("Beginning preparation of DESCRIBE request...")
	easy := curl.EasyInit()
	defer easy.Cleanup()
	if easy != nil {
		fmt.Println("Initialization of CURL OK")
		easy.Setopt(curl.OPT_NOBODY, 1)
		easy.Setopt(curl.OPT_URL, "rtsp://localhost:8554/live.sdp")
		easy.Setopt(curl.OPT_RTSP_STREAM_URI, "rtsp://localhost:8554/live.sdp")
		easy.Setopt(curl.OPT_FOLLOWLOCATION, 0)
		easy.Setopt(curl.OPT_HEADER, 0)
		easy.Setopt(curl.OPT_VERBOSE, 1)
		easy.Setopt(curl.OPT_RTSP_REQUEST, 2)
		fmt.Println("Performing request...")
		easy.Perform()
		fmt.Println("OK")
	}
}

func rtspAnnounce() {
	fmt.Println("Beginning preparation of ANNOUNCE request...")
	easy := curl.EasyInit()
	defer easy.Cleanup()
	if easy != nil {
		fmt.Println("Initialization of CURL OK")
		easy.Setopt(curl.OPT_NOBODY, 1)
		easy.Setopt(curl.OPT_URL, "rtsp://localhost:8554/live.sdp")
		easy.Setopt(curl.OPT_RTSP_STREAM_URI, "rtsp://localhost:8554/live.sdp")
		easy.Setopt(curl.OPT_FOLLOWLOCATION, 0)
		easy.Setopt(curl.OPT_HEADER, 0)
		easy.Setopt(curl.OPT_VERBOSE, 1)
		easy.Setopt(curl.OPT_RTSP_REQUEST, 3)
		fmt.Println("Performing request...")
		easy.Perform()
		fmt.Println("OK")
	}
}

func rtspSetup() string {
	fmt.Println("Beginning preparation of SETUP request...")
	easy := curl.EasyInit()
	defer easy.Cleanup()
	if easy != nil {
		fmt.Println("Initialization of CURL OK")
		easy.Setopt(curl.OPT_NOBODY, 1)
		easy.Setopt(curl.OPT_URL, "rtsp://localhost:8554/live.sdp")
		easy.Setopt(curl.OPT_RTSP_STREAM_URI, "rtsp://localhost:8554/live.sdp")
		easy.Setopt(curl.OPT_RTSPHEADER, []string{"Transport: RTP/AVP;unicast;client_port=5067-5068\n"})
		easy.Setopt(curl.OPT_HEADER, 1)
		easy.Setopt(curl.OPT_FOLLOWLOCATION, 0)
		easy.Setopt(curl.OPT_VERBOSE, 1)
		easy.Setopt(curl.OPT_RTSP_REQUEST, 4)
		fmt.Println("Performing request...")
		err := easy.Perform()
		fmt.Println("OK")
		if err != nil {
			fmt.Println(err)
			os.Exit(0)
		}
		// fmt.Println(ret)
		return "OK"
	}
	return "KO"
}

func rtspPlay(session string) {
	fmt.Println("Beginning preparation of PLAY request...")
	easy := curl.EasyInit()
	defer easy.Cleanup()
	if easy != nil {
		fmt.Println("Initialization of CURL OK")
		easy.Setopt(curl.OPT_NOBODY, 1)
		easy.Setopt(curl.OPT_URL, "rtsp://localhost:8554/live.sdp")
		easy.Setopt(curl.OPT_RTSP_STREAM_URI, "rtsp://localhost:8554/live.sdp")
		easy.Setopt(curl.OPT_FOLLOWLOCATION, 0)
		easy.Setopt(curl.OPT_RTSP_SESSION_ID, 0)
		easy.Setopt(curl.OPT_HEADER, 0)
		easy.Setopt(curl.OPT_VERBOSE, 1)
		easy.Setopt(curl.OPT_RTSP_REQUEST, 5)
		fmt.Println("Performing request...")
		easy.Perform()
		fmt.Println("OK")
	}
}

func rtspPause(session string) {
	fmt.Println("Beginning preparation of PAUSE request...")
	easy := curl.EasyInit()
	defer easy.Cleanup()
	if easy != nil {
		fmt.Println("Initialization of CURL OK")
		easy.Setopt(curl.OPT_NOBODY, 1)
		easy.Setopt(curl.OPT_URL, "rtsp://localhost:8554/live.sdp")
		easy.Setopt(curl.OPT_RTSP_STREAM_URI, "rtsp://localhost:8554/live.sdp")
		easy.Setopt(curl.OPT_FOLLOWLOCATION, 0)
		easy.Setopt(curl.OPT_HEADER, 0)
		easy.Setopt(curl.OPT_VERBOSE, 1)
		easy.Setopt(curl.OPT_RTSP_REQUEST, 6)
		fmt.Println("Performing request...")
		easy.Perform()
		fmt.Println("OK")
	}
}

func rtspTeardown(session string) {
	fmt.Println("Beginning preparation of TEARDOWN request...")
	easy := curl.EasyInit()
	defer easy.Cleanup()
	if easy != nil {
		fmt.Println("Initialization of CURL OK")
		easy.Setopt(curl.OPT_NOBODY, 1)
		easy.Setopt(curl.OPT_URL, "rtsp://localhost:8554/live.sdp")
		easy.Setopt(curl.OPT_RTSP_STREAM_URI, "rtsp://localhost:8554/live.sdp")
		easy.Setopt(curl.OPT_FOLLOWLOCATION, 0)
		easy.Setopt(curl.OPT_RTSP_SESSION_ID, 0)
		easy.Setopt(curl.OPT_HEADER, 0)
		easy.Setopt(curl.OPT_VERBOSE, 1)
		easy.Setopt(curl.OPT_RTSP_REQUEST, 7)
		fmt.Println("Performing request...")
		easy.Perform()
		fmt.Println("OK")
	}
}

func main() {

	for {
		rtspOptions()
		rtspDescribe()
		session := rtspSetup()
		fmt.Println(session)
	}
}
