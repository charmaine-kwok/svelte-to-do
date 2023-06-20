import { browser } from "$app/environment";
import { writable, type Writable } from "svelte/store";

type todoItem = { id: number; desc: string };

export let todos: Writable<todoItem[]> = writable([]);

export let dones: Writable<todoItem[]> = writable([]);

export let todoId: Writable<number> = writable(0);

export let doneId: Writable<number> = writable(0);

if (browser) {
  const todosData = localStorage.getItem("todos");
  if (todosData !== null) {
    todos.set(JSON.parse(todosData));
  }

  const donesData = localStorage.getItem("dones");
  if (donesData !== null) {
    dones.set(JSON.parse(donesData));
  }

  if (localStorage.getItem("todoId")) {
    todoId.set(Number(localStorage.getItem("todoId")));
  }

  if (localStorage.getItem("doneId")) {
    doneId.set(Number(localStorage.getItem("doneId")));
  }

  // Subscribe to changes in the stores and update the localStorage
  todos.subscribe((value) => {
    localStorage.setItem("todos", JSON.stringify(value));
  });

  dones.subscribe((value) => {
    localStorage.setItem("dones", JSON.stringify(value));
  });

  todoId.subscribe((value) => {
    localStorage.setItem("todoId", String(value));
  });

  doneId.subscribe((value) => {
    localStorage.setItem("doneId", String(value));
  });
}
