<script>
  import { page } from '$app/state';
  import InnerNav from '$lib/components/layout/nav/InnerNav.svelte';

  let route = $derived(page.url.pathname);
  let { children } = $props();
</script>

<InnerNav
  label="manage"
  {route}
  parentRoute="manage"
  pages={['formats', 'leaderboards', 'players']} />

<div class="flex flex-col px-1">
  {@render children()}
</div>
