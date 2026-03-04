<script lang="ts">
  import { page } from '$app/state';
  import InnerNav from '$lib/components/layout/nav/InnerNav.svelte';

  let route = $derived(page.url.pathname);
  let pages: string[] = $derived.by(() => {
    let a = ['players, requests'];
    if (data.session) {
      if (data.session.role === 'admin' || data.session.role === 'dev') {
        a = ['formats', 'leaderboards', 'prizepools', 'players', 'requests', 'timeslots'];
      }
    }
    return a;
  });

  let { data, children } = $props();
</script>

<InnerNav label="manage" {route} parentRoute="manage" {pages} />

<div class="flex flex-col px-1">
  {@render children()}
</div>
