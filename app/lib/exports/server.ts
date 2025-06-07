import Welcome from '$lib/components/views/Welcome.svelte'
import Todos from '$lib/components/views/Todos.svelte'
import type {Component} from "svelte";

export const views: Record<string, Component> = {
    "Welcome": Welcome,
    "Todos": Todos,
}