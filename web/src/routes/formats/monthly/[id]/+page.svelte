<script lang="ts">
  import { Client } from '$lib/api/api';
  import { ApiPaths } from '$lib/schema';
</script>

{#await Client.GET(ApiPaths.get_event_kinds, { params: { path: { event_kind: 'monthly' } } })}
  <span></span>
{:then { data: monthlies }}
  <span>under construction</span>
  {#each monthlies as monthly}
    <span>{monthly.event.kind} {monthly.event.kind_id}</span>
  {/each}
{/await}
