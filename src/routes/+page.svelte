<script lang="ts">
  import { Input } from "postcss";
  import { DarkMode } from "flowbite-svelte";
  import Fa from "svelte-fa";
  import { faTrash } from "@fortawesome/free-solid-svg-icons";

  let addTodoItem: string = "";
  let todoItem: { id: number; desc: string };
  let todos: (typeof todoItem)[] = [];
  let todoId: number = 1;
  let dones: (typeof todoItem)[] = [];
  let doneId: number = 1;

  function removetodo(id: number) {
    todos = todos.filter((t) => t.id !== id);
  }

  function removeDone(id: number) {
    dones = dones.filter((t) => t.id !== id);
  }

  function add(todo: string) {
    todos = [...todos, { id: todoId, desc: todo }];
    todoId += 1;
  }

  function doneItem(id: number, todo: string) {
    removetodo(id);
    dones = [...dones, { id: doneId, desc: todo }];
    doneId += 1;
  }
</script>

<div class="text-center">
  <div class="flex justify-center py-4">
    <h1 class="float-left mx-2 text-4xl font-bold dark:text-white">
      Todo List
    </h1>
    <DarkMode class="float-left text-2xl" />
  </div>
  <div class="mx-auto mb-8 w-[70%] rounded-md bg-slate-300 dark:bg-slate-700">
    <input
      on:keyup={(event) => {
        if (event.key === "Enter") {
          add(addTodoItem);
          addTodoItem = "";
        }
      }}
      type="text"
      bind:value={addTodoItem}
      placeholder="what needs to be done?"
      class="my-8 w-[80%] border-white py-2 pl-2 text-xl dark:border-white"
    />
  </div>
  <div>
    {#each todos as todo}
      <label
        class="mx-auto mb-4 w-[56%] border-2 border-green-500 bg-green-300 px-4 py-2 text-xl hover:bg-green-400"
      >
        <div class="flex items-center justify-center">
          <input
            on:click={() => doneItem(todo.id, todo.desc)}
            type="checkbox"
            class="mr-4"
          />
          <div>{todo.desc}</div>
        </div>
        <button
          on:click={() => {
            removetodo(todo.id);
          }}
        >
          <div class="flex items-center">
            <Fa icon={faTrash} />
          </div>
        </button>
      </label>
    {/each}
    {#each dones as doneitem}
      <label
        class="mx-auto mb-4 w-[56%] border-2 border-red-500 bg-red-300 px-4 py-2 text-xl hover:bg-red-400"
      >
        <div class="flex items-center justify-center">
          <div>{doneitem.desc}</div>
        </div>
        <button
          on:click={() => {
            removeDone(doneitem.id);
          }}
        >
          <div class="flex items-center">
            <Fa icon={faTrash} />
          </div>
        </button>
        <span />
      </label>
    {/each}
  </div>
</div>

<style>
  label {
    display: flex;
    justify-content: space-between;
    position: relative;
  }
  span {
    position: absolute;
    bottom: 50%;
    left: 0;
    right: 0;
    border-bottom: 1px solid red;
  }
</style>
