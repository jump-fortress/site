<script lang="ts">
  type Props = {
    label: string;
    href: string;
    subpages?: { label: string; href: string }[] | null;
    route: string;
  };
  let { label, href, route, subpages = null }: Props = $props();

  let currentRoute = $derived(
    (route === '/' && label === 'home') || (href !== '/' && route.startsWith(href))
  );
</script>

<!-- group for subpage hover, peer for dropdown hover -->
<div
  role="navigation"
  class="relative px-2 py-1 {currentRoute
    ? 'text-content'
    : 'text-content/75 hover:text-content'} {subpages ? 'group peer' : ''}">
  <a {href}>
    <span>{label}</span>
    <!-- {#if currentRoute}
      <hr class="absolute bottom-0.5 left-1/6 -z-10 w-2/3 rounded-box border border-primary" />
    {/if} -->
  </a>

  <div class="invisible absolute z-10 flex flex-col pt-1 group-hover:visible hover:visible">
    {#each subpages as subpage}
      <a
        class="relative w-32 rounded-box pl-1 text-content/75 hover:bg-base-800/50 hover:text-content"
        href={subpage.href}>
        <span>{subpage.label}</span>
      </a>
    {/each}
  </div>
</div>
