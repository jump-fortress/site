<script>
  import InnerNav from '$lib/components/InnerNav.svelte';
  import { page } from '$app/state';

  let { children } = $props();
  let route = $derived(page.url.pathname.substring(1));
</script>

<InnerNav {route} parentRoute={'formats'} pages={['monthly', 'motw', 'quest', 'bounty']} />

{@render children?.()}
