<script lang="ts">
  import Flag from '../Flag.svelte';

  import type { Player, PlayerPreview } from '$lib/schema';

  type Props = {
    player: Player | PlayerPreview;
    link?: boolean;
    flag?: boolean;
  };

  let { player, flag = true, link = true }: Props = $props();
</script>

{#if link}
  <a
    class="group flex size-full cursor-pointer items-center gap-1 hover:text-primary hover:underline"
    href="/players/{player.id}">
    {#if flag}
      <Flag code={player.country_code} />
      <div class="w-2"></div>
    {/if}
    <img class="ml-1 size-8 rounded-box" src={player.steam_avatar_url} alt="" loading="lazy" />
    <span class="truncate group-hover:underline">{player.display_name}</span>
  </a>
{:else}
  <div class="group flex size-full cursor-pointer items-center gap-1">
    {#if flag}
      <Flag code={player.country_code} />
    {/if}
    <img class="ml-1 size-8 rounded-box" src={player.steam_avatar_url} alt="" loading="lazy" />
    <span class="truncate group-hover:underline">{player.display_name}</span>
  </div>
{/if}
