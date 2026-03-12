<script lang="ts">
  type Props = {
    href: string;
    label: string;
    newTab?: boolean;
  };

  let { href, label, newTab = false }: Props = $props();
</script>

<a
  class="text-primary/75 hover:text-primary hover:underline"
  {href}
  target={newTab ? '_blank' : '_self'}>{label}</a>
