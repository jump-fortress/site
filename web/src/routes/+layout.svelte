<script lang="ts">
  import './layout.css';
  import favicon from '$lib/assets/favicon.svg';
  import Nav from '$lib/components/layout/nav/Nav.svelte';

  let { children } = $props();
</script>

<svelte:head><link rel="icon" href={favicon} /></svelte:head>

<div class="flex w-full justify-center">
  <div class="bg-base-902 flex w-full max-w-5xl flex-col gap-1">
    <Nav />
    {@render children()}
  </div>
</div>
