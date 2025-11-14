<script lang="ts">
	import '../app.css';
	import favicon from '$lib/assets/favicon.svg';
	import AnimatedLogo from '$lib/components/assets/AnimatedLogo.svelte';
	import Nav from '$lib/components/Nav.svelte';
	import { page } from '$app/state';
	import type { LayoutProps } from './$types';
	import type { Session } from '$lib/schema';
	import { invalidate, invalidateAll } from '$app/navigation';
	import { onMount } from 'svelte';

	let { data, children }: LayoutProps = $props();
	let session: Session | null = $derived(data.session);
	let route = $derived(page.url.pathname.substring(1));

	onMount(() => {
		if (!session) invalidateAll();
	});
</script>

<svelte:head>
	<link rel="icon" href={favicon} />
</svelte:head>

<div class="fixed -bottom-36 -left-36 -z-10">
	<AnimatedLogo />
</div>

<div data-sveltekit-reload>
	<Nav {route} {session} />
</div>

<div class="mt-6 flex w-full flex-col items-center">
	<div class="bg-jfgray-800 w-5xl flex flex-col gap-4 px-8 py-4">
		{@render children?.()}
	</div>
</div>
