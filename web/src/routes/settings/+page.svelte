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
		<label
			for="update_display_name"
			class="border-jfgray-700 focus-within:border-ctp-lavender-50/50 hover:border-ctp-lavender-50/50 relative mt-2 w-80 text-nowrap border-2 transition-colors"
		>
			<span class="bg-jfgray-800 leading-1 absolute -top-0.5 left-2 px-1"
				>request display name change
			</span>
			<div class="flex h-10">
				<input
					class="text-ctp-lavender focus:bg-jfgray-900 bg-jfgray-800 group peer w-full bg-clip-padding p-1 transition-colors"
					id="update_display_name"
					type="text"
				/>
				<!-- svelte-ignore a11y_consider_explicit_label -->
				<button
					class="bg-jfgray-800 peer-focus:bg-jfgray-800 flex h-full w-12 cursor-pointer items-center justify-center transition-colors"
				>
					<span class="icon-[ri--send-plane-line]"></span>
				</button>
			</div>
		</label>

		<div class="h-18 border-jfgray-700 relative mt-2 flex w-fit text-nowrap border-2">
			<span class="bg-jfgray-800 leading-1 absolute -top-1 left-2 z-10 px-1">fav class</span>
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
						: 'opacity-50'} h-full w-20 cursor-pointer object-contain pb-2 pt-3 transition-opacity hover:opacity-100"
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
						: 'opacity-50'} h-full w-20 cursor-pointer object-contain pb-2 pt-3 transition-opacity hover:opacity-100"
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
		@apply opacity-100;
	}
</style>
