<script lang="ts">
  import { faTrash } from "@fortawesome/free-solid-svg-icons";
  import Fa from "svelte-fa";

  import { todos, dones, doneId } from "./store";
  import { send, receive } from "./animation";

  export let todo: {
    id: number;
    desc: string;
  };

  export function removetodo(id: number) {
    $todos = $todos.filter((t) => t.id !== id);
  }

  export function addDoneItem(id: number, todo: string) {
    removetodo(id);
    $dones = [...$dones, { id: $doneId, desc: todo }];
    $doneId += 1;
  }
</script>

<label
  in:receive={{ key: todo.id }}
  out:send={{ key: todo.id }}
  class=" relative mx-auto mb-4 flex w-[56%] justify-between border-2 border-green-500 bg-green-300 px-4 py-2 text-xl hover:bg-green-400"
>
  <div class="flex items-center justify-center">
    <input
      on:click={() => addDoneItem(todo.id, todo.desc)}
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
