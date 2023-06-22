<script lang="ts">
  import { faTrash } from "@fortawesome/free-solid-svg-icons";
  import Fa from "svelte-fa";

  import { send, receive } from "./animation";
  import fetchData from "^/fetctData";
  import { todos, dones } from "~/store";

  export let todo: {
    ID: number;
    subject: string;
  };

  async function refreshData() {
    const { todoData, doneData } = await fetchData();
    $todos = todoData.todos;
    $dones = doneData.dones;
  }

  function removetodo(id: number) {
    fetch(`http://localhost:8080/todo/${id}`, {
      method: "DELETE",
    }).then(async () => {
      await refreshData();
    });
  }

  export function addDoneItem(id: number, todo: string) {
    removetodo(id);
    fetch("http://localhost:8080/done", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({ subject: todo }),
    }).then(async () => {
      await refreshData();
    });
  }
</script>

<label
  in:receive={{ key: todo.ID }}
  out:send={{ key: todo.ID }}
  class=" relative mx-auto mb-4 flex w-[56%] justify-between border-2 border-green-500 bg-green-300 px-4 py-2 text-xl hover:bg-green-400"
>
  <div class="flex items-center justify-center">
    <div>{todo.subject}</div>
  </div>

  <div class="flex flex-row space-x-4">
    <button
      type="button"
      class="flex items-center space-x-2 border border-black bg-orange-300 px-3"
      on:click={() => addDoneItem(todo.ID, todo.subject)}
    >
      <svg
        xmlns="http://www.w3.org/2000/svg"
        class="h-4 w-4"
        viewBox="0 0 20 20"
        fill="currentColor"
      >
        <path
          d="M10 18a8 8 0 100-16 8 8 0 000 16zm3.707-9.293l-4 4a1 1 0 01-1.414 0l-2-2a1 1 0 011.414-1.414L9 10.586l3.293-3.293a1 1 0 011.414 1.414z"
        />
      </svg>
      <span class="text-base">DONE</span>
    </button>
    <button
      on:click={() => {
        removetodo(todo.ID);
      }}
    >
      <div class="flex items-center">
        <Fa icon={faTrash} />
      </div>
    </button>
  </div>
</label>
