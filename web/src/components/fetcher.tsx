import { createSignal } from "solid-js";

export const [playerInfo, setPlayerInfo] = createSignal({
    playername: "",
    position: "",
    status: "",
    volume: "",
    album: "",
    artist: "",
    title: "",
    length: "",
    arturl: "",
});

export function fetchThem() {
    fetch("/api/info")
        .then((response) => {
            if (!response.ok) throw new Error("HTTP error " + response.status);
            return response.json();
        })
        .then((data) => {
            setPlayerInfo(data);
        })
        .catch((err) => console.error("Fetch error:", err));
}
