<script lang="ts">
  import Flag from '../Flag.svelte';

  import type { Player, PlayerPreview } from '$lib/schema';

  type Props = {
    player: Player | PlayerPreview;
  };

  let { player }: Props = $props();
</script>

<div
  class="group flex size-full cursor-pointer items-center gap-2 rounded-box px-2 hover:bg-base-900">
  <Flag code={player.country_code} />
  <img class="ml-2 size-8 rounded-box" src={player.steam_avatar_url} alt="" loading="lazy" />
  <span class="truncate group-hover:underline">{player.display_name}</span>
</div>
