<script lang="ts">
  import { Temporal } from 'temporal-polyfill';
  import DataSection from '$lib/components/DataSection.svelte';
  import PlayerPreview from '$lib/components/PlayerPreview.svelte';
  import Table from '$lib/components/table/Table.svelte';
  import type { Player, Session } from '$lib/schema';
  import Flag from '$lib/components/Flag.svelte';
  import { slide } from 'svelte/transition';
  import ManagePlayer from '$lib/components/ManagePlayer.svelte';

  let { data } = $props();

  // a session is guaranteed here since a redirect happens otherwise
  let session: Session = $derived(data.session as Session);

  // todo: show unset profile
  let selectedPlayer: Player | null = $state(null);
</script>

{#if selectedPlayer}
  <div in:slide>
    <DataSection title={'Manage Player'}>
      <ManagePlayer player={selectedPlayer} role={session.role} />
    </DataSection>
  </div>
{/if}

{#await data.players then players}
  {#if players}
    <DataSection title={'Player List'}>
      <Table data={players}>
        {#snippet header()}
          <th class="w-id">steam id</th>
          <th class="w-flag"></th>
          <th></th>

          <th class="w-division">soldier</th>
          <th class="w-division">demo</th>
          <th class="w-date">join date</th>
        {/snippet}
        {#snippet row(player: Player)}
          <td>{player.id}</td>
          <td><Flag code={player.country_code} /></td>
          <td
            onclick={() => {
              selectedPlayer = player;
            }}
            class="td-hover-preview">
            <PlayerPreview {player} />
          </td>
          <td class="text-division-{player.soldier_division?.toLowerCase()}"
            >{player.soldier_division}</td>
          <td class="text-division-{player.demo_division?.toLowerCase()}"
            >{player.demo_division}</td>
          <th>{Temporal.Instant.from(player.created_at).toLocaleString()}</th>
        {/snippet}
      </Table>
    </DataSection>
  {/if}
{/await}
