<script lang="ts">
  import { Temporal } from 'temporal-polyfill';
  import DataSection from '$lib/components/DataSection.svelte';
  import PlayerPreview from '$lib/components/PlayerPreview.svelte';
  import Table from '$lib/components/table/Table.svelte';
  import type { FullPlayer, PlayerProfile, Session } from '$lib/schema';
  import PlayerHeader from '$lib/components/PlayerHeader.svelte';
  import {
    getPlayerProfile,
    updatePlayerDemoDivision,
    updatePlayerDisplayName,
    updatePlayerSoldierDivision
  } from '$lib/internalApi';
  import Flag from '$lib/components/Flag.svelte';
  import Input from '$lib/components/input/Input.svelte';
  import { divisions } from '$lib/divisions';
  import InputSelect from '$lib/components/input/InputSelect.svelte';
  import { slide } from 'svelte/transition';

  let { data } = $props();

  // a session is guaranteed here since a redirect happens otherwise
  let session: Session = $derived(data.session as Session);
  let fullPlayers: FullPlayer[] | null = $derived(data.fullPlayers);

  // todo: show unset profile
  let selected = $state(false);
  let selectedPlayerID: string = $state('');

  let playerProfile: PlayerProfile | null = $state(null);
</script>

<svelte:boundary {pending}>
  {#if fullPlayers}
    {#if selected && playerProfile}
      <div in:slide>
        <DataSection title={'Manage Player'}>
          <PlayerHeader
            player={playerProfile}
            fullPlayer={fullPlayers.find((player) => {
              return player.id === selectedPlayerID;
            })}
          />
          {#if session.role === 'Mod' || session.role === 'Admin'}
            <DataSection title="Actions">
              <Input
                label={'update display name'}
                placeholder={playerProfile.display_name}
                submitInput={(name: string) => {
                  return updatePlayerDisplayName(selectedPlayerID, name);
                }}
              />
              <InputSelect
                label={'update soldier division'}
                options={divisions}
                placeholder={playerProfile.soldier_division}
                submitOption={async (division: string) => {
                  return updatePlayerSoldierDivision(selectedPlayerID, division);
                }}
              />
              <InputSelect
                label={'update demo division'}
                options={divisions}
                placeholder={playerProfile.demo_division}
                submitOption={async (division: string) => {
                  return updatePlayerDemoDivision(selectedPlayerID, division);
                }}
              />
            </DataSection>
          {/if}
        </DataSection>
      </div>
    {/if}
    <DataSection title={'Player List'}>
      <Table data={fullPlayers}>
        {#snippet header()}
          <th class="w-id">steam id</th>
          <th class="w-flag"></th>
          <th></th>

          <th class="w-division">soldier</th>
          <th class="w-division">demo</th>
          <th class="w-date">join date</th>
        {/snippet}
        {#snippet row(p: FullPlayer)}
          <td>{p.id}</td>
          <td><Flag code={p.country_code} /></td>
          <td
            onclick={async () => {
              selected = true;
              selectedPlayerID = p.id;
              playerProfile = await getPlayerProfile(p.id);
            }}
            class="transition-colors hover:cursor-pointer hover:bg-jfgray-800 hover:underline"
          >
            <PlayerPreview src={p.steam_avatar_url} name={p.display_name} />
          </td>
          <td class="text-division-{p.soldier_division?.toLowerCase()}">{p.soldier_division}</td>
          <td class="text-division-{p.demo_division?.toLowerCase()}">{p.demo_division}</td>
          <th>{Temporal.Instant.from(p.created_at).toLocaleString()}</th>
        {/snippet}
      </Table>
    </DataSection>
  {/if}
</svelte:boundary>

{#snippet pending()}
  todo loading
{/snippet}
