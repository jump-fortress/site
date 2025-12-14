<script lang="ts">
  import { page } from '$app/state';
  import { Temporal } from 'temporal-polyfill';
  import DataSection from '$lib/components/DataSection.svelte';
  import InnerNav from '$lib/components/InnerNav.svelte';
  import PlayerPreview from '$lib/components/PlayerPreview.svelte';
  import Table from '$lib/components/table/Table.svelte';
  import type { FullPlayer, PlayerProfile, Session } from '$lib/schema';
  import Header from '$lib/components/PlayerHeader.svelte';
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

  let route = $derived(page.url.pathname.substring(1));

  let { data, children } = $props();

  // there will always be a session, since we redirect otherwise
  let session: Session = $derived(data.session as Session);
  let fullPlayers: FullPlayer[] | null = $derived(data.fullPlayers);

  // todo: show unset profile
  let selected = $state(false);
  let selectedPlayer: string = $state('');

  let playerProfile: PlayerProfile | null = $state(null);
</script>

{#if session.role === 'Admin'}
  <DataSection title={'Manage Competitions'}>
    <InnerNav {route} parentRoute={'moderation'} pages={['monthly', 'motw', 'quest', 'bounty']} />
  </DataSection>
{/if}

{@render children?.()}

{#if session.role === 'Admin' || session.role === 'Mod'}
  <DataSection title={'Manage Players'}>
    <svelte:boundary {pending}>
      {#if selected && fullPlayers && playerProfile}
        <Header
          player={playerProfile}
          fullPlayer={fullPlayers.find((player) => {
            return player.id === selectedPlayer;
          })}
        />
        <DataSection title="Actions">
          <Input
            label={'update display name'}
            placeholder={playerProfile.display_name}
            submitInput={(name: string) => {
              return updatePlayerDisplayName(selectedPlayer, name);
            }}
          />
          <InputSelect
            label={'update soldier division'}
            options={divisions}
            placeholder={playerProfile.soldier_division}
            submitOption={async (division: string) => {
              return updatePlayerSoldierDivision(selectedPlayer, division);
            }}
          />
          <InputSelect
            label={'update demo division'}
            options={divisions}
            placeholder={playerProfile.demo_division}
            submitOption={async (division: string) => {
              return updatePlayerDemoDivision(selectedPlayer, division);
            }}
          />
        </DataSection>
      {/if}
      <Table data={fullPlayers}>
        {#snippet header()}
          <th class="w-48">steam id</th>
          <th class="w-10"></th>
          <th></th>

          <th class="w-24">soldier</th>
          <th class="w-24">demo</th>
          <th class="w-48">join date</th>
        {/snippet}
        {#snippet row(p: FullPlayer)}
          <td>{p.id}</td>
          <td><Flag code={p.country_code} /></td>
          <td
            onclick={async () => {
              selected = true;
              selectedPlayer = p.id;
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
    </svelte:boundary>
  </DataSection>
{/if}

todo: pending player skeleton.. todo: pending requests.. also separate these to 3 pages

{#snippet pending()}
  todo loading
{/snippet}
