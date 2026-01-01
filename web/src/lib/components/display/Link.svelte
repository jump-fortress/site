<script lang="ts">
  type Props = {
    src: string;
    href: string;
    label: string;
    newTab?: boolean;
  };

  let { src, href, label, newTab = false }: Props = $props();
</script>

<a
  class="flex items-center gap-1 text-primary opacity-75 hover:underline hover:opacity-100"
  {href}
  target={newTab ? '_blank' : ''}>
  <img class="size-6" {src} alt="" />
  <span>{label}</span>
</a>
