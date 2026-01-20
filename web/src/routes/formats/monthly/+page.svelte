<script lang="ts">
  import MonthlyHeader from '$lib/components/display/MonthlyHeader.svelte';
  import Table from '$lib/components/display/table/Table.svelte';
  import TableCompetition from '$lib/components/display/table/TableCompetition.svelte';
  import TableDate from '$lib/components/display/table/TableDate.svelte';
  import TableMap from '$lib/components/display/table/TableMap.svelte';
  import Section from '$lib/components/layout/Section.svelte';

  import type { Monthly } from '$lib/schema';

  let { data } = $props();
</script>

{#await data.monthlies}
  <span></span>
{:then monthlies}
  {#if monthlies}
    {#each monthlies as monthly, i}
      {#if !monthly.competition.complete}
        <MonthlyHeader {monthly} link={true} />
      {/if}
    {/each}

    <!-- completed monthlies -->
    <Section label="past results">
      <Table data={monthlies.filter(({ competition }) => competition.complete)}>
        {#snippet header()}
          <th class="w-40"></th>
          <th></th>
          <th class="w-32"></th>
        {/snippet}
        {#snippet row(m: Monthly)}
          {@const maps = new Set(m.divisions?.map((cd) => cd.map))}
          <td
            ><a href="/formats/monthly/{m.id}"
              ><TableCompetition competition={m.competition} format="monthly" formatId={m.id} /></a
            ></td>
          <td
            ><div class="flex h-full overflow-hidden">
              {#each maps as map}
                <TableMap {map} />
              {/each}
            </div></td>
          <td><TableDate date={m.competition.ends_at} /></td>
        {/snippet}
      </Table>
    </Section>
  {/if}
{/await}
