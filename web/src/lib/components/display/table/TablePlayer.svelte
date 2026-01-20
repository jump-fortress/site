<script lang="ts">
  import Flag from '../Flag.svelte';

  import type { Player, PlayerPreview } from '$lib/schema';

  type Props = {
    player: Player | PlayerPreview;
    flag?: boolean;
  };

  let { player, flag = true }: Props = $props();
</script>

<div class="group flex size-full cursor-pointer items-center gap-1 rounded-box hover:bg-base-900">
  {#if flag}
    <Flag code={player.country_code} />
  {/if}
  <img class="ml-1 size-8 rounded-box" src={player.steam_avatar_url} alt="" loading="lazy" />
  <span class="truncate group-hover:underline">{player.display_name}</span>
</div>
