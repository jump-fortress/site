<script lang="ts">
  import Table from '../Table.svelte';
</script>

<Table data={[{}, {}, {}, {}, {}, {}, {}, {}]}>
  {#snippet header()}
    <th></th>
  {/snippet}
  {#snippet row()}
    <td class="animate-[pulse_3s_infinite]"></td>{/snippet}
</Table>
