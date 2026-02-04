<script lang="ts">
  import type { Snippet } from 'svelte';

  type Props = {
    label: string;
    children: Snippet;
  };

  let { label, children }: Props = $props();
  let visible = $state(false);
</script>

<button
  class="my-1 flex w-fit cursor-pointer items-center gap-1 pb-0.5"
  onclick={() => {
    visible = !visible;
  }}>
  <span class="w-2 text-primary select-none">{visible ? 'v' : '>'}</span>
  <span class="text-nowrap text-primary {!visible ? 'underline' : ''}">{label}</span>
</button>

{#if visible}
  <div class="flex flex-col">
    {@render children()}
  </div>
{/if}
