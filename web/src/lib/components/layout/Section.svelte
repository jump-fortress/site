<script lang="ts">
  import type { Snippet } from 'svelte';

  type Props = {
    label: string;
    indent?: string;
    children: Snippet;
  };

  let { label, indent = '>', children }: Props = $props();
</script>

<!-- label -->
<div class="flex items-center gap-1">
  <span class="text-content/50 select-none">{indent}</span>
  <span class="text-nowrap text-primary">{label}</span>
</div>
<div class="mx-3 flex flex-col">
  {@render children()}
</div>
