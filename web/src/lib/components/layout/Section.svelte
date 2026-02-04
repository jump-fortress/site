<script lang="ts">
  import type { Snippet } from 'svelte';

  type Props = {
    label?: string;
    indent?: string;
    children: Snippet;
  };

  let { label = '', indent = '>', children }: Props = $props();
</script>

<!-- label -->
{#if label}
  <div class="flex items-center gap-1 pt-1 pb-0.5">
    <span class="text-content/50 select-none">{indent}</span>
    <span class="text-nowrap text-primary">{label}</span>
  </div>
{:else}
  <div class="py-1"></div>
{/if}
<div class="mx-3 flex flex-col">
  {@render children()}
</div>
