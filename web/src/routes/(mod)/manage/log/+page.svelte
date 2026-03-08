<script lang="ts">
  import { Client } from '$lib/api/api';
  import Table from '$lib/components/display/table/Table.svelte';
  import TablePlayer from '$lib/components/display/table/TablePlayer.svelte';
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
        <td
          ><div class="flex">
            <TablePlayer player={l.from_player} />
            {l.kind} for <TablePlayer player={l.to_player} /> from {l.from_content} to {l.to_content}
          </div></td>
      {/snippet}
    </Table>
  {:else}
    <span>no logs..</span>
  {/if}
{/await}
