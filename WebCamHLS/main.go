package main

import (
	"fmt"
	"log"
	"net/http"
	"os/exec"
)

func main() {

	http.HandleFunc("/", stringHandler)

	fmt.Println("listening on port 8081")
	log.Fatal(http.ListenAndServe(":8081", nil))
}

// stringHandler returns http respone in string format.
func stringHandler(w http.ResponseWriter, r *http.Request) {
	// For Multiple(360p,480p & 720p resolutions)
	cmd := exec.Command("ffmpeg",
		"-f", "v4l2", "-video_size", "640x480",
		"-i", "/dev/video0",
		"-vf", "format=yuv420p",
		"-c:v", "libx264", "-crf", "21", "-preset", "veryfast",
		"-b:v", "100M", "-b:a", "128k",
		"-f", "hls", "-hls_list_size", "2",
		"-hls_flags", "independent_segments", "-hls_flags", "delete_segments",
		"-hls_segment_type", "mpegts",
		"-hls_segment_filename", "data%02d.ts",
		"-master_pl_name master.m3u8", "out1")

	err := cmd.Run()

	if err != nil {
		log.Println(err)
		return
	}

}
