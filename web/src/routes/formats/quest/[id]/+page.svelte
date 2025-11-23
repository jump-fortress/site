<script lang="ts">
	import DataSection from '$lib/components/DataSection.svelte';
	import Link from '$lib/components/Link.svelte';
	import Table from '$lib/components/table/Table.svelte';
	import TableMap from '$lib/components/TableMap.svelte';
	import rocket from '$lib/assets/static/classes/rocket.png';
	import sticky from '$lib/assets/static/classes/sticky.png';
</script>

<DataSection title={'Details'}>
	<span class="text-base">quest #1</span>
	<Table data={[{}]}>
		{#snippet header()}
			<th class="w-14"></th>
			<th class="w-1/10"></th>
			<th class="w-1/10">amount</th>
			<th class="w-1/12">claimed</th>
			<th>map</th>
			<th class="w-1/10">target time</th>
			<th class="w-1/6">placed by</th>
			<th class="w-1/10">end date</th>
		{/snippet}
		{#snippet row()}
			<td><img src={rocket} class="m-auto size-8" alt="" /></td>
			<td class="text-left text-ctp-teal">Platinum</td>
			<td class="text-shadow-ctp-lavender text-shadow-xs">5 keys</td>
			<td>0/20</td>
			<td class="relative"><TableMap map="jump_beef" /></td>
			<td>completion</td>
			<td>jrjrjr</td>
			<td>10/23/25</td>
		{/snippet}
	</Table>
</DataSection>

<DataSection title={'Results'}>
	<Table data={[{}]}>
		{#snippet header()}
			<th class="w-14">#</th>
			<th class=""></th>
			<th class="w-1/8">run time</th>
			<th class="w-1/10">run date</th>
			<th class="w-14">link</th>
		{/snippet}
		{#snippet row()}
			<td>1</td>
			<td class="text-left text-ctp-sky">mur</td>
			<td>00:54.23</td>
			<td>12/23/25</td>
			<td><Link href="/formats/quest/1" title="results" /></td>
		{/snippet}
	</Table>
</DataSection>
