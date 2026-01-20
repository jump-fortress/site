<script lang="ts">
  import { invalidate, invalidateAll } from '$app/navigation';
  import DivisionTag from '$lib/components/display/DivisionTag.svelte';
  import MonthlyHeader from '$lib/components/display/MonthlyHeader.svelte';
  import Line from '$lib/components/display/charts/ScatterTimes.svelte';
  import Table from '$lib/components/display/table/Table.svelte';
  import TableDate from '$lib/components/display/table/TableDate.svelte';
  import TablePlayer from '$lib/components/display/table/TablePlayer.svelte';
  import TableTime from '$lib/components/display/table/TableTime.svelte';

  import type { TimeWithPlayer } from '$lib/schema';

  let { data } = $props();

  // todo: times don't populate on load.. do server loads not.. update reactive props?

  let selected_times: TimeWithPlayer[] | undefined | null = $state(null);
  let selected_div: string = $state('');

  // svelte-ignore state_referenced_locally
  selected_times = (await data.times)?.at(0)?.times;
  // svelte-ignore state_referenced_locally
  selected_div = (await data.monthly)?.divisions?.at(0)?.division ?? '';

  let selected_index = $state(0);
</script>

{#await data.monthly}
  <span>aaaa</span>
{:then monthly}
  {#if monthly}
    <MonthlyHeader {monthly} header={true} />

    {#await data.times}
      <span>aaaa</span>
    {:then competition_times}
      {#if selected_times}
        {#key selected_index}
          <div class="absolute top-48 left-0 h-fit w-full bg-base-900 py-2">
            <Line data={selected_times} div={selected_div} />
          </div>
          <hr class="h-56" />
        {/key}
      {/if}

      <div class="flex h-10 overflow-hidden rounded-box border border-base-700">
        {#each monthly.divisions as cd, i}
          {@const selected: boolean = selected_div === cd.division}
          <button
            class="flex grow cursor-pointer items-center justify-center {selected
              ? 'bg-base-900'
              : ''}"
            onclick={() => {
              selected_div = cd.division;
              selected_times = competition_times?.at(i)?.times;
            }}>
            <DivisionTag div={cd.division} />
          </button>
        {/each}
      </div>
      {#if selected_times}
        <Table data={selected_times}>
          {#snippet header()}
            <th class="w-rank"></th>
            <th></th>
            <th class="w-2/5"></th>
            <th class="w-date"></th>
          {/snippet}
          {#snippet row(t: TimeWithPlayer, i)}
            <td>#{i}</td>
            <td><TablePlayer player={t.player} flag={false} /></td>
            <td class="text-start"
              ><TableTime run_time={t.time.run_time} tempus_time_id={t.time.tempus_time_id} />
            </td>
            <td class="table-date"><TableDate date={t.time.created_at} /></td>
          {/snippet}
        </Table>
      {/if}
    {/await}
  {/if}
{/await}
