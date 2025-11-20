<script lang="ts">
	import { page } from '$app/state';
	import DataSection from '$lib/components/DataSection.svelte';
	import InnerNav from '$lib/components/InnerNav.svelte';
	import type { FullPlayer, Session } from '$lib/schema';
	import type { LayoutData } from './$types';

	let route = $derived(page.url.pathname.substring(1));

	let { data }: { data: LayoutData } = $props();

	// there will always be a session, since we redirect otherwise
	let session: Session = $derived(data.session as Session);
	let players: FullPlayer[] | null = $derived(data.players);
	$inspect(data);
</script>

{#if session.role === 'admin'}
	<DataSection title={'Manage Competitions'}>
		<InnerNav {route} parentRoute={'moderation'} pages={['monthly', 'motw', 'quest', 'bounty']} />
	</DataSection>
{/if}
{#if session.role === 'admin' || session.role === 'mod'}
	<DataSection title={'Manage Players'}>
		<svelte:boundary {pending}>
			{#if players}
				{#each players as player}
					<span>{player.display_name}</span>
				{/each}
			{/if}
		</svelte:boundary>
	</DataSection>
{/if}

{#snippet pending()}
	todo loading
{/snippet}
