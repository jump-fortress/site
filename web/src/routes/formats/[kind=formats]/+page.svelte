<script lang="ts">
  import TableEvents from '$lib/components/display/table/presets/TableEvents.svelte';
  import { Temporal } from 'temporal-polyfill';
  import type { PageData, RouteParams } from './$types';
  import type { EventWithLeaderboards } from '$lib/schema';
  import Section from '$lib/components/layout/Section.svelte';
  import EventHeader from '$lib/components/display/EventHeader.svelte';
  import Content from '$lib/components/layout/Content.svelte';

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
  const pastEvents: EventWithLeaderboards[] = $derived(
    data.events?.filter(({ event }) => now.since(event.ends_at).seconds > 0) ?? []
  );
</script>

<Content>
  {#if currEvents.length}
    <Section label="current {currEvents.length > 1 ? pluralKind : currEvents.at(0)?.event.kind}">
      {#each currEvents as ewl}
        <EventHeader event={ewl} />
      {/each}
    </Section>
  {/if}

  {#if pastEvents.length}
    <Section label="past {pluralKind}">
      <TableEvents data={pastEvents} href="formats/{params.kind}" onclick={() => {}}></TableEvents>
    </Section>
  {/if}
</Content>
