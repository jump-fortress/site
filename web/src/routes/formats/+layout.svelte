<script>
	import FormatsNav from '$lib/components/FormatsNav.svelte';
	import { page } from '$app/state';

	let { children } = $props();
	let route = $derived(page.url.pathname.substring(1));
</script>

<FormatsNav {route} />

{@render children?.()}
