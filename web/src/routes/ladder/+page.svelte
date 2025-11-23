<script lang="ts">
	import rocket from '$lib/assets/static/classes/rocket.png';
	import sticky from '$lib/assets/static/classes/sticky.png';
	import Table from '$lib/components/table/Table.svelte';

	let sortBy = $state('monthly');
</script>

<div class="flex justify-between">
	<!-- sort by -->
	<div class="flex flex-col justify-between">
		<div class="flex h-full items-center">
			<span class="text-3xl">soldier ladder</span>
		</div>
		<div class="flex items-end transition-colors hover:text-ctp-lavender-50/75">
			<span class="mr-2 text-ctp-lavender">sort by</span>
			<div class="relative flex">
				<hr
					class="absolute right-0 -bottom-0.5 left-1 z-0 m-auto w-17/18 rounded-full border-1 border-jfgray-700"
				/>
				{@render SortButton('last 3 monthlies', 'monthly')}
				{@render SortButton('last 9 MOTWs', 'motw')}
				{@render SortButton('total points', 'points')}
			</div>
		</div>
	</div>

	<!-- class select -->
	<div class="flex h-full">
		<div
			class="flex basis-1/2 cursor-pointer items-center bg-jfgray-900 p-3 pl-4 opacity-100 hover:opacity-100"
		>
			<img
				class="size-max select-none"
				style="filter: brightness(0) saturate(100%) invert(97%) sepia(49%) saturate(6023%) hue-rotate(179deg) brightness(100%) contrast(108%)"
				src={rocket}
				alt=""
				draggable="false"
			/>
		</div>
		<div
			class="flex basis-1/2 cursor-pointer items-center p-3 pr-4 opacity-50 transition-opacity hover:opacity-100"
		>
			<img
				class="size-max"
				style="filter: brightness(0) saturate(100%) invert(97%) sepia(49%) saturate(6023%) hue-rotate(179deg) brightness(100%) contrast(108%)"
				src={sticky}
				alt=""
				draggable="false"
			/>
		</div>
	</div>
</div>

<Table data={[{}]}>
	{#snippet header()}
		<th class="w-1/14">rank</th>
		<th class="w-1/14">div rank</th>
		<th></th>
		<th class="w-1/12 text-ctp-lavender-50/75">monthlies played</th>
		<th class="w-1/12 text-ctp-lavender-50/75">MOTWs played</th>
		<th class="w-1/8">last 3 monthlies</th>
		<th class="w-1/8">last 9 MOTWs</th>
		<th class="w-1/8 text-ctp-lavender-50/75">total points</th>
		<th></th>
	{/snippet}
	{#snippet row()}
		<td>#1</td>
		<td>#1</td>
		<td class="pl-4 text-start text-ctp-sky">vice</td>
		<td class="text-ctp-lavender-50/75">4</td>
		<td class="text-ctp-lavender-50/75">10</td>
		<td>300000</td>
		<td>900000</td>
		<td class="text-ctp-lavender-50/75">48732894</td>
	{/snippet}
</Table>

{#snippet SortButton(name: string, sort: string)}
	<button
		class="relative cursor-pointer px-2 transition-colors hover:text-ctp-lavender-50"
		onclick={() => {
			sortBy = sort;
		}}
	>
		<span>{name}</span>
		{#if sortBy === sort}
			<hr
				class="absolute right-0 left-0 m-auto flex w-5/6 rounded-full border-1 text-center text-ctp-lavender"
			/>
		{/if}
	</button>
{/snippet}
