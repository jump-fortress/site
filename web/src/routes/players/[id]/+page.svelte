<script lang="ts">
  import PlayerHeader from '$lib/components/display/player/PlayerHeader.svelte';

  let { data } = $props();
</script>

{#if data.player}
  <PlayerHeader player={data.player} />
{/if}
