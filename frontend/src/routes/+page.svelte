<script lang="ts">
  import { Input } from "postcss";
  import { DarkMode, Select, Label } from "flowbite-svelte";
  import { onMount } from "svelte";

  import fetchData from "^/fetctData";
  import TodoItem from "~/TodoItem.svelte";
  import DoneItem from "~/DoneItem.svelte";
  import { todos, dones } from "~/store";

  let selected: string = "all";

  let options = [
    { value: "all", name: "All" },
    { value: "done", name: "Done" },
    { value: "todo", name: "Todo" },
  ];

  let selectedFilter = "all";

  async function refreshData() {
    const { todoData, doneData } = await fetchData();
    $todos = todoData.todos;
    $dones = doneData.dones;
  }

  export function add(todo: string) {
    fetch("http://localhost:8080/todo", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({ subject: todo }),
    }).then(async () => {
      await refreshData();
    });
  }

  let addTodoItem: string = "";

  onMount(async () => {
    await refreshData();
  });
</script>

<div class="text-center">
  <div class="flex justify-center py-4">
    <h1 class="float-left mx-2 text-4xl font-bold dark:text-white">
      Todo List
    </h1>
    <DarkMode class="float-left text-2xl" />
  </div>
  <div class="mx-auto w-[70%] rounded-md bg-slate-300 dark:bg-slate-700">
    <input
      on:keyup={(event) => {
        if (event.key === "Enter" && addTodoItem) {
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

  <div class="flex items-center justify-center">
    <div class="my-4 flex w-[56%] items-center justify-start text-start">
      <div class="w-[25%]">
        <Label
          >Filter
          <Select
            class="mt-2 w-[100px]"
            items={options}
            bind:value={selected}
            placeholder="Select Filter..."
          />
        </Label>
      </div>
    </div>
  </div>

  <div>
    {#if selected === "all" || selected === "todo"}
      {#each $todos as todo}
        <TodoItem {todo} />
      {/each}
    {/if}
    {#if selected === "all" || selected === "done"}
      {#each $dones as done}
        <DoneItem {done} />
      {/each}
    {/if}
  </div>
</div>
