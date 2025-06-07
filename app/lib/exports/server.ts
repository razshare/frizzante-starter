import Welcome from "$lib/views/Welcome.svelte"
import Todos from "$lib/views/Todos.svelte"
import type { Component } from "svelte"

export const views: Record<string, Component> = {
    Welcome: Welcome,
    Todos: Todos,
}
