<script lang="ts">
  import PlayerHeader from '$lib/components/display/player/PlayerHeader.svelte';

  import type { PlayerPreview } from '$lib/schema.js';

  let { data } = $props();

  let player: PlayerPreview | null = $state(null);

  // svelte-ignore state_referenced_locally
  player = (await data.playerPreview) ?? null;
</script>

{#await data.playerPreview then _}
  {#if player}
    <PlayerHeader {player} selected={player.preferred_class} />
  {/if}
{/await}
