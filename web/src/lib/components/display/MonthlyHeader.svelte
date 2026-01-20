<script lang="ts">
  import { compareDivisions } from '$lib/src/divisions';

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
    link?: boolean;
    ends_at?: boolean;
  };

  let { monthly, header = false, link = false, ends_at = true }: Props = $props();
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
<!-- todo: link to competition page for divs -->
<div
  class="group grid h-48 w-full items-end justify-center overflow-hidden bg-base-900
  {header ? 'absolute top-0 left-0 rounded-t-layout' : 'relative rounded-layout'}
  {twCols.get(maps.size)}">
  {#each maps as [map, divisions]}
    <!-- map wrapper -->
    <div class="relative flex size-full items-end justify-center">
      <!-- absolute map bg image -->
      {#if link}
        <a class="relative flex size-full overflow-hidden" href="/formats/monthly/{monthly.id}">
          <img
            class="over absolute z-10 h-48 w-full scale-105 object-cover brightness-75 transition-all not-first:mask-x-from-98% not-last:mask-x-from-98% group-hover:brightness-100"
            src="https://tempusplaza.com/map-backgrounds/{map}.jpg"
            alt=""
            draggable="false" />
        </a>
      {:else}
        <div class="relative flex size-full overflow-hidden">
          <img
            class="over absolute z-10 h-48 w-full scale-105 object-cover brightness-75 transition-all not-first:mask-x-from-98% not-last:mask-x-from-98% group-hover:brightness-100"
            src="https://tempusplaza.com/map-backgrounds/{map}.jpg"
            alt=""
            draggable="false" />
        </div>
      {/if}
      <div
        class="absolute z-10 flex flex-col items-center gap-1
      p-2">
        <span class="z-10 truncate text-lg">{map}</span>
        <div class="flex flex-wrap justify-center gap-2">
          {#each divisions as division}
            <DivisionTag div={division} />
          {/each}
        </div>
      </div>
    </div>
  {/each}
  <!-- absolute details container -->
  <div
    class="absolute top-0 flex w-full justify-between text-shadow-xs/100 text-shadow-base-900
    {header ? 'p-6' : 'p-2'}">
    <!-- competition name -->
    <div class="z-10 flex h-12 items-center gap-1">
      <ClassImage selected={monthly.competition.class} />
      <span class="text-lg">monthly #{monthly.id}</span>
    </div>
    <!-- date / prizepool -->
    <div class="z-10 flex flex-col items-end">
      <div class="flex items-center gap-2">
        <span class="relative z-10">
          <TableDate date={monthly.competition.starts_at} fade={false} />
        </span>
        <span class="mt-auto icon-[mdi--calendar-outline]"></span>
      </div>
      {#if ends_at}
        <div class="flex items-center gap-2">
          <span class="relative z-10">
            <TableDate date={monthly.competition.ends_at} fade={false} />
          </span>
          <span class="icon-[mdi--clock-outline]"></span>
        </div>
      {/if}
      {#if monthly.competition.prizepool}
        <div class="flex items-center gap-1">
          <span>{monthly.competition.prizepool} keys</span>
          <span class="icon-[mdi--key]"></span>
        </div>
      {/if}
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
