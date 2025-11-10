<script lang="ts">
	let { map }: { map: string } = $props();
</script>

<span class="text-shadow-sm/50 span-ellipsis relative z-10">{map}</span>
<div
	style={`background-image: url("https://tempusplaza.xyz/map-backgrounds/${map}.jpg")`}
	class="mask-x-from-80% absolute top-0 size-full bg-cover bg-right bg-no-repeat opacity-50"
></div>
