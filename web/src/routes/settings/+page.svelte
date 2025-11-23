<script lang="ts">
	import DataSection from '$lib/components/DataSection.svelte';
	import rocket from '$lib/assets/static/classes/rocket.png';
	import sticky from '$lib/assets/static/classes/sticky.png';
	import type { PageData } from './$types';
	import { updatePreferredClass } from '$lib/internalApi';
	import { classToEnum } from '$lib/enums';

	let { data }: { data: PageData } = $props();
	let player = $derived(data.player);
	let favoriteClass = $derived(player.preferred_class);
</script>

<DataSection title="Profile">
	<div class="flex flex-col gap-2">
		<label for="update_display_name" class="relative mt-2 w-72 text-nowrap">
			<span class="absolute -top-0.5 left-2 bg-jfgray-800 px-1 leading-1"
				>request display name change
			</span>
			<input
				class="w-full border-2 border-ctp-lavender-50/50 p-1 text-ctp-lavender hover:bg-jfgray-700 focus:bg-jfgray-700"
				id="update_display_name"
				type="text"
			/>
		</label>

		<div class="relative mt-2 flex h-18 w-fit border-2 border-ctp-lavender-50/50 text-nowrap">
			<span class="absolute -top-1 left-2 z-10 bg-jfgray-800 px-1 leading-1">fav class</span>
			<button
				onclick={() => {
					// todo: use api response to update on client?
					updatePreferredClass(classToEnum('Soldier'));
					favoriteClass = 'Soldier';
				}}
			>
				<img
					src={rocket}
					alt=""
					class="{favoriteClass === 'Soldier'
						? 'bg-jfgray-900 opacity-100'
						: 'opacity-50'} h-15 w-18 cursor-pointer object-contain pt-2 pb-1 transition-opacity"
				/>
			</button>
			<button
				onclick={() => {
					updatePreferredClass(classToEnum('Demo'));
					favoriteClass = 'Demo';
				}}
			>
				<img
					src={sticky}
					alt=""
					class="{favoriteClass === 'Demo'
						? 'bg-jfgray-900 opacity-100'
						: 'opacity-50'} h-15 w-18 cursor-pointer object-contain pt-2 pb-1 transition-opacity"
				/>
			</button>
		</div>

		<button class="settings-button">update avatar from steam</button>
	</div>
</DataSection>

<DataSection title={'Rank'}>
	{#if !player.soldier_division}
		<button class="settings-button">request soldier placement</button>
	{/if}
	{#if !player.demo_division}
		<button class="settings-button">request demo placement</button>
	{/if}
</DataSection>

<style lang="postcss">
	@reference "../../app.css";

	.settings-button {
		@apply size-fit cursor-pointer rounded-lg border-b-2 border-ctp-lavender-50/50 bg-ctp-lavender-950 px-2 py-1;
	}
</style>
