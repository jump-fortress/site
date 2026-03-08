<script lang="ts">
  import { Client } from '$lib/api/api';
  import Table from '$lib/components/display/table/Table.svelte';
  import TablePlayer from '$lib/components/display/table/TablePlayer.svelte';
  import TemporalDate from '$lib/components/display/TemporalDate.svelte';
  import { ApiPaths, type AuditLog } from '$lib/schema';
</script>

{#await Client.GET(ApiPaths.get_audit_logs)}
  <span></span>
{:then { data: logs }}
  {#if logs}
    <Table data={logs}>
      {#snippet header()}
        <th></th>
      {/snippet}
      {#snippet row(l: AuditLog)}
        <td class="px-2"
          ><div class="flex items-center gap-1">
            <div class="flex min-w-fit">
              <TablePlayer player={l.from_player} flag={false} />
            </div>
            <span>{l.kind} for</span>
            <div class="flex min-w-fit">
              <TablePlayer player={l.to_player} flag={false} />
            </div>
            <span>from {l.from_content} to {l.to_content}</span>
            <div class="ml-auto">
              <TemporalDate datetime={l.created_at} />
            </div>
          </div></td>
      {/snippet}
    </Table>
  {:else}
    <span>no logs..</span>
  {/if}
{/await}
