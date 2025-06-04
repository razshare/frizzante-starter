import Welcome from './Welcome.svelte'
import Todos from './Todos.svelte'
import type {Component} from "svelte";

export const views: Record<string, Component> = {
    "Welcome": Welcome,
    "Todos": Todos,
}