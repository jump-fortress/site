<script lang="ts">
  import demo from '$lib/assets/tf/demo.png';
  import rl_mangler from '$lib/assets/tf/rl_mangler.png';
  import rl_original from '$lib/assets/tf/rl_original.png';
  import rl_stock from '$lib/assets/tf/rl_stock.png';
  import soldier from '$lib/assets/tf/soldier.png';
  import Table from '$lib/components/display/table/Table.svelte';
  import Button from '$lib/components/input/Button.svelte';
  import Input from '$lib/components/input/Input.svelte';
  import Label from '$lib/components/input/Label.svelte';
  import Select from '$lib/components/input/select/Select.svelte';
  import SelectButton from '$lib/components/input/select/SelectButton.svelte';
  import Section from '$lib/components/layout/Section.svelte';
  import {
    createPlayerRequest,
    getAllMapNames,
    updatePreferredClass,
    updatePreferredLauncher,
    updatePreferredMap,
    updateSteamAvatar,
    updateSteamTradeToken,
    updateTempusID
  } from '$lib/src/api.js';
  import { Temporal } from 'temporal-polyfill';

  import type { Player, PlayerRequest } from '$lib/schema';

  let { data } = $props();

  let player: Player | null = $state(null);

  // svelte-ignore state_referenced_locally
  player = (await data.player) ?? null;
</script>

{#await data.player then _}
  {#if player}
    <Section label="profile">
      <SelectButton
        label="fav class"
        options={[
          { src: soldier, value: 'Soldier' },
          { src: demo, value: 'Demo' }
        ]}
        selected={player.preferred_class}
        onsubmit={(value) => {
          updatePreferredClass(value);
        }} />

      <SelectButton
        label="fav launcher"
        options={[
          { src: rl_stock, value: 'Stock' },
          { src: rl_original, value: 'Original' },
          { src: rl_mangler, value: 'Mangler' }
        ]}
        withNone={true}
        selected={player.preferred_launcher}
        onsubmit={(value) => {
          updatePreferredLauncher(value);
        }} />

      <Label label="fav map">
        {#await getAllMapNames() then mapNames}
          {#if mapNames}
            {@const maps = ['none'].concat(mapNames)}
            <Select
              type="text"
              placeholder={player.preferred_map}
              options={maps}
              onsubmit={(value) => {
                return updatePreferredMap(value);
              }} />
          {/if}
        {/await}
      </Label>

      <Button
        onsubmit={() => {
          return updateSteamAvatar();
        }}>
        <span>update avatar from Steam</span>
      </Button>

      <!-- todo: favorite map -->
    </Section>

    <Section label="requests">
      {#await data.requests then requests}
        {#if requests && requests.length}
          <div class="max-w-160 w-full">
            <Table data={requests}>
              {#snippet header()}
                <th class="w-32">request type</th>
                <th class="text-left"></th>
                <th class="w-48">date</th>
              {/snippet}
              {#snippet row(request: PlayerRequest)}
                {@const now = Temporal.Instant.from(request.created_at).toLocaleString()}
                <td>{request.request_type}</td>
                <td class="truncate">{request.request_string}</td>
                <td>{now}</td>
              {/snippet}
            </Table>
          </div>
        {/if}
      {/await}
      <Label label="update display name">
        <Input
          type="text"
          placeholder={player.display_name}
          onsubmit={(value) => {
            return createPlayerRequest('Display Name Change', value);
          }} />
      </Label>

      <Button
        onsubmit={() => {
          return createPlayerRequest('Soldier Placement', 'null');
        }}>
        <span>soldier division placement</span>
      </Button>

      <Button
        onsubmit={() => {
          return createPlayerRequest('Demo Placement', 'null');
        }}>
        <span>demo division placement</span>
      </Button>
    </Section>

    <Section label="connections">
      <Label label="Tempus ID">
        <Input
          type="text"
          placeholder={player.tempus_id?.toString() ?? ''}
          onsubmit={(value) => {
            return updateTempusID(parseInt(value));
          }} />
      </Label>

      <div class="flex flex-col max-w-80 w-full">
        <Label label="Steam Trade URL">
          <Input
            type="text"
            placeholder={player.steam_trade_token}
            onsubmit={(value) => {
              return updateSteamTradeToken(value);
            }} />
        </Label>
        <span class="text-content/50 text-sm">
          your Steam trade token is used to send trade offers for competition winnings.
        </span>
      </div>

      <!-- todo: connect discord -->
    </Section>
  {/if}
{/await}
