<script lang="ts">
  import { invalidateAll } from '$app/navigation';
  import steam_signin from '$lib/assets/components/nav/sits_small.png';
  import { onMount } from 'svelte';

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
    <div class="flex h-full grow items-center transition-colors hover:text-content/75">
      {@render page('home')}
      {@render page('ladder')}
      {@render page('formats')}
      {@render page('help')}
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

{#snippet page(label: string)}
  {@const href: string = label === 'home' ? '/' : `/${label}`}
  {@const currentRoute: boolean = (route === '/' && label === 'home') || (href !== '/' && route.includes(href))}

  <a class="relative px-4 hover:text-content {currentRoute ? 'text-content' : ''}" {href}>
    <span>{label}</span>
    {#if currentRoute}
      <!-- underline -->
      <hr class="absolute bottom-0 left-1/6 w-2/3 rounded-box border border-primary" />
    {/if}
  </a>
{/snippet}
