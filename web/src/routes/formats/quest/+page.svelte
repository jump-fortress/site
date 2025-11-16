<script lang="ts">
	import DataSection from '$lib/components/DataSection.svelte';
	import Table from '$lib/components/Table.svelte';

	import rocket from '$lib/assets/static/classes/rocket.png';
	import sticky from '$lib/assets/static/classes/sticky.png';
	import TableMap from '$lib/components/TableMap.svelte';
	import Link from '$lib/components/Link.svelte';
	import DivisionTag from '$lib/components/DivisionTag.svelte';
</script>

<DataSection title="Active Quests">
	<Table data={[{}]}>
		{#snippet header()}
			<th class="w-12"></th>
			<th class="w-1/10"></th>
			<th class="w-1/10">amount</th>
			<th class="w-1/12">claimed</th>
			<th>map</th>
			<th class="w-1/10">target time</th>
			<th class="w-1/6">placed by</th>
			<th class="w-1/10">end date</th>
			<th class="w-14">results</th>
		{/snippet}
		{#snippet row()}
			<td><img src={rocket} class="m-auto size-8" alt="" /></td>
			<td><div class="flex justify-center"><DivisionTag div="diamond" /></div></td>
			<td class="text-shadow-xs text-shadow-ctp-lavender">5 keys</td>
			<td>0/20</td>
			<td class="relative"><TableMap map="jump_beef" /></td>
			<td>completion</td>
			<td>jrjrjr</td>
			<td>10/23/25</td>
			<td><Link href="ababa" title="ababa" /></td>
		{/snippet}
	</Table>
</DataSection>

<DataSection title="Past Quests">
	<Table data={[{}]}>
		{#snippet header()}
			<th class="w-12"></th>
			<th class="w-1/10"></th>
			<th class="w-1/10">amount</th>
			<th class="w-1/12">claimed</th>
			<th>map</th>
			<th class="w-1/10">target time</th>
			<th class="w-1/6">placed by</th>
			<th class="w-1/10">end date</th>
			<th class="w-14">results</th>
		{/snippet}
		{#snippet row()}
			<td><img src={rocket} class="m-auto size-8" alt="" /></td>
			<td><div class="flex justify-center"><DivisionTag div="diamond" /></div></td>
			<td class="text-shadow-xs text-shadow-ctp-lavender">5 keys</td>
			<td>0/20</td>
			<td class="relative"><TableMap map="jump_beef" /></td>
			<td>completion</td>
			<td>jrjrjr</td>
			<td>10/23/25</td>
			<td><Link href="/formats/quest/1" title="results" /></td>
		{/snippet}
	</Table>
</DataSection>
