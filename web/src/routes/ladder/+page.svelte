<script lang="ts">
  import ClassImage from '$lib/components/display/ClassImage.svelte';
  import ClassSelect from '$lib/components/display/ClassSelect.svelte';
  import Div from '$lib/components/display/Div.svelte';
  import Launcher from '$lib/components/display/Launcher.svelte';
  import Table from '$lib/components/display/table/Table.svelte';
  import TablePlayer from '$lib/components/display/table/TablePlayer.svelte';
  import TemporalDate from '$lib/components/display/TemporalDate.svelte';
  import Section from '$lib/components/layout/Section.svelte';
  import { compareBothDivisions, compareDivisions } from '$lib/helpers/divs';
  import type { Player } from '$lib/schema';
  import type { PageData } from './$types';

  type Props = {
    data: PageData;
  };
  let { data }: Props = $props();
  let sort_class = $state('Soldier');

  // todo: abstractly split class sorting tables / styles
</script>

{#if data.players}
  <Section label="players">
    {#key sort_class}
      <Table
        data={data.players.sort((a, b) =>
          sort_class === 'Soldier'
            ? compareDivisions(a.soldier_div ?? '', b.soldier_div ?? '')
            : compareDivisions(a.demo_div ?? '', b.demo_div ?? '')
        )}>
        {#snippet header()}
          <th class="w-12 text-start">pref</th>
          <th></th>
          {#if sort_class === 'Soldier'}
            <th class="w-div"></th>
          {/if}
          <th
            class="w-div cursor-pointer text-start hover:text-primary {sort_class === 'Soldier'
              ? 'text-primary'
              : ''}"
            onclick={() => {
              sort_class = 'Soldier';
            }}>soldier</th>
          <th
            class="w-div cursor-pointer text-start hover:text-primary {sort_class === 'Demo'
              ? 'text-primary'
              : ''}"
            onclick={() => {
              sort_class = 'Demo';
            }}>demo</th>
          <th class="w-date"></th>
        {/snippet}
        {#snippet row(player: Player, i)}
          <td class="h-8"><ClassImage player_class={player.class_pref} /></td>
          <td><TablePlayer {player} /></td>
          {#if sort_class === 'Soldier'}
            <td class="h-6"><Launcher launcher={player.launcher_pref ?? ''} /></td>
          {/if}
          <td class="text-start"><Div div={player.soldier_div} /></td>
          <td class="text-start"><Div div={player.demo_div} /></td>
          <td class="text-content/75"><TemporalDate datetime={player.created_at} /></td>
        {/snippet}
      </Table>
    {/key}
  </Section>
{/if}
