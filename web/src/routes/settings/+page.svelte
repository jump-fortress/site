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
    createPlayerRequest
  } from '$lib/internalApi';
  import Input from '$lib/components/input/Input.svelte';
  import InputButtons from '$lib/components/input/InputButtons.svelte';
  import InputSelect from '$lib/components/input/InputSelect.svelte';
  import Button from '$lib/components/input/Button.svelte';
  import Table from '$lib/components/table/Table.svelte';
  import { Temporal } from 'temporal-polyfill';
  import type { Player, PlayerRequestPreview } from '$lib/schema';

  let { data }: { data: PageData } = $props();

  let requests: PlayerRequestPreview[] = $state([]);
  let player: Player | null = $state(null);

  player = (await data.player) ?? null;
  requests = (await data.requestPreviews) ?? [];
</script>

{#await data.requestPreviews then _}
  {#if requests.length}
    <DataSection title={'Pending Requests'}>
      <div class="w-fit">
        <Table data={requests}>
          {#snippet header()}
            <th class="w-48">request</th>
            <th class="w-32"></th>
            <th class="w-date">date</th>
          {/snippet}
          {#snippet row(r: PlayerRequestPreview)}
            <td>{r.request_type}</td>
            <td>{r.request_string}</td>
            <td>{Temporal.Instant.from(r.created_at).toLocaleString()}</td>
          {/snippet}
        </Table>
      </div>
    </DataSection>
  {/if}
{/await}

{#await data.player then _}
  {#if player}
    <DataSection title="Profile">
      <InputButtons
        title="fav class"
        options={[
          { src: rocket, value: 'Soldier' },
          { src: sticky, value: 'Demo' }
        ]}
        selectedOption={player.preferred_class}
        onSelect={async (value: string) => {
          // todo: feedback or message for successful update?
          const ok = updatePreferredClass(value);
          player.preferred_class = value;
        }} />

      <InputButtons
        title="fav launcher"
        options={[
          { src: stock, value: 'Stock' },
          { src: original, value: 'Original' },
          { src: mangler, value: 'Mangler' },
          { src: no_option, value: 'None' }
        ]}
        selectedOption={player.preferred_launcher}
        onSelect={async (value: string) => {
          const ok = updatePreferredLauncher(value);
          player.preferred_launcher = value;
        }} />

      <Input
        label={'request display name change'}
        placeholder={player.display_name}
        submitInput={async (name: string) => {
          const response = await createPlayerRequest('Display Name Change', name);
          if (!response.error) {
            requests.push({
              created_at: Temporal.Now.instant().toString(),
              request_string: name,
              request_type: 'Display Name Change'
            });
          }
          return response;
        }} />

      <Button
        onSelect={async () => {
          return updateSteamAvatar();
        }}>update avatar from steam</Button>
    </DataSection>

    <DataSection title={'Rank'}>
      {#if !player.soldier_division}
        <Button
          onSelect={async () => {
            // alternatively, api could return the submitted request.. but this is initially easier to write
            const response = await createPlayerRequest('Soldier Placement', 'null');
            if (!response.error) {
              requests.push({
                created_at: Temporal.Now.instant().toString(),
                request_string: '',
                request_type: 'Soldier Placement'
              });
            }
            return response;
          }}>request soldier placement</Button>
      {/if}
      {#if !player.demo_division}
        <Button
          onSelect={async () => {
            const response = await createPlayerRequest('Demo Placement', 'null');
            if (!response.error) {
              requests.push({
                created_at: Temporal.Now.instant().toString(),
                request_string: '',
                request_type: 'Demo Placement'
              });
            }
            return response;
          }}>request demo placement</Button>
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
        }} />

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
            return updateSteamTradeToken(val);
          }
        }} />

      <span>connect discord (not implemented)</span>
    </DataSection>
  {/if}
{/await}
