<script lang="ts">
  import { Client } from '$lib/api/api';
  import EventHeader from '$lib/components/display/EventHeader.svelte';
  import { ApiPaths, type Time } from '$lib/schema';
  import { onMount } from 'svelte';
  import type { PageData } from './$types';
  import Table from '$lib/components/display/table/Table.svelte';
  import Div from '$lib/components/display/Div.svelte';
  import Section from '$lib/components/layout/Section.svelte';

  type Props = {
    data: PageData;
  };
  let { data }: Props = $props();
  let selectedLeaderboardID = $state(0);

  onMount(() => {
    if (data.ewl) {
      selectedLeaderboardID = data.ewl.leaderboards?.at(0)?.id ?? 0;
    }
  });
</script>

{#if data.ewl}
  <EventHeader event={data.ewl} />

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
      {#if times}
        <Table data={times}>
          {#snippet header()}
            <th></th>
          {/snippet}
          {#snippet row(t: Time)}
            <td>{t.duration}</td>
          {/snippet}
        </Table>
      {/if}
    {/await}
  </Section>
{/if}
