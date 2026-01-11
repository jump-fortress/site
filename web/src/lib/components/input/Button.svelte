<script lang="ts">
  import Response from '$lib/components/input/Response.svelte';

  import type { Snippet } from 'svelte';

  type Props = {
    message?: string;
    table?: boolean;
    children: Snippet;
    onsubmit: () => Promise<InputResponse>;
  };

  let { children, table = false, message = '', onsubmit }: Props = $props();

  let response: Promise<InputResponse> = $derived(
    Promise.resolve({ error: false, message: message })
  );
</script>

<button
  class="relative flex h-10 w-fit cursor-pointer items-center gap-1 rounded-box border-base-700 bg-base-800 px-2 transition-colors hover:border-content/50 hover:bg-base-900
  {table ? 'border' : 'border border-b-content/50'}"
  onclick={() => {
    response = onsubmit();
  }}>
  {@render children()}
  <Response {response} />
</button>
