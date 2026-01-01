<script lang="ts">
  import { slide } from 'svelte/transition';

  import type { Session } from '$lib/schema';

  type Props = {
    session: Session;
  };

  let { session }: Props = $props();
</script>

<div
  in:slide={{ duration: 300 }}
  class="absolute top-0 right-0 -z-10 flex w-full cursor-default flex-col gap-px rounded-b-box border border-t-0 border-content/50 bg-base-900 px-1 py-1 pt-18 backdrop-blur-md"
  data-nav="true">
  {@render page(session.displayName, `/players/${session.id}`)}
  <hr class="relative left-1/24 my-px w-11/12 text-base-700" />
  {#if session.role === 'Admin' || session.role === 'Mod' || session.role === 'Consultant' || session.role === 'Treasurer'}
    {@render page('manage', '/manage')}
    <hr class="relative left-1/24 my-px w-11/12 text-base-700" />
  {/if}
  {@render page('settings', '/settings')}
  {@render page('logout', '/logout')}
</div>

{#snippet page(label: string, href: string)}
  <a
    class="truncate rounded-box pl-2 hover:bg-base-800"
    {href}
    data-nav="true"
    data-sveltekit-preload-data={label === 'logout' ? 'tap' : 'hover'}>
    {label}
  </a>
{/snippet}
