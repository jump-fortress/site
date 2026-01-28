<script lang="ts">
  import type { Snippet } from 'svelte';

  type Props = {
    label: string;
    children: Snippet;
  };

  let { label, children }: Props = $props();
</script>

<!-- label -->
<div class="flex items-center gap-1">
  <span class="text-content/50 select-none">></span>
  <span class="text-nowrap text-primary">{label}</span>
</div>
<div class="ml-3 flex flex-col leading-tight">
  {@render children()}
</div>
