<script lang="ts">
	import type { Snippet } from 'svelte';

	interface Props {
		data: object[];
		header: Snippet;
		row: Snippet<object[]>;
	}

	let { data, header, row }: Props = $props();
</script>

<table class="bg-jfgray-900 table-fixed rounded-t-2xl">
	<thead>
		<tr class="border-ctp-lavender/50 h-6 border-b-2 text-base leading-4">
			{@render header()}
		</tr>
	</thead>

	<tbody>
		{#each data as d}
			<tr class="even:bg-jfgray-800/25 odd:bg-jfgray-900 h-10 text-center text-base">
				{@render row(d)}
			</tr>
		{/each}
	</tbody>
</table>
