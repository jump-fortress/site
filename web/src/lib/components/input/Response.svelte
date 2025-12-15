<script lang="ts">
  import { fade } from 'svelte/transition';

  let { response = $bindable() }: { response: InputError } = $props();
  let fadeMessage = $state(false);

  // fade out success messages
  if (!response.error) {
    setTimeout(() => {
      fadeMessage = true;
    }, 5_000);
  }
</script>

{#if response.message}
  <span
    in:fade
    class="text-base transition-opacity
  {response.error === true ? 'text-ctp-red' : 'text-ctp-lavender'} 
    {fadeMessage ? 'opacity-0' : ''}">{response.message}</span
  >
{/if}
