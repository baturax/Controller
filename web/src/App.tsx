import { createSignal, onCleanup, onMount, Show, type Component } from "solid-js";

import styles from "./App.module.css";

import { Previous, PlayPause, Next, BackFive, NextFive, IncVolume, DecVolume } from "./components/functions";
import { keys } from "./components/keys";

import { fetchThem, playerInfo, setPlayerInfo } from "./components/fetcher";

import PlayButton from "./assets/play.svg";
import PauseButton from "./assets/pause.svg";
import PreviousButton from "./assets/previous.svg";
import NextButton from "./assets/next.svg";
import Prev5Button from "./assets/rewind-5-sec.svg";
import Next5Button from "./assets/forward-5-sec.svg";
import IncVolumeButton from "./assets/volume-up.svg"
import DecVolumeButton from "./assets/volume-down.svg"
import { onStart } from "./components/mount";

const App: Component = () => {

  onStart();

  setInterval(() => {
    fetchThem();
    updateArt();
  }, 500);

  return (
    <div class={styles.App}>
      <header class={styles.header}>
        <img src="./public/art" alt="bai" class={styles.artclass} />

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
          <button onClick={IncVolume}><img src={IncVolumeButton} /></button>
          <button onClick={DecVolume}><img src={DecVolumeButton} /></button>
        </div>
      </header>
    </div>
  );
};

function updateArt() {
  var pic = document.getElementById("art-id") as HTMLImageElement;
  if (pic) {
    pic.src = `./public/art?${Date.now()}`;
  }
}

export default App;
