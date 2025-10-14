import { onCleanup, onMount } from "solid-js";
import { fetchThem } from "./fetcher";
import { keys } from "./keys";

export function onStart() {
    onMount(() => {
        fetchThem();

        window.addEventListener("keydown", keys);

        onCleanup(() => {
            window.removeEventListener("keydown", keys)
        })

    });
}