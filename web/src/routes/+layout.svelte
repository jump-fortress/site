<script lang="ts">
  import './layout.css';
  import favicon from '$lib/assets/favicon.svg';
  import Nav from '$lib/components/layout/nav/Nav.svelte';
  import type { PageData } from './$types';
  import type { Snippet } from 'svelte';
  import Logo from '$lib/components/layout/Logo.svelte';

  type Props = {
    data: PageData;
    children: Snippet;
  };
  let { data, children }: Props = $props();
</script>

<svelte:head><link rel="icon" href={favicon} /></svelte:head>

<div class="flex w-full justify-center">
  <div class="flex w-full max-w-5xl flex-col gap-1 bg-base-900">
    <Nav session={data.session} />
    <main class="flex flex-col bg-base-800">
      {@render children()}
    </main>
  </div>
</div>

<div class="fixed -bottom-36 -left-36 -z-10">
  <Logo />
</div>
