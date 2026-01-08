<script lang="ts">
  import demo from '$lib/assets/tf/demo.png';
  import soldier from '$lib/assets/tf/soldier.png';
  import { compareDivisions } from '$lib/src/divisions';
  import { Temporal } from 'temporal-polyfill';

  import ClassImage from './ClassImage.svelte';
  import DivisionTag from './DivisionTag.svelte';
  import TableDate from './table/TableDate.svelte';

  import type { Monthly, MonthlyInputBody } from '$lib/schema';

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
    monthly: Monthly;
    header?: boolean;
  };

  let { monthly, header = false }: Props = $props();
  let maps = $derived(cdToMaps(monthly.divisions));

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
  class="grid h-48 w-full items-end justify-center bg-base-900 drop-shadow-md/50
  {header ? 'absolute top-0 left-0' : 'relative rounded-layout'}
  {twCols.get(maps.size)}">
  {#each maps as [map, divisions]}
    <!-- map wrapper -->
    <div class="relative flex size-full items-end justify-center">
      <!-- absolute map bg image -->
      <a
        class="relative flex size-full cursor-pointer overflow-hidden"
        href="/formats/monthly/{monthly.id}">
        <img
          class="over absolute z-10 h-48 w-full scale-105 object-cover brightness-90 not-first:mask-x-from-98% not-last:mask-x-from-98%"
          src="https://tempusplaza.xyz/map-backgrounds/{map}.jpg"
          alt=""
          draggable="false" />
      </a>
      <div class="absolute z-10 flex flex-col items-center gap-1 p-2">
        <span class="z-10 truncate text-lg text-shadow-sm/100">{map}</span>
        <div class="flex gap-2">
          {#each divisions as division}
            <DivisionTag div={division} />
          {/each}
        </div>
      </div>
    </div>
  {/each}
  <!-- absolute details container -->
  <div class="absolute top-0 flex w-full justify-between p-2">
    <!-- competition name -->
    <div class="z-10 flex h-12 items-center gap-1">
      <ClassImage selected={monthly.competition.class} />
      <span class="text-lg text-shadow-sm/100">monthly #{monthly.id}</span>
    </div>
    <!-- date / prizepool -->
    <div class="z-10 flex flex-col items-end text-shadow-sm/100">
      <div class="flex items-center gap-1">
        <span class="relative z-10">
          <TableDate date={monthly.competition.starts_at} />
        </span>
        <span class="icon-[mdi--calendar]"></span>
      </div>
      <div class="flex items-center gap-1">
        <span>? keys</span>
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
