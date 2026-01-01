<script lang="ts">
  import { page } from '$app/state';
  import favicon from '$lib/assets/favicon.svg';
  import Logo from '$lib/components/layout/Logo.svelte';
  import Nav from '$lib/components/layout/nav/Nav.svelte';

  import './layout.css';

  import type { Session } from '$lib/schema';

  let { data, children } = $props();
  let route = $derived(page.url.pathname);
  let session: Session | null = $derived(data.session);
</script>

<svelte:head><link rel="icon" href={favicon} /></svelte:head>

<div class="fixed -bottom-36 -left-36 -z-10">
  <Logo />
</div>

<Nav {session} {route} />

<div class="flex w-full flex-col items-center">
  <div class="relative flex w-full max-w-5xl flex-col rounded-layout bg-base-800 p-6 gap-4">
    {@render children()}
  </div>
</div>
