<script lang="ts">
  import jf from '$lib/assets/logo/jf.png';
  import DivisionTag from '$lib/components/display/DivisionTag.svelte';
  import PlayerHeader from '$lib/components/display/player/PlayerHeader.svelte';
  import Table from '$lib/components/display/table/Table.svelte';
  import TablePlayer from '$lib/components/display/table/TablePlayer.svelte';
  import Button from '$lib/components/input/Button.svelte';
  import Input from '$lib/components/input/Input.svelte';
  import Label from '$lib/components/input/Label.svelte';
  import Select from '$lib/components/input/select/Select.svelte';
  import Section from '$lib/components/layout/Section.svelte';
  import {
    updatePlayerDisplayName,
    updatePlayerSoldierDivision,
    updatePlayerDemoDivision,
    resolvePlayerRequest
  } from '$lib/src/api';
  import { divisions } from '$lib/src/divisions';
  import { slide } from 'svelte/transition';
  import { Temporal } from 'temporal-polyfill';

  import type { Player, PlayerWithRequest } from '$lib/schema.js';

  let { data } = $props();

  let requests: PlayerWithRequest[] = $state([]);

  // svelte-ignore state_referenced_locally
  requests = (await data.playersWithRequests) ?? [];

  const placeholder = {
    created_at: '',
    display_name: 'select a request',
    id: '',
    preferred_class: 'Soldier',
    preferred_launcher: 'None',
    role: 'player',
    steam_avatar_url: jf
  };

  let selected: Player = $state(placeholder);
</script>

<PlayerHeader player={selected} selected_class={selected.preferred_class} />

{#if data.session.role === 'Moderator' || (data.session.role === 'Admin' && selected.display_name !== 'select a request')}
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

{#await data.playersWithRequests then _}
  {#if requests.length}
    <Table data={requests}>
      {#snippet header()}
        <th class="w-36">request type</th>
        <th class="text-left"></th>
        <th></th>
        <th class="w-24">soldier</th>
        <th class="w-24">demo</th>
        <th class="w-48">date</th>
        <th class="w-10"></th>
      {/snippet}
      {#snippet row({ player, request }: PlayerWithRequest)}
        <td>{request.request_type}</td>
        <td>{request.request_string}</td>
        <td
          onclick={() => {
            selected = player;
          }}
          ><TablePlayer {player} />
        </td>
        <td><DivisionTag div={player.soldier_division} /></td>
        <td><DivisionTag div={player.demo_division} /></td>
        <td>{Temporal.Instant.from(request.created_at).toLocaleString()}</td>
        <td
          ><Button
            table={true}
            onsubmit={async () => {
              const response = await resolvePlayerRequest(request.id);
              if (response && !response.error) {
                requests = requests.filter(({ request: r }) => {
                  r.id !== request.id;
                });
              }
              return response;
            }}><span class="icon-[mdi--close] w-6"></span></Button
          ></td>
      {/snippet}
    </Table>
  {/if}
{/await}
