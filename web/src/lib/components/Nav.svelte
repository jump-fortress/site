<script lang="ts">
  import { slide } from 'svelte/transition';
  import type { Session } from '$lib/schema';
  import steam_signin_button from '$lib/assets/static/sits_small.png';
  import { onMount } from 'svelte';
  import { invalidateAll } from '$app/navigation';

  type Props = {
    route: string;
    session?: Session | null;
  };

  let { route, session = null }: Props = $props();
  let showNavMenu = $state(false);

  // todo: is this still necessary for updating the navbar when logging in / out?
  onMount(() => {
    if (!session) {
      invalidateAll();
    }
  });
</script>

<svelte:document
  onmousedown={(e) => {
    // @ts-ignore
    if (!e.target.dataset['nav']) {
      showNavMenu = false;
    }
  }}
/>

<nav
  class="fixed z-50 flex h-16 w-full justify-center bg-jfgray-900/75 backdrop-blur-sm select-none"
>
  <div class="flex w-5xl items-center justify-between">
    <!-- left nav -->
    <div class="flex h-full grow hover:text-ctp-lavender-50/75">
      {@render NavPage('')}
      {@render NavPage('ladder')}
      {@render NavPage('formats')}
      {@render NavPage('help')}
    </div>
    <!-- right nav -->
    <div class="relative flex h-16 flex-row-reverse items-center" data-nav="true">
      {#if !session}
        <a href="http://localhost:8000/internal/session/steam/discover">
          <img class="cursor-pointer" src={steam_signin_button} alt="" data-nav="true" />
        </a>
      {:else}
        <button
          class="group relative cursor-pointer px-4"
          onmousedown={() => {
            showNavMenu = true;
          }}
          data-nav="true"
        >
          <img
            class="size-14 rounded-full border-2 border-jfgray-800 group-hover:border-ctp-lavender/50"
            src={session.steamAvatarUrl}
            alt=""
            data-nav="true"
          />
          {#if showNavMenu}
            <ul
              in:slide
              class="absolute -top-2 right-0 -z-10 flex size-fit w-44 cursor-default flex-col gap-1 rounded-b-lg border-2 border-ctp-lavender/50 bg-jfgray-900/90 p-2 pt-16 text-start delay-150 starting:border-b-ctp-lavender/0"
              data-nav="true"
            >
              <a
                href="/players/{session.id}"
                data-nav="true"
                class="span-ellipsis rounded-lg pl-2 hover:bg-jfgray-700">{session.displayName}</a
              >

              <hr class="border-1 border-jfgray-700" />

              {#if session.role === 'Admin' || session.role === 'Mod'}
                <a
                  href="/moderation"
                  data-nav="true"
                  class="span-ellipsis rounded-lg pl-2 hover:bg-jfgray-700">moderation</a
                >

                <hr class="border-1 border-jfgray-700" />
              {/if}

              <a href="/settings" data-nav="true" class="rounded-lg pl-2 hover:bg-jfgray-700"
                >settings</a
              >
              <a
                href="/logout"
                data-sveltekit-preload-data="tap"
                data-nav="true"
                class="rounded-lg pl-2 hover:bg-jfgray-700"
              >
                logout
              </a>
            </ul>
          {/if}
        </button>
        <a
          class="flex items-center px-2 text-ctp-lavender-50 opacity-75 transition-opacity hover:opacity-100"
          href="/support"
          title="support"
        >
          <span class="icon-[ri--heart-line] size-6"></span>
        </a>
        <a
          class="flex items-center px-2 text-ctp-lavender-50 opacity-75 transition-opacity hover:opacity-100"
          href="https://discord.gg/tusBc64wnv"
          title="discord"
        >
          <span class="icon-[ri--discord-fill] size-6"></span>
        </a>
      {/if}
    </div>
  </div>
</nav>

<hr class="h-14 opacity-0" />

{#snippet NavPage(name: string)}
  <a
    href="/{name === '' ? '' : name}"
    class="flex h-full items-center px-4 hover:text-ctp-lavender-50"
    ><div class="relative">
      <span>{name === '' ? 'home' : name}</span>
      {#if (route.includes(name) && name !== '') || (route === '' && name === '')}
        <hr
          class="absolute right-0 left-0 m-auto flex w-11/12 rounded-full border-1 text-ctp-lavender"
        />
      {/if}
    </div></a
  >
{/snippet}
