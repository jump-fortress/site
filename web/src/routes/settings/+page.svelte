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
    updateTempusID
  } from '$lib/internalApi';
  import Input from '$lib/components/input/Input.svelte';
  import SelectButtons from '$lib/components/input/SelectButtons.svelte';
  import type { FullPlayer } from '$lib/schema';

  let { data }: { data: PageData } = $props();
  let player: FullPlayer | null = $derived(data.fullPlayer);
  $inspect(player);
  let favoriteClass = $derived(player?.preferred_class ?? '');
  let favoriteLauncher = $derived(player?.preferred_launcher ?? '');
</script>

<DataSection title="Profile">
  <SelectButtons
    title="fav class"
    options={[
      { src: rocket, value: 'Soldier' },
      { src: sticky, value: 'Demo' }
    ]}
    selectedOption={favoriteClass}
    onSelect={async (value: string) => {
      // todo: feedback or message for successful update
      const updated = updatePreferredClass(value);
      favoriteClass = value;
    }}
  />

  <SelectButtons
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

  <div class="flex flex-col gap-2">
    <Input
      label={'request display name change'}
      submitInput={async (val: string) => {
        if (val === '') {
          // todo: update with api request
          return {
            error: true,
            message: 'empty input'
          };
        } else {
          return {
            error: false,
            message: 'request sent! (is what this would say if it were implemented..)'
          };
        }
      }}
    />

    <button class="settings-button">update avatar from steam</button>
  </div>
</DataSection>

<DataSection title={'Rank'}>
  {#if !player?.soldier_division}
    <button class="settings-button">request soldier placement</button>
  {/if}
  {#if !player?.demo_division}
    <button class="settings-button">request demo placement</button>
  {/if}
</DataSection>

<DataSection title="Connections">
  <Input
    label="Tempus ID"
    initialMessage={player?.tempus_id ? `set to ${player.tempus_id}` : ''}
    submitInput={async (val: string) => {
      if (val === '') {
        return {
          error: true,
          message: 'empty input'
        };
      } else {
        return await updateTempusID(parseInt(val));
      }
    }}
  />

  <Input
    label="Steam Trade URL"
    initialMessage={player?.steam_trade_token ? `set to ${player.steam_trade_token}` : ''}
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
