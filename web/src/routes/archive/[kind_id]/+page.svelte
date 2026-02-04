<script lang="ts">
  import { Client } from '$lib/api/api';
  import EventHeader from '$lib/components/display/EventHeader.svelte';
  import { ApiPaths, type Leaderboard, type Player, type TimeWithPlayer } from '$lib/schema';
  import { onMount } from 'svelte';
  import type { PageData } from './$types';
  import Table from '$lib/components/display/table/Table.svelte';
  import Div from '$lib/components/display/Div.svelte';
  import Section from '$lib/components/layout/Section.svelte';
  import TablePlayer from '$lib/components/display/table/TablePlayer.svelte';
  import { filterBestTimes, twTableGradients } from '$lib/helpers/times';
  import TableTime from '$lib/components/display/table/TableTime.svelte';
  import TemporalDate from '$lib/components/display/TemporalDate.svelte';
  import Content from '$lib/components/layout/Content.svelte';

  type Props = {
    data: PageData;
  };
  let { data }: Props = $props();

  let selectedLeaderboardID: number = $state(0);
  let playerLeaderboard: Leaderboard | undefined = $state();
  let prPlayer: Player | undefined = $state();

  onMount(async () => {
    if (data.ewl) {
      selectedLeaderboardID = data.ewl.leaderboards?.at(0)?.id ?? 0;

      // set matching player leaderboard ID
      if (data.session) {
        const { data: player } = await Client.GET(ApiPaths.get_player, {
          params: { path: { player_id: data.session.id } }
        });
        if (player) {
          prPlayer = player;
          const div =
            data.ewl.event.player_class === 'Soldier' ? player.soldier_div : player.demo_div;
          for (const l of data.ewl.leaderboards ?? []) {
            if (l.div === div || !l.div) {
              playerLeaderboard = l;
              break;
            }
          }
        }
      }
    }
  });
</script>

{#if data.ewl}
  <EventHeader event={data.ewl} />

  <Content>
    {#if data.session}
      <Section>
        {#await Client.GET( ApiPaths.get_leaderboard_pr, { params: { path: { event_id: data.ewl.event.id } } } )}
          <span></span>
        {:then { data: pr }}
          {#if prPlayer && pr}
            <Table data={[pr]}>
              {#snippet header()}
                <th class="w-rank"></th>
                <th class="w-32"></th>
                <th class=""></th>
                <th class="w-date"></th>
              {/snippet}
              {#snippet row({ _, leaderboard, time })}
                <td class={twTableGradients.get(`r${leaderboard.div?.toLowerCase()}`)}>PR</td>
                <td class={twTableGradients.get(`t${leaderboard.div?.toLowerCase()}`)}
                  ><TableTime {time} /></td>
                <td><TablePlayer player={prPlayer as Player} link={true} /></td>
                <td class="table-date"><TemporalDate datetime={time.created_at} /></td>
              {/snippet}
            </Table>
          {/if}
        {/await}
      </Section>
    {/if}
    <Section label={'leaderboards'}>
      <!-- select leaderboard buttons -->
      <div class="flex rounded-box border border-base-700">
        {#if data.ewl.leaderboards?.length}
          {#each data.ewl.leaderboards as leaderboard}
            <button
              class="items-center} flex h-9 grow cursor-pointer items-center justify-center rounded-box
          {selectedLeaderboardID === leaderboard.id
                ? 'border-b border-b-content/50 bg-base-900'
                : 'bg-base-800 opacity-50 hover:opacity-100'}"
              onclick={() => {
                selectedLeaderboardID = leaderboard.id;
              }}>
              <Div div={leaderboard.div} />
            </button>
          {/each}
        {/if}
      </div>

      {#await Client.GET( ApiPaths.get_leaderboard_times, { params: { path: { leaderboard_id: selectedLeaderboardID } } } )}
        <span>under construction..</span>
      {:then { data: times }}
        {@const bestTimes = filterBestTimes(times ?? [])}
        <Table data={bestTimes}>
          {#snippet header()}
            <th class="w-rank"></th>
            <th class="w-32"></th>
            <th class=""></th>
            <th class="w-date"></th>
          {/snippet}
          {#snippet row({ player, time }: TimeWithPlayer, i)}
            <td class={twTableGradients.get(`r${i}`)}>{i}</td>
            <td class={twTableGradients.get(`t${i}`)}><TableTime {time} /></td>
            <td><TablePlayer {player} link={true} /></td>
            <td class="table-date"><TemporalDate datetime={time.created_at} /></td>
          {/snippet}
        </Table>
      {/await}
    </Section>
  </Content>
{/if}
