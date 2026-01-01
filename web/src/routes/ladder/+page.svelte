<script lang="ts">
  import DivisionTag from '$lib/components/display/DivisionTag.svelte';
  import RocketLauncher from '$lib/components/display/RocketLauncher.svelte';
  import Table from '$lib/components/display/table/Table.svelte';
  import TablePlayer from '$lib/components/display/table/TablePlayer.svelte';
  import Section from '$lib/components/layout/Section.svelte';

  import type { PlayerPreview } from '$lib/schema.js';

  let { data } = $props();
</script>

<!-- todo: class swap -->
<Section label="ladder (soldier)">
  <!-- todo: sort by.. -->
  <span></span>
</Section>
{#await data.playerPreviews then players}
  {#if players}
    <Table data={players}>
      {#snippet header()}
        <th class="w-16">rank</th>
        <th class="w-16">div</th>
        <th class="w-24"></th>
        <th></th>
        <th class="w-22"></th>
        <th class="w-16"># monthly</th>
        <th class="w-16"># motw</th>
      {/snippet}
      {#snippet row(player: PlayerPreview)}
        <td>1</td>
        <td>1</td>
        <td><DivisionTag div={player.soldier_division ?? ''} /></td>
        <td><TablePlayer {player} /></td>
        <td>
          {#if player.preferred_launcher !== 'None'}
            <div class="h-6 flex justify-center">
              <RocketLauncher launcher={player.preferred_launcher} />
            </div>
          {/if}
        </td>
        <td>3</td>
        <td>9</td>
      {/snippet}
    </Table>
  {/if}
{/await}
