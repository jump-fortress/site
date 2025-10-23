<script lang="ts">
	import type { Snippet } from 'svelte';

	let { children, title }: { children: Snippet; title: string } = $props();
</script>

<div class="flex flex-col gap-1">
	<span class="text-ctp-lavender">{title}</span>
	{@render children()}
</div>
