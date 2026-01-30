<script lang="ts">
  type Props = {
    map: string;
  };

  let { map }: Props = $props();
</script>

<div
  class="relative grid size-full scale-x-105 place-content-center text-shadow-xs/100 text-shadow-base-900 not-first:mask-l-from-97% not-last:mask-r-from-97%">
  <span class="relative z-10 truncate">{map}</span>
  <img
    class="absolute size-full object-cover brightness-75"
    src="https://tempusplaza.com/map-backgrounds/{map}.jpg"
    alt=""
    draggable="false"
    loading="lazy" />
</div>
