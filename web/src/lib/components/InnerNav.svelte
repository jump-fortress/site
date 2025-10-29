<script lang="ts">
	let { route, parentRoute, pages }: { route: string; parentRoute: string; pages: string[] } =
		$props();
</script>

<div class="bg-jfgray-800/90 sticky top-16 z-20 backdrop-blur-sm">
	<div class="hover:text-ctp-lavender-50/75 relative flex h-12 w-fit transition-colors">
		<hr
			class="border-1 border-jfgray-700 absolute bottom-3 left-0 right-0 m-auto w-11/12 rounded-full"
		/>
		{#each pages as name}
			{@render FormatPage(name)}
		{/each}
	</div>
</div>

{#snippet FormatPage(name: string)}
	<a
		href="/{parentRoute}/{name}"
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
