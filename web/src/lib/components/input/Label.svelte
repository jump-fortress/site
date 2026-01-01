<script lang="ts">
  import type { Snippet } from 'svelte';

  type Props = {
    label: string;
    children: Snippet;
  };

  let { label, children }: Props = $props();
</script>

<!-- container for input & response -->
<label
  class="relative mt-2 flex w-full max-w-80 rounded-box border border-base-700 bg-base-800 transition-colors focus-within:border-content/50 focus-within:bg-base-900 hover:border-content/50">
  <!-- label -->
  <span class="absolute -top-1 left-2 bg-base-800 px-1 text-sm leading-1">{label}</span>
  <!-- input container -->
  <div class="relative flex h-10 w-full">
    {@render children()}
  </div>
</label>
