<script lang="ts">
  import { page } from '$app/state';
  import InnerNav from '$lib/components/InnerNav.svelte';
  import type { Session } from '$lib/schema';
  import type { LayoutProps } from './$types';

  let { data, children }: LayoutProps = $props();
  let route = $derived(page.url.pathname.substring(1));
</script>

<InnerNav
  {route}
  parentRoute={'manage'}
  pages={['competitions', 'players', 'requests', 'payouts']} />

{@render children()}
