import { createSignal, onMount, Show, type Component } from "solid-js";

import styles from "./App.module.css";

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
    fetch("http://192.168.1.161:31313/api/info")
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
  });

  setInterval(() => {
    bai();
  }, 1000);

  return (
    <div class={styles.App}>
      <header class={styles.header}>
        <button id="play-pause-btn" onClick={PlayPause}>
          Play/Pause
        </button>

        <div id="media-box">
          <h1>{playerInfo().title}</h1>
          <div id="extra-media">
            <Show when={playerInfo().artist}>
              <h3>Artist: {playerInfo().artist}</h3>
            </Show>
            <Show when={playerInfo().album}>
              <h3>Album: {playerInfo().album}</h3>
            </Show>
          </div>
        </div>

        <div id="controller">
          <button onClick={Previous}>Previous</button>
          <button onClick={BackFive}>BackFive</button>
          <p>{playerInfo().position}</p>
          <p>----------------------</p>
          <p>{playerInfo().length}</p>
          <button onClick={Next}>Next</button>
          <button onClick={NextFive}>NextFive</button>
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
