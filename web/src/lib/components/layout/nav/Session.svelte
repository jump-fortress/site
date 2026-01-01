<script lang="ts">
  import Dropdown from './Dropdown.svelte';

  import type { Session } from '$lib/schema';

  type Props = {
    session: Session;
  };

  let { session }: Props = $props();

  let dropdown = $state(false);

  // @ts-ignore
  function onmousedown(event) {
    if (!event.target.dataset['nav']) {
      dropdown = false;
    }
  }
</script>

<!-- hide nav if it's not clicked on -->
<svelte:window {onmousedown} />

{#if dropdown}
  <Dropdown {session} />
{/if}

<button
  class="group relative cursor-pointer px-4"
  onclick={() => {
    dropdown = true;
  }}
  data-nav="true">
  <img
    class="size-14 rounded-full border-2 border-base-700 group-hover:border-content/50"
    src={session.steamAvatarUrl}
    alt=""
    data-nav="true" />
</button>

<!-- svelte-ignore a11y_consider_explicit_label -->
<a
  class="flex items-center px-2 text-content/75 transition-colors hover:text-content"
  href="/support">
  <span class="icon-[mdi--heart-outline] size-6"></span>
</a>

<!-- extra left padding for dropdown -->
<!-- svelte-ignore a11y_consider_explicit_label -->
<a
  class="flex items-center pr-2 pl-4 text-content/75 transition-colors hover:text-content"
  href="https://discord.gg/tusBc64wnv"
  target="_blank">
  <span class="icon-[mdi--discord] size-6"></span>
</a>
