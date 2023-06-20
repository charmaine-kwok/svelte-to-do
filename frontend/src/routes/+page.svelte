<script lang="ts">
  import { Input } from "postcss";
  import { DarkMode } from "flowbite-svelte";

  import TodoItem from "~/TodoItem.svelte";
  import DoneItem from "~/DoneItem.svelte";
  import { todos, todoId, dones } from "~/store";

  let addTodoItem: string = "";

  function add(todo: string) {
    $todos = [...$todos, { id: $todoId, desc: todo }];
    $todoId += 1;
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
    {#each $todos as todo}
      <TodoItem {todo} />
    {/each}
    {#each $dones as doneitem}
      <DoneItem {doneitem} />
    {/each}
  </div>
</div>
