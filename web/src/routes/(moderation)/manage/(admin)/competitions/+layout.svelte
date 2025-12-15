<script lang="ts">
  import { page } from '$app/state';
  import DataSection from '$lib/components/DataSection.svelte';
  import InnerNav from '$lib/components/InnerNav.svelte';
  import type { LayoutProps } from './$types';

  let route = $derived(page.url.pathname.substring(1));

  let { children }: LayoutProps = $props();
</script>

<DataSection title={'Manage Competitions'}>
  <InnerNav
    {route}
    parentRoute={'manage/competitions'}
    pages={['monthly', 'motw', 'quest', 'bounty']}
  />
</DataSection>

{@render children()}
