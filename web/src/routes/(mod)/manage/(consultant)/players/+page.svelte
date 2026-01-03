<script lang="ts">
  import jf from '$lib/assets/logo/jf.png';
  import DivisionTag from '$lib/components/display/DivisionTag.svelte';
  import PlayerHeader from '$lib/components/display/player/PlayerHeader.svelte';
  import Table from '$lib/components/display/table/Table.svelte';
  import TablePlayer from '$lib/components/display/table/TablePlayer.svelte';
  import Input from '$lib/components/input/Input.svelte';
  import Label from '$lib/components/input/Label.svelte';
  import Select from '$lib/components/input/select/Select.svelte';
  import Section from '$lib/components/layout/Section.svelte';
  import {
    updatePlayerDemoDivision,
    updatePlayerDisplayName,
    updatePlayerSoldierDivision
  } from '$lib/src/api.js';
  import { divisions } from '$lib/src/divisions.js';
  import { slide } from 'svelte/transition';
  import { Temporal } from 'temporal-polyfill';

  import type { Player } from '$lib/schema.js';

  let { data } = $props();

  let players: Player[] | [] = $state([]);

  // svelte-ignore state_referenced_locally
  players = (await data.players) ?? [];

  const placeholder = {
    created_at: '',
    display_name: 'select a player',
    id: '',
    preferred_class: 'Soldier',
    preferred_launcher: 'None',
    role: 'player',
    steam_avatar_url: jf
  };

  let selected: Player = $state(placeholder);
</script>

<PlayerHeader player={selected} selected_class={selected.preferred_class} />

{#if data.session.role === 'Moderator' || (data.session.role === 'Admin' && selected.display_name !== 'select a player')}
  <div in:slide>
    <Section label="manage player">
      <Label label="update display name">
        <Input
          type="text"
          placeholder={selected.display_name}
          onsubmit={(value) => {
            return updatePlayerDisplayName(selected.id, value);
          }} />
      </Label>

      <Label label="update soldier div">
        <Select
          type="text"
          placeholder={selected.soldier_division}
          options={divisions.concat(['None'])}
          onsubmit={(value) => {
            return updatePlayerSoldierDivision(selected.id, value);
          }} />
      </Label>

      <Label label="update demo div">
        <Select
          type="text"
          placeholder={selected.demo_division}
          options={divisions.concat(['None'])}
          onsubmit={(value) => {
            return updatePlayerDemoDivision(selected.id, value);
          }} />
      </Label>
    </Section>
  </div>
{/if}

{#await data.players then _}
  {#if players.length}
    <Table data={players}>
      {#snippet header()}
        <th class="w-36">steam id</th>
        <th></th>
        <th class="w-24">soldier</th>
        <th class="w-24">demo</th>
        <th class="w-48">join date</th>
      {/snippet}
      {#snippet row(player: Player)}
        <td>{player.id}</td>
        <td
          onclick={() => {
            selected = player;
          }}
          ><TablePlayer {player} />
        </td>
        <td><DivisionTag div={player.soldier_division} /></td>
        <td><DivisionTag div={player.demo_division} /></td>
        <td>{Temporal.Instant.from(player.created_at).toLocaleString()}</td>
      {/snippet}
    </Table>
  {/if}
{/await}
