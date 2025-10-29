<script lang="ts">
	import rocket from '$lib/assets/static/classes/rocket.png';
	import sticky from '$lib/assets/static/classes/sticky.png';
	import plaza from '$lib/assets/static/players/plaza.png';
	import tempus from '$lib/assets/static/players/tempus.png';
	import zigzagoon from '$lib/assets/static/players/zigzagoon.gif';
	import DataSection from '$lib/components/DataSection.svelte';
	import DivisionTag from '$lib/components/DivisionTag.svelte';
	import FullDivider from '$lib/components/FullDivider.svelte';
	import Table from '$lib/components/Table.svelte';

	let showSoldier = $state(true);
	let showDemo = $state(false);
	let me = $state(true);
</script>

<!-- top profile -->
<div class="h-46 relative flex gap-4">
	<img
		class="size-46"
		src="https://avatars.akamai.steamstatic.com/55f57bf80f543b1e7261f9e0a17b8e12b992de28_full.jpg"
		alt=""
	/>
	<img
		style="image-rendering: pixelated"
		src={zigzagoon}
		alt=""
		class="scale-200 absolute bottom-0"
	/>

	<div class="flex h-full grow flex-col justify-between">
		<!-- name and div -->
		<div class="flex flex-col">
			<span class="text-6xl/10">mur</span>
			<span class="text-division-diamond text-4xl/6">Platinum Soldier</span>
		</div>

		<!-- points -->
		<div class="flex flex-col gap-2">
			<div class="text-ctp-blue/50 flex">
				<a
					href="https://tempusplaza.xyz/players/254279"
					class="hover:text-ctp-blue flex items-end gap-1 pr-2 transition-colors"
				>
					<img src={plaza} class="size-6" alt="" />
					<span class="flex">Plaza</span>
				</a>
				<a
					href="https://tempus2.xyz/players/254279"
					class="hover:text-ctp-blue flex items-end gap-1 pl-2 transition-colors"
				>
					<img src={tempus} class="size-6" alt="" />
					<span class="flex">Tempus</span>
				</a>
			</div>

			<div class="flex w-full gap-2">
				<div class="bg-jfgray-900 flex flex-col px-2 py-1">
					<span class="text-2xl/4 opacity-75">last 3 monthlies</span>
					<div class="flex items-center gap-2">
						<span>555555</span>
						<span class="ml-auto text-2xl/4">(#51)</span>
					</div>
				</div>
				<div class="bg-jfgray-900 flex flex-col px-2 py-1">
					<span class="text-2xl/4 opacity-75">last 9 MOTWs</span>
					<div class="flex items-center gap-2">
						<span>123456</span>
						<span class="ml-auto text-2xl/4">(#12)</span>
					</div>
				</div>
				<div class="bg-jfgray-900 flex flex-col px-2 py-1">
					<span class="text-2xl/4 opacity-75">total points</span>
					<div class="flex items-center gap-2">
						<span>4738291</span>
						<span class="ml-auto text-2xl/4">(#21)</span>
					</div>
				</div>
			</div>
		</div>
	</div>

	<!-- class select -->
	<div class="flex h-full flex-col">
		<div
			class="bg-jfgray-900 flex basis-1/2 cursor-pointer items-center p-3 pt-4 opacity-100 hover:opacity-100"
		>
			<img
				class="size-16 select-none"
				style="filter: brightness(0) saturate(100%) invert(97%) sepia(49%) saturate(6023%) hue-rotate(179deg) brightness(100%) contrast(108%)"
				src={rocket}
				alt=""
				draggable="false"
			/>
		</div>
		<div
			class="flex basis-1/2 cursor-pointer items-center p-3 pb-4 opacity-50 transition-opacity hover:opacity-100"
		>
			<img
				class="size-16 select-none"
				style="filter: brightness(0) saturate(100%) invert(97%) sepia(49%) saturate(6023%) hue-rotate(179deg) brightness(100%) contrast(108%)"
				src={sticky}
				alt=""
				draggable="false"
			/>
		</div>
	</div>
</div>

<FullDivider />

<DataSection title="Bounties Claimed">
	<div class="flex flex-wrap gap-x-4 gap-y-3">
		{#each { length: 1 }}
			{@render testBounty()}
		{/each}
	</div>
</DataSection>

<FullDivider />

<DataSection title="Trophies">
	<div class="flex flex-wrap gap-2">
		{#each { length: 1 }}
			{@render testTrophy()}
		{/each}
	</div>
</DataSection>

<FullDivider />

<DataSection title="Quests Claimed">
	<div class="flex flex-wrap gap-2">
		{#each { length: 1 }}
			{@render testQuest()}
		{/each}
	</div>
</DataSection>

<FullDivider />

<DataSection title="Competition History">
	<Table data={[{}]}>
		{#snippet header()}
			<th class="w-12"></th>
			<th class="w-1/10"></th>
			<th class="w-1/8">format</th>
			<th>map</th>
			<th class="w-1/10">time</th>
			<th class="w-1/16">place</th>
			<th class="w-1/12">points</th>
			<th class="w-1/10">date</th>
			<th class="w-14">results</th>
		{/snippet}
		{#snippet row()}
			<td><img src={rocket} class="m-auto size-8" alt="" /></td>
			<td><div class="flex justify-center"><DivisionTag div="platinum" /></div></td>
			<td>monthly #1</td>
			<td class="relative">
				{@render TableMap('jump_escape_rc4')}
			</td>
			<td>3:53.00</td>
			<td>2nd</td>
			<td>9000</td>
			<td>10/24/25</td>
			<td>
				<a
					href="/formats/monthly/1"
					title="link"
					class="opacity-50 transition-opacity hover:opacity-100"
				>
					<span class="icon-[ri--link-m] size-6"></span>
				</a>
			</td>
		{/snippet}
	</Table>
</DataSection>

{#snippet testTrophy()}
	<div class="bg-jfgray-900 relative flex h-16 w-48 shrink-0 overflow-hidden">
		<span class="icon-[ant-design--trophy-outlined] text-ctp-teal absolute -left-8 z-0 size-16"
		></span>
		<div class="absolute z-10 size-full">
			<div class="ml-8 flex h-full flex-col justify-between p-1 pr-2">
				<div class="flex flex-col">
					<div class="flex items-center gap-2">
						<span class="text-ctp-yellow text-shadow-2xs text-shadow-ctp-yellow">1st</span>
						<span class="overflow-ellipsis text-nowrap text-2xl/4">jump_escape_rc4</span>
					</div>
					<span class="font-semibold">3:53.00</span>
				</div>
				<span class="text-2xl/4 opacity-75">monthly #1</span>
			</div>
		</div>
		<div
			class="mask-l-from-50% mask-l-to-85% size-full bg-[url(https://tempusplaza.xyz/map-backgrounds/jump_escape_rc4.jpg)] bg-cover bg-right bg-no-repeat opacity-50"
		></div>
	</div>
{/snippet}

{#snippet testBounty()}
	<div class="bg-jfgray-900 w-59 relative flex h-24 shrink-0 items-center overflow-hidden">
		<span
			class="icon-[ri--star-line] bg-linear-to-b from-ctp-lavender-950 to-ctp-lavender-50 absolute -left-10 z-0 size-24 bg-clip-content"
		></span>
		<div class="absolute z-10 size-full">
			<div class="ml-6 flex h-full flex-col p-1 pr-2">
				<div class="flex items-center gap-2">
					<span class="text-shadow-xs text-shadow-ctp-lavender">100 keys</span>
					<span class="text-2xl/4 opacity-90">for sub 4:00</span>
				</div>

				<span class="ml-9 overflow-ellipsis text-nowrap text-2xl/4">jump_escape_rc4</span>
				<span class="text-ctp-lavender ml-9 mt-3 text-4xl/6 font-semibold">3:53.00</span>
			</div>
		</div>
		<div
			class="mask-l-from-50% mask-l-to-75% size-full bg-[url(https://tempusplaza.xyz/map-backgrounds/jump_escape_rc4.jpg)] bg-cover bg-right bg-no-repeat opacity-50"
		></div>
	</div>
{/snippet}

{#snippet testQuest()}
	<div class="bg-jfgray-900 relative flex h-12 w-48 shrink-0 overflow-hidden">
		<span class="icon-[ri--treasure-map-line] text-ctp-teal absolute -left-4 z-0 size-12"></span>
		<div class="absolute z-10 size-full">
			<div class="flex h-full flex-col justify-between p-1 pr-2">
				<div class="ml-8 flex flex-col text-2xl/4">
					<span class="overflow-ellipsis text-nowrap">jump_escape_rc4</span>
					<span>completion</span>
				</div>
			</div>
		</div>
		<div
			class="mask-l-from-50% mask-l-to-75% size-full bg-[url(https://tempusplaza.xyz/map-backgrounds/jump_escape_rc4.jpg)] bg-cover bg-right bg-no-repeat opacity-50"
		></div>
	</div>
{/snippet}

{#snippet TableMap(map: string)}
	<span class="relative z-10 overflow-ellipsis text-nowrap">{map}</span>
	<div
		style={`background-image: url("https://tempusplaza.xyz/map-backgrounds/${map}.jpg")`}
		class="mask-x-from-80% absolute top-0 size-full bg-cover bg-right bg-no-repeat opacity-50"
	></div>
{/snippet}
