<script lang="ts">
  import DataSection from '$lib/components/DataSection.svelte';
  import { Temporal } from 'temporal-polyfill';
  import type { PageData } from './$types';
  import Table from '$lib/components/table/Table.svelte';
  import type { FullPlayer, PlayerProfile, PlayerRequest, Session } from '$lib/schema';
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
  import Button from '$lib/components/input/Button.svelte';

  let { data }: { data: PageData } = $props();

  let session = $derived(data.session as Session);
  let requests = $derived(data.request);

  let selected = $state(false);
  let selectedPlayerID: string = $state('');
  let selectedPlayer: FullPlayer | null = $state(null);
  let selectedPlayerProfile: PlayerProfile | null = $state(null);
</script>

{#if selected && selectedPlayer && selectedPlayerProfile}
  <div in:slide>
    <DataSection title={'Manage Player'}>
      <PlayerHeader player={selectedPlayerProfile} fullPlayer={selectedPlayer} />
      {#if session.role === 'Mod' || session.role === 'Admin'}
        <DataSection title="Actions">
          <Input
            label={'update display name'}
            placeholder={selectedPlayer.display_name}
            submitInput={(name: string) => {
              return updatePlayerDisplayName(selectedPlayerID, name);
            }}
          />
          <InputSelect
            label={'update soldier division'}
            options={divisions}
            placeholder={selectedPlayer.soldier_division}
            submitOption={async (division: string) => {
              return updatePlayerSoldierDivision(selectedPlayerID, division);
            }}
          />
          <InputSelect
            label={'update demo division'}
            options={divisions}
            placeholder={selectedPlayer.demo_division}
            submitOption={async (division: string) => {
              return updatePlayerDemoDivision(selectedPlayerID, division);
            }}
          />
        </DataSection>
      {/if}
    </DataSection>
  </div>
{/if}

<DataSection title={'Pending Requests'}>
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
      {#snippet row({ request, player }: PlayerRequest)}
        <td>{request.request_type}</td>
        <td>{request.request_string}</td>
        <td
          onclick={async () => {
            selected = true;
            selectedPlayerID = player.id;
            selectedPlayer = player;
            selectedPlayerProfile = await getPlayerProfile(player.id);
          }}
          class="transition-colors hover:cursor-pointer hover:bg-jfgray-800 hover:underline"
        >
          <PlayerPreview src={player.steam_avatar_url} name={player.display_name} />
        </td>
        <td class="text-division-{player.soldier_division?.toLowerCase()}"
          >{player.soldier_division}</td
        >
        <td class="text-division-{player.demo_division?.toLowerCase()}">{player.demo_division}</td>
        <td>{Temporal.Instant.from(request.created_at).toLocaleString()}</td>
        <td
          ><Button
            onSelect={async () => {
              resolvePlayerRequest(request.id);

              // remove this issue from the list when clicked
              requests = requests.filter(({ request: entry }) => {
                entry.id !== request.id;
              });

              // always return empty error to not disturb table layout
              return { error: false, message: '' };
            }}
            ><span class="icon-[ri--check-line]"></span>
          </Button>
        </td>
      {/snippet}
    </Table>
  {:else}
    <span>no requests!</span>
  {/if}
</DataSection>
