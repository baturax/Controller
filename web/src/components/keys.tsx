import { Previous, PlayPause, Next, BackFive, NextFive } from "./functions";

export const keys = (e: KeyboardEvent) => {
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