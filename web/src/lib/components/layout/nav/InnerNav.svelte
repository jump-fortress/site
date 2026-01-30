<script lang="ts">
  import Section from '../Section.svelte';
  import NavPage from './NavPage.svelte';

  type Props = {
    label: string;
    route: string;
    parentRoute: string;
    pages: string[];
  };

  let { label, route, parentRoute, pages }: Props = $props();
</script>

<nav class="flex flex-col bg-base-900/50 px-1 hover:text-content/75">
  <Section {label} indent="/">
    <div class="flex items-end">
      {#each pages as label}
        <NavPage {label} href="/{parentRoute}/{label}" {route} />
      {/each}
    </div>
  </Section>
  <!-- <hr class="absolute bottom-4 left-1/24 w-11/12 rounded-box text-base-700" /> -->
</nav>
