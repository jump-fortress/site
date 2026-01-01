<script lang="ts">
  type Props = {
    map: string;
  };

  let { map }: Props = $props();
</script>

<div
  class="grid size-full place-content-center mask-x-from-90% bg-cover bg-center"
  style:background-image={`url("https://tempusplaza.xyz/map-backgrounds/${map}.jpg")`}>
  <span class="truncate text-shadow-base-900 text-shadow-sm">{map}</span>
</div>
