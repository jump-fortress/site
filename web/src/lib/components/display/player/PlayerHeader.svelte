<script lang="ts">
  import plaza from '$lib/assets/components/profile/plaza.png';
  import tempus from '$lib/assets/components/profile/tempus.png';
  import no_map from '$lib/assets/no_map.png';
  import Flag from '$lib/components/display/Flag.svelte';
  import Link from '$lib/components/display/Link.svelte';
  import SelectClass from '$lib/components/input/select/SelectClass.svelte';

  import RocketLauncher from '../RocketLauncher.svelte';
  import Points from './PlayerPoints.svelte';

  import type { Player, PlayerPoints, PlayerPreview } from '$lib/schema';

  type Props = {
    player: Player | PlayerPreview;
    points?: PlayerPoints | null;
    selected: string;
  };

  let { player, points = null, selected = $bindable(player.preferred_class) }: Props = $props();
</script>

<!-- container -->
<div
  class="absolute top-0 left-0 flex h-64 w-full flex-col overflow-hidden rounded-t-layout bg-base-900">
  <!-- absolute avatar -->
  <img
    class="absolute top-1/3 left-6 z-10 size-32 rounded-layout drop-shadow-md/50"
    src={player.steam_avatar_url}
    alt="" />
  <!-- absolute display name -->
  <div class="absolute top-36 left-0 z-10 ml-42">
    <span class="rounded-layout bg-base-900 px-2 py-1 pb-12 text-2xl text-primary"
      >{player.display_name}</span>
  </div>
  <!-- absolute bottom, class select & points -->
  <div class="absolute right-6 bottom-1 z-20 flex flex-row-reverse items-end">
    <div class="flex">
      <SelectClass
        selected={player.preferred_class}
        onsubmit={(value) => {
          selected = value;
        }} />
    </div>
    {#if points}
      {@const selected_points = selected === 'Soldier' ? points.soldier : points.demo}
      <Points points={selected_points} />
      <div class="relative bottom-5 left-7 -rotate-60">
        <RocketLauncher launcher={'Stock'} />
      </div>
    {/if}
  </div>

  <!-- top, map bg image -->
  {#if player.preferred_map}
    <img
      class="h-2/3 w-full mask-b-from-95% object-cover object-center"
      src="https://tempusplaza.xyz/map-backgrounds/{player.preferred_map}.jpg"
      alt="" />
  {:else}
    <div class="h-2/3 w-full mask-b-from-95%">
      <div
        class="filter-purelavender bg-contain bg-repeat size-full scale-300 rotate-6 opacity-50 relative left-48 top-4"
        style:background-image={`url(${no_map})`}>
      </div>
    </div>
  {/if}
  <!-- middle, country & division -->
  <div class="relative z-10 ml-38 h-1/6 px-3 pt-3">
    <div class="flex items-center gap-1 text-content/75">
      {#if player.country_code}
        <Flag code={player.country_code} country={player.country} />
        <span class="icon-[mdi--circle-medium] size-3.5"></span>
      {/if}
      {#if selected === 'Soldier'}
        <span class="text-division-{player.soldier_division?.toLowerCase() ?? ''}">
          {player.soldier_division ?? 'Divisionless'} Soldier
        </span>
        <!-- demo -->
      {:else}
        <span class="text-division-{player.demo_division?.toLowerCase() ?? ''}">
          {player.demo_division ?? 'Divisionless'} Demo
        </span>
      {/if}
    </div>
  </div>
  <!-- bottom, links -->
  <div class="ml-6 flex grow items-center gap-2 text-sm">
    {#if player.tempus_id}
      <Link
        label="Tempus"
        src={tempus}
        href={`https://tempus2.xyz/players/${player.tempus_id}`}
        newTab={true} />
      <Link
        label="Plaza"
        src={plaza}
        href={`https://tempusplaza.xyz/players/${player.tempus_id}`}
        newTab={true} />
    {/if}
  </div>
</div>

<hr class="invisible mt-4 h-64 w-full" />
