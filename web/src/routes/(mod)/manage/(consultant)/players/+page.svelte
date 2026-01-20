<script lang="ts">
  import jf from '$lib/assets/logo/jf.png';
  import DivisionTag from '$lib/components/display/DivisionTag.svelte';
  import PlayerHeader from '$lib/components/display/player/PlayerHeader.svelte';
  import Table from '$lib/components/display/table/Table.svelte';
  import TableDate from '$lib/components/display/table/TableDate.svelte';
  import TablePlayer from '$lib/components/display/table/TablePlayer.svelte';
  import Input from '$lib/components/input/Input.svelte';
  import Label from '$lib/components/input/Label.svelte';
  import Select from '$lib/components/input/select/Select.svelte';
  import Section from '$lib/components/layout/Section.svelte';
  import {
    updatePlayerDemoDivision,
    updatePlayerDisplayName,
    updatePlayerRole,
    updatePlayerSoldierDivision
  } from '$lib/src/api.js';
  import { compareBothDivisions, compareDivisions, divisions } from '$lib/src/divisions.js';
  import { slide } from 'svelte/transition';

  import type { Player } from '$lib/schema.js';

  let { data } = $props();

  let players: Player[] | [] = $state([]);

  // svelte-ignore state_referenced_locally
  players = ((await data.players) ?? []).sort((a, b) =>
    compareBothDivisions(
      a.soldier_division ?? '',
      a.demo_division ?? '',
      b.soldier_division ?? '',
      b.demo_division ?? ''
    )
  );

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

{#await data.session then session}
  {#if session}
    {#if (session.role === 'Moderator' || session.role === 'Admin') && selected.display_name !== 'select a player'}
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

          {#if session.role === 'Admin'}
            <Label label="update role">
              <Select
                type="text"
                placeholder={selected.role}
                options={['Admin', 'Moderator', 'Consultant', 'Treasurer', 'Player']}
                onsubmit={(value) => {
                  return updatePlayerRole(selected.id, value);
                }} />
            </Label>
          {/if}
        </Section>
      </div>
    {/if}
  {/if}
{/await}

{#await data.players}
  <span></span>
{:then _}
  {#if players.length}
    <Table data={players}>
      {#snippet header()}
        <th class="w-div">role</th>
        <th></th>
        <th class="w-div text-left">soldier</th>
        <th class="w-div text-left">demo</th>
        <th class="w-date">join date</th>
      {/snippet}
      {#snippet row(player: Player)}
        <td>{player.role === 'Player' ? '' : player.role}</td>
        <td
          onclick={() => {
            selected = player;
          }}
          ><TablePlayer {player} link={false} />
        </td>
        <td class="text-left"><DivisionTag div={player.soldier_division} /></td>
        <td class="text-left"><DivisionTag div={player.demo_division} /></td>
        <td class="table-date"><TableDate date={player.created_at} /></td>
      {/snippet}
    </Table>
  {/if}
{/await}
