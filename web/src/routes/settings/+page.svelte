<script lang="ts">
  import DataSection from '$lib/components/DataSection.svelte';
  import rocket from '$lib/assets/static/classes/rocket.png';
  import sticky from '$lib/assets/static/classes/sticky.png';
  import stock from '$lib/assets/static/rocketlaunchers/stock.png';
  import original from '$lib/assets/static/rocketlaunchers/original.png';
  import mangler from '$lib/assets/static/rocketlaunchers/mangler.png';
  import type { PageData } from './$types';
  import { updatePreferredClass } from '$lib/internalApi';
  import { classToEnum } from '$lib/enums';
  import Input from '$lib/components/input/Input.svelte';
  import SelectButtons from '$lib/components/input/SelectButtons.svelte';
  import type { PlayerProfile } from '$lib/schema';

  let { data }: { data: PageData } = $props();
  let player: PlayerProfile | null = $derived(data.player);
  let favoriteClass = $derived(player?.preferred_class ?? '');
  // todo: update with preferred_launcher
  let favoriteLauncher = $state('stock');
</script>

<DataSection title="Profile">
  <SelectButtons
    title="fav class"
    options={[
      { src: rocket, value: 'Soldier' },
      { src: sticky, value: 'Demo' }
    ]}
    selectedOption={favoriteClass}
    onSelect={(value: string) => {
      updatePreferredClass(classToEnum(value as 'Soldier' | 'Demo'));
      favoriteClass = value;
    }}
  />

  <SelectButtons
    title="fav launcher"
    options={[
      { src: stock, value: 'Stock' },
      { src: original, value: 'Original' },
      { src: mangler, value: 'Mangler' }
    ]}
    selectedOption={favoriteLauncher}
    onSelect={(value: string) => {
      // todo: update launcher
      favoriteLauncher = value;
    }}
  />

  <div class="flex flex-col gap-2">
    <Input
      label={'request display name change'}
      submitInput={(val: string) => {
        if (val === '') {
          return {
            error: true,
            message: 'empty input'
          };
        } else {
          return {
            error: false,
            message: 'request sent!'
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
    submitInput={(val: string) => {
      if (val === '') {
        return {
          error: true,
          message: 'empty input'
        };
      } else {
        return {
          error: false,
          message: 'request sent!'
        };
      }
    }}
  />

  <Input
    label="Steam Trade URL"
    submitInput={(val: string) => {
      if (val === '') {
        return {
          error: true,
          message: 'empty input'
        };
      } else {
        return {
          error: false,
          message: 'request sent!'
        };
      }
    }}
  />

  <span>connect discord (not implemented)</span>
</DataSection>
