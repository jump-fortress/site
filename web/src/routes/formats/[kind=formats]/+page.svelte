<script lang="ts">
  import TableEvents from '$lib/components/display/table/presets/TableEvents.svelte';
  import { Temporal } from 'temporal-polyfill';
  import type { PageData, RouteParams } from './$types';
  import type { EventWithLeaderboards } from '$lib/schema';
  import Section from '$lib/components/layout/Section.svelte';
  import EventHeader from '$lib/components/display/EventHeader.svelte';

  type Props = {
    data: PageData;
    params: RouteParams;
  };
  let { data, params }: Props = $props();
  const pluralKind = $derived(`${params.kind.replace('y', 'ie')}s`);

  const now = Temporal.Now.instant();
  const currEvents: EventWithLeaderboards[] = $derived(
    data.events?.filter(({ event }) => now.since(event.ends_at).seconds <= 0) ?? []
  );
  const prevEvents: EventWithLeaderboards[] = $derived(
    data.events?.filter(({ event }) => now.since(event.ends_at).seconds > 0) ?? []
  );
</script>

{#if currEvents.length}
  <Section label="ongoing {pluralKind}">
    {#each currEvents as ewl}
      <EventHeader event={ewl} />
    {/each}
  </Section>
{/if}

{#if prevEvents.length}
  <Section label="past {pluralKind}">
    <TableEvents data={prevEvents} link={true} onclick={() => {}}></TableEvents>
  </Section>
{/if}
