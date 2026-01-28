<script lang="ts">
  import './layout.css';
  import favicon from '$lib/assets/favicon.svg';
  import Nav from '$lib/components/layout/nav/Nav.svelte';

  let { data, children } = $props();
</script>

<svelte:head><link rel="icon" href={favicon} /></svelte:head>

<div class="flex w-full justify-center">
  <div class="flex w-full max-w-5xl flex-col gap-1 bg-base-900">
    <Nav session={data.session} />
    <main class="bg-base-800">
      {@render children()}
    </main>
  </div>
</div>
