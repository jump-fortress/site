<script lang="ts">
  import { Client } from '$lib/api/api';
  import Div from '$lib/components/display/Div.svelte';
  import PlayerHeader from '$lib/components/display/player/PlayerHeader.svelte';
  import Table from '$lib/components/display/table/Table.svelte';
  import TablePlayer from '$lib/components/display/table/TablePlayer.svelte';
  import TemporalDate from '$lib/components/display/TemporalDate.svelte';
  import Section from '$lib/components/layout/Section.svelte';
  import { ApiPaths, type Player } from '$lib/schema';
  import jf from '$lib/assets/logo/jf.png';
  import Input from '$lib/components/input/Input.svelte';
  import Errors from '$lib/components/input/Errors.svelte';
  import Select from '$lib/components/input/Select.svelte';
  import { divs } from '$lib/helpers/divs';
  import type { PageData } from './$types';

  type Props = {
    data: PageData;
  };

  let player: Player = $state({
    id: '0',
    role: 'player',
    alias: 'select a player',
    avatar_url: jf,
    class_pref: 'Soldier',
    created_at: '2026-01-01T00:00:00Z'
  });

  let oerror: OpenAPIError = $state(undefined);

  let { data }: Props = $props();
</script>

<PlayerHeader {player} />
{#if player.id !== '0'}
  <Section label="update player">
    <Errors {oerror} />
    <Input
      label="alias"
      type="text"
      placeholder={player.alias}
      onsubmit={async (value) => {
        const resp = await Client.POST(ApiPaths.update_alias, {
          params: { path: { player_id: player.id, alias: value } }
        });
        oerror = resp.error;
        if (resp.response.ok) {
          player.alias = value;
        }
        return resp.response.ok;
      }} />
    <Select
      label="soldier div"
      type="text"
      placeholder={player.soldier_div}
      options={divs.concat('none')}
      onsubmit={async (value) => {
        const resp = await Client.POST(ApiPaths.update_div, {
          params: { path: { player_id: player.id, player_class: 'Soldier', div: value } }
        });
        oerror = resp.error;
        if (resp.response.ok) {
          player.soldier_div = value;
        }
        return resp.response.ok;
      }} />
    <Select
      label="demo div"
      type="text"
      placeholder={player.demo_div}
      options={divs.concat('none')}
      onsubmit={async (value) => {
        const resp = await Client.POST(ApiPaths.update_div, {
          params: { path: { player_id: player.id, player_class: 'Demo', div: value } }
        });
        oerror = resp.error;
        if (resp.response.ok) {
          player.demo_div = value;
        }
        return resp.response.ok;
      }} />
    {#if data.session?.role === 'admin' || data.session?.role === 'dev'}
      <span>under construction</span>
      <Select
        label="role"
        type="text"
        placeholder={player.role}
        options={['player', 'mod', 'admin']}
        onsubmit={async (value) => {
          const resp = await Client.GET(ApiPaths.readyz);
          oerror = resp.error;
          if (resp.response.ok) {
            //player.alias = value;
          }
          return resp.response.ok;
        }} />
    {/if}
  </Section>
{/if}

<Section label="player list">
  {#await Client.GET(ApiPaths.get_players)}
    <span></span>
  {:then { data: players }}
    {#if players}
      <Table data={players}>
        {#snippet header()}
          <th class="w-div">role</th>
          <th></th>
          <th class="w-div">soldier</th>
          <th class="w-div">demo</th>
          <th class="w-date">joined</th>
        {/snippet}
        {#snippet row(p: Player)}
          <td>{p.role === 'player' ? '' : p.role}</td>
          <td
            onclick={() => {
              player = p;
              // jump to top
              window.scrollTo({ top: 0, behavior: 'smooth' });
            }}><TablePlayer player={p} link={false} /></td>
          <td><Div div={p.soldier_div} /></td>
          <td><Div div={p.demo_div} /></td>
          <td class="table-date"><TemporalDate datetime={p.created_at} /></td>
        {/snippet}
      </Table>
    {/if}
  {/await}
</Section>
