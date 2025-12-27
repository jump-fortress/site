<script lang="ts">
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
    class="relative flex h-64 w-full items-end justify-center rounded-2xl bg-cover bg-center py-2 text-shadow-sm/50"
    style:background-image={`url(https://tempusplaza.xyz/map-backgrounds/${maps[0]!.name}.jpg)`}>
    <!-- competition info -->
    <div class="absolute z-10 flex size-full justify-between">
      <div class="size-32 bg-ctp-lavender/50"></div>
      <div class="size-32 bg-ctp-lavender/50"></div>
    </div>
    <!-- maps & divisions -->
    <div class="relative flex flex-col items-center">
      <span>{maps[0]!.name}</span>
      <div class="flex gap-2">
        {#each maps[0]!.divisions as div}
          <DivisionTag {div} />
        {/each}
      </div>
    </div>
  </div>
  <!-- 2-6 maps -->
{:else}
  <div
    class="relative grid h-64 w-full grid-rows-1 items-center overflow-hidden rounded-2xl bg-jfgray-700 text-shadow-sm/50
     {twCols.get(maps.length)}">
    <!-- competition info -->
    <div class="absolute z-10 flex size-full justify-between">
      <div class="size-32 bg-ctp-lavender/50"></div>
      <div class="size-32 bg-ctp-lavender/50"></div>
    </div>
    <!-- maps & divisions -->
    {#each maps as map, i}
      <div class="relative flex size-full h-10/9 items-center justify-start">
        <div
          class="absolute flex h-6/4 rotate-10 items-center justify-center overflow-hidden border-2 border-ctp-lavender-50/25
          {i === 0 ? ' -left-6 w-5/4' : i === mapsLength - 1 ? ' w-5/4' : ' z-10 w-full'}">
          <!-- slight clipping without extra height -->
          <div
            class="absolute flex h-65 w-4/3 -rotate-10 items-end justify-center bg-cover bg-center py-2"
            style:background-image={`url(https://tempusplaza.xyz/map-backgrounds/${map.name}.jpg)`}>
            <!-- inside stuff, it is centered but left offset is needed since content is at the bottom of the slant -->
            <!-- note: width is hard set to a fraction of this component's assumed width, not relative (960px) -->
            <!-- 50px is subtracted due to the slant -->
            <div class="relative -left-4 flex w-full flex-col items-center">
              <span>{map.name}</span>
              <div
                class="relative flex flex-wrap justify-center gap-2"
                style:width={`${(1 / maps.length) * 960 - 25}px`}>
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
