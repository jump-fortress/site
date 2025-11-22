<script lang="ts">
	type Props = {
		src: string;
		name: string;
	};

	let { src, name } = $props();
</script>

<div class="flex items-center gap-2">
	<img {src} alt="" class="size-8 rounded-lg" />
	<span class="span-ellipsis">{name}</span>
</div>
