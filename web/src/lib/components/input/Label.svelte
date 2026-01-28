<script lang="ts">
  import type { Snippet } from 'svelte';

  type Props = {
    label: string;
    max_width: string;
    children: Snippet;
  };

  let { label, max_width, children }: Props = $props();
</script>

<label
  class="relative flex w-full {max_width} mt-1.5 rounded-box border border-base-700 bg-base-800 transition-colors focus-within:border-content/50 focus-within:bg-base-900 hover:border-content/50">
  <!-- label -->
  <span class="absolute -top-1 left-1 bg-base-800 px-2 text-sm leading-1">{label}</span>
  <!-- input container -->
  <div class="relative flex h-9 w-full">
    {@render children()}
  </div>
</label>
