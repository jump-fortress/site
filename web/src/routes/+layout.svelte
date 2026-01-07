<script lang="ts">
  import { page } from '$app/state';
  import favicon from '$lib/assets/favicon.svg';
  import Logo from '$lib/components/layout/Logo.svelte';
  import Nav from '$lib/components/layout/nav/Nav.svelte';

  import './layout.css';

  let { data, children } = $props();
  let route = $derived(page.url.pathname);
</script>

<svelte:head><link rel="icon" href={favicon} /></svelte:head>

<div class="fixed -bottom-36 -left-36 -z-10">
  <Logo />
</div>

{#await data.session then session}
  <Nav {session} {route} />
{/await}

<div class="flex w-full flex-col items-center">
  <div class="relative flex w-full max-w-5xl flex-col gap-4 rounded-layout bg-base-800 p-6">
    {@render children()}
  </div>
</div>
