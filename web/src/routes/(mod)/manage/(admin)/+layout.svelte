<script lang="ts">
  import { page } from '$app/state';
  import InnerNav from '$lib/components/layout/InnerNav.svelte';

  let { children } = $props();

  let route = $derived(page.url.pathname);
</script>

<div class="relative h-16">
  <InnerNav {route} parentRoute={'manage/formats'} pages={['monthly', 'motw', 'bounty', 'quest']} />
</div>
{@render children()}
