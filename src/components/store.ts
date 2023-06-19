import { writable, type Writable } from "svelte/store";

type todoItem = { id: number; desc: string };

export let todos: Writable<todoItem[]> = writable([]);

export let dones: Writable<todoItem[]> = writable([]);

export let todoId: Writable<number> = writable(0);

export let doneId: Writable<number> = writable(0);
