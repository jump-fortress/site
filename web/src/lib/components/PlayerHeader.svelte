<script lang="ts">
  import plaza from '$lib/assets/static/players/plaza.png';
  import steam from '$lib/assets/static/players/steam.png';
  import tempus from '$lib/assets/static/players/tempus.png';
  import ClassSelect from '$lib/components/ClassSelect.svelte';
  import Points from '$lib/components/Points.svelte';
  import type { PlayerPreview, Player, PlayerPoints } from '$lib/schema';
  import Flag from '$lib/components/Flag.svelte';
  import ExternalLink from '$lib/components/ExternalLink.svelte';

  type Props = {
    player: PlayerPreview | Player;
    points?: PlayerPoints | null;
    modView?: boolean;
  };

  let { player, points = null, modView = false }: Props = $props();
  let selected_class = $derived(player.preferred_class);

  // a Player was passed if we're in modView
  let p: Player = $derived(player);
</script>

<div class="relative flex h-46 gap-4 text-shadow-sm/50">
  <img class="size-46" src={player.steam_avatar_url} alt="" />
  <!--
	<img
		style="image-rendering: pixelated"
		src={zigzagoon}
		alt=""
		class="scale-200 absolute bottom-0"
	/>
  -->

  <div class="flex h-full flex-col justify-between">
    <!-- name, div, country -->
    <div class="flex flex-col">
      <div class="flex items-end gap-2">
        {#if modView}
          <a href="/players/{player.id}" class="w-fit"
            ><span class="text-3xl/8 hover:underline">{player.display_name}</span></a>
        {:else}
          <span class="text-3xl/8">{player.display_name}</span>
        {/if}
        {#if modView && p.discord_id}
          <div class="flex items-center gap-1 opacity-75">
            <span class="icon-[ri--discord-fill] size-6"></span>
            <span>{p.discord_id}</span>
          </div>
        {/if}
      </div>
      {#if selected_class === 'Soldier'}
        <!-- no division -->
        {#if player.soldier_division}
          <span class="text-division-{player.soldier_division.toLowerCase()}">
            {player.soldier_division} Soldier
          </span>
        {:else}
          <span class="text-ctp-lavender-50/75">Divisionless Soldier</span>
        {/if}
      {:else}
        <!-- no division -->
        {#if player.demo_division}
          <span class="text-division-{player.demo_division.toLowerCase()}">
            {player.demo_division} Demo
          </span>
        {:else}
          <span class="text-ctp-lavender-50/75">Divisionless Demo</span>
        {/if}
      {/if}

      {#if player.country}
        <div class="flex gap-1">
          <Flag code={player.country_code} />
          <span class="text-base text-ctp-lavender-50/75">{player.country}</span>
        </div>
      {/if}
    </div>

    <!-- links and points -->
    <div class="flex flex-col gap-1">
      {#if player.tempus_id}
        <div class="flex">
          <ExternalLink
            src={plaza}
            href={`https://tempusplaza.xyz/players/${player.tempus_id}`}
            name="Plaza" />
          <ExternalLink
            src={tempus}
            href={`https://tempus2.xyz/players/${player.tempus_id}`}
            name="Tempus" />
          {#if modView}
            <ExternalLink
              src={steam}
              href={`https://steamcommunity.com/profiles/${p.id}`}
              name="Steam" />
          {/if}
        </div>
      {/if}

      {#if points}
        <Points {selected_class} {points} launcher={player.preferred_launcher} />
      {/if}
    </div>
  </div>

  <!-- class select -->
  <ClassSelect bind:selected_class />
</div>
