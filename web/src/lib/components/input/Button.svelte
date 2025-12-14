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
    class="settings-button"
    onclick={async () => {
      response = onSelect();
    }}>{label}</button
  >

  {#await response then response}
    <Response {response} />
  {/await}
</div>
