<script lang="ts">
  type Props = {
    map: string;
  };

  let { map }: Props = $props();
</script>

<div class="relative grid size-full scale-x-105 place-content-center mask-x-from-97%">
  <span class="relative z-10 truncate text-shadow-base-900 text-shadow-sm">{map}</span>
  <img
    class="absolute size-full object-cover"
    src="https://tempusplaza.xyz/map-backgrounds/{map}.jpg"
    alt=""
    loading="lazy" />
</div>
