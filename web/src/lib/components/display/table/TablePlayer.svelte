<script lang="ts">
  import Flag from '../Flag.svelte';

  import type { Player } from '$lib/schema';

  type Props = {
    player: Player;
    link?: boolean;
    flag?: boolean;
  };

  let { player, flag = true, link = false }: Props = $props();
</script>

{#if link}
  <a
    class="group flex size-full cursor-pointer items-center gap-1 hover:text-primary hover:underline"
    href="/players/{player.id}">
    {#if flag}
      <Flag code={player.country_code} />
      <div class="w-0.5"></div>
    {/if}
    <img class="ml-1 size-8 rounded-box" src={player.avatar_url} alt="" loading="lazy" />
    <span class="truncate group-hover:underline">{player.alias}</span>
  </a>
{:else}
  <div class="group flex size-full cursor-pointer items-center gap-1">
    {#if flag}
      <Flag code={player.country_code} />
      <div class="w-0"></div>
    {/if}
    <img class="ml-1 size-8 rounded-box" src={player.avatar_url} alt="" loading="lazy" />
    <span class="truncate group-hover:underline">{player.alias}</span>
  </div>
{/if}
