<script lang="ts">
	let { map }: { map: string } = $props();
</script>

<span class="relative z-10 overflow-ellipsis text-nowrap">{map}</span>
<div
	style={`background-image: url("https://tempusplaza.xyz/map-backgrounds/${map}.jpg")`}
	class="mask-x-from-80% absolute top-0 size-full bg-cover bg-right bg-no-repeat opacity-50"
></div>
