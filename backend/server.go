package backend

import (
	"embed"
	"fmt"
	"io"
	"io/fs"
	"log"
	"net"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

//go:embed frontend/*.*
var frontend embed.FS

func HandleAll() {
	i, p, c := readConfig()

	ip := net.JoinHostPort(i, p)

	webFS, err := fs.Sub(frontend, "frontend")
	if err != nil {
		log.Fatalf("Error reading embedded frontend: %v", err)
	}

	if c != "" {
		log.Println("Using Custom UI provided in config:", c)
		http.Handle("/", http.FileServer(http.Dir(c)))
	} else {
		log.Println("Using Built-in UI")
		http.Handle("/", http.FileServer(http.FS(webFS)))

	}

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

	err = http.ListenAndServe(ip, nil)
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
    out, e := exec.Command("playerctl", "metadata", "-f", `{"playername":"{{playerName}}","position":"{{duration(position)}}","status":"{{status}}","volume":"{{volume}}","album":"{{xesam:album}}","artist":"{{xesam:artist}}","title":"{{xesam:title}}", "length": "{{duration(mpris:length)}}","arturl":"{{mpris:artUrl}}"}`).Output()
    if e != nil {
        log.Printf("Error running playerctl %v", e)
    }

    data := string(out)

    start := strings.Index(data, `"arturl":"file://`)
    if start != -1 {
        start += len(`"arturl":"file://`)
        end := strings.Index(data[start:], `"`)
        if end != -1 {
            tmpPath := data[start : start+end]

            filename := "art"

            _, _, c := readConfig()
            var publicPath string
            if c != "" {
                publicPath = filepath.Join(c, "public", filename)
            } else {
                publicPath = filepath.Join("frontend", "public", filename)
            }

            if err := copyFile(tmpPath, publicPath); err != nil {
                log.Printf("Arturl copy error: %v", err)
            }

            data = data[:start] + "/public/" + filename + data[start+end:]
        }
    }

    w.Header().Set("Content-Type", "application/json")
    w.Write([]byte(data))
}

func handleForward5Sec(w http.ResponseWriter, r *http.Request) {
	e := playerctl("Got 5 seconds forward from server", "position", "5+")
	if e != nil {
		log.Printf("Error running playerctl %v", e)
	}
}

func handleRewind5Sec(w http.ResponseWriter, r *http.Request) {
	e := playerctl("Got 5 Second Backwards from server", "position", "5-")
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

func copyFile(src, dst string) error {
    src = strings.TrimPrefix(src, "file://")

    if err := os.MkdirAll(filepath.Dir(dst), 0755); err != nil {
        return err
    }

    from, err := os.Open(src)
    if err != nil {
        return err
    }
    defer from.Close()

    to, err := os.Create(dst)
    if err != nil {
        return err
    }
    defer to.Close()

    _, err = io.Copy(to, from)
    return err
}
