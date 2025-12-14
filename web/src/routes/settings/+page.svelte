<script lang="ts">
  import DataSection from '$lib/components/DataSection.svelte';
  import rocket from '$lib/assets/static/classes/rocket.png';
  import sticky from '$lib/assets/static/classes/sticky.png';
  import stock from '$lib/assets/static/rocketlaunchers/stock.png';
  import original from '$lib/assets/static/rocketlaunchers/original.png';
  import mangler from '$lib/assets/static/rocketlaunchers/mangler.png';
  import no_option from '$lib/assets/static/no_option.png';
  import type { PageData } from './$types';
  import {
    updatePreferredClass,
    updatePreferredLauncher,
    updateSteamTradeToken,
    updateSteamAvatar,
    updateTempusID,
    createPlayerRequest,
    getPlayerRequests
  } from '$lib/internalApi';
  import Input from '$lib/components/input/Input.svelte';
  import InputButtons from '$lib/components/input/InputButtons.svelte';
  import type { FullPlayer, SelfPlayerRequest } from '$lib/schema';
  import InputSelect from '$lib/components/input/InputSelect.svelte';
  import Button from '$lib/components/input/Button.svelte';
  import Table from '$lib/components/table/Table.svelte';
  import { Temporal } from 'temporal-polyfill';

  let { data }: { data: PageData } = $props();
  let player: FullPlayer | null = $derived(data.fullPlayer);
  let favoriteClass = $derived(player?.preferred_class ?? '');
  let favoriteLauncher = $derived(player?.preferred_launcher ?? '');
</script>

<!-- todo: boundary for await -->
{#if player}
  {#await getPlayerRequests() then requests}
    {#if requests.length}
      <DataSection title={'Pending Requests'}>
        <Table data={requests}>
          {#snippet header()}
            <th class="w-64">request</th>
            <th></th>
            <th class="w-48">submitted</th>
          {/snippet}
          {#snippet row(r: SelfPlayerRequest)}
            <td>{r.request_type}</td>
            <td class="flex h-10 items-center"><span>{r.request_string}</span></td>
            <td>{Temporal.Instant.from(r.created_at).toLocaleString()}</td>
          {/snippet}
        </Table>
      </DataSection>
    {/if}
  {/await}

  <DataSection title="Profile">
    <InputButtons
      title="fav class"
      options={[
        { src: rocket, value: 'Soldier' },
        { src: sticky, value: 'Demo' }
      ]}
      selectedOption={favoriteClass}
      onSelect={async (value: string) => {
        // todo: feedback or message for successful update?
        const updated = updatePreferredClass(value);
        favoriteClass = value;
      }}
    />

    <InputButtons
      title="fav launcher"
      options={[
        { src: stock, value: 'Stock' },
        { src: original, value: 'Original' },
        { src: mangler, value: 'Mangler' },
        { src: no_option, value: 'None' }
      ]}
      selectedOption={favoriteLauncher}
      onSelect={async (value: string) => {
        const updated = updatePreferredLauncher(value);
        favoriteLauncher = value;
      }}
    />

    <Input
      label={'request display name change'}
      placeholder={player.display_name}
      submitInput={async (name: string) => {
        if (name === '') {
          // todo: update with api request
          return {
            error: true,
            message: 'empty input'
          };
        } else {
          return createPlayerRequest('Display Name Change', name);
        }
      }}
    />

    <Button
      label={'update avatar from steam'}
      onSelect={async () => {
        return updateSteamAvatar();
      }}
    />
  </DataSection>

  <DataSection title={'Rank'}>
    {#if !player.soldier_division}
      <Button
        label={'request soldier placement'}
        onSelect={async () => {
          return createPlayerRequest('Soldier Placement', 'null');
        }}
      />
    {/if}
    {#if !player.demo_division}
      <Button
        label={'request soldier placement'}
        onSelect={async () => {
          return createPlayerRequest('Demo Placement', 'null');
        }}
      />
    {/if}
  </DataSection>

  <DataSection title="Connections">
    <Input
      label="Tempus ID"
      placeholder={player.tempus_id ? player.tempus_id.toString() : ''}
      submitInput={async (val: string) => {
        if (val === '') {
          return {
            error: true,
            message: 'empty input'
          };
        } else {
          return updateTempusID(parseInt(val));
        }
      }}
    />

    <Input
      label="Steam Trade URL"
      placeholder={player.steam_trade_token}
      submitInput={async (val: string) => {
        if (val === '') {
          return {
            error: true,
            message: 'empty input'
          };
        } else {
          return await updateSteamTradeToken(val);
        }
      }}
    />

    <span>connect discord (not implemented)</span>
  </DataSection>
{:else}
  <!-- todo: error for missing player (redirect?) -->
  <span></span>
{/if}
