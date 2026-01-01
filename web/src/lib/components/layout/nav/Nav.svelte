<script lang="ts">
  import { invalidateAll } from '$app/navigation';
  import steam_signin from '$lib/assets/components/nav/sits_small.png';
  import { onMount } from 'svelte';

  import Page from './Page.svelte';
  import NavSession from './Session.svelte';

  import type { Session } from '$lib/schema';

  type Props = {
    route: string;
    session: Session | null;
  };

  let { route, session }: Props = $props();

  // todo: invalidateAll() needed to refresh Nav when user logs in / out, is onMount() the right thing for this?
  onMount(() => {
    if (!session) {
      invalidateAll();
    }
  });
</script>

<nav
  class="sticky z-50 flex h-18 w-full justify-center bg-base-900/50 backdrop-blur-md select-none">
  <div class="flex w-full max-w-5xl items-center justify-between text-lg">
    <!-- left nav -->
    <div class="flex h-full items-center grow transition-colors hover:text-content/75">
      <Page label={'home'} {route} />
      <Page label={'ladder'} {route} />
      <Page label={'formats'} subpages={['monthly', 'motw', 'bounties', 'quests']} {route} />
      <Page label={'help'} subpages={['resources', 'faq', 'rules']} {route} />
      <!-- dropdown -->
      <div
        class="peer-hover:visible invisible absolute top-0 left-0 -z-20 flex h-44 w-full cursor-default flex-col gap-px border-b border-primary/50 bg-base-900/95 px-1 py-1 backdrop-blur-none transition-all shadow-sm shadow-primary/50">
      </div>
    </div>
    <!-- right nav -->
    <div class="relative flex h-full flex-row-reverse items-center" data-nav="true">
      {#if session}
        <NavSession {session} />
      {:else}
        <a href="http://localhost:8000/internal/session/steam/discover">
          <img class="cursor-pointer" src={steam_signin} alt="" />
        </a>
      {/if}
    </div>
  </div>
</nav>
