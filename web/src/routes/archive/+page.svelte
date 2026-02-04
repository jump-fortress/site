<script lang="ts">
  import TableEvents from '$lib/components/display/table/presets/TableEvents.svelte';
  import { Temporal } from 'temporal-polyfill';
  import type { PageData, RouteParams } from './$types';
  import Section from '$lib/components/layout/Section.svelte';
  import Content from '$lib/components/layout/Content.svelte';

  type Props = {
    data: PageData;
    params: RouteParams;
  };
  let { data }: Props = $props();
</script>

<Content>
  {#if data.events}
    <Section label="archived events">
      <TableEvents data={data.events} href={'archive'} onclick={() => {}}></TableEvents>
    </Section>
  {/if}

  <Section label="third party archives">
    <span>under construction..</span>
  </Section>
</Content>
