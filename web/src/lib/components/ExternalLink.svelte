<script lang="ts">
  type Props = {
    src: string;
    href: string;
    name: string;
  };

  let { src, href, name } = $props();
</script>

<a
  {href}
  class="flex items-end gap-1 pl-2 text-base text-ctp-blue/50 decoration-1 transition-colors first:pl-0 hover:text-ctp-blue hover:underline">
  <img {src} class="size-6" alt="" />
  <span class="flex">{name}</span>
</a>
