<script lang="ts">
  import { Client } from '$lib/api/api.js';
  import PlayerHeader from '$lib/components/display/player/PlayerHeader.svelte';
  import Table from '$lib/components/display/table/Table.svelte';
  import TableEvent from '$lib/components/display/table/TableEvent.svelte';
  import TableMap from '$lib/components/display/table/TableMap.svelte';
  import TableTime from '$lib/components/display/table/TableTime.svelte';
  import TemporalDate from '$lib/components/display/TemporalDate.svelte';
  import Content from '$lib/components/layout/Content.svelte';
  import Section from '$lib/components/layout/Section.svelte';
  import { formatPosition, twTableGradients } from '$lib/helpers/times.js';
  import { ApiPaths, type EventLeaderboardTime } from '$lib/schema.js';

  let { data } = $props();

  let selected_class = $derived(data.player?.class_pref ?? 'Soldier');
</script>

{#if data.player}
  <PlayerHeader player={data.player} bind:class_pref={selected_class} />

  <Content>
    {#await Client.GET( ApiPaths.get_player_prs, { params: { path: { player_id: data.player.id } } } )}
      <span></span>
    {:then { data: prs }}
      {@const monthlies =
        prs?.filter(
          ({ event }) => event.player_class === selected_class && event.kind === 'monthly'
        ) ?? []}
      {@const archives =
        prs?.filter(
          ({ event }) => event.player_class === selected_class && event.kind === 'archive'
        ) ?? []}
      {#if monthlies.length}
        <Section label="past monthlies">
          <Table data={monthlies}>
            {#snippet header()}
              <th class="w-rank"></th>
              <th class="w-time"></th>
              <th class=""></th>
              <th class="w-event"></th>
              <th class="w-0"></th>
              <th class="w-date"></th>
            {/snippet}
            {#snippet row({ event, leaderboard, time, position, prize }: EventLeaderboardTime)}
              <td class={twTableGradients.get(`r${position}`)}>{formatPosition(position)}</td>
              <td class={twTableGradients.get(`t${position}`)}><TableTime {time} /></td>
              <td class="overflow-hidden"
                ><TableMap map={leaderboard.map} div={leaderboard.div} /></td>
              <td><TableEvent {event} href={'formats/monthly'} /></td>
              {#if prize.keys !== 0}
                <td>{prize.keys} keys</td>
              {:else}
                <td></td>
              {/if}
              <td class="table-date"><TemporalDate datetime={time.created_at} /></td>
            {/snippet}
          </Table>
        </Section>
      {/if}

      {#if archives.length}
        <Section label="past archives">
          <Table data={archives}>
            {#snippet header()}
              <th class="w-rank"></th>
              <th class="w-time"></th>
              <th class=""></th>
              <th class="w-event"></th>
              <th class="w-0"></th>
              <th class="w-date"></th>
            {/snippet}
            {#snippet row({ event, leaderboard, time, position, prize }: EventLeaderboardTime)}
              <td class={twTableGradients.get(`r${position}`)}>{formatPosition(position)}</td>
              <td class={twTableGradients.get(`t${position}`)}><TableTime {time} /></td>
              <td class="overflow-hidden"
                ><TableMap map={leaderboard.map} div={leaderboard.div} /></td>
              <td><TableEvent {event} href={'formats/archive'} /></td>
              {#if prize.keys !== 0}
                <td>{prize.keys} keys</td>
              {:else}
                <td></td>
              {/if}
              <td class="table-date"><TemporalDate datetime={time.created_at} /></td>
            {/snippet}
          </Table>
        </Section>
      {/if}
    {/await}
  </Content>
{/if}
