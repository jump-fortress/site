<script lang="ts">
  import Response from '$lib/components/input/Response.svelte';
  import type { Snippet } from 'svelte';
  import { fade } from 'svelte/transition';

  type Props = {
    onSelect: () => Promise<InputError>;
  };

  let { onSelect }: Props = $props();

  // response is a promise so #await can be used to handle pending state
  let response: Promise<InputError> = $state(
    Promise.resolve({
      error: false,
      message: ''
    })
  );
</script>

<div class="flex items-center gap-2">
  <button
    class=" size-10 cursor-pointer border-b-2 border-b-ctp-lavender/50 bg-jfgray-900 text-base transition-colors hover:bg-jfgray-700"
    onclick={async () => {
      response = onSelect();
    }}>
    {#await response}
      <span in:fade class="icon-[ri--loader-3-line] animate-spin text-ctp-lavender"></span>
    {:then response}
      <span
        in:fade
        class={response.error === true
          ? 'icon-[ri--close-line] text-ctp-red'
          : 'icon-[ri--check-line] text-ctp-lavender'}></span>
    {/await}
  </button>
</div>
