<script lang="ts">
  import { page } from '$app/state';
  import { Client } from '$lib/api/api';
  import PlayerHeader from '$lib/components/display/player/PlayerHeader.svelte';
  import { ApiPaths } from '$lib/schema';

  const player_id = $derived(page.params.id);
</script>

{#if player_id}
  {#await Client.GET(ApiPaths.get_player, { params: { path: { player_id: player_id } } })}
    <span></span>
  {:then { data: player }}
    {#if player}
      <PlayerHeader {player} />
    {:else}
      <span>no player :(</span>
    {/if}
  {/await}
{/if}
