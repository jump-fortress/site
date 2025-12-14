<script lang="ts">
  import Response from '$lib/components/input/Response.svelte';

  type Props = {
    label: string;
    responseMessage?: string;
    onSelect: () => Promise<InputError>;
  };

  let { label, responseMessage = '', onSelect }: Props = $props();

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
    class="size-fit cursor-pointer rounded-lg border-2 border-jfgray-700 border-b-ctp-lavender/50 bg-jfgray-800 px-2 py-1 text-base transition-colors hover:bg-ctp-lavender-950/50"
    onclick={async () => {
      response = onSelect();
    }}>{label}</button
  >

  {#await response then response}
    <Response {response} />
  {/await}
</div>
