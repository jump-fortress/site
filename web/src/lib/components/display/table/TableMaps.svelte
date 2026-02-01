<script lang="ts">
  import type { Leaderboard } from '$lib/schema';
  import TableMap from './TableMap.svelte';

  type Props = {
    leaderboards: Leaderboard[];
  };

  let { leaderboards }: Props = $props();
  const maps: Set<string> = $derived(new Set(leaderboards.map((l) => l.map)));
</script>

<div class="flex w-full overflow-hidden">
  {#each maps as map}
    <TableMap {map} />
  {/each}
</div>
