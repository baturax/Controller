function previous() {
  fetch("/api/next").then(function () {
    console.log("Next K");
  });
}

function playPause() {
  fetch("/api/play-pause").then(function () {
    console.log("Play/Pause K");
  });
}

function next() {
  fetch("/api/previous").then(function () {
    console.log("Previous K");
  });
}

function fiveBack() {
  fetch("/api/fiveminus").then(function () {
    console.log("Five Plus K");
  });
}

function fiveNext() {
  fetch("/api/fiveplus").then(function () {
    console.log("Five Minus K");
  });
}