<script>
  import { page } from '$app/state';
  import NavPage from '$lib/components/layout/nav/NavPage.svelte';
  import steam_signin from '$lib/assets/components/nav/sits_small.png';
  import { ApiPaths } from '$lib/schema';
  import { Client } from '$lib/api/api';
  import NavSession from './NavSession.svelte';

  let route = $state(page.url.pathname);
</script>

<nav class="z-40 mt-2 flex justify-between">
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
        { label: 'bounty', href: '/formats/bounty' },
        { label: 'quest', href: '/formats/quest' }
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
        { label: 'rules', href: '/help/rules' },
        { label: 'resources', href: '/help/resources' }
      ]}
      {route} />
    <!-- subpages dropdown background -->
    <div
      class="invisible absolute top-0 left-0 -z-20 flex h-42 w-full cursor-default flex-col gap-px border-b border-primary/75 bg-base-900/95 px-1 py-1 shadow-sm shadow-primary/50 backdrop-blur-none peer-hover:visible">
    </div>
  </div>
  <!-- right nav -->
  <div class="relative flex flex-row-reverse items-end">
    {#await Client.GET(ApiPaths.get_session)}
      <div
        class="group relative mx-2 size-14 rounded-full border border-base-900 bg-primary/50 group-hover:border-content/50">
      </div>

      <!-- extra left padding for dropdown -->
      <!-- svelte-ignore a11y_consider_explicit_label -->
      <a
        class="flex items-center px-2 py-1 text-content/75 hover:text-content"
        href="https://discord.gg/tusBc64wnv"
        target="_blank">
        <span class="icon-[mdi--discord]"></span>
      </a>

      <!-- svelte-ignore a11y_consider_explicit_label -->
      <a class="flex items-center p-1 pl-4 text-content/75 hover:text-content" href="/support">
        <span class="icon-[mdi--heart-outline]"></span>
      </a>
    {:then { data: session }}
      {#if session}
        <NavSession {session} />
      {:else}
        <a href="http://localhost:8000/internal/steam/discover">
          <img class="cursor-pointer" src={steam_signin} alt="" />
        </a>
      {/if}
    {/await}
  </div>
</nav>
