<script lang="ts">
	let { route }: { route: string } = $props();
</script>

<div class="hover:text-ctp-lavender-50/75 relative flex h-16 w-fit transition-colors">
	<hr
		class="border-1 border-jfgray-700 absolute bottom-5 left-0 right-0 m-auto w-11/12 rounded-full"
	/>
	{@render FormatPage('monthly')}
	{@render FormatPage('motw')}
	{@render FormatPage('quest')}
	{@render FormatPage('bounty')}
</div>

{#snippet FormatPage(name: string)}
	<a
		href="/formats/{name}"
		class="hover:text-ctp-lavender-50 flex h-full items-center px-2 transition-colors first:pl-0"
		><div class="relative">
			<span>{name}</span>
			{#if route.includes(name)}
				<hr
					class="text-ctp-lavender border-1 absolute left-0 right-0 m-auto flex w-11/12 rounded-full"
				/>
			{/if}
		</div></a
	>
{/snippet}
