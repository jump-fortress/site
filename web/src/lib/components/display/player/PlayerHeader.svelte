<script lang="ts">
  import { ApiPaths, type Player } from '$lib/schema';
  import ClassSelect from '$lib/components/display/ClassSelect.svelte';
  import Div from '$lib/components/display/Div.svelte';
  import ExternalLink from '$lib/components/display/ExternalLink.svelte';
  import Flag from '$lib/components/display/Flag.svelte';
  import tempus from '$lib/assets/components/profile/tempus.png';
  import plaza from '$lib/assets/components/profile/plaza.png';

  type Props = {
    player: Player;
    class_pref?: string;
  };

  let { player, class_pref = $bindable(player.class_pref) }: Props = $props();
</script>

<div class="relative h-56 flex-col bg-base-800 inset-shadow-sm inset-shadow-base-700">
  <div
    class="h-36 w-full mask-b-from-98% bg-cover bg-center"
    style:background-image={`url("https://tempusplaza.com/map-backgrounds/jump_escape_rc4.jpg")`}>
  </div>
  <!-- avatar -->
  <img
    class="absolute top-22 left-4 z-10 h-24 rounded-box object-cover"
    src={player.avatar_url}
    alt="" />
  <div class="relative -top-4 flex flex-col gap-1">
    <span class="w-fit rounded-tr-box bg-base-800 pr-2 pl-30 text-lg">{player.alias}</span>
    <div class="ml-30 flex items-center gap-2">
      <!-- div -->
      {#if class_pref === 'Soldier'}
        {#if player.soldier_div}
          <Div div={player.soldier_div} playerClass="Soldier" />
        {:else}
          <Div div="Divless" playerClass="Soldier" />
        {/if}
      {:else if player.demo_div}
        <Div div={player.demo_div} playerClass="Demo" />
      {:else}
        <Div div="Divless" playerClass="Demo" />
      {/if}
      <div class="flex gap-1">
        <Flag code={player.country_code} country={player.country} />
      </div>
    </div>
    {#if player.tempus_id}
      <div class="mt-2 ml-4 flex gap-2">
        <ExternalLink
          label="Tempus"
          src={tempus}
          href={`https://tempus2.xyz/players/${player.tempus_id}`}
          newTab={true} />
        <ExternalLink
          label="Plaza"
          src={plaza}
          href={`https://tempusplaza.com/players/${player.tempus_id}`}
          newTab={true} />
      </div>
    {/if}

    <!-- class select -->
    <div class="absolute top-4 right-2 ml-auto flex">
      <ClassSelect bind:player_class={class_pref} />
    </div>
  </div>
</div>
