<script lang="ts">
	import type { Snippet } from 'svelte';

	let { children, title }: { children: Snippet; title: string } = $props();
</script>

<div class="flex flex-col gap-1">
	<div class="flex items-center gap-2">
		<span class="text-ctp-lavender">{title}</span>
		<div class="border-jfgray-700 border-y-1 grow"></div>
	</div>
	{@render children()}
</div>
