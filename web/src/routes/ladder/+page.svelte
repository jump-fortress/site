<script lang="ts">
  import { Client } from '$lib/api/api';
  import { ApiPaths } from '$lib/schema';
</script>

{#await Client.GET(ApiPaths.get_players)}
  <span></span>
{:then { data: players }}
  <span>under construction</span>
  {#each players as player}
    <span>{player.alias}</span>
  {/each}
{/await}
