<script lang="ts">
	import { page } from '$app/state';
	import DataSection from '$lib/components/DataSection.svelte';
	import InnerNav from '$lib/components/InnerNav.svelte';
	import type { Session } from '$lib/schema';
	import type { PageData } from './$types';

	let route = $derived(page.url.pathname.substring(1));

	let { data: pageData }: { data: PageData } = $props();
	let session: Session | null = $derived(pageData.session);
</script>

{#if session.role === 'admin'}
	<DataSection title={'Manage Formats'}>
		<InnerNav {route} parentRoute={'moderation'} pages={['monthly', 'motw', 'quest', 'bounty']} />
	</DataSection>
{/if}

{#if session.role === 'admin' || session.role === 'mod'}
	<DataSection title={'Manage Players'}>
		<span></span>
	</DataSection>
{/if}
