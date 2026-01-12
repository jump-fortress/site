<script lang="ts">
  import MonthlyHeader from '$lib/components/display/MonthlyHeader.svelte';
  import Table from '$lib/components/display/table/Table.svelte';
  import TableCompetition from '$lib/components/display/table/TableCompetition.svelte';
  import TableDate from '$lib/components/display/table/TableDate.svelte';
  import TableMap from '$lib/components/display/table/TableMap.svelte';

  import type { Monthly } from '$lib/schema';

  let { data } = $props();
</script>

{#await data.monthlies then monthlies}
  {#if monthlies}
    {#each monthlies as monthly, i}
      {#if !monthly.competition.complete}
        <MonthlyHeader {monthly} header={i === 0} link={true} />
      {/if}
    {/each}

    <!-- completed monthlies -->
    <Table data={monthlies.filter(({ competition }) => competition.complete)}>
      {#snippet header()}
        <th></th>
        <th class="w-48"></th>
        <th class="w-48"></th>
        <th class="w-32">end date</th>
      {/snippet}
      {#snippet row(m: Monthly)}
        <td><TableCompetition competition={m.competition} format="monthly" formatId={m.id} /></td>
        <td><TableMap map="jump_beef" /></td>
        <td><TableMap map="jump_escape_rc4" /></td>
        <td><TableDate date={m.competition.ends_at} /></td>
      {/snippet}
    </Table>
  {/if}
{/await}
