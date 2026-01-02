<script lang="ts">
  type Props = {
    map: string;
  };

  let { map }: Props = $props();
</script>

<div
  class="relative grid size-full place-content-center mask-x-from-90% object-cover object-center">
  <span class="relative z-10 truncate text-shadow-base-900 text-shadow-sm">{map}</span>
  <img
    class="absolute size-full"
    src="https://tempusplaza.xyz/map-backgrounds/{map}.jpg"
    alt=""
    loading="lazy" />
</div>
