package backend

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"os/exec"
)

func HandleAll() {
	i, p := readConfig()

	ip := net.JoinHostPort(i, p)

	server := http.FileServer(http.Dir("./frontend"))

	http.Handle("/", server)

	routes := map[string]http.HandlerFunc{
		"/api/play-pause":    handlePlayPause,
		"/api/next":          handleNext,
		"/api/previous":      handlePrevious,
		"/api/forward-5-sec": handleForward5Sec,
		"/api/rewind-5-sec":  handleRewind5Sec,
		"/api/volume-up-5":   handleVolumeUp,
		"/api/volume-down-5": handleVolumeDown,
		"/api/info":          sendInfo,
	}

	for path, handler := range routes {
		http.HandleFunc(path, cors(handler))
	}

	fmt.Println("Started at: ", ip)

	err := http.ListenAndServe(ip, nil)
	if err != nil {
		fmt.Printf("Error starting server: %s\n", err)
	}
}

func cors(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		if r.Method == "OPTIONS" {
			return
		}

		h(w, r)
	}
}

func handleVolumeUp(w http.ResponseWriter, r *http.Request) {
	e := exec.Command("wpctl", "set-volume", "@DEFAULT_AUDIO_SINK@", "0.05+").Run()
	if e != nil {
		log.Printf("Error running wpctl %v", e)
	}
	notifier("Got Volume Up from server")
}

func handleVolumeDown(w http.ResponseWriter, r *http.Request) {
	e := exec.Command("wpctl", "set-volume", "@DEFAULT_AUDIO_SINK@", "0.05-").Run()
	if e != nil {
		log.Printf("Error running wpctl %v", e)
	}
	notifier("Got Volume Down from server")
}

func sendInfo(w http.ResponseWriter, r *http.Request) {
	out, e := exec.Command("playerctl", "metadata", "-f", `{"playername":"{{playerName}}","position":"{{duration(position)}}","status":"{{status}}","volume":"{{volume}}","album":"{{xesam:album}}","artist":"{{xesam:artist}}","title":"{{xesam:title}}", "length": "{{duration(mpris:length)}}"}`).Output()
	if e != nil {
		log.Printf("Error running playerctl %v", e)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(out)
}

func handleForward5Sec(w http.ResponseWriter, r *http.Request) {
	e := playerctl("Got 5 seconds forward from server", "position", "5-")
	if e != nil {
		log.Printf("Error running playerctl %v", e)
	}
}

func handleRewind5Sec(w http.ResponseWriter, r *http.Request) {
	e := playerctl("Got 5 Second Backwards from server", "position", "5+")
	if e != nil {
		log.Printf("Error running playerctl %v", e)
	}
}

func handlePrevious(w http.ResponseWriter, r *http.Request) {
	e := playerctl("Got Previous from server", "previous")
	if e != nil {
		log.Printf("Error running playerctl %v", e)
	}
}

func handleNext(w http.ResponseWriter, r *http.Request) {
	e := playerctl("Got Next from server", "next")
	if e != nil {
		log.Printf("Error running playerctl %v", e)
	}
}

func handlePlayPause(w http.ResponseWriter, r *http.Request) {
	e := playerctl("Got Play/Pause from server", "play-pause")
	if e != nil {
		log.Printf("Error running playerctl %v", e)
	}
}

func playerctl(message string, args ...string) error {
	cmd := exec.Command("playerctl", args...)
	notifier(message)
	return cmd.Run()
}

func notifier(message string) {
	log.Println(message)
}
