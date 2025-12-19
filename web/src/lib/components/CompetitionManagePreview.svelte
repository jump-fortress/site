<script lang="ts">
  import escape from '$lib/assets/static/maps/jump_escape_rc4.jpg';
  import type { Map } from '$lib/schema';
  import DivisionTag from './DivisionTag.svelte';

  type Props = {
    maps: MapWithDivisions[];
  };

  let { maps }: Props = $props();
  let mapsLength = $derived(maps.length);

  let twCols = new Map([
    [1, 'grid-cols-1'],
    [2, 'grid-cols-2'],
    [3, 'grid-cols-3'],
    [4, 'grid-cols-4'],
    [5, 'grid-cols-5'],
    [6, 'grid-cols-6'],
    [7, 'grid-cols-7'],
    [8, 'grid-cols-8']
  ]);
</script>

<!--
@component
preview for competitions
- 1 to 6 maps recommended
-->
{#if mapsLength === 1}
  <div
    class="flex h-64 w-full items-end justify-center rounded-2xl bg-cover bg-center text-shadow-sm/50"
    style:background-image={`url(https://tempusplaza.xyz/map-backgrounds/${maps[0]!.name}.jpg)`}
  >
    <!-- inside stuff -->
    <div class="relative flex flex-col items-center">
      <span>{maps[0]!.name}</span>
      <div class="flex gap-2">
        {#each maps[0]!.divisions as div}
          <DivisionTag {div} />
        {/each}
      </div>
    </div>
  </div>
{:else}
  <div
    class="relative grid grid-rows-1 bg-jfgray-700 text-shadow-sm/50 {twCols.get(
      maps.length
    )} relative h-64 w-full items-center overflow-hidden rounded-2xl"
  >
    {#each maps as map, i}
      <div class="relative flex size-full h-10/9 items-center justify-start">
        <div
          class="h-6/4 rotate-10 {i === 0
            ? '-left-6 w-5/4'
            : i === mapsLength - 1
              ? 'w-5/4'
              : 'z-10 w-full'} absolute flex items-center justify-center overflow-hidden border-2"
        >
          <!-- slight clipping without extra height -->
          <div
            class="absolute flex h-65 w-4/3 -rotate-10 items-end justify-center bg-cover bg-center py-2"
            style:background-image={`url(https://tempusplaza.xyz/map-backgrounds/${map.name}.jpg)`}
          >
            <!-- inside stuff, it is centered but left offset is needed since content is at the bottom of the slant -->
            <div class="relative -left-4 flex flex-col items-center">
              <span>{map.name}</span>
              <div class="flex gap-2">
                {#each map.divisions as div}
                  <DivisionTag {div} />
                {/each}
              </div>
            </div>
          </div>
        </div>
      </div>
    {/each}
  </div>
{/if}
