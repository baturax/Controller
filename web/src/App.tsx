import { createSignal, onCleanup, onMount, Show, type Component } from "solid-js";

import styles from "./App.module.css";

import PlayButton from "./assets/play.svg";
import PauseButton from "./assets/pause.svg";
import PreviousButton from "./assets/previous.svg";
import NextButton from "./assets/next.svg";
import Prev5Button from "./assets/replay5.svg";
import Next5Button from "./assets/forward5.svg";

const App: Component = () => {
  const [playerInfo, setPlayerInfo] = createSignal({
    playername: "",
    position: "",
    status: "",
    volume: "",
    album: "",
    artist: "",
    title: "",
    length: "",
  });

  function bai() {
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

  onMount(() => {
    bai();

    const keys = (e: KeyboardEvent) => {
      switch (e.code) {
        case "ArrowLeft":
          e.preventDefault();
          Previous();
          break;

        case "Space":
          e.preventDefault();
          PlayPause();
          break;
        case "ArrowRight":
          e.preventDefault();
          Next();
          break;
        case "ArrowDown":
          e.preventDefault();
          BackFive();
          break;
        case "ArrowUp":
          e.preventDefault();
          NextFive();
          break;
      }
    }

    window.addEventListener("keydown", keys);

    onCleanup(() => {
      window.removeEventListener("keydown", keys)
    })

  });


  setInterval(() => {
    bai();
  }, 500);

  return (
    <div class={styles.App}>
      <header class={styles.header}>
        <Show when={playerInfo().status == "Playing"}>
          <button id="play-pause-btn" onClick={PlayPause}>
            <img src={PauseButton} />
          </button>
        </Show>

        <Show when={playerInfo().status == "Paused"}>
          <button id="play-pause-btn" onClick={PlayPause}>
            <img src={PlayButton} />
          </button>
        </Show>

        <div id="media-box">
          <h1 id="title-h3">{playerInfo().title}</h1>
          <div id="extra-media">
            <Show when={playerInfo().artist}>
              <h3 id="artist-h3">Artist: {playerInfo().artist}</h3>
            </Show>
            <Show when={playerInfo().album}>
              <h3 id="album-h3">Album: {playerInfo().album}</h3>
            </Show>
          </div>
        </div>

        <div class={styles.controller}>
          <div class={styles["controller-left"]}>
            <button onClick={Previous}><img src={PreviousButton} /></button>
            <button onClick={BackFive}><img src={Prev5Button} /></button>
          </div>

          <div class={styles["controller-center"]}>
            <p class={styles.position}>{playerInfo().position}</p>
            <div class={styles.progressBar}></div>
            <p class={styles.length}>{playerInfo().length}</p>
          </div>

          <div class={styles["controller-right"]}>
            <button onClick={NextFive}><img src={Next5Button} /></button>
            <button onClick={Next}><img src={NextButton} /></button>
          </div>
        </div>
      </header>
    </div>
  );
};

function Previous() {
  fetch("/api/previous").then(function () {
    console.log("Previous K");
  });
}

function PlayPause() {
  fetch("/api/play-pause").then(function () {
    console.log("Play/Pause K");
  });
}

function Next() {
  fetch("/api/next").then(function () {
    console.log("Next K");
  });
}

function BackFive() {
  fetch("/api/fiveminus").then(function () {
    console.log("5 Seconds Minus K");
  });
}

function NextFive() {
  fetch("/api/fiveplus").then(function () {
    console.log("5 Seconds Plus K");
  });
}

export default App;
