<script lang="ts">
  type Props = {
    route: string;
    parentRoute: string;
    pages: string[];
  };

  let { route, parentRoute, pages }: Props = $props();
</script>

<nav
  class="absolute top-0 left-0 flex h-16 w-full rounded-t-layout bg-base-900/50 px-6 text-lg transition-colors select-none hover:text-content/75">
  <div class="relative flex items-center">
    {#each pages as label}
      {@render page(label)}
    {/each}
    <hr class="absolute bottom-4 left-1/24 w-11/12 rounded-box text-base-700" />
  </div>
</nav>

<hr class="invisible h-10 w-full" />

{#snippet page(label: string)}
  {@const currentRoute: boolean = (route.includes(label))}
  <a
    class="relative px-2 hover:text-content {currentRoute ? 'text-content' : ''}"
    href="/{parentRoute}/{label}">
    <span>{label}</span>
    {#if currentRoute}
      <hr class="absolute -bottom-0.5 left-1/12 z-10 w-5/6 rounded-box border border-primary" />
    {/if}
  </a>
{/snippet}
