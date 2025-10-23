<script lang="ts">
	import type { Snippet } from 'svelte';

	interface Props {
		data: object[];
		header: Snippet;
		row: Snippet<object[]>;
	}

	let { data, header, row }: { data: object[]; header: Snippet; row: Snippet<[any]> } = $props();
</script>

<table class="bg-jfgray-900 mt-2 table-fixed rounded-t-lg">
	<thead>
		<tr class="border-ctp-lavender/50 h-8 border-b-2 text-2xl/4">
			{@render header()}
		</tr>
	</thead>

	<tbody>
		<tr class="even:bg-jfgray-800/75 odd:bg-jfgray-900 h-12 text-center">
			{#each data as d}
				{@render row(d)}
			{/each}
		</tr>
	</tbody>
</table>
