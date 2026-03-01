<script lang="ts">
  type Props = {
    src?: string;
    href: string;
    label: string;
    newTab?: boolean;
  };

  let { src = '', href, label, newTab = false }: Props = $props();
</script>

<a
  class="flex items-center gap-1 text-sm text-primary opacity-75 hover:underline hover:opacity-100"
  {href}
  target={newTab ? '_blank' : ''}>
  {#if src}
    <img class="size-6" {src} alt="" />
  {/if}
  <span class="text-nowrap">{label}</span>
</a>
