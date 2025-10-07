import type { Component } from "solid-js";

import logo from "./logo.svg";
import styles from "./App.module.css";

const App: Component = () => {
  return (
    <div class={styles.App}>
      <header class={styles.header}>
        <button onClick={Previous}>Previous</button>
        <button onClick={PlayPause}>Play/Pause</button>
        <button onClick={Next}>Next</button>
        <br />
        <button onClick={BackFive}>BackFive</button>
        <h1>Song Info Placeholder</h1>
        <button onClick={NextFive}>NextFive</button>
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
