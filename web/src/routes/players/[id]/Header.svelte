<script lang="ts">
	import rocket from '$lib/assets/static/classes/rocket.png';
	import sticky from '$lib/assets/static/classes/sticky.png';
	import plaza from '$lib/assets/static/players/plaza.png';
	import tempus from '$lib/assets/static/players/tempus.png';
	import zigzagoon from '$lib/assets/static/players/zigzagoon.gif';
	import Link from '$lib/components/Link.svelte';
	import ClassSelect from '$lib/components/ClassSelect.svelte';
	import Points from './Points.svelte';
	import type { PlayerProfile } from '$lib/schema';

	let { player }: { player: PlayerProfile } = $props();
	let selected_class = $derived(player.preferred_class);
</script>

<div class="h-46 relative flex gap-4">
	<img class="size-46" src={player.steam_avatar_url} alt="" />
	<!--
	<img
		style="image-rendering: pixelated"
		src={zigzagoon}
		alt=""
		class="scale-200 absolute bottom-0"
	/>
  -->

	<div class="flex h-full flex-col justify-between">
		<!-- name and div -->
		<div class="flex flex-col">
			<span class="text-5xl/8">{player.display_name}</span>
			{#if selected_class === 'Soldier'}
				<!-- no division -->
				{#if player.soldier_division}
					<span class="text-division-{player.soldier_division.toLowerCase()}">
						{player.soldier_division} Soldier
					</span>
				{:else}
					<span>Unranked Soldier</span>
				{/if}
			{:else}
				<!-- no division -->
				{#if player.demo_division}
					<span class="text-division-{player.demo_division.toLowerCase()}">
						{player.demo_division} Demo
					</span>
				{:else}
					<span>Unranked Demo</span>
				{/if}
			{/if}
		</div>

		<!-- links and points -->
		<div class="flex flex-col gap-2">
			{#if player.tempus_id}
				<div class="text-ctp-blue/50 flex">
					<a
						href="https://tempusplaza.xyz/players/{player.tempus_id}"
						class="hover:text-ctp-blue flex items-end gap-1 pr-2 decoration-1 transition-colors hover:underline"
					>
						<img src={plaza} class="size-6" alt="" />
						<span class="flex">Plaza</span>
					</a>
					<a
						href="https://tempus2.xyz/players/{player.tempus_id}"
						class="hover:text-ctp-blue flex items-end gap-1 pl-2 decoration-1 transition-colors hover:underline"
					>
						<img src={tempus} class="size-6" alt="" />
						<span class="flex">Tempus</span>
					</a>
				</div>
			{/if}

			<Points {selected_class} soldier={player.soldier_points} demo={player.demo_points} />
		</div>
	</div>

	<!-- class select -->
	<ClassSelect bind:selected_class />
</div>
