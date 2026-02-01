<script lang="ts">
  import Table from '$lib/components/display/table/Table.svelte';
  import TableEvent from '$lib/components/display/table/TableEvent.svelte';
  import TableMaps from '$lib/components/display/table/TableMaps.svelte';
  import TemporalDate from '$lib/components/display/TemporalDate.svelte';
  import type { EventWithLeaderboards } from '$lib/schema';

  type Props = {
    data: EventWithLeaderboards[];
    link?: boolean;
    onclick: (ewl: EventWithLeaderboards) => void;
  };

  let { data, link = false, onclick }: Props = $props();
</script>

<Table {data}>
  {#snippet header()}
    <th class="w-event"></th>
    <th></th>
    <th class="w-date"></th>
  {/snippet}
  {#snippet row({ event, leaderboards }: EventWithLeaderboards)}
    <td
      onclick={() => {
        onclick({ event, leaderboards });
      }}><TableEvent {event} {link} /></td>
    <td class="flex"><TableMaps leaderboards={leaderboards ?? []} /></td>
    <td class="table-date"><TemporalDate datetime={event.starts_at} /></td>
  {/snippet}
</Table>
