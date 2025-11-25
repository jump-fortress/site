<script lang="ts">
  let { href, title }: { href: string; title: string } = $props();
</script>

<a {href} {title} class="opacity-50 transition-opacity hover:opacity-100">
  <span class="icon-[ri--link-m] size-6"></span>
</a>
