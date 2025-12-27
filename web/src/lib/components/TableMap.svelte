<script lang="ts">
  let { map }: { map: string } = $props();
</script>

<span class="span-ellipsis relative z-10 text-shadow-sm/50">{map}</span>
<div
  style={`background-image: url("https://tempusplaza.xyz/map-backgrounds/${map}.jpg")`}
  class="absolute top-0 size-full mask-x-from-80% bg-cover bg-right bg-no-repeat opacity-50">
</div>
