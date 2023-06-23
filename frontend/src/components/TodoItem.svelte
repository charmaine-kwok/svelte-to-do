<script lang="ts">
  import { faTrash, faPenToSquare } from "@fortawesome/free-solid-svg-icons";
  import Fa from "svelte-fa";

  import { send, receive } from "./animation";
  import fetchData from "^/fetctData";
  import { todos, dones } from "~/store";

  export let todo: {
    ID: number;
    subject: string;
  };

  let isEditing = false;

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

  function updatetodo(id: number, todo: string) {
    fetch(`http://localhost:8080/todo/${id}`, {
      method: "PATCH",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({ subject: todo }),
    }).then(async () => {
      await refreshData();
    });
  }

  function addDoneItem(id: number, todo: string) {
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
  class=" relative mx-auto mb-4 flex h-[3rem] w-[56%] justify-between border-2 border-green-500 bg-green-300 px-4 py-2 text-xl hover:bg-green-400"
>
  <div class="flex w-full items-center justify-start">
    {#if isEditing}
      <input
        type="text"
        bind:value={todo.subject}
        class="h-[2.5rem] w-[90%]"
        on:keyup={(event) => {
          if (event.key === "Enter") {
            isEditing = false;
            updatetodo(todo.ID, todo.subject);
          }
        }}
      />
    {:else}
      <div>{todo.subject}</div>
    {/if}
  </div>

  <div class="flex flex-row items-center justify-center space-x-4">
    <button
      on:click={() => {
        isEditing = !isEditing;
      }}
    >
      <Fa icon={faPenToSquare} />
    </button>

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
      <Fa icon={faTrash} />
    </button>
  </div>
</label>
