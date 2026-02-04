<script lang="ts">
  import type { Snippet } from 'svelte';

  type Props = {
    table?: boolean;
    children: Snippet;
    onsubmit: () => Promise<boolean>;
  };

  let { children, table = false, onsubmit }: Props = $props();

  let valid: Promise<boolean> = $state(Promise.resolve(true));
</script>

<button
  class="relative flex h-9 max-w-fit cursor-pointer items-center justify-between gap-1 rounded-box border-base-700 bg-base-800 px-2 text-nowrap transition-colors hover:border-content/50 hover:bg-base-900
  {table ? 'border' : 'min-w-40 border border-b-content/50'}"
  onclick={() => {
    valid = onsubmit();
  }}>
  {@render children()}
</button>
