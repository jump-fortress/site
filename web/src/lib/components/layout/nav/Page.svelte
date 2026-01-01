<script lang="ts">
  type Props = {
    label: string;
    subpages?: string[] | null;
    route: string;
  };
  let { label, route, subpages = null }: Props = $props();

  let href = $derived(label === 'home' ? '/' : `/${label}`);
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
      <hr class="absolute bottom-0 left-1/6 w-2/3 rounded-box border border-primary -z-10" />
    {/if}
  </a>

  <div class="absolute pl-2 flex flex-col z-10 hover:visible invisible group-hover:visible pt-2">
    {#each subpages as subpage}
      {@const href: string = `/${subpage}`}
      <a
        class="relative hover:text-content text-content/75 rounded-box pl-2 hover:bg-base-800/50 w-32"
        {href}>
        <span>{subpage}</span>
      </a>
    {/each}
  </div>
</div>
