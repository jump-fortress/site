<script lang="ts">
  import DivisionTag from '$lib/components/display/DivisionTag.svelte';
  import RocketLauncher from '$lib/components/display/RocketLauncher.svelte';
  import Table from '$lib/components/display/table/Table.svelte';
  import TableMap from '$lib/components/display/table/TableMap.svelte';
  import TablePlayer from '$lib/components/display/table/TablePlayer.svelte';
  import Section from '$lib/components/layout/Section.svelte';

  import type { PlayerWithPoints } from '$lib/schema.js';

  let { data } = $props();
</script>

<!-- todo: class swap -->
<Section label="ladder (soldier)">
  <!-- todo: sort by.. -->
  <span></span>
</Section>
{#await data.playersWithPoints}
  <span></span>
{:then pwp}
  {#if pwp}
    <Table data={pwp}>
      {#snippet header()}
        <th class="w-10"></th>
        <th></th>
        <th class="w-24 text-left"></th>
        <th class="w-24"></th>
        <th class="w-24">total points</th>
        <th class="w-24">last 3 monthly</th>
        <th class="w-24">last 9 motw</th>
      {/snippet}
      {#snippet row({ player, points }: PlayerWithPoints, i)}
        <td>#{i}</td>
        <td class="pl-10"><TablePlayer {player} /></td>
        <td class="text-left"><DivisionTag div={player.soldier_division ?? ''} /></td>
        <td>
          <div class="flex h-6 justify-center">
            <RocketLauncher launcher={player.preferred_launcher} />
          </div>
        </td>
        <td>{points.soldier.total}</td>
        <td>{points.soldier.last_3_monthly}</td>
        <td>{points.soldier.last_9_motw}</td>
      {/snippet}
    </Table>
  {/if}
{/await}
