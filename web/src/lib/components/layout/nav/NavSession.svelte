<script lang="ts">
  import NavDropdown from '$lib/components/layout/nav/NavDropdown.svelte';

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
  <NavDropdown {session} />
{/if}

<button
  class="group relative cursor-pointer px-2"
  onclick={() => {
    dropdown = true;
  }}
  data-nav="true">
  <img
    class="size-14 rounded-full border border-base-900 group-hover:border-content/50"
    src={session.avatar_url}
    alt=""
    data-nav="true" />
</button>

<!-- extra left padding for dropdown -->
<!-- svelte-ignore a11y_consider_explicit_label -->
<a class="flex items-center px-2 py-1 text-content/75 hover:text-content" href="/" target="_blank">
  <span class="icon-[mdi--discord]"></span>
</a>

<!-- svelte-ignore a11y_consider_explicit_label -->
<a class="flex items-center p-1 pl-4 text-content/75 hover:text-content" href="/support">
  <span class="icon-[mdi--heart-outline]"></span>
</a>
