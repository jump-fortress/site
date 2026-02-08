<script lang="ts">
  import type { Leaderboard } from '$lib/schema';
  import Div from '../display/Div.svelte';

  type Props = {
    leaderboards: Leaderboard[];
    selected: number;
  };

  let { leaderboards, selected = $bindable() }: Props = $props();
</script>

<div class="flex overflow-hidden rounded-t-box">
  {#each leaderboards as l}
    <button
      class="flex h-9 grow cursor-pointer items-center justify-center transition-all
          {selected === l.id
        ? `bg-div-${l.div?.toLowerCase()}/25`
        : 'bg-base-800 opacity-50 hover:opacity-100'}"
      onclick={() => {
        selected = l.id;
      }}>
      <Div div={l.div} />
    </button>
  {/each}
</div>
