<script lang="ts">
  import demo from '$lib/assets/tf/demo.png';
  import rl_mangler from '$lib/assets/tf/rl_mangler.png';
  import rl_original from '$lib/assets/tf/rl_original.png';
  import rl_stock from '$lib/assets/tf/rl_stock.png';
  import soldier from '$lib/assets/tf/soldier.png';
  import ClassImage from '$lib/components/display/ClassImage.svelte';
  import Table from '$lib/components/display/table/Table.svelte';
  import TableDate from '$lib/components/display/table/TableDate.svelte';
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

  import type { Player, PlayerRequest } from '$lib/schema';

  let { data } = $props();

  let player: Player | undefined = $derived(undefined);

  // svelte-ignore state_referenced_locally
  player = await data.player;
</script>

{#await data.player}
  <span></span>
{:then _}
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
        <span class="icon-[mdi--image-outline]"></span>
        <span>update Steam avatar</span>
      </Button>

      <!-- todo: favorite map -->
    </Section>

    <Section label="requests">
      {#await data.requests then requests}
        {#if requests && requests.length}
          <div class="w-full max-w-160">
            <Table data={requests}>
              {#snippet header()}
                <th class="w-40">request type</th>
                <th class="text-left"></th>
                <th class="w-48">date</th>
              {/snippet}
              {#snippet row(request: PlayerRequest)}
                <td>{request.request_type}</td>
                <td class="truncate">{request.request_string}</td>
                <td><TableDate date={request.created_at} /></td>
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
        <span class="h-6"><ClassImage selected={'Soldier'} /></span>
        <span>soldier placement</span>
      </Button>

      <Button
        onsubmit={() => {
          return createPlayerRequest('Demo Placement', 'null');
        }}>
        <span class="h-6"><ClassImage selected={'Demo'} /></span>
        <span>demo placement</span>
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

      <div class="flex w-full max-w-80 flex-col">
        <Label label="Steam Trade URL">
          <Input
            type="text"
            placeholder={player.steam_trade_token}
            onsubmit={(value) => {
              return updateSteamTradeToken(value);
            }} />
        </Label>
        <span class="text-sm text-content/50">
          your Steam trade token is used to send trade offers for competition winnings.
        </span>
      </div>

      <!-- todo: connect discord -->
    </Section>
  {/if}
{/await}
