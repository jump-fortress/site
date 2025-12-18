<script lang="ts">
  import type { Player, PlayerPreview } from '$lib/schema';

  let { player }: { player: Player | PlayerPreview } = $props();
</script>

<div class="flex items-center gap-2">
  <img src={player.steam_avatar_url} alt="" class="size-8 rounded-lg" />
  <span class="span-ellipsis">{player.display_name}</span>
</div>
