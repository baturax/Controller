export function Previous() {
  fetch("/api/previous").then(function () {
    console.log("Previous K");
  });
}

export function PlayPause() {
  fetch("/api/play-pause").then(function () {
    console.log("Play/Pause K");
  });
}

export function Next() {
  fetch("/api/next").then(function () {
    console.log("Next K");
  });
}

export function BackFive() {
  fetch("/api/rewind-5-sec").then(function () {
    console.log("5 Seconds Minus K");
  });
}

export function NextFive() {
  fetch("/api/forward-5-sec").then(function () {
    console.log("5 Seconds Plus K");
  });
}

export function IncVolume() {
  fetch("/api/volume-up-5").then(function(){
    console.log("+5 volume okay");
  })
}

export function DecVolume() {
  fetch("/api/volume-down-5").then(function(){
    console.log("-5 volume okay");
  })
}