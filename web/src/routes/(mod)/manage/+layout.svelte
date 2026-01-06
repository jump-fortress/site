<script lang="ts">
  import { page } from '$app/state';
  import InnerNav from '$lib/components/layout/InnerNav.svelte';

  let { data, children } = $props();

  let route = $derived(page.url.pathname);
</script>

{#await data.session then session}
  {#if session}
    {#if session.role === 'Admin'}
      <InnerNav
        {route}
        parentRoute="manage"
        pages={['competitions', 'players', 'requests', 'payouts']} />
    {:else if session.role === 'Mod' || session.role === 'Consultant'}
      <InnerNav {route} parentRoute="manage" pages={['players', 'requests']} />
    {:else if session.role === 'Treasurer'}
      <InnerNav {route} parentRoute="manage" pages={['payouts']} />
    {/if}
  {/if}
{/await}

{@render children()}
