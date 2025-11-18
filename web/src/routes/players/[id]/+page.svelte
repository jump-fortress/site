<script lang="ts">
	import { page } from '$app/state';
	import rocket from '$lib/assets/static/classes/rocket.png';
	import sticky from '$lib/assets/static/classes/sticky.png';
	import DataSection from '$lib/components/DataSection.svelte';
	import DivisionTag from '$lib/components/DivisionTag.svelte';
	import Table from '$lib/components/Table.svelte';
	import TableMap from '$lib/components/TableMap.svelte';
	import { Client } from '$lib/internalApi.js';
	import type { PlayerProfile, Session } from '$lib/schema';
	import type { PageData } from './$types';
	import Header from './Header.svelte';

	let { data: pageData }: { data: PageData } = $props();
	let session: Session | null = $derived(pageData.session);

	const playerID = Number.parseInt(page.params.id as string);
	const { data: playerData } = await Client.GET('/internal/players/profile/{id}', {
		fetch: fetch,
		params: {
			path: { id: playerID }
		}
	});

	let player: PlayerProfile | undefined = $derived(playerData);
</script>

<svelte:boundary {pending}>
	{#if player}
		<Header {player} />
	{:else}
		todo no player
	{/if}
</svelte:boundary>

<DataSection title="Bounties Claimed">
	<div class="grid w-full grid-cols-3 gap-2">
		{#each { length: 1 }}
			{@render testBounty()}
		{/each}
	</div>
</DataSection>

<DataSection title="Trophies">
	<div class="grid w-full grid-cols-4 gap-2">
		{#each { length: 1 }}
			{@render testTrophy()}
		{/each}
	</div>
</DataSection>

<DataSection title="Quests Claimed">
	<div class="grid w-full grid-cols-4 gap-2">
		{#each { length: 1 }}
			{@render testQuest()}
		{/each}
	</div>
</DataSection>

<DataSection title="Competition History">
	<Table data={[{}]}>
		{#snippet header()}
			<th class="text-ctp-lavender-50/75 w-12"></th>
			<th class="w-1/10 text-ctp-lavender-50/75"></th>
			<th class="w-1/8 text-ctp-lavender-50/75">format</th>
			<th class="text-ctp-lavender-50/75">map</th>
			<th class="w-1/10 text-ctp-lavender-50/75">time</th>
			<th class="w-1/16 text-ctp-lavender-50/75">place</th>
			<th class="text-ctp-lavender-50/75 w-1/12">points</th>
		{/snippet}
		{#snippet row()}
			<td><img src={rocket} class="m-auto size-8" alt="" /></td>
			<td><div class="flex justify-center"><DivisionTag div="platinum" /></div></td>
			<td class="cursor-pointer decoration-1 hover:underline">monthly #1</td>
			<td class="relative">
				<TableMap map={'jump_escape_rc4'} />
			</td>
			<td>3:53.00</td>
			<td>2nd</td>
			<td>9000</td>
		{/snippet}
	</Table>
</DataSection>

{#snippet testTrophy()}
	<div class="bg-jfgray-900 relative flex h-20 overflow-hidden">
		<span class="icon-[ant-design--trophy-outlined] text-ctp-teal absolute -left-10 z-0 size-20"
		></span>
		<div class="absolute z-10 size-full p-1">
			<div class="ml-10 flex h-full flex-col justify-between">
				<div class="flex flex-col">
					<div class="flex items-center gap-2">
						<span class="text-ctp-yellow text-shadow-2xs text-shadow-ctp-yellow">1st</span>
						<span class="span-ellipsis text-2xl/4">jump_escape_rc4</span>
					</div>
					<span class="font-semibold">3:53.00</span>
				</div>
				<span class="text-2xl/4 opacity-75">monthly #1</span>
			</div>
		</div>
		<div
			style={`background-image: url("https://tempusplaza.xyz/map-backgrounds/jump_escape_rc4.jpg")`}
			class="mask-l-from-50% mask-l-to-75% size-full bg-cover bg-right bg-no-repeat opacity-50"
		></div>
	</div>
{/snippet}

{#snippet testBounty()}
	<div class="bg-jfgray-900 relative flex h-28 items-center overflow-hidden">
		<span
			class="icon-[ri--star-line] bg-linear-to-b from-ctp-lavender-950 to-ctp-lavender-50 absolute -left-12 z-0 size-28"
		></span>
		<div class="absolute z-10 size-full p-2">
			<div class="ml-6 flex h-full flex-col gap-0">
				<div class="flex items-center gap-2">
					<span class="text-shadow-xs text-shadow-ctp-lavender">100 keys</span>
					<span class="text-2xl/4 opacity-90">for sub 4:00</span>
				</div>
				<span class="span-ellipsis ml-9 text-2xl/4">jump_escape_rc4</span>
				<span class="text-ctp-lavender ml-9 mt-3 text-4xl/6 font-semibold">3:53.00</span>
			</div>
		</div>
		<div
			style={`background-image: url("https://tempusplaza.xyz/map-backgrounds/jump_escape_rc4.jpg")`}
			class="mask-l-from-50% mask-l-to-75% size-full bg-cover bg-right bg-no-repeat opacity-50"
		></div>
	</div>
{/snippet}

{#snippet testQuest()}
	<div class="bg-jfgray-900 relative flex h-14 overflow-hidden">
		<span class="icon-[ri--treasure-map-line] text-ctp-teal absolute -left-6 z-0 size-14"></span>
		<div class="absolute z-10 size-full p-1">
			<div class="flex h-full flex-col">
				<div class="ml-10 flex flex-col text-2xl/4">
					<span class="span-ellipsis">jump_escape_rc4</span>
					<span>completion</span>
				</div>
			</div>
		</div>
		<div
			class="mask-l-from-50% mask-l-to-75% size-full bg-[url(https://tempusplaza.xyz/map-backgrounds/jump_escape_rc4.jpg)] bg-cover bg-right bg-no-repeat opacity-50"
		></div>
	</div>
{/snippet}

{#snippet pending()}
	todo loading
{/snippet}
