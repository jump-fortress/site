<script lang="ts">
  import { page } from '$app/state';
  import NavPage from '$lib/components/layout/nav/NavPage.svelte';
  import steam_signin from '$lib/assets/components/nav/sits_small.png';
  import NavSession from './NavSession.svelte';
  import type { Session } from '$lib/schema';
  import { onMount } from 'svelte';
  import { invalidateAll } from '$app/navigation';

  type Props = {
    session: Session | undefined;
  };
  let { session }: Props = $props();

  let route = $state(page.url.pathname);

  onMount(() => {
    if (!session) {
      invalidateAll();
    }
  });
</script>

<nav class="z-40 flex h-20 justify-between">
  <!-- left nav -->
  <div class="hover flex grow items-end hover:text-content/75">
    <NavPage label={'home'} href={'/'} {route} />
    <NavPage label={'ladder'} href={'/ladder'} {route} />
    <NavPage
      label={'formats'}
      href={'/formats/monthly'}
      subpages={[
        { label: 'monthly', href: '/formats/monthly' },
        { label: 'motw', href: '/formats/motw' },
        { label: 'bounty', href: '/formats/bounty' }
      ]}
      {route} />
    <NavPage
      label={'invitationals'}
      href={'/invitationals'}
      subpages={[
        { label: 'world cup', href: '/invitationals/jwc' },
        { label: 'playoffs', href: '/invitationals/playoffs' }
      ]}
      {route} />
    <NavPage label="archive" href="/archive" {route} />
    <NavPage
      label={'help'}
      href={'/help/faq'}
      subpages={[
        { label: 'faq', href: '/help/faq' },
        { label: 'resources', href: '/help/resources' }
      ]}
      {route} />
    <!-- subpages dropdown background -->
    <div
      class="invisible absolute top-0 left-0 -z-20 flex h-38 w-full cursor-default flex-col gap-px border-b border-primary/75 bg-base-900/90 px-1 py-1 shadow-sm shadow-primary/50 backdrop-blur-none peer-hover:visible">
    </div>
  </div>
  <!-- right nav -->
  <div class="relative flex flex-row-reverse items-end">
    {#if session}
      <NavSession {session} />
    {:else}
      <a class="mb-1" href="http://localhost:8000/internal/steam/discover">
        <img class="cursor-pointer" src={steam_signin} alt="" />
      </a>
    {/if}
  </div>
</nav>
