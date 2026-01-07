<script lang="ts">
  type Props = {
    label: string;
    href: string;
    subpages?: { label: string; href: string }[] | null;
    route: string;
  };
  let { label, href, route, subpages = null }: Props = $props();

  let currentRoute = $derived(
    (route === '/' && label === 'home') || (href !== '/' && route.includes(href))
  );
</script>

<!-- group for subpage hover, peer for dropdown hover -->
<div role="navigation" class="relative hover:text-content {subpages ? 'group peer' : ''}">
  <a class="px-4 {currentRoute ? 'text-content' : ''}" {href}>
    <span>{label}</span>
    {#if currentRoute}
      <!-- underline -->
      <hr class="absolute bottom-0 left-1/6 -z-10 w-2/3 rounded-box border border-primary" />
    {/if}
  </a>

  <div class="invisible absolute z-10 flex flex-col pt-2 pl-2 group-hover:visible hover:visible">
    {#each subpages as subpage}
      <a
        class="relative w-32 rounded-box pl-2 text-content/75 hover:bg-base-800/50 hover:text-content"
        href={subpage.href}>
        <span>{subpage.label}</span>
      </a>
    {/each}
  </div>
</div>
