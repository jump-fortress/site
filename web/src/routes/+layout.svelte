<script lang="ts">
	import '../app.css';
	import favicon from '$lib/assets/favicon.svg';
	import AnimatedLogo from '$lib/components/assets/AnimatedLogo.svelte';
	import Nav from '$lib/components/Nav.svelte';
	import { page } from '$app/state';

	let { children } = $props();
	let route = $derived(page.url.pathname.substring(1));
</script>

<svelte:head>
	<link rel="icon" href={favicon} />
</svelte:head>

<div class="fixed -bottom-36 -left-36 -z-10">
	<AnimatedLogo />
</div>

<Nav {route} />

<div class="mt-6 flex w-full flex-col items-center">
	<div class="bg-jfgray-800 w-5xl flex flex-col gap-4 px-8 py-4">
		{@render children?.()}
	</div>
</div>
