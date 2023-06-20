<script lang="ts">
  import Fa from "svelte-fa";
  import { faTrash } from "@fortawesome/free-solid-svg-icons";

  import { send, receive } from "./animation";
  import fetchData from "^/fetctData";
  import { dones } from "~/store";

  export let done: {
    ID: number;
    subject: string;
  };

  async function refreshData() {
    const { doneData } = await fetchData();
    $dones = doneData.dones;
  }

  function removeDone(id: number) {
    fetch(`http://localhost:8080/done/${id}`, {
      method: "DELETE",
    }).then(async () => {
      await refreshData();
    });
  }
</script>

<label
  in:receive={{ key: done.ID }}
  out:send={{ key: done.ID }}
  class="relative mx-auto mb-4 flex w-[56%] justify-between border-2 border-red-500 bg-red-300 px-4 py-2 text-xl hover:bg-red-400"
>
  <div class="flex items-center justify-center">
    <div>{done.subject}</div>
  </div>
  <button
    on:click={() => {
      removeDone(done.ID);
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
