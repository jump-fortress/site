<script lang="ts">
  import DataSection from '$lib/components/DataSection.svelte';
  import { Temporal } from 'temporal-polyfill';
  import type { PageData } from './$types';
  import Table from '$lib/components/table/Table.svelte';
  import type {
    Player,
    PlayerProfile,
    PlayerRequest,
    PlayerWithRequest,
    Session
  } from '$lib/schema';
  import PlayerPreview from '$lib/components/PlayerPreview.svelte';
  import Input from '$lib/components/input/Input.svelte';
  import { slide } from 'svelte/transition';
  import PlayerHeader from '$lib/components/PlayerHeader.svelte';
  import InputSelect from '$lib/components/input/InputSelect.svelte';
  import {
    getPlayerProfile,
    resolvePlayerRequest,
    updatePlayerDemoDivision,
    updatePlayerDisplayName,
    updatePlayerSoldierDivision
  } from '$lib/internalApi';
  import { divisions } from '$lib/divisions';
  import CheckButton from '$lib/components/input/CheckButton.svelte';
  import ManagePlayer from '$lib/components/ManagePlayer.svelte';

  let { data }: { data: PageData } = $props();

  let session = $derived(data.session as Session);

  let selectedPlayer: Player | null = $state(null);
</script>

{#if selectedPlayer}
  <div in:slide>
    <DataSection title={'Manage Player'}>
      <ManagePlayer player={selectedPlayer} role={session.role} />
    </DataSection>
  </div>
{/if}

<DataSection title={'Pending Requests'}>
  {#await data.playersWithRequests then requests}
    {#if requests}
      <Table data={requests}>
        {#snippet header()}
          <th class="w-48">request</th>
          <th class="w-32">content</th>
          <th></th>
          <th class="w-division">soldier</th>
          <th class="w-division">demo</th>
          <th class="w-date">date</th>
          <th class="w-0"></th>
        {/snippet}
        {#snippet row({ request, player }: PlayerWithRequest)}
          <td>{request.request_type}</td>
          <td>{request.request_string}</td>
          <td
            onclick={() => {
              selectedPlayer = player;
            }}
            class="transition-colors hover:cursor-pointer hover:bg-jfgray-800 hover:underline"
          >
            <PlayerPreview src={player.steam_avatar_url} name={player.display_name} />
          </td>
          <td class="text-division-{player.soldier_division?.toLowerCase()}"
            >{player.soldier_division}</td
          >
          <td class="text-division-{player.demo_division?.toLowerCase()}">{player.demo_division}</td
          >
          <td>{Temporal.Instant.from(request.created_at).toLocaleString()}</td>
          <td
            ><CheckButton
              onSelect={async () => {
                return resolvePlayerRequest(request.id);
              }}
            />
          </td>
        {/snippet}
      </Table>
    {:else}
      <span>no requests!</span>
    {/if}
  {/await}
</DataSection>
