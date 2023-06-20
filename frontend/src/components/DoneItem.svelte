<script lang="ts">
  import Fa from "svelte-fa";
  import { faTrash } from "@fortawesome/free-solid-svg-icons";

  import { dones } from "./store";
  import { send, receive } from "./animation";

  export let doneitem: {
    id: number;
    desc: string;
  };

  function removeDone(id: number) {
    $dones = $dones.filter((t) => t.id !== id);
  }
</script>

<label
  in:receive={{ key: doneitem.id }}
  out:send={{ key: doneitem.id }}
  class="relative mx-auto mb-4 flex w-[56%] justify-between border-2 border-red-500 bg-red-300 px-4 py-2 text-xl hover:bg-red-400"
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
    </div></button
  ><span />
</label>

<style>
  span {
    position: absolute;
    bottom: 50%;
    left: 0;
    right: 0;
    border-bottom: 1px solid red;
  }
</style>
