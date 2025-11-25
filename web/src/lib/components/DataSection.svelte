<script lang="ts">
  import type { Snippet } from 'svelte';

  let { children, title }: { children: Snippet; title: string } = $props();
</script>

<div class="flex flex-col gap-1">
  <div class="flex items-center gap-2">
    <span class="text-ctp-lavender">{title}</span>
    <div class="grow border-y-1 border-jfgray-700"></div>
  </div>
  {@render children()}
</div>
