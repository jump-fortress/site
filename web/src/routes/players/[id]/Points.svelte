<script lang="ts">
	import type { PlayerPoints } from '$lib/schema';

	type Props = {
		selected_class: string;
		soldier: PlayerPoints;
		demo: PlayerPoints;
	};

	let { selected_class, soldier, demo }: Props = $props();
</script>

<div class="flex gap-2 leading-6">
	<div class="flex flex-col bg-jfgray-900 px-2 py-1">
		<span class="text-base opacity-75">last 3 monthlies</span>
		<div class="flex items-center gap-2">
			{#if selected_class === 'Soldier'}
				<span>{soldier.last_3_monthly}</span>
				<span class="ml-auto">(#1)</span>
			{:else}
				<span>{demo.last_3_monthly}</span>
				<span class="ml-auto">(#1)</span>
			{/if}
		</div>
	</div>
	<div class="flex flex-col bg-jfgray-900 px-2 py-1">
		<span class="text-base opacity-75">last 9 MOTWs</span>
		<div class="flex items-center gap-2">
			{#if selected_class === 'Soldier'}
				<span>{soldier.last_9_motw}</span>
				<span class="ml-auto">(#1)</span>
			{:else}
				<span>{demo.last_9_motw}</span>
				<span class="ml-auto">(#1)</span>
			{/if}
		</div>
	</div>
	<div class="flex flex-col bg-jfgray-900 px-2 py-1">
		<span class="text-base opacity-75">total points</span>
		<div class="flex items-center gap-2">
			{#if selected_class === 'Soldier'}
				<span>{soldier.total}</span>
				<span class="ml-auto">(#1)</span>
			{:else}
				<span>{demo.total}</span>
				<span class="ml-auto">(#1)</span>
			{/if}
		</div>
	</div>
</div>
