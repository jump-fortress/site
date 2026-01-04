<script lang="ts">
  import soldier from '$lib/assets/tf/soldier.png';
  import { compareDivisions } from '$lib/src/divisions';

  import DivisionTag from './DivisionTag.svelte';

  import type { MonthlyInputBody } from '$lib/schema';

  // a <"map name", "divisions[]"> map is needed for the preview
  function cdToMaps(cd: MonthlyInputBody['divisions']) {
    const maps: Map<string, string[]> = new Map();
    if (!cd) return maps;
    for (const { division, map } of cd) {
      const divisions = maps.get(map) ?? [];
      divisions.push(division);
      maps.set(map, divisions);
    }

    // sort divisions
    for (const [map, divisions] of maps) {
      maps.set(map, divisions.sort(compareDivisions));
    }
    return maps;
  }

  type Props = {
    monthly: MonthlyInputBody;
    header?: boolean;
  };

  let { monthly, header = false }: Props = $props();
  let maps = $derived(cdToMaps(monthly.divisions));
  $inspect(maps);

  const twCols = new Map([
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

<!-- container -->
<!-- todo: link to competition page for maps -->
<div
  class="grid h-48 w-full cursor-pointer items-end justify-center overflow-hidden bg-base-900 drop-shadow-md/50
  {header ? 'absolute top-70 left-0 rounded-t-layout' : 'relative rounded-layout'}
  {twCols.get(maps.size)}">
  {#each maps as [map, divisions]}
    <!-- map wrapper -->
    <div class="relative flex size-full items-end justify-center">
      <!-- absolute map bg image -->
      <img
        class="absolute z-10 h-48 w-full scale-105 object-cover brightness-90 not-first:mask-x-from-98% not-last:mask-x-from-98%"
        src="https://tempusplaza.xyz/map-backgrounds/{map}.jpg"
        alt=""
        draggable="false" />
      <div class="relative z-10 flex flex-col items-center gap-1 p-2">
        <span class="truncate text-lg text-shadow-sm/100">{map}</span>
        <div class="flex gap-2">
          {#each divisions as division}
            <DivisionTag div={division} />
          {/each}
        </div>
      </div>
    </div>
  {/each}
  <!-- absolute details container -->
  <div class="absolute top-0 z-10 flex w-full justify-between p-2">
    <!-- competition name -->
    <div class="flex items-center gap-1">
      <img class="filter-lavender w-10" src={soldier} alt="" draggable="false" />
      <span class="text-lg text-shadow-sm/100">monthly #1</span>
    </div>
    <!-- date / prizepool -->
    <div class="flex flex-col items-end text-shadow-sm/100">
      <div class="flex items-center gap-1">
        <span>12/12/26 12:00 PM UTC-7</span>
        <span class="icon-[mdi--calendar]"></span>
      </div>
      <div class="flex items-center gap-1">
        <span>400 keys</span>
        <span class="icon-[mdi--key]"></span>
      </div>
    </div>
  </div>
</div>

{#if header}
  <hr class="invisible h-48 w-full" />
{/if}

<style>
  span {
    cursor: text;
  }
</style>
