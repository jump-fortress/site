<script lang="ts">
	let { selected_class = $bindable() } = $props();

	import rocket from '$lib/assets/static/classes/rocket.png';
	import sticky from '$lib/assets/static/classes/sticky.png';
	import { classToEnum } from '$lib/enums';
	import { Client } from '$lib/internalApi';
	import type { PathsInternalPlayersPreferredclassClassPutParametersPathClass as Class } from '$lib/schema.d.ts';

	const style =
		'filter: brightness(0) saturate(100%) invert(97%) sepia(49%) saturate(6023%) hue-rotate(179deg) brightness(100%) contrast(108%)';

	async function handleUpdatePreferredClass(selectedClass: Class) {
		const { response } = await Client.PUT('/internal/players/preferredclass/{class}', {
			fetch: fetch,
			params: {
				path: { class: selectedClass }
			}
		});

		// if (response.ok) {
		// 	pageStore.preferredClass.set = true;
		// 	setTimeout(() => {
		// 		pageStore.preferredClass.set = false;
		// 	}, 2000);
		// }
	}
</script>

<div class="ml-auto flex h-full flex-col">
	<button
		onclick={() => {
			selected_class = 'Soldier';
			handleUpdatePreferredClass(classToEnum('Soldier'));
		}}
		class="{selected_class === 'Soldier'
			? 'bg-jfgray-900 opacity-100'
			: 'opacity-50'} flex basis-1/2 cursor-pointer items-center p-3 pt-4 transition-all hover:opacity-100"
	>
		<img class="size-16 select-none" {style} src={rocket} alt="" draggable="false" />
	</button>
	<button
		onclick={() => {
			selected_class = 'Demo';
			handleUpdatePreferredClass(classToEnum('Demo'));
		}}
		class="{selected_class === 'Demo'
			? 'bg-jfgray-900 opacity-100'
			: 'opacity-50'} flex basis-1/2 cursor-pointer items-center p-3 pb-4 transition-all hover:opacity-100"
	>
		<img class="size-16 select-none" {style} src={sticky} alt="" draggable="false" />
	</button>
</div>
