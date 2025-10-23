<script lang="ts">
</script>

<div class="flex flex-col gap-1">
	<span class="text-ctp-lavender">Active Quests</span>
	{@render TestActiveQuests()}
</div>

<hr class="border-1 border-jfgray-700 rounded-full" />

<div class="flex flex-col gap-1">
	<span class="text-ctp-lavender">Past Quests</span>
	{@render TestPastQuests()}
</div>
<hr class="border-1 border-jfgray-700 rounded-full" />

{#snippet TestActiveQuests()}
	<table class="bg-jfgray-900 mt-2 table-fixed rounded-t-lg">
		<thead>
			<tr class="border-ctp-lavender/50 h-8 border-b-2">
				<th>a column</th>
			</tr>
		</thead>

		<tbody>
			<tr class="even:bg-jfgray-800/75 odd:bg-jfgray-900 h-12 text-center">
				<td>a value</td>
			</tr>
		</tbody>
	</table>
{/snippet}

{#snippet TestPastQuests()}
	<table class="bg-jfgray-900 mt-2 table-fixed rounded-t-lg">
		<thead>
			<tr class="border-ctp-lavender/50 h-8 border-b-2">
				<th>a column</th>
			</tr>
		</thead>

		<tbody>
			<tr class="even:bg-jfgray-800/75 odd:bg-jfgray-900 h-12 text-center">
				<td>a value</td>
			</tr>
		</tbody>
	</table>
{/snippet}
