<script lang="ts">
  import TableEvents from '$lib/components/display/table/presets/TableEvents.svelte';
  import type { PageData } from './$types';

  type Props = {
    data: PageData;
  };
  let { data }: Props = $props();
</script>

{#if data.events}
  <TableEvents data={data.events} link={true} onclick={() => {}}></TableEvents>
{/if}
