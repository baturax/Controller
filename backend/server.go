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

	http.HandleFunc("/api/play-pause", cors(handlePlayPause))
	http.HandleFunc("/api/next", cors(handleNext))
	http.HandleFunc("/api/previous", cors(handlePrev))
	http.HandleFunc("/api/fiveplus", cors(handleFivePlus))
	http.HandleFunc("/api/fiveminus", cors(handleFiveMinus))
	http.HandleFunc("/api/info", cors(sendInfo))

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


func sendInfo(w http.ResponseWriter, r *http.Request) {
	out, err := exec.Command("playerctl", "metadata", "-f", `{"playername":"{{playerName}}","position":"{{duration(position)}}","status":"{{status}}","volume":"{{volume}}","album":"{{xesam:album}}","artist":"{{xesam:artist}}","title":"{{xesam:title}}", "length": "{{duration(mpris:length)}}"}`).Output()
	if err != nil {
		log.Print("Error running playerctl ", err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(out)
}

func handleFiveMinus(w http.ResponseWriter, r *http.Request) {
	e := playerctl("5 seconds backwards?? from server", "position", "5-")
	if e != nil {
		log.Print("Error running playerctl ", e)
	}
}

func handleFivePlus(w http.ResponseWriter, r *http.Request) {
	e := playerctl("5 seconds forward?? from server", "position", "5+")
	if e != nil {
		log.Print("Error running playerctl ", e)
	}
}

func handlePrev(w http.ResponseWriter, r *http.Request) {
	e := playerctl("previous from server", "previous")
	if e != nil {
		log.Print("Error running playerctl ", e)
	}
}

func handleNext(w http.ResponseWriter, r *http.Request) {
	e := playerctl("next from server", "next")
	if e != nil {
		log.Print("Error running playerctl ", e)
	}
}

func handlePlayPause(w http.ResponseWriter, r *http.Request) {
	e := playerctl("play-pause from server", "play-pause")
	if e != nil {
		log.Print("Error running playerctl ", e)
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
