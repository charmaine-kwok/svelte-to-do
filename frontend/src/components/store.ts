import { writable, type Writable } from "svelte/store";

type todoItem = { ID: number; subject: string };

export let todos: Writable<todoItem[]> = writable([]);

export let dones: Writable<todoItem[]> = writable([]);
