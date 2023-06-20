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
    <input
      on:click={() => addDoneItem(todo.ID, todo.subject)}
      type="checkbox"
      class="mr-4"
    />
    <div>{todo.subject}</div>
  </div>
  <button
    on:click={() => {
      removetodo(todo.ID);
    }}
  >
    <div class="flex items-center">
      <Fa icon={faTrash} />
    </div>
  </button>
</label>
