<script lang="ts">
  import type { Snippet } from 'svelte';

  type Props = {
    label: string;
    children: Snippet;
    indent?: number;
  };

  let { label, indent = 0, children }: Props = $props();
</script>

<div class="flex flex-col gap-1">
  <!-- label -->
  <div class="relative flex items-center gap-1">
    <span class="text-lg text-content/50">></span>
    <!-- <hr class="mt-1 rounded-box text-base-700" style:width={`${indent / 4}rem`} /> -->
    <span class="text-lg text-nowrap text-primary">{label}</span>
    <hr class="mt-1 w-full rounded-box text-base-700" />
  </div>

  <!-- inner content -->
  {@render children()}
</div>
