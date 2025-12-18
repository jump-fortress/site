<script lang="ts">
  import type { Snippet } from 'svelte';
  import { fade, slide } from 'svelte/transition';

  interface Props {
    data: object[];
    header: Snippet;
    row: any; // Snippet<object>
  }

  let { data, header, row }: Props = $props();
</script>

<table class="table-fixed rounded-t-2xl bg-jfgray-900">
  <thead>
    <tr class="h-6 border-b-2 border-ctp-lavender/50 text-base leading-4">
      {@render header()}
    </tr>
  </thead>

  <tbody>
    {#each data as d}
      <tr out:fade class="h-10 text-center text-base odd:bg-jfgray-900 even:bg-jfgray-800/25">
        {@render row(d)}
      </tr>
    {/each}
  </tbody>
</table>
