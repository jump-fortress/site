<script lang="ts">
  import { fade } from 'svelte/transition';

  let { response }: { response: InputError } = $props();
</script>

<span
  in:fade
  class="text-base {response.error === true
    ? 'text-ctp-red'
    : 'text-ctp-lavender'} transition-colors">{response.message}</span
>
