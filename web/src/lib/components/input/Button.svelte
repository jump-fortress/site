<script lang="ts">
  import Response from '$lib/components/input/Response.svelte';
  import type { Snippet } from 'svelte';

  type Props = {
    children: Snippet;
    responseMessage?: string;
    onSelect: () => Promise<InputError>;
  };

  let { children, responseMessage = '', onSelect }: Props = $props();

  // response is a promise so #await can be used to handle pending state
  let response: Promise<InputError> = $derived(
    Promise.resolve({
      error: false,
      message: responseMessage
    })
  );
</script>

<div class="flex items-center gap-2">
  <button
    class="border-jfgray-700 border-b-ctp-lavender/50 bg-jfgray-800 hover:bg-ctp-lavender-950/50 size-fit cursor-pointer rounded-lg border-2 px-2 py-1 text-base transition-colors"
    onclick={async () => {
      response = onSelect();
    }}>{@render children()}</button
  >

  {#await response}
    <Response response={{ error: false, message: '...' }} />
  {:then response}
    <Response {response} />
  {/await}
</div>
