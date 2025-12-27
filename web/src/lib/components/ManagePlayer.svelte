<script lang="ts">
  import { divisions } from '$lib/divisions';

  import {
    updatePlayerDisplayName,
    updatePlayerSoldierDivision,
    updatePlayerDemoDivision
  } from '$lib/internalApi';

  import type { Player } from '$lib/schema';
  import DataSection from '$lib/components/DataSection.svelte';
  import Input from '$lib/components/input/Input.svelte';
  import InputSelect from '$lib/components/input/InputSelect.svelte';
  import PlayerHeader from '$lib/components/PlayerHeader.svelte';

  type Props = {
    player: Player;
    role: string;
  };

  let { player, role }: Props = $props();
</script>

<PlayerHeader {player} modView={true} />
{#if role === 'Mod' || role === 'Admin'}
  <DataSection title="Actions">
    <Input
      label={'update display name'}
      placeholder={player.display_name}
      submitInput={async (name: string) => {
        return updatePlayerDisplayName(player.id, name);
      }} />
    <InputSelect
      label={'update soldier division'}
      options={divisions.concat('None')}
      placeholder={player.soldier_division}
      submitOption={async (division: string) => {
        return updatePlayerSoldierDivision(player.id, division);
      }} />
    <InputSelect
      label={'update demo division'}
      options={divisions.concat('None')}
      placeholder={player.demo_division}
      submitOption={async (division: string) => {
        return updatePlayerDemoDivision(player.id, division);
      }} />
  </DataSection>
{/if}
