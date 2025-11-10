<script lang="ts">
	import type { Snippet } from 'svelte';

	interface Props {
		data: object[];
		header: Snippet;
		row: Snippet<object[]>;
	}

	let { data, header, row }: Props = $props();
</script>

<table class="bg-jfgray-900 table-fixed overflow-hidden rounded-2xl">
	<thead>
		<tr class="border-ctp-lavender/50 h-6 border-b-2 text-2xl/4">
			{@render header()}
		</tr>
	</thead>

	<tbody>
		<tr class="even:bg-jfgray-800/75 odd:bg-jfgray-900 h-10 text-center text-2xl/4">
			{#each data as d}
				{@render row(d)}
			{/each}
		</tr>
	</tbody>
</table>
