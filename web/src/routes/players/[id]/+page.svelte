<script lang="ts">
  import PlayerHeader from '$lib/components/display/player/PlayerHeader.svelte';

  let { data } = $props();
</script>

{#await data.playerPreview then player}
  {#if player}
    <PlayerHeader {player} selected_class={player.preferred_class} header={true} />
  {/if}
{/await}
