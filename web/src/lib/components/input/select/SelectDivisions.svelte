<script lang="ts">
  import DivisionTag from '$lib/components/display/DivisionTag.svelte';
  import { divisions } from '$lib/src/divisions';

  import type { SvelteSet } from 'svelte/reactivity';

  type Props = {
    label: string;
    withNone?: boolean;
    selected: SvelteSet<string>;
    onsubmit: (value: SvelteSet<string>) => void;
  };

  let options = divisions;

  let { label, withNone = false, selected, onsubmit }: Props = $props();
</script>

<!-- container for input -->
<label
  class="relative mt-2 flex w-fit rounded-box border border-base-700 bg-base-800 transition-colors hover:border-content/50">
  <!-- label -->
  <span class="absolute -top-1 left-2 bg-base-800 px-1 text-sm leading-1">{label}</span>
  <!-- input container -->
  <div class="flex h-10 w-full">
    {#each options as div}
      {@const isSelected: boolean = selected.has(div)}
      <button
        class="w-24 cursor-pointer rounded-box
         {isSelected ? 'bg-base-900' : 'bg-base-800'}"
        onclick={() => {
          if (selected.has(div)) {
            selected.delete(div);
          } else {
            selected.add(div);
          }
          onsubmit(selected);
        }}>
        <DivisionTag {div} />
      </button>
    {/each}
    <!-- reset option -->
    {#if withNone}
      <!-- svelte-ignore a11y_consider_explicit_label -->
      <button
        class="grid w-10 cursor-pointer place-content-center rounded-box"
        onclick={() => {
          selected.clear();
          onsubmit(selected);
        }}>
        <span class="icon-[mdi--close] size-6"></span>
      </button>
    {/if}
  </div>
</label>
