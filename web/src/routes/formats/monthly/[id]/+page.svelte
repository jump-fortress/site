<script lang="ts">
  import DivisionTag from '$lib/components/display/DivisionTag.svelte';
  import MonthlyHeader from '$lib/components/display/MonthlyHeader.svelte';
  import Line from '$lib/components/display/charts/ScatterTimes.svelte';
  import Table from '$lib/components/display/table/Table.svelte';
  import TableDate from '$lib/components/display/table/TableDate.svelte';
  import TablePlayer from '$lib/components/display/table/TablePlayer.svelte';
  import { dateToMs } from '$lib/src/temporal.js';
  import { formatRunTime } from '$lib/src/times.js';

  import type { TimeWithPlayer } from '$lib/schema';

  let { data } = $props();

  // todo: times don't populate on load.. do server loads not.. update reactive props?
</script>

{#await data.monthly}
  <span>aaaa</span>
{:then monthly}
  {#if monthly}
    <MonthlyHeader {monthly} />

    {#await data.times}
      <span>aaaa</span>
    {:then competition_times}
      {#each competition_times as cd_times}
        {#if cd_times.times}
          {@const div = monthly.divisions?.find((cd) => cd.id === cd_times.id)?.division}
          <DivisionTag {div} />
          <div class="flex w-full items-start gap-2">
            <div class="flex basis-2/5">
              <Table data={cd_times.times}>
                {#snippet header()}
                  <th class="w-8"></th>
                  <th></th>
                  <th></th>
                  <th class="w-20"></th>
                  <th class="w-32"></th>
                {/snippet}
                {#snippet row(t: TimeWithPlayer, i)}
                  <td class="text-end">{i + 1}</td>
                  <td>https://tempus2.xyz/records/{t.time.tempus_time_id}</td>
                  <td><TablePlayer player={t.player} flag={false} /></td>
                  <td class="text-start">{formatRunTime(t.time.run_time)}</td>
                  <td class="pr-3 text-end"><TableDate date={t.time.created_at} /></td>
                {/snippet}
              </Table>
            </div>
            <div class="h-full basis-3/5 bg-base-700">
              <Line data={cd_times.times} />
            </div>
          </div>
        {/if}
      {/each}
    {/await}
  {/if}
{/await}
