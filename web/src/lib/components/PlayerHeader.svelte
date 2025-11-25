<script lang="ts">
  import rocket from '$lib/assets/static/classes/rocket.png';
  import sticky from '$lib/assets/static/classes/sticky.png';
  import plaza from '$lib/assets/static/players/plaza.png';
  import steam from '$lib/assets/static/players/steam.png';
  import tempus from '$lib/assets/static/players/tempus.png';
  import zigzagoon from '$lib/assets/static/players/zigzagoon.gif';
  import Link from '$lib/components/Link.svelte';
  import ClassSelect from './ClassSelect.svelte';
  import Points from '../../routes/players/[id]/Points.svelte';
  import type { FullPlayer, PlayerProfile } from '$lib/schema';

  type Props = {
    player: PlayerProfile;
    fullPlayer?: FullPlayer | null;
  };

  let { player, fullPlayer = null }: Props = $props();
  let selected_class = $derived(player.preferred_class);
</script>

<div class="relative flex h-46 gap-4">
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
    <!-- name and div -->
    <div class="flex flex-col">
      {#if fullPlayer}
        <a href="/players/{fullPlayer.id}" class="w-fit"
          ><span class="text-3xl/8 hover:underline">{player.display_name}</span></a
        >
      {:else}
        <span class="text-3xl/8">{player.display_name}</span>
      {/if}
      {#if selected_class === 'Soldier'}
        <!-- no division -->
        {#if player.soldier_division}
          <span class="text-division-{player.soldier_division.toLowerCase()}">
            {player.soldier_division} Soldier
          </span>
        {:else}
          <span>Unranked Soldier</span>
        {/if}
      {:else}
        <!-- no division -->
        {#if player.demo_division}
          <span class="text-division-{player.demo_division.toLowerCase()}">
            {player.demo_division} Demo
          </span>
        {:else}
          <span>Unranked Demo</span>
        {/if}
      {/if}
      {#if fullPlayer && fullPlayer.discord_id}
        <div class="flex items-center gap-1 opacity-75">
          <span class="icon-[ri--discord-fill] size-6"></span>
          <span>{fullPlayer.discord_id}</span>
        </div>
      {/if}
    </div>

    <!-- links and points -->
    <div class="flex flex-col gap-2">
      {#if player.tempus_id}
        <div class="flex">
          {@render externalLink(
            plaza,
            `https://tempusplaza.xyz/players/${player.tempus_id}`,
            'Plaza'
          )}
          {@render externalLink(
            tempus,
            `https://tempus2.xyz/players/${player.tempus_id}`,
            'Tempus'
          )}
          {#if fullPlayer}
            {@render externalLink(
              steam,
              `https://steamcommunity.com/profiles/${fullPlayer.steam_id64}`,
              'Steam'
            )}
          {/if}
        </div>
      {/if}

      <Points {selected_class} soldier={player.soldier_points} demo={player.demo_points} />
    </div>
  </div>

  <!-- class select -->
  <ClassSelect bind:selected_class />
</div>

{#snippet externalLink(src: string, href: string, name: string)}
  <a
    {href}
    class="flex items-end gap-1 pl-2 text-ctp-blue/50 decoration-1 transition-colors first:pl-0 hover:text-ctp-blue hover:underline"
  >
    <img {src} class="size-6" alt="" />
    <span class="flex">{name}</span>
  </a>
{/snippet}
